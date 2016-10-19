package mock

import (
	"github.com/mibzman/CourseCorrect/scheduling"
)

var IOSClasses = []scheduling.Class{
	scheduling.Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       scheduling.MustParseTime("13:10:00"),
		EndTime:         scheduling.MustParseTime("14:00:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var IOSCourse = scheduling.Course{
	CourseId:  1,
	Priority:  9,
	Manditory: true,
	Classes:   IOSClasses,
	Name:      "iOS Application Development: From Zero to Hero",
}

var DataStructuresClasses = []scheduling.Class{
	scheduling.Class{
		ClassId:         2,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("14:15:00"),
		EndTime:         scheduling.MustParseTime("15:05:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	scheduling.Class{
		ClassId:         3,
		CourseId:        2,
		StartTime:       scheduling.MustParseTime("17:10:00"),
		EndTime:         scheduling.MustParseTime("18:25:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var DataStructuresCourse = scheduling.Course{
	CourseId:  2,
	Priority:  9,
	Manditory: true,
	Classes:   DataStructuresClasses,
	Name:      "An Introduction to Structuring Data",
}

var OOPClasses = []scheduling.Class{
	scheduling.Class{
		ClassId:         4,
		CourseId:        3,
		StartTime:       scheduling.MustParseTime("15:15:00"),
		EndTime:         scheduling.MustParseTime("16:30:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}
var OOPCourse = scheduling.Course{
	CourseId:  3,
	Priority:  9,
	Manditory: true,
	Classes:   OOPClasses,
	Name:      "Object Oriented Programming Concepts I",
}

var WebClasses = []scheduling.Class{
	scheduling.Class{
		ClassId:         5,
		CourseId:        4,
		StartTime:       scheduling.MustParseTime("13:45:00"),
		EndTime:         scheduling.MustParseTime("15:00:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	scheduling.Class{
		ClassId:         6,
		CourseId:        4,
		StartTime:       scheduling.MustParseTime("17:10:00"),
		EndTime:         scheduling.MustParseTime("18:25:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var WebCourse = scheduling.Course{
	CourseId:  4,
	Priority:  9,
	Manditory: true,
	Classes:   WebClasses,
	Name:      "Introduction to HTTP, JavaScript and NPM",
}

var StastisticsClasses = []scheduling.Class{
	scheduling.Class{
		ClassId:         7,
		CourseId:        5,
		StartTime:       scheduling.MustParseTime("14:05:00"),
		EndTime:         scheduling.MustParseTime("15:05:00"),
		MeetingDays:     "MTWHF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	scheduling.Class{
		ClassId:         8,
		CourseId:        5,
		StartTime:       scheduling.MustParseTime("18:05:00"),
		EndTime:         scheduling.MustParseTime("19:45:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var StatisticsCourse = scheduling.Course{
	CourseId:  5,
	Priority:  9,
	Manditory: true,
	Classes:   StastisticsClasses,
	Name:      "Probability and Statistics for Engineering",
}

var Courses = []scheduling.Course{
	IOSCourse,
	DataStructuresCourse,
	OOPCourse,
	WebCourse,
	StatisticsCourse,
}
