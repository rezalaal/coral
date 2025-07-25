package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlpha(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"سلام", true},               // نام فارسی معتبر
		{"Hello", true},              // نام انگلیسی معتبر
		{"سلام Hello", true},         // ترکیب فارسی و انگلیسی
		{"سلام123", false},           // حروف با اعداد ترکیب شده
		{"Hello123", false},          // حروف با اعداد ترکیب شده
		{"", false},                  // رشته خالی
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsAlpha(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestIsValidName(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"سلام", true},                     // نام فارسی معتبر
		{"Hello", true},                    // نام انگلیسی معتبر
		{"سلام Hello", true},               // ترکیب فارسی و انگلیسی
		{"", false},                        // نام خالی
		{"This is a very long name that exceeds the hundred character limit to check the validation", false}, // نام طولانی تر از 100 کاراکتر
		{"نام123", false},                  // ترکیب حروف با اعداد
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsValidName(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestIsValidMobile(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"09121234567", true},      // شماره موبایل معتبر
		{"0912123456", false},      // شماره موبایل با طول اشتباه
		{"091212345678", false},    // شماره موبایل با طول اشتباه
		{"0912abc567", false},      // شماره موبایل با حروف
		{"", false},                // شماره موبایل خالی
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsValidMobile(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
