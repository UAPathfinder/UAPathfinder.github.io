package tests

import (
	"github.com/mibzman/CourseCorrect/scheduling"
	"testing"
)

type TestAccessor struct {
	Classes []scheduling.Class
}

func (accessor *TestAccessor) GetClasses(courseIdentifier string) []scheduling.Class {
	results := []scheduling.Class{}
	for _, class := range accessor.Classes {
		if class.Course == courseIdentifier {
			results = append(results, class)
		}
	}

	return results
}

func (accessor *TestAccessor) GetCourse(courseIdentifier string) scheduling.Course {
	return scheduling.Course{}
}

func TestFindSchedules(t *testing.T) {
	accessor := &TestAccessor{
		[]scheduling.Class{
			scheduling.Class{
				Identifier: "A:1",
				Course:     "A",
				Times: scheduling.Times{
					Monday:       true,
					RawStartTime: 6 * 60 * 60 * 60,  // 6 AM
					RawEndTime:   10 * 60 * 60 * 60, // 10 AM
				},
			},
			scheduling.Class{
				Identifier: "B:1",
				Course:     "B",
				Times: scheduling.Times{
					Monday:       true,
					RawStartTime: 11 * 60 * 60 * 60, // 11 AM
					RawEndTime:   13 * 60 * 60 * 60, // 1 PM
				},
			},
		},
	}

	schedules := scheduling.FindSchedules(
		[]string{"A", "B"}, // Courses
		nil,                // Use zero value of EventPropery. Everything is optional and has a weight of 0.
		accessor,
	)

	if len(schedules) != 1 || len(schedules[0].Calendar.Events) != 2 {
		t.Log("Invalid number of schedules found")
		t.Fail()
	}
}
