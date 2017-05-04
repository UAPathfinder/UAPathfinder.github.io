package scheduling

import (
	// "database/sql"
	// "fmt"
	"os"
	"strconv"
	"testing"
	"time"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/mock"
)

//returns the default
func GenerateRequest(NumClasses int, NumCourses int) (scheduleRequest ScheduleRequest) {

	for i := 0; i < NumCourses; i++ {
		course1 := Course{}

		for j := 0; j < NumClasses; j++ {
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

			course1.Classes = append(course1.Classes, class1)
		}

		course1.Title.String = fake.JobTitle()

		scheduleRequest.Courses = append(scheduleRequest.Courses, course1)
	}

	return
}

func GenerateNoConflicts(NumClasses int, NumCourses int) (scheduleRequest ScheduleRequest) {

	scheduleRequest = GenerateRequest(NumClasses, NumCourses)

	HourCounter := 0
	MinuteCounter := 0
	for i := 0; i < NumCourses; i++ {
		course1 := Course{}

		for j := 0; j < NumClasses; j++ {
			if MinuteCounter > 59 {
				MinuteCounter = 0
				HourCounter++
			}

			MinuteString := ""
			if MinuteCounter < 10 {
				MinuteString = "0" + strconv.Itoa(MinuteCounter)
			} else {
				MinuteString = strconv.Itoa(MinuteCounter)
			}

			var class1 Class
			class1.StartTime = MustParseTime(strconv.Itoa(HourCounter) + ":" + MinuteString)
			class1.EndTime = MustParseTime(strconv.Itoa(HourCounter) + ":" + MinuteString)
			class1.Monday = false
			class1.Tuesday = false
			class1.Wednesday = false
			class1.Thursday = false
			class1.Friday = false
			class1.Saturday = false
			class1.Sunday = false

			course1.Classes = append(course1.Classes, class1)
		}

		course1.Title.String = fake.JobTitle()

		scheduleRequest.Courses = append(scheduleRequest.Courses, course1)
	}

	return
}

func TestFindASchedule(t *testing.T) {
	result := FindSchedules(GenerateRequest(1, 1))

	assert.Equal(t, 1, len(result), "more than one scheudle was returned")
	assert.Equal(t, 1, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}

func TestFindTwoSchedules(t *testing.T) {
	result := FindSchedules(GenerateRequest(2, 1))

	assert.Equal(t, 2, len(result), "something other than two schedules")
	assert.Equal(t, 1, len(result[0].Classes), "more than one class is in the scheudle")

	// t.Fail()
}


func TestFindTwoCoursesSchedules(t *testing.T) {
	result := FindSchedules(GenerateRequest(1, 2))

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
}

func TimeSchedule(NumClasses int, NumCourses int) time.Duration {
	scheduleRequest := GenerateNoConflicts(NumClasses, NumCourses)

	start := time.Now()
	_ = FindSchedules(scheduleRequest)
	elapsed := time.Since(start)

	return elapsed

	// fmt.Print(NumClasses)
	// fmt.Print(", ")
	// fmt.Print(NumCourses)
	// fmt.Print(" ")
	// fmt.Print(elapsed)
	// fmt.Print("\n")

	// assert.Equal(t, 0, len(result), "conflicting schedule")

func FloatToString(InputNum float64) string {
	// to convert a float number to a string}

	return strconv.FormatFloat(InputNum, 'f', 6, 64)
}

func TestTimeSchedule(t *testing.T) {
	MaxClasses := 7
	// MaxCourse := 4

	f, err := os.OpenFile("test_output.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// for i := 0; i < MaxCourse; i++ {
	for j := 0; j < MaxClasses; j++ {
		elapsed := TimeSchedule(j, j)

		output := strconv.Itoa(j)
		output += ", "
		output += strconv.Itoa(j)
		output += ", "
		output += FloatToString(elapsed.Seconds())
		// output += "\n"

		if _, err = f.WriteString(output); err != nil {
			panic(err)
		}

		if _, err = f.WriteString("\r\n"); err != nil {
			panic(err)
		}
	}
	// }
}
