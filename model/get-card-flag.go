package model

import "strings"

func CardFlagByNumber(cardNumber string) string {
	number := normalizeCardNumber(cardNumber)
	length := len(number)

	if length == 0 {
		return "Unknown"
	}

	if strings.HasPrefix(number, "4") && (length == 13 || length == 16 || length == 19) {
		return "Visa"
	}

	if length == 16 {
		prefix2 := number[:2]
		prefix4 := number[:4]

		if prefix2 >= "51" && prefix2 <= "55" {
			return "Mastercard"
		}

		if prefix4 >= "2221" && prefix4 <= "2720" {
			return "Mastercard"
		}
	}

	if length == 15 {
		prefix2 := number[:2]
		if prefix2 == "34" || prefix2 == "37" {
			return "American Express"
		}
	}

	if (length == 16 || length == 19) && strings.HasPrefix(number, "606282") {
		return "Hipercard"
	}

	if (length == 16 || length == 19) &&
		(strings.HasPrefix(number, "6011") ||
			strings.HasPrefix(number, "65") ||
			(number[:3] >= "644" && number[:3] <= "649")) {
		return "Discover"
	}

	if length == 14 {
		prefix3 := number[:3]
		prefix2 := number[:2]

		if (prefix3 >= "300" && prefix3 <= "305") || prefix2 == "36" || (prefix2 >= "38" && prefix2 <= "39") {
			return "Diners Club"
		}
	}

	if (length == 16 || length == 19) &&
		(strings.HasPrefix(number, "5067") ||
			strings.HasPrefix(number, "509") ||
			strings.HasPrefix(number, "627780") ||
			strings.HasPrefix(number, "636297") ||
			strings.HasPrefix(number, "650") ||
			strings.HasPrefix(number, "6516") ||
			strings.HasPrefix(number, "6550")) {
		return "Elo"
	}

	return "Unknown"
}

func normalizeCardNumber(cardNumber string) string {
	replacer := strings.NewReplacer(" ", "", "-", "")
	return replacer.Replace(strings.TrimSpace(cardNumber))
}
