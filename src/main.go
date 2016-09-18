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

type Class struct { //A single class
	//ex Data Structures starting at 3 PM in room 301 with professor x
	ClassId         int
	CourseId        int
	StartTime       time.Time
	EndTime         time.Time
	MeetingDays     string //e.x. MWF, TH (Tuesday tHursday)
	ProfessorName   string
	MeetingLocation string
}

type Course struct { //ex 3960:401 Data Structures
	CourseId  int
	Priority  int //from 1 to 10
	Manditory bool
	Classes   []Class //initializing an empty slice
}

type Combo struct {
	Classes []Class
	Score   int
}

type Criteria struct {
}

type ByScore []Combo                 //implements sort.Interface for []Combo based on Score
func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func main() {
}


func NumCombos(courses []Course) int {
	var total int = 1
	for i := range courses {
		total *= len(courses[i].Classes)
	}
	return total
}

func ScoreCombo(combo Combo) {

}

func GenerateCombos(courses []Course, result *[]Combo, depth int, current Combo) {
//There is almost certiantly a better way to do this
	_ = "breakpoint"
	if depth == len(courses) {
		*result = append(*result, current)
		_ = "breakpoint"
	}else{
		for i := 0; i < len(courses[depth].Classes); i++ {
			var tempCurrent Combo
			tempCurrent.Classes = append(current.Classes, courses[depth].Classes[i])
			GenerateCombos(courses, result, depth + 1, tempCurrent)
		}
	}
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
