import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';
import { Course } from './mock/course';

@Injectable()
export class MockCourseService {
	getCourses(): Observable<Course[]> {
		return Observable.create((observer) => {
			observer.next([
				{CourseId: "MATH-1620", Name: "Calculus II"},
				{CourseId: "CHEM-1010", Name: "Introduction to Inorganic Chemistry"},
		});
	}
}
