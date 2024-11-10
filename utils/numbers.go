package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func PersianToEnglishDigits(persianNum string) (int, error) {

	persianNum = strings.Trim(persianNum, " ")
	// Mapping of Persian digits to English digits
	persianToEnglish := map[rune]rune{
		'۰': '0', '۱': '1', '۲': '2', '۳': '3', '۴': '4',
		'۵': '5', '۶': '6', '۷': '7', '۸': '8', '۹': '9',
	}

	// Remove any commas
	cleanedNum := strings.ReplaceAll(persianNum, "٬", "")
	cleanedNum = strings.ReplaceAll(cleanedNum, ",", "")

	// Convert Persian digits to English digits
	var englishNum string
	for _, ch := range cleanedNum {

		if unicode.IsDigit(ch) {
			if engDigit, ok := persianToEnglish[ch]; ok {

				englishNum += string(engDigit)
			} else {
				return 0, fmt.Errorf("invalid character found: %c", ch)
			}
		} else {
			return 0, fmt.Errorf("invalid character found: %c", ch)
		}
	}

	// Convert the English number string to an integer
	num, err := strconv.Atoi(englishNum)

	if err != nil {
		return 0, err
	}
	return num, nil
}
