package scheduling_test

import (
	"." //this is really dumb, something wrong is happening here.
	"sort"
	"testing"
	//"time"
)

func TestSortByStartTime(t *testing.T) {
	var class1 scheduling.Class
	class1.StartTime = scheduling.MustParseTime("07:00")

	var class2 scheduling.Class
	class2.StartTime = scheduling.MustParseTime("08:00")

	var class3 scheduling.Class
	class3.StartTime = scheduling.MustParseTime("09:00")

	arr := []scheduling.Class{class3, class1, class2}
	result := arr

	sort.Sort(scheduling.ByEndTime(result))

	if result[0] != class3 {
		t.Fatal("Returned wrong class", result[0], arr[0])
	}
}
