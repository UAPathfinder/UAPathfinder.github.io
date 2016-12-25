// Request made from the client for requesting the possible combinations.
export interface CombinationsRequest {
	Courses: Array<CoursesRequest>;
    // TODO: Some event stuff
}

export interface CoursesRequest {
    Course: string;

    // TODO: Event properties
}

