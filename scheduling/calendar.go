package scheduling

import (
// "time"
)

type Schedule struct {
	Classes []Class
	Score   int
}

type BySchedule []Schedule

func (s BySchedule) Len() int           { return len(s) }
func (s BySchedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySchedule) Less(i, j int) bool { return s[i].Score < s[j].Score }

// Returns whether this addition will cause a conflict of events and returns the
// conflicting event.
// func (cal *Schedule) DoesConflict(Class Class) Class {
// 	// start := Class.StartTime
// 	// end := Class.EndTime

// 	// // Look For Conflicts
// 	// // TODO: (Complexity: O(n)) Could be simplified if we didn't have to return
// 	// // the conflict. Or Add complexity would be worse.
// 	// // TODO: Handle Multiple Conflicts
// 	// for _, possibleConflict := range cal.Events {
// 	// 	conflictStart := possibleConflict.StartTime()
// 	// 	conflictEnd := possibleConflict.EndTime()

// 	// 	if (start.After(conflictStart) && start.Before(conflictEnd)) ||
// 	// 		(end.After(conflictStart) && end.Before(conflictEnd)) {

// 	// 		return possibleConflict
// 	// 	}
// 	// }

// 	return Class{}
// }
