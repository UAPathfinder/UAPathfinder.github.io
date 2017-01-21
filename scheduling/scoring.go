package scheduling

// func ScoreSchedule(sched Schedule, criteria Criteria) int {
// 	return ScoreBreaks(combo, criteria)
// }
//
// // Note: I don't really know how to do these, so some of them may have some inherent bias
// func ScoreBreaks(combo Combo, criteria Criteria) int {
// 	output := 0
//
// 	for i := 0; i < len(combo.Classes)-1; i++ {
// 		minutes := MinuteDiff(combo.Classes[i].EndTime, combo.Classes[i+1].StartTime)
// 		if minutes > 15 {
// 			// this promotes a 15 minute gap between every class
// 			// maybe later make it a user-configurable option
// 			if criteria.Breaks.Maximize {
// 				output += (minutes - 15)
// 			} else {
// 				output -= (minutes - 15)
// 			}
// 		} else {
// 			output -= (15 - minutes)
// 		}
// 	}
// 	return output * criteria.Breaks.Weight
// }
