package luhn

import (
	"strings"
	"unicode"
)

// Valid is a card number valid?
func Valid(card string) bool {
	// fmt.Print("Valid? ", card, "\t")

	sum := 0

	// remove spaces
	card = strings.ReplaceAll(card, " ", "")

	// count down instead of up
	count := len(card)
	if count < 2 {
		return false
	}

	for _, ch := range card {

		// fail on non digit
		if !unicode.IsDigit(ch) {
			return false
		}

		// int('0') - int('0') = 0, '1' - '0' = 1
		n := int(ch) - int('0')

		// double every second digit
		// starting from the right
		isSecond := count%2 == 0
		if isSecond {

			// If doubling the number results greater than 9
			// then subtract 9 from the product
			prod := 2 * n
			if prod > 9 {
				prod -= 9
			}
			//sum += prod
			n = prod
		}
		sum += n
		count--
	}

	result := sum%10 == 0
	return result
}
