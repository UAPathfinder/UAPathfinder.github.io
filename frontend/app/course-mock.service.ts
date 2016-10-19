import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Observer } from 'rxjs/Observer';
import { Course } from './models/course';

@Injectable()
export class MockCourseService {
	getCourses(): Observable<Course[]> {
		return Observable.create((observer: Observer<Course[]>) => {
			observer.next(<Course[]>[
				{
					// CourseId: "MATH-1620",
					CourseId: 1,
					Name: "Calculus II",
				},
				{
					// CourseId: "CHEM-1010",
					// TODO: Use strings instead of numbers.
					CourseId: 2,
					Name: "Introduction to Inorganic Chemistry",
				},
			]);
		});
	}
}
