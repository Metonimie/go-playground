package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	h int
	m int
}

// normalizeHours Normalizes the hours of a clock.
func normalizeHours(hours int) int {
	hours = hours % 24
	if hours < 0 {
		hours = 24 + hours
	}
	return hours
}

// calculateMinutes returns the amount of minutes and hours from minutes
func calculateMinutes(minutes int) (int, int) {
	hours := (minutes / 60) % 24
	minutes = minutes % 60

	if minutes < 0 {
		minutes = 60 + minutes
		hours -= 1
	}

	return hours, minutes
}

// New creates a new Clock
func New(hour, minute int) Clock {

	hours, totalMinutes := calculateMinutes(minute) // Retrieve the amount of minutes and hours
	totalHours := normalizeHours(hour + hours)      // Add the extra hours to the existing hours and normalize

	return Clock{totalHours, totalMinutes}
}

// Returns the string representation of the clock, in the form hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Adds minutes to a clock.
func (c Clock) Add(minutes int) Clock {
	newClock := New(c.h, c.m+minutes)
	return newClock
}
