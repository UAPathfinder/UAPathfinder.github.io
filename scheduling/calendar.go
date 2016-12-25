package scheduling

import (
	"time"
)

type Event interface {
	StartTime() time.Time
	EndTime() time.Time
	Properties() EventProperties
}

type EventProperties struct {
	Weight   int
	Optional bool
}

// Holds information about which times are free and occupied. Only considers
// days and times. It is not aware about the difference between two days on two
// different weeks.
type Calendar struct {
	Events []Event
}

func (cal *Calendar) Add(evt ...Event) {
	cal.Events = append(cal.Events, evt...)
}

// Returns whether this addition will cause a conflict of events and returns the
// conflicting event.
func (cal *Calendar) DoesConflict(evt Event) Event {
	start := evt.StartTime()
	end := evt.EndTime()

	// Look For Conflicts
	// TODO: (Complexity: O(n)) Could be simplified if we didn't have to return
	// the conflict. Or Add complexity would be worse.
	// TODO: Handle Multiple Conflicts
	for _, possibleConflict := range cal.Events {
		conflictStart := possibleConflict.StartTime()
		conflictEnd := possibleConflict.EndTime()

		if (start.After(conflictStart) && start.Before(conflictEnd)) ||
			(end.After(conflictStart) && end.Before(conflictEnd)) {

			return possibleConflict
		}
	}

	return nil
}
