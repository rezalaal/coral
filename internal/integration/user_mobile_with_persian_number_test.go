package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"coral/internal/integration"
	"coral/internal/user/models"
)

func TestCreateUser_WithPersianNumbers(t *testing.T) {
	server, _, teardown := integration.SetupTestServer(t)
	defer teardown()

	// ساخت کاربر با شماره موبایل به فارسی
	userPayload := map[string]string{
		"name":          "Integration User",
		"mobile":        "۰۹۱۲۱۲۳۴۵۶۷", // شماره موبایل با اعداد فارسی
		"password_hash": "hashedpass",
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
	assert.Equal(t, "09121234567", createdUser.Mobile, "Mobile number mismatch")

	// بستن body پاسخ پس از استفاده
	resp.Body.Close()
}
