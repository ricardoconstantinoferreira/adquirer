package validation

import "strings"

func IsValidLuhn(card string) bool {
	cleanCard := strings.ReplaceAll(card, " ", "")
	if len(cleanCard) == 0 {
		return false
	}

	sum := 0
	shouldDouble := false

	for i := len(cleanCard) - 1; i >= 0; i-- {
		digit := cleanCard[i]
		if digit < '0' || digit > '9' {
			return false
		}

		n := int(digit - '0')
		if shouldDouble {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += n
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}
