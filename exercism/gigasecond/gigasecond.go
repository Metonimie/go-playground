// Package gigasecond implements utility functions related to time
package gigasecond

// import path for the time package from the standard library
import (
	"time"
	"math"
	"fmt"
)

const testVersion = 4

// AddGigasecond adds a gigasecond to the given time and returns the total time
func AddGigasecond(ct time.Time) time.Time {
	var gigasecond int64 = int64(math.Pow(10, 9))

	// Ignore errors
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", gigasecond))

	return ct.Add(duration)
}
