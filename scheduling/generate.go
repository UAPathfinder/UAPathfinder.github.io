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

	// 	if Depth == len(Courses) {
	// 		// Deepest Case, Called After Builder Cases for 0..(Depth - 1)
	// 		*Result = append(*Result, Current)
	// 		return
	// 	}

	// 	// Builder Case
	// 	Course := Request.Courses[Depth]
	// 	Course.Optional = !Course.ma
	// 	// Get Classes for Course
	// 	Classes := Course.Classes

	// ClassesLoop:
	// 	for _, class := range Classes {
	// 		//THIS IS IMPORTANT
	// 		//only make changes to Current if you want them applied to every tree branch after this point
	// 		workingCurrent := Current

	// 		workingCurrent.Classes = append(workingCurrent.Classes, class)

	// 		// Container to hold potential deletions.
	// 		//this is something that martin did, idk why, it seems to work
	// 		var pendingDeletions []Class

	// 		cost := 0

	// 		for _, Class := range workingCurrent.Classes {

	// 			conflictingClass := workingCurrent.Calendar.DoesConflict(Class)

	// 			if conflictingClass != nil {

	// 				if !conflictingProps.Optional && !Course.Optional {
	// 					// Both Required
	// 					continue ClassesLoop
	// 				} else if conflictingProps.Optional && !Course.Optional {
	// 					// Ours Required
	// 					// Pend Deletion of Other, Keep Ours
	// 					pendingDeletions = append(pendingDeletions, conflictingClass)
	// 					cost += conflictingProps.Weight
	// 				} else if !conflictingProps.Optional && Course.Optional {
	// 					// Other Required, Ours Optional
	// 					cost += Course.Weight
	// 				} else {
	// 					// Both Optional, Skip Lower Priority
	// 					if conflictingProps.Weight < Course.Weight {
	// 						cost += conflictingProps.Weight
	// 						pendingDeletions = append(pendingDeletions, conflictingClass)
	// 					} else {
	// 						cost += Course.Weight
	// 					}
	// 				}
	// 			}
	// 		}

	// 		// Incur Cost
	// 		workingCurrent.Score -= cost

	// 		for _, deletion := range pendingDeletions {
	// 			switch deletion := deletion.(type) {
	// 			case ClassClass:
	// 				// Remove All Elements of This Class
	// 				j := 0
	// 				for _, Class := range workingCurrent.Calendar.Classs {
	// 					Class, ok := Class.(ClassClass)
	// 					if !ok {
	// 						continue
	// 					}

	// 					if Class.Class.Identifier != deletion.Class.Identifier {
	// 						workingCurrent.Calendar.Classs[j] = Class
	// 						j++
	// 					}
	// 				}
	// 				workingCurrent.Calendar.Classs = workingCurrent.Calendar.Classs[:j]
	// 			}
	// 		}

	// 		for _, event := range events {
	// 			workingCurrent.Calendar.Add(event)
	// 		}

	// 		RecursiveFindSchedules(Courses, props, accessor, Result, Depth+1, workingCurrent)

	// 		//     Add Class to Schedule
	// 		//     Calculate Score of Schedule
	// 		//     Descide Whether Sub-Tree Is Viable
	// 		//     Continue Building By Recursion (Depth + 1)

	// 		// If Hard Constraint Violated, Discard All Sub-Schedules
	// 		//     Return
	// 		// Else, Continue Recursing:
	// 		//     RecursiveFindSchedules(Courses, Result, Depth + 1, Current)
	// 	}
}
