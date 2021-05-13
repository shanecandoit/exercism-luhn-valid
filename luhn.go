package luhn

import (
	"fmt"
	"strings"
	"unicode"
)

// ReverseString reverse a given string
func ReverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Valid is a card number valid?
func Valid(card string) bool {
	fmt.Print("Valid? ", card, "\t")

	sum := 0

	// remove spaces
	// ReplaceAll(s, old, new string)
	card = strings.ReplaceAll(card, " ", "")
	// fmt.Println("no spaces", card)

	// count down instead of up
	//count := 1
	count := len(card)
	if count < 2 {
		// fmt.Println("too few digits", count)
		fmt.Println("=", false)
		return false
	}

	// str := strings.Split(card, "")
	for _, ch := range card {
		// fmt.Println("ch", ch, string(ch))

		if unicode.IsDigit(ch) {
			//if ch <= '0' && ch >= '9' {
			n := int(ch) - int('0')

			// double every second digit
			// starting from the right
			isSecond := count%2 == 0
			if isSecond {

				//If doubling the number results greater than 9
				// then subtract 9 from the product
				prod := 2 * n
				if prod > 9 {
					prod -= 9
				}

				sum += prod
			} else {
				sum += n
			}
			// fmt.Println("digit", ch, n, "sum", sum, "sec?", isSecond)

			//count++
			count--
		} else {
			// fmt.Println("FAIL. non-digit non-space", ch)
			fmt.Println("=", false)
			return false
		}
	}

	// fmt.Println("sum", sum, sum)
	// fmt.Println("count", count)

	result := false
	if count <= 2 {
		result = false
	}

	if sum%10 == 0 {
		result = true
	}

	fmt.Println("=", result)
	return result
}
