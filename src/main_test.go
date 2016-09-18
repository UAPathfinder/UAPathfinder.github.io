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

type Flex struct { //allowes for a mixed array
	f byte
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
	OrderCombos(comboArr)

	if !CompareCombos(combo1, combo1) {
		t.Fatalf("CompareCombos is broke")
	}

	if !(CompareCombos(comboArr[0], combo1) && CompareCombos(comboArr[1], combo2) && CompareCombos(comboArr[2], combo3)) {
		t.Fatalf("OrderCombos Failed.")
	}

}

func TestDoesOverlap(t *testing.T) {
	Class1Start := SimpleParse("09:00:00")
	Class1End := SimpleParse("10:00:00")
	Class2Start := SimpleParse("09:30:00")
	Class2End := SimpleParse("10:30:00")
	if !DoesOverlap(Class1Start, Class1End, Class2Start, Class2End) {
		t.Fatalf("DoesOverlap Failed")
	}
	if !DoesOverlap(Class2Start, Class2End, Class1Start, Class1End) {
		t.Fatalf("DoesOverlap Failed")
	}

}


func TestGenerateCombos(t *testing.T) {
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
        var courses = []Course{course1, course2}

	result := make([]Combo, 0)
	var current Combo

	GenerateCombos(courses, &result, 0, current)
	t.Logf("result number = %d", len(result))
	//t.Logf("depth is %d"), depth)
	for i := range result {
		for x := range result[i].Classes {
			t.Logf("%d,", result[i].Classes[x].ClassId)
		}
		t.Logf("\n")
	}

	if (len(result) != NumCombos(courses)){
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
