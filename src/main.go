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

func ScoreCombo(combo Combo, criteria Criteria) int{
	return ScoreBreaks(combo, criteria) + ScoreProfs(combo, criteria) + ScoreEarliestClass(combo, criteria) + ScoreLatestClass(combo, criteria) + ScoreDays(combo, criteria)
}

func ScoreBreaks(combo Combo, criteria Criteria) int{
	return 0
}

func ScoreProfs(combo Combo, criteria Criteria) int{
//will have to do some sort of query to rateMyProfessor
	return 0 //temp
}

func ScoreEarliestClass(combo Combo, criteria Criteria) int{
	return 0
}

func ScoreLatestClass(combo Combo, criteria Criteria) int{
	return 0
}

func ScoreDays(combo Combo, criteria Criteria) int{
	return 0
}

func GenerateCombos(courses []Course, result *[]Combo, depth int, current Combo) {
//There is almost certiantly a better way to do this
	if depth == len(courses) {
		if !DoesHaveOverlap(current){
			*result = append(*result, current)
		}

	}else{
		for i := 0; i < len(courses[depth].Classes); i++ {
			var tempCurrent Combo
			tempCurrent.Classes = append(current.Classes, courses[depth].Classes[i])
			GenerateCombos(courses, result, depth + 1, tempCurrent)
		}
	}
}

func DoesHaveOverlap(combo Combo) bool{
	for i := 0; i < len(combo.Classes); i++ {
		for j := i + 1; j < len(combo.Classes); j++{
			if DoesOverlap(combo.Classes[i].StartTime, combo.Classes[i].EndTime, combo.Classes[j].StartTime, combo.Classes[j].EndTime) {
				return true
                        }
                }
        }
	return false

}

func OrderCombos(combos []Combo) {
	sort.Reverse(ByScore(combos))
}

func FillCourses(courses map[int]int) { //has to wait for sql stuff
	for i := range courses {
		courses[i]++
	}
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
