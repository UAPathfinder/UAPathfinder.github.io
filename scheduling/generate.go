package scheduling

import (
	"fmt"
)

// An interface to get classes for FindSchedules. Used for mocking.
type Accessor interface {
	// Given a course identifier, returns a list of classes. Called once for
	// each course for each execution of FindSchedules.
	GetClasses(courseIdentifier string) []Class

	GetCourse(courseIdentifier string) Course
}

type Schedule struct {
	Calendar
	Score int
}

type BySchedule []Schedule

func (s BySchedule) Len() int           { return len(s) }
func (s BySchedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySchedule) Less(i, j int) bool { return s[i].Score < s[j].Score }

// Given a list of Courses, returns the best combination of classes.
// Requiments:
// - Many Soft Constraints:
//   - Slight Overlaps Not Rejected
// - Search Space (Akron: 88*55*45*45*41*33*27*26 @ 1 op/ns = 2.6 hours)
// Solutions:
// - Genetic Algorithm
// - Simulated Annealing
func FindSchedules(courses []string, props map[string]EventProperties, accessor Accessor) []Schedule {
	fmt.Println("Starting algo")
	result := []Schedule{}
	RecursiveFindSchedules(courses, props, accessor, &result, 0, nil)
	fmt.Println("Ending algo")
	return result
}

// http://stackoverflow.com/questions/17192796/generate-all-combinations-from-multiple-lists
func RecursiveFindSchedules(courses []string, props map[string]EventProperties, accessor Accessor, result *[]Schedule, depth int, current *Schedule) {
	if depth == len(courses) {
		// Deepest Case, Called After Builder Cases for 0..(depth - 1)
		*result = append(*result, *current)
		return
	}

	if depth == 0 {
		// First Builder Case
		current = &Schedule{
			Calendar: Calendar{
				Events: make([]Event, 0),
			},
		}
	}

	// Builder Case
	course := courses[depth]

	// Get Classes for Course
	classes := accessor.GetClasses(course)

	// Get parameters for the course. If does not exist, returns zero value of
	// EventProperties.
	courseProps := props[course]

classesLoop:
	for _, class := range classes {

		// Try Adding Classes From course to Schedule
		events := class.Events(&courseProps)

		// Container to hold potential deletions.
		var pendingDeletions []Event

		cost := 0
		shouldntAdd := false

		for _, event := range events {
			conflictingEvent := current.Calendar.DoesConflict(event)

			if conflictingEvent != nil {
				// Found A Conflict. Fail If Possible.
				conflictingProps := conflictingEvent.Properties()

				if !conflictingProps.Optional && !courseProps.Optional {
					// Both Required
					continue classesLoop
				} else if conflictingProps.Optional && !courseProps.Optional {
					// Ours Required
					// Pend Deletion of Other, Keep Ours.
					cost += conflictingProps.Weight
				} else if !conflictingProps.Optional && courseProps.Optional {
					// Other Required, Ours Optional
					// Leave Other in, Don't Add Ours
					cost += courseProps.Weight
					shouldntAdd = true
				} else {
					// Both Optional, Skip Lower Priority
					if conflictingProps.Weight < courseProps.Weight {
						cost += conflictingProps.Weight
						pendingDeletions = append(pendingDeletions, conflictingEvent)
					} else {
						cost += courseProps.Weight
						shouldntAdd = true
					}
				}
			}
		}

		// Incur Cost
		current.Score -= cost

		for _, deletion := range pendingDeletions {
			switch deletion := deletion.(type) {
			case ClassEvent:
				// Remove All Elements of This Class
				j := 0

				// TODO: O(deletions * events)
				for _, event := range current.Calendar.Events {
					event, ok := event.(ClassEvent)
					if !ok {
						// Not a class event. Not a conflicting class event,
						// leave it be.
						// TODO: Copy this through.
					}

					if event.Class.Identifier != deletion.Class.Identifier {
						current.Calendar.Events[j] = event
						j++
					}
				}
				current.Calendar.Events = current.Calendar.Events[:j]

				// TODO: Handle normal event deletion. Isn't an issue until we
				// add events which aren't classes to schedules.
			}
		}

		if !shouldntAdd {
			for _, event := range events {
				current.Calendar.Add(event)
			}
		}

		// Send Copy of Current Down
		RecursiveFindSchedules(courses, props, accessor, result, depth+1, &(*current))

		//     Add Class to Schedule
		//     Calculate Score of Schedule
		//     Descide Whether Sub-Tree Is Viable
		//     Continue Building By Recursion (depth + 1)

		// If Hard Constraint Violated, Discard All Sub-Schedules
		//     Return
		// Else, Continue Recursing:
		//     RecursiveFindSchedules(courses, result, depth + 1, current)
	}
}
