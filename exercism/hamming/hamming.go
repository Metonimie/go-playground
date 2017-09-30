//Package hamming implements methods related to DNA
package hamming

import "fmt"

const testVersion = 6

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return  -1, fmt.Errorf("The length of the DNA strands is not the same! a: %d b: %d", len(a), len(b))
	}

	var distance int = 0

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}

	return distance, nil
}
