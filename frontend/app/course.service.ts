import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import { Course } from './models/course';
import { Combination } from './models/combination';
import { CombinationsRequest } from './models/combinations-request'

@Injectable()
export class CourseService {
	private courseEndpoint = '/api/v0/courses';
	private combosEndpoint = '/api/v0/schedules';

	constructor(private http: Http) {}

	getCourses(): Observable<Course[]> {
		return this.http.get(this.courseEndpoint)
			.map((res: Response) => res.json())
			// TODO: Handle Error
	}

	getSchedules(request: CombinationsRequest): Observable<Combination[]> {
		return this.http.post(this.combosEndpoint, request)
			.map((res: Response) => res.json());
			// TODO: Handle Error
	}
}

