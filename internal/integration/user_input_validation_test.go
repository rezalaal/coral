package integration_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/rezalaal/coral/internal/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_InputValidation(t *testing.T) {
	server, _, teardown := integration.SetupTestServer(t)
	defer teardown()

	tests := []struct {
		name       string
		payload    string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "نام معتبر نیست",
			payload:    `{"name": "1234", "mobile": "09351234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "نام معتبر نیست",
		},
		{
			name:       "شماره موبایل نامعتبر",
			payload:    `{"name": "علی", "mobile": "0935"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "شماره موبایل نامعتبر است",
		},
		{
			name:       "ورودی ناقص (فیلد نام گم شده)",
			payload:    `{"mobile": "09351234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند",
		},
		{
			name:       "ورودی ناقص (فیلد موبایل گم شده)",
			payload:    `{"name": "علی"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند",
		},
		{
			name:       "نام بسیار طولانی",
			payload:    `{"name": "این یک نام بسیار طولانی است که از حد مجاز عبور می‌کند و باید خطا بدهد", "mobile": "09351234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "نام معتبر نیست",
		},
		{
			name:       "نام بسیار طولانی",
			payload:    `{"name": "این یک نام بسیار طولانی است که از حد مجاز عبور می‌کند و باید خطا بدهد", "mobile": "09121234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "نام معتبر نیست",
		},
		{
			name:       "نام با کاراکتر غیرمجاز",
			payload:    `{"name": "Ali123", "mobile": "09121234567"}`, // کاراکتر عددی در نام
			wantStatus: http.StatusBadRequest,
			wantBody:   "نام معتبر نیست",
		},
		{
			name:       "شماره موبایل کوتاه",
			payload:    `{"name": "علی", "mobile": "0912"}`, // موبایل کوتاه‌تر از حد مجاز
			wantStatus: http.StatusBadRequest,
			wantBody:   "شماره موبایل نامعتبر است",
		},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Post(server.URL+"/users/create", "application/json", bytes.NewReader([]byte(tt.payload)))
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			body, _ := io.ReadAll(resp.Body)
			assert.Contains(t, string(body), tt.wantBody)
		})
	}
}
