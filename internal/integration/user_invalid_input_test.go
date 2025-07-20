package integration_test

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"github.com/rezalaal/coral/internal/integration"
)

func TestCreateUser_InvalidInput(t *testing.T) {
	server, teardown := integration.SetupTestServer(t)
	defer teardown()

	tests := []struct {
		name       string
		payload    string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "ناقص بودن JSON",
			payload:    `{"name": "علی", "mobile": "0935...`, // JSON ناقص
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطای تجزیه‌ی JSON",
		},
		{
			name:       "مقدار اشتباه برای name",
			payload:    `{"name": 123, "mobile": "09351234567"}`, // نوع نامعتبر
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطای تجزیه‌ی JSON",
		},
		{
			name:       "عدم وجود فیلد name",
			payload:    `{"mobile": "09351234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند",
		},
		{
			name:       "عدم وجود فیلد mobile",
			payload:    `{"name": "کاربر بدون موبایل"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Post(server.URL+"/users/create", "application/json", strings.NewReader(tt.payload))
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			assert.Contains(t, string(body), tt.wantBody)
		})
	}
}
