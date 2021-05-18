package luhn

import (
	"strings"
	"unicode"
)

// Valid is a card number valid?
func Valid(card string) bool {

	card = strings.ReplaceAll(card, " ", "")
	if len(card) < 2 {
		return false
	}

	isSecond := len(card)%2 == 0
	sum := 0
	for _, ch := range card {

		if !unicode.IsDigit(ch) {
			return false
		}

		n := int(ch - '0')
		// double every second digit
		if isSecond {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		isSecond = !isSecond
	}

	return sum%10 == 0
}
