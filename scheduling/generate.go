package scheduling

import (
	"sort"
)

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

	sort.Sort(ByEndTime(Classes))

	for _, ThisClass := range Classes {
		//THIS IS IMPORTANT
		//only make changes to Current if you want them applied to every tree branch after this point
		workingCurrent := Current

		workingCurrent.Classes = append(workingCurrent.Classes, ThisClass)

		// Container to hold potential deletions.
		//this is something that martin did, idk why, it seems to work
		// var pendingDeletions []Class

		// DoesConfict, conflictingClass := workingCurrent.DoesConflict(ThisClass)9

		// if DoesConfict {
		// 	pendingDeletions = append(pendingDeletions, conflictingClass)
		// }

		// for _, Class := range Classes {
		// 	if !Class.ExistsIn(pendingDeletions) {
		// 		workingCurrent.Classes = append(workingCurrent.Classes, Class)
		// 	}
		// }

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
