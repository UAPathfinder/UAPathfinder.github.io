package scheduling

var MockIOSClasses = []Class{
	Class{
		ClassId:         1,
		CourseId:        1,
		StartTime:       MustParseTime("13:10:00"),
		EndTime:         MustParseTime("14:00:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var MockIOSCourse = Course{CourseId: 1,
	Priority:  9,
	Manditory: true,
	Classes:   MockIOSClasses,
}

var MockDataStructuresClasses = []Class{
	Class{
		ClassId:         2,
		CourseId:        2,
		StartTime:       MustParseTime("14:15:00"),
		EndTime:         MustParseTime("15:05:00"),
		MeetingDays:     "MWF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	Class{
		ClassId:         3,
		CourseId:        2,
		StartTime:       MustParseTime("17:10:00"),
		EndTime:         MustParseTime("18:25:00"),
		MeetingDays:     "MW",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var MockDataStructuresCourse = Course{CourseId: 2,
	Priority:  9,
	Manditory: true,
	Classes:   MockDataStructuresClasses,
}

var MockOOPClasses = []Class{
	Class{
		ClassId:         4,
		CourseId:        3,
		StartTime:       MustParseTime("15:15:00"),
		EndTime:         MustParseTime("16:30:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}
var MockOOPCourse = Course{
	CourseId:  3,
	Priority:  9,
	Manditory: true,
	Classes:   MockOOPClasses,
}

var MockWebClasses = []Class{
	Class{
		ClassId:         5,
		CourseId:        4,
		StartTime:       MustParseTime("13:45:00"),
		EndTime:         MustParseTime("15:00:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	Class{
		ClassId:         6,
		CourseId:        4,
		StartTime:       MustParseTime("17:10:00"),
		EndTime:         MustParseTime("18:25:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var MockWebCourse = Course{
	CourseId:  4,
	Priority:  9,
	Manditory: true,
	Classes:   MockWebClasses,
}

var MockStastisticsClasses = []Class{
	Class{
		ClassId:         7,
		CourseId:        5,
		StartTime:       MustParseTime("14:05:00"),
		EndTime:         MustParseTime("15:05:00"),
		MeetingDays:     "MTWHF",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
	Class{
		ClassId:         8,
		CourseId:        5,
		StartTime:       MustParseTime("18:05:00"),
		EndTime:         MustParseTime("19:45:00"),
		MeetingDays:     "TH",
		ProfessorName:   "Bob Jones",
		MeetingLocation: "Leigh Hall",
	},
}

var MockStatisticsCourse = Course{
	CourseId:  5,
	Priority:  9,
	Manditory: true,
	Classes:   MockStastisticsClasses,
}

var MockCourses = []Course{
	MockIOSCourse,
	MockDataStructuresCourse,
	MockOOPCourse,
	MockWebCourse,
	MockStatisticsCourse,
}
