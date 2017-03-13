import { Course } from './course';
import * as moment from 'moment';

// Request made from the client for requesting the possible combinations.
export interface CombinationsRequest {
	// StartTime: string;
	// EndTime: string;
	// Courses: Array<Course>;
	//
	// // TODO: Change this to an enum or something.
	// Days: string;

	Courses: Array<CoursesRequest>;
}

export interface CoursesRequest {
	Course: string;
	Weight: number;
	Optional: boolean;
}
