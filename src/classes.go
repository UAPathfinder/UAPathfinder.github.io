package main

import (
	"time"
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
	OrCourses []Course //allowes 'oring' of a large number of classes
}

type Combo struct {
        Classes []Class
        Score   int
}

type Criteria struct {
        Breaks Criterion //distance between classes
        Professor Criterion
        EarliestClass Criterion
        LatestClass Criterion
        Days Criterion //if Max is true, listed days are days off
}

type Criterion struct{ //singular of Criteria, basically an advanced key/value pair
        Maximize bool
        Manditory bool
        Weight int
        Time time.Time //for time-related criteria
        Other string //mostly for days
}

type ByScore []Combo                 //implements sort.Interface for []Combo based on Score
func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

type ByStartTime []Class
func (a ByStartTime) Len() int           { return len(a) }
func (a ByStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStartTime) Less(i, j int) bool { return a[j].StartTime.After(a[i].StartTime) }

