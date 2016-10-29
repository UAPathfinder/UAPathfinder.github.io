package scheduling

import (
	"time"
)

// Utility function to parse time in format (hh:mm:ss). Panics on failure.
func MustParseTime(input string) time.Time {
	output, err := time.Parse("15:04:05", input)
	if err != nil {
		panic(err)
	}

	return output
}

func CompareCombos(input1, input2 Combo) bool {
	for i := range input1.Classes {
		if input1.Classes[i] != input2.Classes[i] {
			return false
		}
	}
	return true
}
