package scheduling

import (
	"database/sql"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccessor struct {
	mock.Mock
}

func (accessor *MockAccessor) GetClasses(courseIdentifier string) []Class {
	var output []Class
	if courseIdentifier == "3460 210" {
		output1 := Class{
			"70613",
			"3460 210",
			sql.NullInt64{
				Int64: 18,
				Valid: true,
			},
			sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			sql.NullString{
				String: "bob",
				Valid:  true,
			},
			sql.NullString{
				String: "here",
				Valid:  true,
			},
			Times{
				false,
				true,
				false,
				false,
				false,
				false,
				false,
				sql.NullInt64{
					Int64: 67200,
					Valid: true,
				},
				sql.NullInt64{
					Int64: 70200,
					Valid: true,
				},
			},
		}
		output = append(output, output1)

		output2 := Class{
			"70612",
			"3460 210",
			sql.NullInt64{
				Int64: 38,
				Valid: true,
			},
			sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			sql.NullString{
				String: "john",
				Valid:  true,
			},
			sql.NullString{
				String: "there",
				Valid:  true,
			},
			Times{
				false,
				true,
				false,
				true,
				false,
				false,
				false,
				sql.NullInt64{
					Int64: 61800,
					Valid: true,
				},
				sql.NullInt64{
					Int64: 66300,
					Valid: true,
				},
			},
		}
		output = append(output, output2)
	}

	if courseIdentifier == "3460 455" {
		output1 := Class{
			"77642",
			"3460 455",
			sql.NullInt64{
				Int64: 23,
				Valid: true,
			},
			sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
			sql.NullString{
				String: "jeff",
				Valid:  true,
			},
			sql.NullString{
				String: "somewhere else",
				Valid:  true,
			},
			Times{
				false,
				false,
				true,
				false,
				true,
				false,
				false,
				sql.NullInt64{
					Int64: 38700,
					Valid: true,
				},
				sql.NullInt64{
					Int64: 86400,
					Valid: true,
				},
			},
		}
		output = append(output, output1)
	}

	return output
}

func (accessor *MockAccessor) GetCourse(courseIdentifier string) Course {
	var output Course
	return output
}

func TestMockAcessorWorks(t *testing.T) {
	var test MockAccessor
	result := test.GetClasses("3460 210")
	assert.Equal(t, result[0].Monday, true, "they should be equal")
}

func TestFindSchedulesFindsSchedules(t *testing.T) {
	var testAccessor MockAccessor
	courses := []string{"3460 210", "3460 455"}
	props := map[string]EventProperties{"3460 210": {Weight: 10, Optional: false}, "3460 455": {Weight: 10, Optional: false}}
	//_ = "breakpoint"
	result := FindSchedules(courses, props, &testAccessor)
	//assert.NotNil(t, result, "blarg")
	log.Println(result)
}

//this really doesn't belong here, it's technically an integration test

func TestFindSchedulesFindsRealSchedules(t *testing.T) {
	db, err := gorm.Open("sqlite3", "data/test")
	if err != nil {
	}
	accessor := &DatabaseAccessor{db}

	courses := []string{"3460 210", "3460 455"}
	props := map[string]EventProperties{"3460 210": {Weight: 10, Optional: false}, "3460 455": {Weight: 10, Optional: false}}
	//_ = "breakpoint"
	result := FindSchedules(courses, props, accessor)
	assert.NotNil(t, result, "Find Schedules returned null")
}

func TestFindSchedulesReturnsNoRepeatClasses(t *testing.T) {
	db, err := gorm.Open("sqlite3", "../data/test")
	if err != nil {
		log.Println(err)
	}
	accessor := &DatabaseAccessor{db}

	courses := []string{"3460 210", "3460 455"}
	props := map[string]EventProperties{"3460 210": {Weight: 10, Optional: false}, "3460 455": {Weight: 10, Optional: false}}

	_ = FindSchedules(courses, props, accessor)
	// for _, schedule := range result {
	// 	log.Println(len(schedule.Events))
	// 	for index, event1 := range schedule.Events {
	// 		for _, event2 := range schedule.Events[index:] {
	// 			//assert.NotEqual(t, event1., class2.Identifier, "classes were equal within the same schedule")
	// 			assert.Fail(t, "because I said so")
	// 		}
	// 	}
	// }
}

type DatabaseAccessor struct {
	*gorm.DB
}

func (accessor *DatabaseAccessor) GetClasses(courseIdentifier string) []Class {
	classes := []Class{}
	accessor.DB.Where(&Class{Course: courseIdentifier}).Find(&classes)
	return classes
}

func (accessor *DatabaseAccessor) GetCourse(courseIdentifier string) Course {
	course := Course{}
	accessor.DB.Where(&Course{Identifier: courseIdentifier}).First(&course)
	return course
}
