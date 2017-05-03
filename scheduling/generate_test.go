package scheduling

import (
	// "database/sql"
	// "log"
	"testing"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

func TestFindASchedule(t *testing.T) {
	var class1 Class
	class1.StartTime = MustParseTime("7:00")
	class1.EndTime = MustParseTime("15:00")
	class1.Monday = false
	class1.Tuesday = false
	class1.Wednesday = false
	class1.Thursday = false
	class1.Friday = false
	class1.Saturday = false
	class1.Sunday = false

	course1 := Course{}
	course1.Classes = []Class{class1}

	var scheduleRequest ScheduleRequest
	scheduleRequest.Courses = []Course{course1}

	result := FindSchedules(scheduleRequest)

	assert.Equal(t, 1, len(result), "more than one scheudle was returned")
	assert.Equal(t, 1, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}

func TestFindTwoSchedules(t *testing.T) {
	var class1 Class
	class1.StartTime = MustParseTime("7:00")
	class1.EndTime = MustParseTime("15:00")
	class1.Monday = true
	class1.Tuesday = false
	class1.Wednesday = true
	class1.Thursday = false
	class1.Friday = true
	class1.Saturday = false
	class1.Sunday = false

	var class2 Class
	class1.StartTime = MustParseTime("7:00")
	class1.EndTime = MustParseTime("15:00")
	class1.Monday = true
	class1.Tuesday = true
	class1.Wednesday = true
	class1.Thursday = true
	class1.Friday = true
	class1.Saturday = false
	class1.Sunday = false

	course1 := Course{}
	course1.Classes = []Class{class1, class2}

	var scheduleRequest ScheduleRequest
	scheduleRequest.Courses = []Course{course1}

	result := FindSchedules(scheduleRequest)

	assert.Equal(t, 2, len(result), "something other than two schedules")
	assert.Equal(t, 1, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}

func TestFindTwoCoursesSchedules(t *testing.T) {
	var class1 Class
	class1.StartTime = MustParseTime("7:00")
	class1.EndTime = MustParseTime("15:00")
	class1.Monday = true
	class1.Tuesday = false
	class1.Wednesday = true
	class1.Thursday = false
	class1.Friday = true
	class1.Saturday = false
	class1.Sunday = false

	var class2 Class
	class2.StartTime = MustParseTime("7:00")
	class2.EndTime = MustParseTime("15:00")
	class2.Monday = false
	class2.Tuesday = true
	class2.Wednesday = false
	class2.Thursday = true
	class2.Friday = false
	class2.Saturday = false
	class2.Sunday = false

	course1 := Course{}
	course1.Classes = []Class{class1}

	course2 := Course{}
	course2.Classes = []Class{class2}

	var scheduleRequest ScheduleRequest
	scheduleRequest.Courses = []Course{course1, course2}

	result := FindSchedules(scheduleRequest)

	assert.Equal(t, 1, len(result), "something other than two schedules")
	assert.Equal(t, 2, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}

func TestFindTwoConflictingSchedules(t *testing.T) {
	var class1 Class
	class1.StartTime = MustParseTime("7:00")
	class1.EndTime = MustParseTime("15:00")
	class1.Monday = true
	class1.Tuesday = false
	class1.Wednesday = true
	class1.Thursday = false
	class1.Friday = true
	class1.Saturday = false
	class1.Sunday = false

	var class2 Class
	class2.StartTime = MustParseTime("8:00")
	class2.EndTime = MustParseTime("14:00")
	class2.Monday = true
	class2.Tuesday = false
	class2.Wednesday = true
	class2.Thursday = true
	class2.Friday = false
	class2.Saturday = false
	class2.Sunday = false

	course1 := Course{}
	course1.Classes = []Class{class1}

	course2 := Course{}
	course2.Classes = []Class{class2}

	var scheduleRequest ScheduleRequest
	scheduleRequest.Courses = []Course{course1, course2}

	result := FindSchedules(scheduleRequest)

	assert.Equal(t, 0, len(result), "conflicting schedule")
	// assert.Equal(t, 2, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}
