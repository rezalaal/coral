package utils

// تبدیل اعداد فارسی به انگلیسی
func ConvertPersianToEnglishNumbers(input string) string {
	replacement := map[rune]rune{
		'۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4',
		'۵': '5', '۶': '6', '۷': '7', '۸': '8', '۹': '9',
	}

	var result []rune
	for _, char := range input {
		if englishChar, exists := replacement[char]; exists {
			result = append(result, englishChar)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
