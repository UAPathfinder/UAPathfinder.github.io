import { Class } from './class';
export class Course {

	// A human readable unique identifier for the course.
	CourseId: number = 0;

	// A descriptive title for the course.
	Identifier: string = "";

	Title: {
		String: string;
		Valid: boolean;
	} = {
		String: "",
		Valid: true
	};

	Classes: Array<Class> = [];

	// A property of the view. This is true when the course is selected to be
	// combined.
	isSelected: boolean = true;
}
