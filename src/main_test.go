package main

import (
	"testing"
	"time"
	//"fmt"
)

func TestGetCourses(t *testing.T) {
	output := GetCourses()
	value := output[1]
	if value != 2 {
		t.Fatalf("Get Courses is not returning the testing output: output[1] = 2")
	}
}

func TestFillCourses(t *testing.T) {
	input := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	FillCourses(input)
	if input[1] != 2 {
		t.Fatalf("output is not correct.  output[1] = %d", input[1])
	}
}

func TestNumCombos(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := Class{ClassId: 2, CourseId: 1, StartTime: SimpleParse("11:30:00"), EndTime: SimpleParse("13:00:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("05:05:00"), EndTime: SimpleParse("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("09:25:00"), EndTime: SimpleParse("10:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("07:45:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class6 := Class{ClassId: 6, CourseId: 2, StartTime: SimpleParse("18:00:00"), EndTime: SimpleParse("19:45:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	var arr1 = []Class{class1, class2, class3}
	var arr2 = []Class{class4, class5, class6}

	course1 := Course{CourseId: 1, Priority: 9, Manditory: true, Classes: arr1}
	course2 := Course{CourseId: 2, Priority: 7, Manditory: false, Classes: arr2}
	var arr3 = []Course{course1, course2}
	output := NumCombos(arr3)
	if output != 9 {
		t.Fatalf("NumCombos is not correct.  output should be 9, is: %d", output)
	}

}

func TestOrderCombos(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("05:05:00"), EndTime: SimpleParse("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("09:25:00"), EndTime: SimpleParse("10:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("07:45:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	var arr1 = []Class{class1, class5}
	var arr2 = []Class{class3, class4}
	var arr3 = []Class{class1, class4}
	var arr4 = []Class{class3, class5}

	combo1 := Combo{Classes: arr1, Score: 45}
	combo2 := Combo{Classes: arr2, Score: 145}
	combo3 := Combo{Classes: arr3, Score: 4}
	combo4 := Combo{Classes: arr4, Score: 85}

	var comboArr = []Combo{combo1, combo2, combo3, combo4}
	OrderCombos(&comboArr)

	if !CompareCombos(combo1, combo1) {
		t.Fatalf("CompareCombos is broke")
	}

	if !(CompareCombos(comboArr[0], combo1) && CompareCombos(comboArr[1], combo2) && CompareCombos(comboArr[2], combo3)) {
		t.Fatalf("OrderCombos Failed.")
	}

}

func TestOrderClasses(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class2 := Class{ClassId: 2, CourseId: 1, StartTime: SimpleParse("11:30:00"), EndTime: SimpleParse("13:00:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("10:05:00"), EndTime: SimpleParse("11:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("11:25:00"), EndTime: SimpleParse("12:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("13:45:00"), EndTime: SimpleParse("14:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	classes := []Class{class1, class2, class3, class5, class4}
	combo := Combo{Classes: classes}

	OrderClasses(&combo)

	if !(combo.Classes[0] == class1 && combo.Classes[1] == class3 && combo.Classes[2] == class4 && combo.Classes[3] == class2 && combo.Classes[4] == class5) {
		t.Log(combo)
		t.Fatalf("OrderClasses failed.")
	}

}

func TestDoesHaveOverlap(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	//class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("05:05:00"), EndTime: SimpleParse("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("09:25:00"), EndTime: SimpleParse("10:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("10:45:00"), EndTime: SimpleParse("11:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	var noOverlap = []Class{class1, class5}
	//var arr2 = []Class{class3, class4}
	var hasOverlap = []Class{class1, class4}
	//var arr4 = []Class{class3, class5}

	combo1 := Combo{Classes: noOverlap, Score: 45}
	combo2 := Combo{Classes: hasOverlap, Score: 145}
	//combo3 := Combo{Classes: arr3, Score: 4}
	//combo4 := Combo{Classes: arr4, Score: 85}

	output1, _, _ := DoesHaveOverlap(combo1)
	output2, _, _ := DoesHaveOverlap(combo2)

	if output1 {
		t.Fatalf("no Overlap")
	}
	if !output2 {
		t.Fatalf("has overlap")
	}
}

func TestDoesOverlap(t *testing.T) {
	Class1Start := SimpleParse("09:00:00")
	Class1End := SimpleParse("10:00:00")
	Class2Start := SimpleParse("09:30:00")
	Class2End := SimpleParse("10:30:00")
	//Class3End := SimpleParse("11:30:00")
	//output := DoesOverlap(Class1Start, Class1End, Class1End, Class3End)
	//t.Fatalf("for same class: %t", output)
	if !DoesOverlap(Class1Start, Class1End, Class2Start, Class2End) {
		t.Fatalf("DoesOverlap Failed")
	}
	if DoesOverlap(Class1Start, Class2Start, Class1End, Class2End) {
		t.Fatalf("DoesOverlap Failed")
	}

}

func TestGenerateCombos(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := Class{ClassId: 2, CourseId: 1, StartTime: SimpleParse("11:30:00"), EndTime: SimpleParse("13:00:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("05:05:00"), EndTime: SimpleParse("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("10:25:00"), EndTime: SimpleParse("11:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("07:45:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class6 := Class{ClassId: 6, CourseId: 2, StartTime: SimpleParse("18:00:00"), EndTime: SimpleParse("19:45:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	var arr1 = []Class{class1, class2, class3}
	var arr2 = []Class{class4, class5, class6}

	course1 := Course{CourseId: 1, Priority: 9, Manditory: true, Classes: arr1}
	course2 := Course{CourseId: 2, Priority: 7, Manditory: false, Classes: arr2}
	var courses = []Course{course1, course2}

	result := make([]Combo, 0)
	var current Combo

	AllCourses = append(AllCourses, course1)
	AllCourses = append(AllCourses, course2)

	GenerateCombos(courses, &result, 0, current)
	t.Logf("result number = %d", len(result))
	//t.Logf("depth is %d"), depth)
	for i := range result {
		for x := range result[i].Classes {
			t.Logf("%d,", result[i].Classes[x].ClassId)
		}
		t.Logf("\n")
	}

	if len(result) != 9 {
		t.Fatalf("is not right size.  is : %d", len(result))
	}

	//t.FailNow()

}

func TestGenerateCombos2(t *testing.T) {
	class1 := Class{ClassId: 1, CourseId: 1, StartTime: SimpleParse("09:00:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MWF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	//The sample data is all the same prof and location because we're not testing those portions right now
	class2 := Class{ClassId: 2, CourseId: 1, StartTime: SimpleParse("11:30:00"), EndTime: SimpleParse("13:00:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class3 := Class{ClassId: 3, CourseId: 1, StartTime: SimpleParse("05:05:00"), EndTime: SimpleParse("06:15:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class4 := Class{ClassId: 4, CourseId: 2, StartTime: SimpleParse("10:25:00"), EndTime: SimpleParse("11:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class5 := Class{ClassId: 5, CourseId: 2, StartTime: SimpleParse("07:45:00"), EndTime: SimpleParse("10:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class6 := Class{ClassId: 6, CourseId: 2, StartTime: SimpleParse("18:00:00"), EndTime: SimpleParse("19:45:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class7 := Class{ClassId: 7, CourseId: 3, StartTime: SimpleParse("20:25:00"), EndTime: SimpleParse("21:45:00"), MeetingDays: "THF", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class8 := Class{ClassId: 8, CourseId: 3, StartTime: SimpleParse("20:45:00"), EndTime: SimpleParse("21:00:00"), MeetingDays: "MW", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}
	class9 := Class{ClassId: 9, CourseId: 3, StartTime: SimpleParse("20:00:00"), EndTime: SimpleParse("21:45:00"), MeetingDays: "TH", ProfessorName: "Bob Jones", MeetingLocation: "Leigh Hall"}

	var arr1 = []Class{class1, class2, class3}
	var arr2 = []Class{class4, class5, class6}
	var arr3 = []Class{class7, class8, class9}

	course3 := Course{CourseId: 3, Priority: 9, Manditory: false, Classes: arr3}
	arr4 := make([]Course, 0)
	arr4 = append(arr4, course3)

	course1 := Course{CourseId: 1, Priority: 9, Manditory: true, Classes: arr1}
	course2 := Course{CourseId: 2, Priority: 7, Manditory: false, Classes: arr2, OrCourses: arr4}
	var courses = []Course{course1, course2}

	result := make([]Combo, 0)
	var current Combo

	AllCourses = append(AllCourses, course1)
	AllCourses = append(AllCourses, course2)

	GenerateCombos(courses, &result, 0, current)
	t.Logf("result number = %d", len(result))
	//t.Logf("depth is %d"), depth)
	for i := range result {
		for x := range result[i].Classes {
			t.Logf("%d,", result[i].Classes[x].ClassId)
		}
		t.Logf("\n")
	}

	if len(result) != 18 {
		t.Fatalf("is not right size.  is : %d", len(result))
	}

	//t.FailNow()

}

func CompareCombos(input1, input2 Combo) bool {
	for i := range input1.Classes {
		if input1.Classes[i] != input2.Classes[i] {
			return false
		}
	}
	return true
}

func SimpleParse(input string) time.Time {
	output, _ := time.Parse("15:04:05", input)
	return output
}

func TestTime(t *testing.T) {
	test := SimpleParse("10:00:00")
	if test.After(test) {
		t.Fatalf("no")
	}
}
