package scheduling

import (
	//"bufio"
	//"testing"
	"os"
	"strconv"
	//"strings"
	"text/tabwriter"
	"fmt"

	//"github.com/mibzman/CourseCorrect/ 
)

func DiagnosticPrintCombo(combo  Combo) {
	fmt.Printf("Schedule ------------------------Score: %v \n", combo.Score)
	for i := range combo.Classes {
		//fmt.Print("Schedule %d -----------------------------------", i)
		DiagnosticPrintClass(combo.Classes[i])
	}
}

func DiagnosticPrintCourse(course  Course) {
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 10, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "Name\tCourse-Id\tPriority\tManditory\t")
	fmt.Fprintln(w, ""+course.Name+"\t"+strconv.Itoa(course.CourseId)+"\t"+strconv.Itoa(course.Priority)+"\t"+strconv.FormatBool(course.Manditory)+"\t")
	fmt.Fprintln(w)
	w.Flush()

	for j := range course.Classes {
		DiagnosticPrintClass(course.Classes[j])
	}

	for i := range course.OrCourses {
		DiagnosticPrintCourse(course.OrCourses[i])
	}
}

func DiagnosticPrintClass(class  Class) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "     CourseId\tClassId\tStart Time\tEnd Time\tDays\t")
	fmt.Fprintln(w, "    "+strconv.Itoa(class.CourseId)+" \t"+strconv.Itoa(class.ClassId)+" \t"+class.StartTime.Format("15:04")+"\t"+class.EndTime.Format("15:04")+"\t"+class.MeetingDays+"\t")
	fmt.Fprintln(w)
	w.Flush()
}

/*
func DiagnosticCriteriaHolder() [] Criteria {
	var ynresponse string
	reader := bufio.NewReader(os.Stdin)
	output := make([] Criteria, 0)

	output = append(output, InputCriteria())
	fmt.Print("Would you like to enter another set of criteria? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output = append(output, DiagnosticCriteriaHolder()...)
	}
	return output

}

func InputCriteria()  Criteria {
	fmt.Print("Enter Course \n")
	var output  Criteria
	var ynresponse string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Would you like to maxamize breaks between classes? (y to maxamize, n to minimize): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.Breaks.Maximize = true
	} else {
		output.Breaks.Maximize = false
	}

	fmt.Print("Enter Break's Priority: ")
	tempString, _ := reader.ReadString('\n')
	output.Breaks.Weight, _ = strconv.Atoi(tempString)

	fmt.Print("What would be the earliest acceptable start time for your earliest class?: ")
	otherInput, _ := reader.ReadString('\n')
	output.EarliestClass.Time =  MustParseTime(otherInput + ":00")

	fmt.Print("Is this manditory? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.EarliestClass.Manditory = true
	} else {
		output.EarliestClass.Manditory = false
	}

	fmt.Print("Enter Earliest Class's Priority: ")
	tempString, _ = reader.ReadString('\n')
	output.EarliestClass.Weight, _ = strconv.Atoi(tempString)

	fmt.Print("What would be the latest acceptable end time for your latest class?: ")
	otherInput, _ = reader.ReadString('\n')
	output.LatestClass.Time =  MustParseTime(otherInput + ":00")

	fmt.Print("Is this manditory? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.LatestClass.Manditory = true
	} else {
		output.LatestClass.Manditory = false
	}

	fmt.Print("Enter Latest Class's Priority: ")
	tempString, _ = reader.ReadString('\n')
	output.LatestClass.Weight, _ = strconv.Atoi(tempString)

	fmt.Print("What days would you prefer not to have any classes? (if none, leave blank): ")
	output.Days.Other, _ = reader.ReadString('\n')

	fmt.Print("Is this manditory? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.Days.Manditory = true
	} else {
		output.Days.Manditory = false
	}

	fmt.Print("Enter Day's Priority: ")
	tempString, _ = reader.ReadString('\n')
	output.Days.Weight, _ = strconv.Atoi(tempString)

	return output
}

func CourseHolder() [] Course {
	var ynresponse string
	reader := bufio.NewReader(os.Stdin)
	output := make([] Course, 0)

	output = append(output, InputCourse())
	fmt.Print("Would you like to enter another course? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output = append(output, CourseHolder()...)
	}
	return output

}

func InputCourse()  Course {
	var output  Course
	var ynresponse string
	reader := bufio.NewReader(os.Stdin)

	output = GnenericInputCourse()

	fmt.Print("Would you like to 'or' this course with another? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.OrCourses = OrCourseHolder()
	}
	return output

}

func OrCourseHolder() [] Course {
	output := make([] Course, 0)

	var ynresponse string
	reader := bufio.NewReader(os.Stdin)

	output = append(output, GnenericInputCourse())
	fmt.Print("Would you like to 'or' this course with another? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output = append(output, OrCourseHolder()...)
	}
	return output
}

func GnenericInputCourse()  Course {
	fmt.Print("Enter Course \n")
	var output  Course
	var ynresponse string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Course Id: ")
	_, _ = fmt.Scanf("%d", output.CourseId)

	//fmt.Print("Enter Course Name: ")
	//output.Name, _ = reader.ReadString('\n\n')

	fmt.Print("Enter Course Priority: ")
	tempString, _ := reader.ReadString('\n')
	fmt.Print("Enter Course Id: ")
	output.CourseId, _ = strconv.Atoi(tempString)

	fmt.Print("Enter Course Name: ")
	output.Name, _ = reader.ReadString('\n')

	fmt.Print("Enter Course Priority: ")
	tempString, _ = reader.ReadString('\n')
	output.Priority, _ = strconv.Atoi(tempString)

	fmt.Print("Is this course manditory? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output.Manditory = true
	} else {
		output.Manditory = false
	}

	output.Classes = ClassHolder(output.CourseId)
	return output
}

func ClassHolder(id int) [] Class {
	output := make([] Class, 0)

	var ynresponse string
	reader := bufio.NewReader(os.Stdin)

	output = append(output, GenericInputClass(id))
	fmt.Print("Would you like to add another class to this course? (y/n): ")
	ynresponse, _ = reader.ReadString('\n')
	if strings.Compare(ynresponse, "y\n") == 0 {
		output = append(output, ClassHolder(id)...)
	}
	//fmt.Print(strings.Compare(ynresponse, "y\n"))
	return output
}

func GenericInputClass(id int)  Class {
	fmt.Print("Enter Class \n")
	var otherInput string
	var output  Class
	output.CourseId = id
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Class Id: ")
	tempString, _ := reader.ReadString('\n')
	output.ClassId, _ = strconv.Atoi(tempString)

	fmt.Print("Enter Start Time: ")
	otherInput, _ = reader.ReadString('\n')
	output.StartTime =  MustParseTime(otherInput + ":00")

	fmt.Print("Enter End Time: ")
	otherInput, _ = reader.ReadString('\n')
	output.EndTime =  MustParseTime(otherInput + ":00")

	fmt.Print("Enter Meeting Days: ")
	output.MeetingDays, _ = reader.ReadString('\n')
	return output
}
*/
