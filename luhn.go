package luhn

import (
	"strings"
	"unicode"
)

// Valid is a card number valid?
func Valid(card string) bool {

	card = strings.ReplaceAll(card, " ", "")

	count := len(card)
	if count < 2 {
		return false
	}

	sum := 0
	for _, ch := range card {

		if !unicode.IsDigit(ch) {
			return false
		}

		n := int(ch) - int('0')
		// double every second digit
		isSecond := count%2 == 0
		if isSecond {

			prod := 2 * n
			if prod > 9 {
				prod -= 9
			}
			n = prod
		}
		sum += n
		count--
	}

	return sum%10 == 0
}
