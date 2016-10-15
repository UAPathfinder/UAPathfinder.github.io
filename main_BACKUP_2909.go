package main

import (
<<<<<<< HEAD
//"bufio"
//"fmt"
//"sort"
//"time"
//"strings"
//"os"
//"strconv"
//mysql --local-infile -uroot -pyourpwd yourdbname/"database/sql" //non-functional MySQL imports
// _ "/mysql"
=======
	"fmt"
	"sort"
	"time"
>>>>>>> 0850cf763d6222cd09bd7d75284f0a871c408757
)

var AllCourses []Course

func main() {
<<<<<<< HEAD
	/*fmt.Print("Welcome to CourseCorrect!  more information goes here! \n ")
	AllCourses = CourseHolder()
	AllCriteria := CriteriaHolder()
=======
	fmt.Println("Welcome to CourseCorrect! More information goes here!")
	AllCourses = CourseHolder()
	AllCriteria := CriteriaHolder()
}

// Total possiable combos, includes ones with class confilicts
func NumCombos(courses []Course) int {
	total := 1
	for i := range courses {
		total *= len(courses[i].Classes)
	}
	return total
}

func ScoreCombo(combo Combo, criteria Criteria) int { //assumes OrderClasses has been called on the combo
	return ScoreBreaks(combo, criteria) + ScoreProfs(combo, criteria) + ScoreEarliestClass(combo, criteria) + ScoreLatestClass(combo, criteria) + ScoreDays(combo, criteria)
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

func ScoreDays(combo Combo, criteria Criteria) int { //could use some refactoring
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

func GenerateCombos(courses []Course, result *[]Combo, depth int, current Combo) { //eventuially this should be modified to not add a class if that class isn't manditory and then also start the score of that combo slightly lowered since it dropped a class
	if depth == len(courses) {
		hasOverlap, issue1, issue2 := DoesHaveOverlap(current)
		if !hasOverlap {
			*result = append(*result, current)
		} else {
			course1 := GetCourse(current.Classes[issue1].CourseId)
			course2 := GetCourse(current.Classes[issue2].CourseId)
			if !course1.Manditory && !course2.Manditory {
				if course1.Priority < course2.Priority {
					current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
					current.Score -= course1.Priority
					GenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				} else if course1.Priority > course2.Priority {
					current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
					current.Score -= course2.Priority
					GenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				} //otherwise don't append
			} else if course1.Manditory {
				current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
				current.Score -= course2.Priority
				GenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
			} else if course2.Manditory {
				current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
				current.Score -= course2.Priority
				GenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
			} //otherwise don't append
		}

	} else {
		currentCourse := courses[depth]
		for i := 0; i < len(currentCourse.Classes); i++ {
			var tempCurrent Combo
			tempCurrent.Classes = append(current.Classes, currentCourse.Classes[i])
			GenerateCombos(courses, result, depth+1, tempCurrent)
			//yes, it's recursive.  it only goes [depth] layers deep before returning so the stack shouldn't overflow for a reasonable number of courses
		}
		if len(currentCourse.OrCourses) != 0 {
			for p := 0; p < len(currentCourse.OrCourses); p++ {
				for i := 0; i < len(currentCourse.OrCourses[p].Classes); i++ {
					var tempCurrent Combo
					tempCurrent.Classes = append(current.Classes, currentCourse.OrCourses[p].Classes[i])
					GenerateCombos(courses, result, depth+1, tempCurrent)
				}
>>>>>>> 0850cf763d6222cd09bd7d75284f0a871c408757

	for i := range AllCourses {
		PrintCourse(AllCourses[i])
	}
	_ = AllCriteria*/

	//RunFullTest()
}
<<<<<<< HEAD
=======

func MinuteDiff(first, second time.Time) int {
	diff := first.Sub(second)
	var minutes = diff.Minutes()
	return int(minutes)
}

func GetCourses() map[int]int { //has to wait for sql stuff
	//retrives list of course id's with priorities
	output := make(map[int]int)
	output[1] = 2 //tempoary output for testing
	return output
}

func DoesOverlap(Class1Start, Class1End, Class2Start, Class2End time.Time) bool {
	return !(Class2Start.After(Class1End) || Class1Start.After(Class2End))
}

func SimpleParse(input string) time.Time {
	output, _ := time.Parse("15:04:05", input)
	return output
}
>>>>>>> 0850cf763d6222cd09bd7d75284f0a871c408757
