package integration_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rezalaal/coral/internal/integration"
)

func TestSQLInjection(t *testing.T) {
	server, teardown := integration.SetupTestServer(t)
	defer teardown()

	// تست SQL Injection برای نام کاربر
	userPayload := map[string]string{
		"name":          "' OR 1=1 --", // SQL Injection در فیلد نام
		"mobile":        "09121234567",
		"password_hash": "hashedpass",
	}
	payloadBytes, _ := json.Marshal(userPayload)

	// ارسال درخواست POST برای ساخت کاربر
	resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader(payloadBytes))
	require.NoError(t, err)

	// بررسی اینکه پاسخ باید خطا باشد چون SQL Injection پذیرفته نمی‌شود
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Expected SQL Injection to be blocked")

	// بررسی اینکه آیا خطای مشخصی برگشت داده شده است
	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), "نام معتبر نیست", "Expected specific error message")
}
