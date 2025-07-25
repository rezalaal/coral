package validation

import "regexp"

// بررسی اینکه نام تنها شامل حروف انگلیسی و فارسی باشد
func IsAlpha(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z\x{0600}-\x{06FF}\s]+$`).MatchString(s)
}

// بررسی اعتبار نام
func IsValidName(name string) bool {
	if !IsAlpha(name) {
		return false
	}
	if len(name) == 0 || len(name) > 80 {
		return false
	}
	return true
}


// بررسی اعتبار شماره موبایل
func IsValidMobile(mobile string) bool {
	return regexp.MustCompile(`^\d{11}$`).MatchString(mobile)
}
