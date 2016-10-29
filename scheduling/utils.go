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

//returns true when it finds a duplicate combo
func FindDuplicateCombos(combos []Combo) (bool, int, int) {
	for i := 0; i < len(combos); i++ {
		for j := i + 1; j < len(combos); j++ {
			if CompareCombos(combos[i], combos[j]) {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

//returns true if equal
func CompareCombos(input1, input2 Combo) bool {
	if len(input1.Classes) != len(input2.Classes) {
		return false
	}
	for i := range input1.Classes {
		if input1.Classes[i] != input2.Classes[i] {
			return false
		}

	}
	return true
}
