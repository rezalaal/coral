package integration_test

import (
	"bytes"
	"encoding/json"
	
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rezalaal/coral/internal/integration"
	"github.com/rezalaal/coral/internal/user/models"
)

func TestUserIntegration(t *testing.T) {
	server, teardown := integration.SetupTestServer(t)
	defer teardown()

	// ساخت کاربر با داده‌های معتبر
	userPayload := map[string]string{
		"name":          "Integration User", // نام معتبر
		"mobile":        "09121234567",      // شماره موبایل معتبر
		"password_hash": "hashedpass",       // پسورد هَش‌شده معتبر
	}
	payloadBytes, _ := json.Marshal(userPayload)

	// ارسال درخواست POST برای ساخت کاربر
	resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader(payloadBytes))
	require.NoError(t, err)

	// بررسی وضعیت پاسخ
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected HTTP Status: 201 Created")

	var createdUser models.User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	require.NoError(t, err)

	// بررسی داده‌های کاربر ایجاد شده
	assert.Equal(t, "Integration User", createdUser.Name, "User name mismatch")

	// بستن body پاسخ پس از استفاده
	resp.Body.Close()

	// دریافت لیست کاربران با دقت بیشتر در تشخیص خطا
	resp, err = http.Get(server.URL + "/users")
	require.NoError(t, err)

	// بررسی وضعیت پاسخ دریافت کاربران
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected HTTP Status: 200 OK")

	var users []models.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	require.NoError(t, err)

	// بستن body پس از خواندن
	resp.Body.Close()

	// بررسی اینکه حداقل یک کاربر وجود دارد
	assert.GreaterOrEqual(t, len(users), 1, "Expected at least one user")
}
