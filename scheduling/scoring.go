package scheduling

import (
	"sort"
	"time"
)

func FullTest() {
	// result := GenerateCombos(MockCourses)
	// for i := range result {
	// 	PrintCombo(result[i])
	// }
}

// Total possible combos, includes ones with class confilicts
func NumCombos(courses []Course) int {
	total := 1
	for i := range courses {
		total *= len(courses[i].Classes)
	}
	return total
}

// Assumes OrderClasses has been called on the combo
func ScoreCombo(combo Combo, criteria Criteria) int {
	return ScoreBreaks(combo, criteria) +
		ScoreProfs(combo, criteria) +
		ScoreEarliestClass(combo, criteria) +
		ScoreLatestClass(combo, criteria) +
		ScoreDays(combo, criteria)
}

//note: I don't really know how to do these, so some of them may have some inherent bias
func ScoreBreaks(combo Combo, criteria Criteria) int {
	output := 0
	minutes := 0
	for i := 0; i < len(combo.Classes)-1; i++ {
		minutes = MinuteDiff(combo.Classes[i].EndTime, combo.Classes[i+1].StartTime)
		if minutes > 15 { //this promotes a 15 minute gap between every class
			//maybe later make it a user-configurable option
			if criteria.Breaks.Maximize {
				output += (minutes - 15)
			} else {
				output -= (minutes - 15)
			}
		} else {
			output -= (15 - minutes)
		}
	}
	return output * criteria.Breaks.Weight
}

func ScoreProfs(combo Combo, criteria Criteria) int {
	//will have to do some sort of query to rateMyProfessor
	return 0 //temp
}

func ScoreEarliestClass(combo Combo, criteria Criteria) int {
	output := 0
	earliestStart := combo.Classes[0].StartTime
	earliestWanted := criteria.EarliestClass.Time
	diff := earliestStart.Sub(earliestWanted)
	var minutes = diff.Minutes()
	if earliestStart.After(earliestWanted) {
		output += int(minutes)
	} else {
		if criteria.EarliestClass.Manditory {
			return -100000
		} else {
			output -= int(minutes)
		}
	}
	output *= criteria.EarliestClass.Weight
	return output
}

func ScoreLatestClass(combo Combo, criteria Criteria) int {
	output := 0
	latestEnd := combo.Classes[len(combo.Classes)-1].EndTime
	latestWanted := criteria.LatestClass.Time
	diff := latestEnd.Sub(latestWanted)
	var minutes = diff.Minutes()
	if latestEnd.Before(latestWanted) {
		output += int(minutes)
	} else {
		if criteria.EarliestClass.Manditory {
			return -100000
		} else {
			output -= int(minutes)
		}
	}

	return output * criteria.EarliestClass.Weight
}

// TODO: Could use some refactoring
func ScoreDays(combo Combo, criteria Criteria) int {
	output := 100
	for class := range combo.Classes {
		for ClassDay := range combo.Classes[class].MeetingDays {
			for CriteriaDay := range criteria.Days.Other {
				if combo.Classes[class].MeetingDays[ClassDay] == criteria.Days.Other[CriteriaDay] {
					if criteria.Days.Manditory {
						output = -10000000
					} else {
						output -= 10
					}
				}
			}
		}
	}
	output *= criteria.Days.Weight
	return output
}

func GenerateCombos(courses []Course) []Combo {
	var current Combo
	result := make([]Combo, 0)
	RecursiveGenerateCombos(courses, &result, 0, current)
	return result
}

//the insperation for this algorithm comes from here: 
//http://stackoverflow.com/questions/17192796/generate-all-combinations-from-multiple-lists
//we need recursion because the number of courses is going to vary between users
func RecursiveGenerateCombos(courses []Course, result *[]Combo, depth int, current Combo) {
	if depth == len(courses) {
		hasOverlap, issue1, issue2 := DoesHaveOverlap(current)
		if !hasOverlap {
			//horray, no overlaps!  apend the current combo to the result
			*result = append(*result, current)
		} else {
			//get the course objects so we can get the manditory and priority members
			//later this will be a database call
			course1 := GetCourse(courses, current.Classes[issue1].CourseId)
			course2 := GetCourse(courses, current.Classes[issue2].CourseId)
			//if neither conflicting courses are manditory
			if !course1.Manditory && !course2.Manditory {
				//we drop the one that has a lower priority
				if course1.Priority < course2.Priority {
					current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
					current.Score -= course1.Priority
					RecursiveGenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				} else {
					current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
					current.Score -= course2.Priority
					RecursiveGenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				} 
			//the cases that one is manditory but the other isn't
			} else if course1.Manditory {
				current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
				current.Score -= course2.Priority
				//kicks it back with the same depth to check for overlaps again
				RecursiveGenerateCombos(courses, result, depth, current)
			} else if course2.Manditory {
				current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
				current.Score -= course2.Priority
				//kicks it back with the same depth to check for overlaps again
				RecursiveGenerateCombos(courses, result, depth, current) 
			} //otherwise don't append because the whole combo does not fit the user's requirements
		}

	} else {
		//this loop handles the 'root course'
		currentCourse := courses[depth]
		for i := 0; i < len(currentCourse.Classes); i++ {
			var tempCurrent Combo
			//append the first class in the current course to the current combo, then fires GenerateCombos again
			tempCurrent.Classes = append(current.Classes, currentCourse.Classes[i])
			RecursiveGenerateCombos(courses, result, depth+1, tempCurrent)
			//yes, it's recursive.  it only goes [depth] layers deep before returning so the stack shouldn't overflow for a reasonable number of courses
		}
		//this loop goes through all the 'ored' courses within the 'root' course
		if len(currentCourse.OrCourses) != 0 {
			for p := 0; p < len(currentCourse.OrCourses); p++ {
				for i := 0; i < len(currentCourse.OrCourses[p].Classes); i++ {
					//the exact same loop as above but for the 'ored' course
					var tempCurrent Combo
					tempCurrent.Classes = append(current.Classes, currentCourse.OrCourses[p].Classes[i])
					RecursiveGenerateCombos(courses, result, depth+1, tempCurrent)
				}

			}
		}
	}
}

func DoesHaveOverlap(combo Combo) (bool, int, int) {
	for i := 0; i < len(combo.Classes); i++ {
		for j := i + 1; j < len(combo.Classes); j++ {
			if DoesOverlap(combo.Classes[i], combo.Classes[j]) {
				return true, i, j
			}
		}
	}
	return false, -1, -1

}

func GetCourse(courses []Course, id int) Course {
	for _, course := range courses {
		if course.CourseId == id {
			return course
		}
	}

	output := Course{}
	return output
}

func OrderCombos(combos *[]Combo) {
	sort.Reverse(ByScore(*combos))
}

func OrderClasses(combo *Combo) {
	sort.Sort(ByStartTime(combo.Classes))
}

func MinuteDiff(first, second time.Time) int {
	diff := first.Sub(second)
	var minutes = diff.Minutes()
	return int(minutes)
}

func DoesOverlap(Class1, Class2 Class) bool {
	counter := 0
	for i := range Class1.MeetingDays {
		for j := range Class2.MeetingDays {
			if Class1.MeetingDays[i] == Class2.MeetingDays[j] {
				counter += 1
			}
		}
	}
	if counter <= 0 {
		return false
	}
	return !(Class2.StartTime.After(Class1.EndTime) || Class1.StartTime.After(Class2.EndTime))
}
