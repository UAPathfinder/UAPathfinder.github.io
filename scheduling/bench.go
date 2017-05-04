/*
 Bench.go benchmarks 1->100 courses and their optimal schedules based on our algorithm;
 1.) Generates a New Course by generating an arbitrary # of classes;
 2.) For each class, it specifies a random start & end time, as well as days;
 3.) Calculates the nanoseconds taken from running the algorithm and the expanding schedule.
 4.) Further documentation available upon request
*/

package scheduling

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func FindBench() {
	runs := 0
	var cont []Course
	var scheduleRequest ScheduleRequest
	for runs != 100 {
		newCourse := genCourse()
		cont = append(cont, newCourse)
		scheduleRequest.Courses = cont
		start := time.Now()
		result := FindSchedules(scheduleRequest)
		elapsed := time.Since(start)
		taken := elapsed.Nanoseconds()
		printTime(runs+1, int(taken))
		runs++
		if result != nil {
			//have to use this somewhere...
		}
	}
}

func genTime() (string, string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	min, max := 7, 19
	p := r.Perm(max - min + 1)
	a, b := p[0], p[1]
	var start, end string
	if a < b {
		start = strconv.Itoa(a)
		end = strconv.Itoa(b)
	} else {
		start = strconv.Itoa(b)
		end = strconv.Itoa(a)
	}
	end = fmt.Sprintf("%s%s", end, ":00")
	start = fmt.Sprintf("%s%s", start, ":00")
	return start, end
}

//seven days of the week, all represented by flags (see abstraction[s])
func genDays() []bool {
	rand.Seed(time.Now().UTC().UnixNano())
	var days []bool
	min, max := 0, 1 //either true or false.
	i := 0
	for i != 7 {
		happens := rand.Intn(max-min) + min
		days = append(days, happens != 0)
		i++
	}
	return days
}

func printTime(courses int, taken int) {
	runs := strconv.Itoa(courses)
	needed := strconv.Itoa(taken)
	fmt.Println(runs, " ", needed)
}

func genCourse() Course {
	rand.Seed(time.Now().UTC().UnixNano())
	var arbitrary Course
	var options []Class
	min, max := 1, 5
	classes := rand.Intn(max-min) + min
	index := 0
	for index != classes {
		newClass := genClass()
		options = append(options, newClass)
		index++
	}
	arbitrary.Classes = options
	return arbitrary
}

func genClass() Class {
	start, end := genTime()
	days := genDays()
	var arbitrary Class
	arbitrary.StartTime = MustParseTime(start)
	arbitrary.EndTime = MustParseTime(end)
	arbitrary.Monday = days[0]
	arbitrary.Tuesday = days[1]
	arbitrary.Wednesday = days[2]
	arbitrary.Thursday = days[3]
	arbitrary.Friday = days[4]
	arbitrary.Saturday = days[5]
	arbitrary.Sunday = days[6]
	return arbitrary
}
