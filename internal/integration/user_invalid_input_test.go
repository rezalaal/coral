package integration

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/rezalaal/coral/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateUser_InvalidJSON(t *testing.T) {
	server := test.NewTestServer()
	defer server.Close()

	tests := []struct {
		name       string
		payload    string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "ناقص بودن JSON",
			payload:    `{"name": "علی", "mobile": "0935...`, // ناقص و بسته نشده
			wantStatus: http.StatusBadRequest,
			wantBody:   "نامعتبر",
		},
		{
			name:       "مقدار اشتباه برای name",
			payload:    `{"name": 123, "mobile": "09351234567"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   "نامعتبر",
		},
		{
			name:       "عدم وجود فیلد name",
			payload:    `{"mobile": "09351234567"}`,
			wantStatus: http.StatusInternalServerError, // چون فرض ما اینه در repo باید validate بشه
			wantBody:   "خطا در ایجاد کاربر",
		},
		{
			name:       "عدم وجود فیلد mobile",
			payload:    `{"name": "کاربر بدون موبایل"}`,
			wantStatus: http.StatusInternalServerError,
			wantBody:   "خطا در ایجاد کاربر",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Post(server.URL+"/users", "application/json", strings.NewReader(tt.payload))
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, tt.wantStatus, resp.StatusCode)

			data, _ := io.ReadAll(resp.Body)
			assert.Contains(t, string(data), tt.wantBody)
		})
	}
}
