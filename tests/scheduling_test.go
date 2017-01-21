package tests

import (
	"github.com/mibzman/CourseCorrect/scheduling"
	"strconv"
	"testing"
)

func TestNumCombos(t *testing.T) {
	class1 := scheduling.Class{ClassId: 1, CourseId: 1, StartTime: scheduling.MustParseTime("09:00:00"), EndTime: scheduling.MustParseTime("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := scheduling.Class{ClassId: 2, CourseId: 1, StartTime: scheduling.MustParseTime("11:30:00"), EndTime: scheduling.MustParseTime("13:00:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class3 := scheduling.Class{ClassId: 3, CourseId: 1, StartTime: scheduling.MustParseTime("05:05:00"), EndTime: scheduling.MustParseTime("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := scheduling.Class{ClassId: 4, CourseId: 2, StartTime: scheduling.MustParseTime("09:25:00"), EndTime: scheduling.MustParseTime("10:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := scheduling.Class{ClassId: 5, CourseId: 2, StartTime: scheduling.MustParseTime("07:45:00"), EndTime: scheduling.MustParseTime("10:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class6 := scheduling.Class{ClassId: 6, CourseId: 2, StartTime: scheduling.MustParseTime("18:00:00"), EndTime: scheduling.MustParseTime("19:45:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

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
	var arr3 = []scheduling.Course{course1, course2}
	output := scheduling.NumCombos(arr3)
	if output != 9 {
		t.Fatalf("NumCombos is not correct. Output should be 9, is: %d", output)
	}
}

func TestOrderCombos(t *testing.T) {
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

	combo1 := scheduling.Combo{Classes: arr1, Score: 45}
	combo2 := scheduling.Combo{Classes: arr2, Score: 145}
	combo3 := scheduling.Combo{Classes: arr3, Score: 4}
	combo4 := scheduling.Combo{Classes: arr4, Score: 85}

	var comboArr = []scheduling.Combo{combo1, combo2, combo3, combo4}
	scheduling.OrderCombos(&comboArr)

	if !scheduling.CompareCombos(combo1, combo1) {
		t.Fatalf("scheduling.CompareCombos: combo1 != combo1")
	}

	if scheduling.CompareCombos(combo1, combo2) {
		t.Fatalf("scheduling.CompareCombos: combo1 == combo2")
	}

	if !scheduling.CompareCombos(comboArr[0], combo2) && scheduling.CompareCombos(comboArr[1], combo4) && scheduling.CompareCombos(comboArr[2], combo1) {

		t.Logf(strconv.Itoa(comboArr[0].Score))
		t.Logf(strconv.Itoa(comboArr[1].Score))
		t.Logf(strconv.Itoa(comboArr[2].Score))
		t.Logf(strconv.Itoa(comboArr[3].Score))
		t.Fatalf("OrderCombos Failed.")

		//for _, combo := range comboArr {
		//PrintCombo(combo)
		//}
	}

}

func TestOrderClasses(t *testing.T) {
	class1 := scheduling.Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("09:00:00"),
		EndTime:         scheduling.MustParseTime("10:00:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class2 := scheduling.Class{
		ClassId:         2,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("11:30:00"),
		EndTime:         scheduling.MustParseTime("13:00:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class3 := scheduling.Class{
		ClassId:         3,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("10:05:00"),
		EndTime:         scheduling.MustParseTime("11:15:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class4 := scheduling.Class{
		ClassId:         4,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("11:25:00"),
		EndTime:         scheduling.MustParseTime("12:45:00"),
		MeetingDays:     "THF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	class5 := scheduling.Class{
		ClassId:         5,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("13:45:00"),
		EndTime:         scheduling.MustParseTime("14:00:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}

	classes := []scheduling.Class{class4, class2, class3, class5, class1}
	combo := scheduling.Combo{Classes: classes}

	scheduling.OrderClasses(&combo)

	if !(combo.Classes[0] == class1 && combo.Classes[1] == class3 && combo.Classes[2] == class4 && combo.Classes[3] == class2 && combo.Classes[4] == class5) {
		t.Log(combo)
		t.Fatalf("OrderClasses failed.")
	}

}

func TestDoesHaveOverlap(t *testing.T) {
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
	//class3 := scheduling.Class{ClassId: 3,
	// CourseId: 1,
	// StartTime: scheduling.MustParseTime("05:05:00"),
	// EndTime: scheduling.MustParseTime("06:15:00"),
	// MeetingDays: "MW",
	// ProfessorName: "Bob Jones",
	// MeetingLocation: "Leigh Hall", }

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
		StartTime:       scheduling.MustParseTime("10:45:00"),
		EndTime:         scheduling.MustParseTime("11:00:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}

	var noOverlap = []scheduling.Class{class1,
		class5}
	//var arr2 = []scheduling.Class{class3, class4}
	var hasOverlap = []scheduling.Class{class1, class4}
	//var arr4 = []scheduling.Class{class3, class5}

	combo1 := scheduling.Combo{Classes: noOverlap, Score: 45}
	combo2 := scheduling.Combo{Classes: hasOverlap, Score: 145}
	//combo3 := scheduling.Combo{Classes: arr3, Score: 4}
	//combo4 := scheduling.Combo{Classes: arr4, Score: 85}

	output1, _, _ := scheduling.DoesHaveOverlap(combo1)
	output2, _, _ := scheduling.DoesHaveOverlap(combo2)

	if output1 {
		t.Fatalf("no Overlap")
	}
	if !output2 {
		t.Fatalf("has overlap")
	}
}

func TestDoesOverlap(t *testing.T) {
	Class1Start := scheduling.MustParseTime("09:00:00")
	Class1End := scheduling.MustParseTime("10:00:00")
	Class2Start := scheduling.MustParseTime("09:30:00")
	Class2End := scheduling.MustParseTime("10:30:00")
	iosClass1 := scheduling.Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       Class1Start,
		EndTime:         Class1End,
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	iosClass2 := scheduling.Class{
		ClassId:         2,
		CourseId:        1,
		StartTime:       Class2Start,
		EndTime:         Class2End,
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	iosClass3 := scheduling.Class{
		ClassId:         2,
		CourseId:        1,
		StartTime:       Class2Start,
		EndTime:         Class2End,
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	}
	//Class3End := scheduling.MustParseTime("11:30:00")
	//output := scheduling.DoesOverlap(Class1Start,
	// Class1End,
	// Class1End,
	// Class3End)
	//t.Fatalf("for same class: %t", output)
	if scheduling.DoesOverlap(iosClass1, iosClass2) {
		t.Fatalf("scheduling.DoesOverlap Failed")
	}

	if !scheduling.DoesOverlap(iosClass1, iosClass3) {
		t.Fatalf("scheduling.DoesOverlap Failed2")
	}

}

/*
func scheduling.CompareCombos(input1, input2 scheduling.Combo) bool {
	for i := range input1.Classes {
		if input1.Classes[i] != input2.Classes[i] {
			return false
		}
	}
	return true
}*/

func TestTime(t *testing.T) {
	test := scheduling.MustParseTime("10:00:00")
	if test.After(test) {
		t.Fatalf("no")
	}
}
