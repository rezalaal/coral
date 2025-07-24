// internal/utils/utils_test.go
package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConvertPersianToEnglishNumbers(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"۰۱۲۳۴۵۶۷۸۹", "0123456789"},
		{"Hello 1234", "Hello 1234"},
		{"سلام ۱۲۳۴", "سلام 1234"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// فراخوانی تابع از پکیج utils
			result := ConvertPersianToEnglishNumbers(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
