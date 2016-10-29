package testing

import (
	"strconv"
	"testing"
	"github.com/mibzman/CourseCorrect/scheduling"
)

//TODO: not only is this not implemented but this test is disabled

func TestScoreDays(t *testing.T) {
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
		StartTime:       scheduling.MustParseTime("09:25:00"),
		EndTime:         scheduling.MustParseTime("10:45:00"),
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

	var arr1 = []scheduling.Class{class1, class5}
	var arr2 = []scheduling.Class{class3, class4}
	var arr3 = []scheduling.Class{class1, class4}
	var arr4 = []scheduling.Class{class3, class5}

	combo1 := scheduling.Combo{Classes: arr1}
	combo2 := scheduling.Combo{Classes: arr2, Score: 145}
	combo3 := scheduling.Combo{Classes: arr3, Score: 4}
	combo4 := scheduling.Combo{Classes: arr4, Score: 85}
	var _ = []scheduling.Combo{combo2, combo3, combo4}

	criteria := scheduling.Criteria{
		Days: scheduling.Criterion{
			Other:     "SFS",
			Manditory: true,
			Weight:    10,
		},
	}
	combo1.Score = scheduling.ScoreCombo(combo1, criteria)
	t.Logf(strconv.Itoa(combo1.Score))
	if combo1.Score != -100000000 {
		t.Fatalf("Test ScoreDays")
	}

}
