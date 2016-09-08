package main

import (
//"bufio"
"fmt"
//"os"
//"database/sql" //non-functional MySQL imports
// _ "/mysql"
)

type Class struct { //A single class
	//ex Data Structures starting at 3 PM in room 301 with professor x
	ClassId         int
	CourseId        int
	StartTime       string //ISO 8601 (hh:mm:ss) e.x. 12:15:00
	EndTime         string //see above
	MeetingDays     string //e.x. MWF, TH (Tuesday tHursday)
	ProfessorName   string
	MeetingLocation string
}

type Course struct { //ex 3960:401 Data Structures
	CourseId  int
	Priority  int //from 1 to 10
	Manditory bool
	Classes []Class//initializing an empty slice
}

func main() {
	
}

func FillCourses(courses map[int]int){
	for i := range courses {
		fmt.Println(courses[i])
	}
}

/* So I don't forget the syntax
func courseThing(c Course){
	fmt.Println(c.CourseId)
	c.Classes = append(c.Classes,  Class{ClassId: 4444444})
	fmt.Println(c.Classes[0].ClassI
}
*/

func MySQLQuery(input string) string {
	//this will return the MySQL query string, right now it doesn't
	//con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
	//defer con.Close()
	return "not initialized"
}

func GetCourses() map[int]int {
	//retrives list of course id's with priorities
	output := make(map[int]int)
	output[1] = 2 //tempoary output for testing
	return output
}

//func ComputeCombinations //multiplies all course.Classes.length by each other
