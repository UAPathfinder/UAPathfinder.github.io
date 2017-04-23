package scheduling

// An interface to get Classes for FindSchedules. Mockable for testing the
// algorithm against custom data.

// Given a list of Courses, returns the best combination of Classes.
// Requiments:
// - Many Soft Constraints:
//   - Slight Overlaps Not Rejected
// - Search Space (Akron: 88*55*45*45*41*33*27*26 @ 1 op/ns = 2.6 hours)
// Solutions:
// - Genetic Algorithm
// - Simulated Annealing
func FindSchedules(Request ScheduleRequest) []Schedule {
	Result := []Schedule{}
	RecursiveFindSchedules(Request, &Result, 0, Schedule{})
	return Result
}

// http://stackoverflow.com/questions/17192796/generate-all-combinations-from-multiple-lists
func RecursiveFindSchedules(Request ScheduleRequest, Result *[]Schedule, Depth int, Current Schedule) {

	if Depth == len(Request.Courses) {
		// Deepest Case, Called After Builder Cases for 0..(Depth - 1)
		*Result = append(*Result, Current)
		return
	}

	// Builder Case
	Course := Request.Courses[Depth]

	// Get Classes for Course
	Classes := Course.Classes

ClassesLoop:
	for _, ThisClass := range Classes {
		ThisClass.Optional = !ThisClass.Manditory
		//THIS IS IMPORTANT
		//only make changes to Current if you want them applied to every tree branch after this point
		workingCurrent := Current

		workingCurrent.Classes = append(workingCurrent.Classes, ThisClass)

		// Container to hold potential deletions.
		//this is something that martin did, idk why, it seems to work
		var pendingDeletions []Class

		cost := 0

		DoesConfict, conflictingClass := workingCurrent.DoesConflict(ThisClass)

		if DoesConfict {

			if !conflictingClass.Optional && !ThisClass.Optional {
				// Both Required
				continue ClassesLoop
			} else if conflictingClass.Optional && !ThisClass.Optional {
				// Ours Required
				// Pend Deletion of Other, Keep Ours
				pendingDeletions = append(pendingDeletions, conflictingClass)
				cost += conflictingClass.Priority
			} else if !conflictingClass.Optional && ThisClass.Optional {
				// Other Required, Ours Optional
				cost += ThisClass.Priority
			} else {
				// Both Optional, Skip Lower Priority
				if conflictingClass.Priority < ThisClass.Priority {
					cost += conflictingClass.Priority
					pendingDeletions = append(pendingDeletions, conflictingClass)
				} else {
					cost += ThisClass.Priority
				}
			}
		}

		// Incur Cost
		workingCurrent.Score -= cost

		//TODO: make this not suck
		for _, Class := range Classes {
			if !Class.ExistsIn(pendingDeletions) {
				workingCurrent.Classes = append(workingCurrent.Classes, Class)
			}
		}

		RecursiveFindSchedules(Request, Result, Depth+1, workingCurrent)

		//     Add Class to Schedule
		//     Calculate Score of Schedule
		//     Descide Whether Sub-Tree Is Viable
		//     Continue Building By Recursion (Depth + 1)

		// If Hard Constraint Violated, Discard All Sub-Schedules
		//     Return
		// Else, Continue Recursing:
		//     RecursiveFindSchedules(Courses, Result, Depth + 1, Current)
	}
}
