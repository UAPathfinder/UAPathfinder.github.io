package scheduling

import (
	"fmt"
	"log"
)

// An interface to get classes for FindSchedules. Mockable for testing the
// algorithm against custom data.
type Accessor interface {
	// Given a course identifier, returns a list of classes. Called once for
	// each course for each execution of FindSchedules.
	GetClasses(courseIdentifier string) []Class

	GetCourse(courseIdentifier string) Course
}

type Schedule struct {
	Calendar
	Classes []Class
	Score   int
}

type BySchedule []Schedule

func (s BySchedule) Len() int           { return len(s) }
func (s BySchedule) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BySchedule) Less(i, j int) bool { return s[i].Score < s[j].Score }

// Given a list of Courses, returns the best combination of classes.
// Requiments:
// - Many Soft Constraints:
//   - Slight Overlaps Not Rejected
// - Search Space (Akron: 88*55*45*45*41*33*27*26 @ 1 op/ns = 2.6 hours)
// Solutions:
// - Genetic Algorithm
// - Simulated Annealing
func FindSchedules(courses []string, props map[string]EventProperties, accessor Accessor) []Schedule {
	result := []Schedule{}
	RecursiveFindSchedules(courses, props, accessor, &result, 0, &Schedule{})
	return result
}

// http://stackoverflow.com/questions/17192796/generate-all-combinations-from-multiple-lists
func RecursiveFindSchedules(courses []string, props map[string]EventProperties, accessor Accessor, result *[]Schedule, depth int, current *Schedule) {
	//_ = "breakpoint"
	log.Println("recurse, depth: ", depth)
	log.Println(len(courses))
	if depth == len(courses) {
		// Deepest Case, Called After Builder Cases for 0..(depth - 1)
		*result = append(*result, *current)
		return
	}

	// Builder Case
	course := courses[depth]

	log.Println("course ident:", course)
	// Get Classes for Course
	classes := accessor.GetClasses(course)
	_ = "breakpoint"
	log.Println("GetClasses returns", len(classes), "classes")

	// Get parameters for the course. If does not exist, returns zero value of
	// CourseParam.
	courseProps := props[course]

classesLoop:
	for index, class := range classes {
		log.Println("classesloop")
		if index == 5249 {
			log.Println("It's every class")
		}

		// Try Adding Classes From course to Schedule
		events := class.Events(&courseProps)

		//log.Println("events: ", events)

		// Container to hold potential deletions.
		var pendingDeletions []Event

		cost := 0

		for _, event := range events {
			log.Println("event loop.")
			conflictingEvent := current.Calendar.DoesConflict(event)
			//log.Println("got conflicts")

			if conflictingEvent != nil {
				log.Println("Found A Conflict. Fail If Possible.")
				conflictingProps := conflictingEvent.Properties()

				if !conflictingProps.Optional && !courseProps.Optional {
					// Both Required
					continue classesLoop
				} else if conflictingProps.Optional && !courseProps.Optional {
					// Ours Required
					// Pend Deletion of Other, Keep Ours
					pendingDeletions = append(pendingDeletions, conflictingEvent)
					cost += conflictingProps.Weight
				} else if !conflictingProps.Optional && courseProps.Optional {
					// Other Required, Ours Optional
					cost += courseProps.Weight
					// TODO: WTF?
				} else {
					// Both Optional, Skip Lower Priority
					if conflictingProps.Weight < courseProps.Weight {
						cost += conflictingProps.Weight
						pendingDeletions = append(pendingDeletions, conflictingEvent)
					} else {
						cost += courseProps.Weight
					}
				}
			}
		}

		// Incur Cost
		current.Score -= cost

		for _, deletion := range pendingDeletions {
			switch deletion := deletion.(type) {
			case ClassEvent:
				// Remove All Elements of This Class
				j := 0
				for _, event := range current.Calendar.Events {
					event, ok := event.(ClassEvent)
					if !ok {
						continue
					}

					if event.Class.Identifier != deletion.Class.Identifier {
						current.Calendar.Events[j] = event
						fmt.Println(event)
						j++
					}
				}
				current.Calendar.Events = current.Calendar.Events[:j]
			}
		}

		for _, event := range events {
			current.Calendar.Add(event)
			log.Println("added event: ", event)
		}
		log.Println("past event loop")

		RecursiveFindSchedules(courses, props, accessor, result, depth+1, current)

		//     Add Class to Schedule
		//     Calculate Score of Schedule
		//     Descide Whether Sub-Tree Is Viable
		//     Continue Building By Recursion (depth + 1)

		// If Hard Constraint Violated, Discard All Sub-Schedules
		//     Return
		// Else, Continue Recursing:
		//     RecursiveFindSchedules(courses, result, depth + 1, current)
	}
}
