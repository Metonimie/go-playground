// Package raindrops converts numbers into the raindrop language.
package raindrops

import (
	"strconv"
	"bytes"
)

const testVersion = 3

// Convert converts a number into the Raindrop Language
func Convert(n int) string {
	var retval bytes.Buffer

	if n%3 == 0 {
		retval.WriteString("Pling")
	}
	if n%5 == 0 {
		retval.WriteString("Plang")
	}
	if n%7 == 0 {
		retval.WriteString("Plong")
	}

	if len(retval.String()) == 0 {
		retval.WriteString(strconv.Itoa(n))
	}
	return retval.String()
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
