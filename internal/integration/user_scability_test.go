package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
	"database/sql"
	_ "github.com/lib/pq" // برای اتصال به PostgreSQL

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rezalaal/coral/internal/integration"
)

func generateRandomName(length int) string {
	// تعیین مجموعه کاراکترهایی که می‌خواهیم در نام تصادفی باشند
	allowedChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // حروف انگلیسی	

	var name strings.Builder
	rand.Seed(time.Now().UnixNano()) // برای تولید تصادفی واقعی

	// ایجاد نام تصادفی با طول مشخص
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(allowedChars)) // انتخاب تصادفی یک کاراکتر
		name.WriteByte(allowedChars[randomIndex]) // افزودن کاراکتر به نام
	}

	return name.String()
}

func generateRandomMobile(db *sql.DB) string {
	rand.Seed(time.Now().UnixNano()) // برای تولید تصادفی واقعی
	mobile := "09" // شروع شماره موبایل با "09" برای مطابقت با ایران
	for i := 0; i < 9; i++ {
		mobile += fmt.Sprintf("%d", rand.Intn(10)) // تولید یک رقم تصادفی
	}

	// اطمینان از اینکه شماره موبایل منحصر به فرد است
	for !isMobileUnique(mobile, db) {
		// اگر شماره موبایل تکراری بود، دوباره تولید می‌کنیم
		mobile = "09"
		for i := 0; i < 9; i++ {
			mobile += fmt.Sprintf("%d", rand.Intn(10))
		}
	}
	return mobile
}

// تابع برای بررسی اینکه شماره موبایل در پایگاه داده موجود است یا خیر
func isMobileUnique(mobile string, db *sql.DB) bool {
	// ایجاد کوئری برای جستجوی شماره موبایل در پایگاه داده
	var count int
	query := `SELECT COUNT(*) FROM users WHERE mobile = $1`
	err := db.QueryRow(query, mobile).Scan(&count)
	if err != nil {
		// اگر مشکلی در اجرا پیش بیاید، خطا را ثبت می‌کنیم و فرض می‌کنیم که شماره موبایل یکتا است
		return true
	}
	// اگر نتیجه بیش از صفر بود، یعنی شماره موبایل تکراری است
	return count == 0
}

func TestScalabilityWithLargeData(t *testing.T) {
	server, _, teardown := integration.SetupTestServer(t)
	defer teardown()

	// ایجاد اتصال به پایگاه داده برای بررسی یکتایی شماره موبایل
	db, err := sql.Open("postgres", "your_connection_string_here")
	require.NoError(t, err)
	defer db.Close()

	// تعداد داده‌ها (کاربران) که می‌خواهیم تست کنیم
	numUsers := 1000 // یا می‌توانید از مقادیر بزرگتر استفاده کنید، مانند 10000

	// تست زمان شروع
	startTime := time.Now()

	for i := 0; i < numUsers; i++ {
		name := generateRandomName(10)
		randomMobile := generateRandomMobile(db) // شماره موبایل منحصر به فرد

		// ساختن داده‌های کاربر
		userPayload := map[string]string{
			"name":          name,
			"mobile":        randomMobile,
			"password_hash": "hashedpass",
		}

		// لاگ ورودی‌ها
		t.Logf("Payload being sent: %v", userPayload)

		payloadBytes, _ := json.Marshal(userPayload)

		// ارسال درخواست POST برای ساخت کاربر
		resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader(payloadBytes))
		require.NoError(t, err)

		// بررسی وضعیت پاسخ
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("Failed to create user %s, got status %d", name, resp.StatusCode)
		}

		// بستن body پاسخ پس از استفاده
		resp.Body.Close()
	}

	// زمان صرف‌شده برای انجام تست
	duration := time.Since(startTime)

	// نمایش زمان اجرای تست
	t.Logf("Time taken to create %d users: %s", numUsers, duration)

	// بررسی اینکه زمان تست منطقی است (مثلا، اگر خیلی طول کشید، باید بهینه‌سازی صورت گیرد)
	assert.Less(t, duration.Seconds(), float64(numUsers*2), "Test took too long. Consider optimizing.")
}
