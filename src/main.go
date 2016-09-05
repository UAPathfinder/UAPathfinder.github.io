package main
import (
 //"bufio"
 //"fmt"
 //"os"
 //"database/sql" //non-functional MySQL imports
 // _ "/mysql"
)
type Class struct { //A single class
//ex Data Structures starting at 3 PM in room 301 with professor x
	ClassId int
	CourseId int
	StartTime string //ISO 8601 (hh:mm:ss) e.x. 12:15:00
	EndTime string //see above
	MeetingDays string //e.x. MWF, TH (Tuesday tHursday)
	ProfessorName string
	MeetingLocation string
}

type Course struct { //ex 3960:401 Data Structures
	CourseId int
	Priority int //from 1 to 10
	Manditory bool = false
	//List<Class> Classes, but that's not a thing in Go.  
		//Maybe a slice of an arbatrarially large array?
		//or implement c++ vector
}

var Courses map[int]i

func main(){
	_courses = GetCourses();//or replace with whatever list thing we figure out
}

func MySQLQuery (input string) string{
        //this will return the MySQL query string, right now it doesn't
        //con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
        //defer con.Close()
	return "not initialized"
}

func GetCourses()  map[int]int{
	//retrives list of course id's with priorities
	var output map[int]int
	output[1] = 2
	return output
}

//func ComputeCombinations //multiplies all course.Classes.length by each other


