package scheduling

import (
	"time"
)

// TODO: Write a time library fore our custom time stuff.

// Utility function to parse time in format (hh:mm:ss). Panics on failure.
func MustParseTime(input string) time.Time {
	output, err := time.Parse("15:04", input)
	if err != nil {
		panic(err)
	}

	return output
}

func MinuteDiff(first, second time.Time) int {
	diff := first.Sub(second)
	var minutes = diff.Minutes()
	return int(minutes)
}
