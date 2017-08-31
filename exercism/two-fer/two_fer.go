// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns the string "One for you, one for me." if name is empty, otherwise
// it returns "One for {name}, one for me."
func ShareWith(name string) string {
	if len(name) == 0 {
		return "One for you, one for me."
	} else {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
}
