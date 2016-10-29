package scheduling

import (
	"sort"
	"time"
)

// Total possible combos, includes ones with class confilicts
func NumCombos(courses []Course) int {
	total := 1
	for i := range courses {
		total *= len(courses[i].Classes)
	}
	return total
}

func DoesHaveOverlap(combo Combo) (bool, int, int) {
	for i := 0; i < len(combo.Classes); i++ {
		for j := i + 1; j < len(combo.Classes); j++ {
			if DoesOverlap(combo.Classes[i], combo.Classes[j]) {
				return true, i, j
			}
		}
	}
	return false, -1, -1

}

func GetCourse(courses []Course, id int) Course {
	for _, course := range courses {
		if course.CourseId == id {
			return course
		}
	}

	output := Course{}
	return output
}

func OrderCombos(combos *[]Combo) {
	sort.Sort(sort.Reverse(ByScore(*combos)))
}

func OrderClasses(combo *Combo) {
	sort.Sort(ByStartTime(combo.Classes))
}

func MinuteDiff(first, second time.Time) int {
	diff := first.Sub(second)
	var minutes = diff.Minutes()
	return int(minutes)
}

func DoesOverlap(Class1, Class2 Class) bool {
	counter := 0
	for i := range Class1.MeetingDays {
		for j := range Class2.MeetingDays {
			if Class1.MeetingDays[i] == Class2.MeetingDays[j] {
				counter += 1
			}
		}
	}
	if counter <= 0 {
		return false
	}
	return !(Class2.StartTime.After(Class1.EndTime) || Class1.StartTime.After(Class2.EndTime))
}
