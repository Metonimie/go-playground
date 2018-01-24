// Package raindrop converts numbers into the raindrop language.
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

func raindropLanguage(factors []int) string {
	var buffer bytes.Buffer

	for _, v := range factors {
		switch v {
		case 3:
			buffer.WriteString("Pling")
		case 5:
			buffer.WriteString("Plang")
		case 7:
			buffer.WriteString("Plong")
		}
	}

	return buffer.String()
}

// Convert converts a number into the Raindrop Language
func Convert(n int) string {
	factors := make([]int, 0, 20)

	// Small function to check if factors contains the num.
	contains := func(num int) bool {
		for _, v := range factors {
			if v == num {
				return true
			}
		}
		return false
	}

	// Using trial division algorithm.
	number := n
	factor := 2
	for number > 1 {
		if number%factor == 0 {
			if contains(factor) != true {
				factors = append(factors, factor)
			}
			number = number / factor
		} else {
			factor = factor + 1
		}
	}

	raindropLangString := raindropLanguage(factors)
	if len(raindropLangString) == 0 {
		raindropLangString = strconv.Itoa(n)
	}
	return raindropLangString
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
