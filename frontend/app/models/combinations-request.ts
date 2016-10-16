import { Course } from './course';

// Request made from the client for requesting the possible combinations.
export interface CombinationsRequest {
	StartTime: Date;
	EndTime: Date;
	Courses: Array<Course>;

	// TODO: Change this to an enum or something.
	Days: string;
}

