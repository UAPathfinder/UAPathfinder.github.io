package scheduling

//strictly speaking this doesn't need to be its own file, but its big and important and will likely get bigger.
//plus the testing should be seperated

func GenerateCombos(courses []Course) []Combo {
	var current Combo
	result := make([]Combo, 0)
	RecursiveGenerateCombos(courses, &result, 0, current)
	return result
}

//the insperation for this algorithm comes from here:
//http://stackoverflow.com/questions/17192796/generate-all-combinations-from-multiple-lists
//we need recursion because the number of courses is going to vary between users
func RecursiveGenerateCombos(courses []Course, result *[]Combo, depth int, current Combo) {
	if depth == len(courses) {
		hasOverlap, issue1, issue2 := DoesHaveOverlap(current)
		if !hasOverlap {
			//horray, no overlaps!  apend the current combo to the result
			*result = append(*result, current)
		} else {
			//get the course objects so we can get the manditory and priority members
			//later this will be a database call

			course1 := GetCourse(courses, current.Classes[issue1].CourseId)
			course2 := GetCourse(courses, current.Classes[issue2].CourseId)
			//if neither conflicting courses are manditory
			if !course1.Manditory && !course2.Manditory {
				//we drop the one that has a lower priority
				if course1.Priority < course2.Priority {
					current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
					current.Score -= course1.Priority
					RecursiveGenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				} else {
					current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
					current.Score -= course2.Priority
					RecursiveGenerateCombos(courses, result, depth, current) //kicks it back with the same depth to check for overlaps again
				}
				//the cases that one is manditory but the other isn't
			} else if !course2.Manditory {
				current.Classes = append(current.Classes[:issue2], current.Classes[issue2+1:]...)
				current.Score -= course2.Priority
				//kicks it back with the same depth to check for overlaps again
				RecursiveGenerateCombos(courses, result, depth, current)
			} else if !course1.Manditory {
				current.Classes = append(current.Classes[:issue1], current.Classes[issue1+1:]...)
				current.Score -= course2.Priority
				//kicks it back with the same depth to check for overlaps again
				RecursiveGenerateCombos(courses, result, depth, current)
			} //otherwise don't append because the whole combo does not fit the user's requirements
		}

	} else {
		//this loop handles the 'root course'
		currentCourse := courses[depth]
		for i := 0; i < len(currentCourse.Classes); i++ {
			var tempCurrent Combo
			//append the first class in the current course to the current combo, then fires GenerateCombos again
			tempCurrent.Classes = append(current.Classes, currentCourse.Classes[i])
			RecursiveGenerateCombos(courses, result, depth+1, tempCurrent)
			//yes, it's recursive.  it only goes [depth] layers deep before returning so the stack shouldn't overflow for a reasonable number of courses
		}
		//this loop goes through all the 'ored' courses within the 'root' course
		if len(currentCourse.OrCourses) > 0 {
			for p := 0; p < len(currentCourse.OrCourses); p++ {
				for i := 0; i < len(currentCourse.OrCourses[p].Classes); i++ {
					//the exact same loop as above but for the 'ored' course
					var tempCurrent Combo
					tempCurrent.Classes = append(current.Classes, currentCourse.OrCourses[p].Classes[i])
					RecursiveGenerateCombos(courses, result, depth+1, tempCurrent)
				}

			}
		}
	}
}
