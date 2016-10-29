package testing

import (
	"fmt"
	"github.com/mibzman/CourseCorrect/IO"
	"github.com/mibzman/CourseCorrect/mock"
	"github.com/mibzman/CourseCorrect/scheduling"
	"sort"
	"testing"
)

func TestGenerateCombos(t *testing.T) {
	class1 := scheduling.Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("09:00:00"),
		EndTime:         scheduling.MustParseTime("10:00:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := scheduling.Class{
		ClassId:         2,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("11:30:00"),
		EndTime:         scheduling.MustParseTime("13:00:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class3 := scheduling.Class{
		ClassId:         3,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("05:05:00"),
		EndTime:         scheduling.MustParseTime("06:15:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class4 := scheduling.Class{
		ClassId:         4,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("10:25:00"),
		EndTime:         scheduling.MustParseTime("11:45:00"),
		MeetingDays:     "THF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class5 := scheduling.Class{
		ClassId:         5,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("07:45:00"),
		EndTime:         scheduling.MustParseTime("10:00:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class6 := scheduling.Class{
		ClassId:         6,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("18:00:00"),
		EndTime:         scheduling.MustParseTime("19:45:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}

	var arr1 = []scheduling.Class{class1, class2, class3}
	var arr2 = []scheduling.Class{class4, class5, class6}

	course1 := scheduling.Course{
		CourseId:  1,
		Priority:  9,
		Manditory: true,
		Classes:   arr1,
	}
	course2 := scheduling.Course{
		CourseId:  2,
		Priority:  7,
		Manditory: false,
		Classes:   arr2,
	}
	var courses = []scheduling.Course{course1, course2}

	result := scheduling.GenerateCombos(courses)
	t.Logf("result number = %d",
		len(result))
	//t.Logf("depth is %d"), depth)
	for i := range result {
		for x := range result[i].Classes {
			t.Logf("%d,", result[i].Classes[x].
				ClassId)
		}
		t.Logf("\n")
	}

	if len(result) != 9 {
		t.Fatalf("is not right size.  is : %d",
			len(result))
	}
}

func TestGenerateCombos2(t *testing.T) {
	class1 := scheduling.Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("09:00:00"),
		EndTime:         scheduling.MustParseTime("10:00:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := scheduling.Class{
		ClassId:         2,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("11:30:00"),
		EndTime:         scheduling.MustParseTime("13:00:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class3 := scheduling.Class{
		ClassId:         3,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("05:05:00"),
		EndTime:         scheduling.MustParseTime("06:15:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class4 := scheduling.Class{
		ClassId:         4,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("10:25:00"),
		EndTime:         scheduling.MustParseTime("11:45:00"),
		MeetingDays:     "THF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class5 := scheduling.Class{
		ClassId:         5,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("07:45:00"),
		EndTime:         scheduling.MustParseTime("10:00:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class6 := scheduling.Class{
		ClassId:         6,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("18:00:00"),
		EndTime:         scheduling.MustParseTime("19:45:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class7 := scheduling.Class{
		ClassId:         7,
		CourseId:        3,
		StartTime:       scheduling.MustParseTime("20:25:00"),
		EndTime:         scheduling.MustParseTime("21:45:00"),
		MeetingDays:     "THF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class8 := scheduling.Class{
		ClassId:         8,
		CourseId:        3,
		StartTime:       scheduling.MustParseTime("20:45:00"),
		EndTime:         scheduling.MustParseTime("21:00:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class9 := scheduling.Class{
		ClassId:         9,
		CourseId:        3,
		StartTime:       scheduling.MustParseTime("20:00:00"),
		EndTime:         scheduling.MustParseTime("21:45:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}

	var arr1 = []scheduling.Class{class1, class2, class3}
	var arr2 = []scheduling.Class{class4, class5, class6}
	var arr3 = []scheduling.Class{class7, class8, class9}

	course3 := scheduling.Course{
		CourseId:  3,
		Priority:  9,
		Manditory: false,
		Classes:   arr3,
	}
	arr4 := make([]scheduling.Course, 0)
	arr4 = append(arr4, course3)

	course1 := scheduling.Course{
		CourseId:  1,
		Priority:  9,
		Manditory: true,
		Classes:   arr1,
	}
	course2 := scheduling.Course{
		CourseId:  2,
		Priority:  7,
		Manditory: false,
		Classes:   arr2,
		OrCourses: arr4,
	}
	var courses = []scheduling.Course{course1, course2}

	result := scheduling.GenerateCombos(courses)
	t.Logf("result number = %d", len(result))
	//t.Logf("depth is %d"),depth)
	for i := range result {
		for x := range result[i].Classes {
			t.Logf("%d,", result[i].Classes[x].ClassId)
		}
		t.Logf("\n")
	}

	if len(result) != 18 {
		t.Fatalf("is not right size.  is : %d", len(result))
	}
}

func TestGenerateCombos3(t *testing.T) {
	criteria := scheduling.Criteria{
		EarliestClass: scheduling.Criterion{
			Time:      scheduling.MustParseTime("07:00:00"),
			Manditory: true,
			Weight:    10,
		},
		LatestClass: scheduling.Criterion{
			Time:      scheduling.MustParseTime("17:00:00"),
			Manditory: true,
			Weight:    10,
		},
		Days: scheduling.Criterion{
			Other:     "SFS",
			Manditory: true,
			Weight:    10,
		},
	}
	combos := scheduling.GenerateCombos(mock.S3Courses)
	for i := range combos {
		combo := &combos[i]
		sort.Sort(scheduling.ByStartTime(combo.Classes))
		combo.Score = scheduling.ScoreCombo(*combo, criteria)
	}
	sort.Sort(scheduling.ByScore(combos))
	sort.Reverse(scheduling.ByScore(combos))
	for i := range combos {
		fmt.Print(i)
		fmt.Print(" ")
		IO.PrintCombo(combos[i])
	}

}
