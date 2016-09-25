package main

import (
	//"bufio"
	//"fmt"
	"sort"
	"time"
	//"os"
	//mysql --local-infile -uroot -pyourpwd yourdbname/"database/sql" //non-functional MySQL imports
	// _ "/mysql"
)

func main() {
}

func NumCombos(courses []Course) int { //total possiable combos, includes ones with class confilicts
	var total int = 1
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
		if !DoesHaveOverlap(current) {
			*result = append(*result, current)
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
                        for p :=0; p < len(currentCourse.OrCourses); p++ {
                                 for i := 0; i < len(currentCourse.OrCourses[p].Classes); i++ {
                                        var tempCurrent Combo
                                        tempCurrent.Classes = append(current.Classes, currentCourse.OrCourses[p].Classes[i])
                                        GenerateCombos(courses, result, depth+1, tempCurrent)
                                }

                        }
                }
	}
}

func DoesHaveOverlap(combo Combo) bool {
	for i := 0; i < len(combo.Classes); i++ {
		for j := i + 1; j < len(combo.Classes); j++ {
			if DoesOverlap(combo.Classes[i].StartTime, combo.Classes[i].EndTime, combo.Classes[j].StartTime, combo.Classes[j].EndTime) {
				return true
			}
		}
	}
	return false

}

func OrderCombos(combos *[]Combo) {
	sort.Reverse(ByScore(*combos))
}

func OrderClasses(combo *Combo) {
	sort.Sort(ByStartTime(combo.Classes))
}

func FillCourses(courses map[int]int) { //has to wait for sql stuff
	for i := range courses {
		courses[i]++
	}

}

func MinuteDiff(first, second time.Time) int {
	diff := first.Sub(second)
	var minutes = diff.Minutes()
	return int(minutes)
}
func MySQLQuery(input string) string {
	//this will return the MySQL query string, right now it doesn't
	//con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
	//defer con.Close()
	return "not initialized"
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
