export class Course {
	// A human readable unique identifier for the course.
	CourseId: number;

	// A descriptive title for the course.
	Name: string;

	// A property of the view. This is true when the course is selected to be
	// combined.
	isSelected: boolean = false;

	Priority: number = 0;

	// TODO: Rename this soon.
	Manditory: boolean;
}

