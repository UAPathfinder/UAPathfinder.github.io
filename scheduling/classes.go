package scheduling

import (
	"database/sql"
	// "log"
	"time"
)

// A singular class. A class is something you could attend. There are often
// many classes for each course.
// Example: Data Structures starting at 3 PM in room 301 with professor x
type Class struct {
	// Class identifier string. Human readable.
	Identifier string //unused, old DB stuff

	// Course identifier string. Human readable.
	Course string //title of parent course

	// TODO: Omit Information About Validity In JSON
	Capicity   sql.NullInt64  //unused
	Registered sql.NullInt64  //unused
	Professor  sql.NullString //unused
	Location   sql.NullString //unused

	Priority  int  //unused
	Manditory bool //unused
	//the oposite of manditory, becuase I don't feel like inverting the already
	//written logic in generate.  sue me
	Optional bool //unused

	Times
}

func (Class Class) ExistsIn(Arr []Class) bool {
	for i := range Arr {
		if Arr[i] == Class {
			return true
		}
	}

	return false
}

type Times struct {
	// TODO: How are null values handled without nullable?
	Sunday    bool
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool

	RawStartTime string //ui only
	RawEndTime   string //ui only

	StartTime time.Time
	EndTime   time.Time
}

type ByStartTime []Class

func (s ByStartTime) Len() int           { return len(s) }
func (s ByStartTime) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByStartTime) Less(i, j int) bool { return s[i].StartTime.Before(s[j].StartTime) }

// A group of classes which share some common characteristics. For example,
// 3960:401 Data Structures
type Course struct {
	Classes []Class

	Title sql.NullString

	//note: we don't actually ask for these
	//TODO: add to UI
}

type ScheduleRequest struct {
	Courses []Course

	Times //we can ignore this, I'll remove it from the ui closer to the demo
}

func (Request *ScheduleRequest) ParseTime() {
	for CourseID, ThisCourse := range Request.Courses {
		for ClassID, ThisClass := range ThisCourse.Classes {
			//have I mentioned how much I hate this?
			Request.Courses[CourseID].Classes[ClassID].StartTime = SimpleParse(ThisClass.RawStartTime)
			Request.Courses[CourseID].Classes[ClassID].EndTime = SimpleParse(ThisClass.RawEndTime)
		}
	}

	Request.StartTime = SimpleParse(Request.RawStartTime)
	Request.EndTime = SimpleParse(Request.RawEndTime)
}

func SimpleParse(input string) (output time.Time) {
	output, _ = time.Parse("15:04", input)
	return
}
