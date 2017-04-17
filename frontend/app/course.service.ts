import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';

import { Observable } from 'rxjs/Observable';

import { Course } from './models/course';
import { Schedule } from './models/schedule';
import { CombinationsRequest } from './models/combinations-request'

@Injectable()
export class CourseService {
	//poop
	private endpoint = "http://localhost:8080/api/v0/";
	private courseEndpoint = this.endpoint + 'courses';
	private combosEndpoint = this.endpoint + 'schedules';

	constructor(private http: Http) {}

	getCourses(): Observable<Course[]> {
		return this.http.get(this.courseEndpoint)
			.map((res: Response) => res.json())
			// TODO: Handle Error
	}

	// getCombos(request: CombinationsRequest): Observable<Combination[]> {
	// 	console.log(request);
	// 	return this.http.post(this.combosEndpoint, request)
	// 		.map((res: Response) => res.json());
	// 		// TODO: Handle Error
	// }

	getSchedules(request: CombinationsRequest): Observable<Schedule[]> {
		//console.log(request);
		return this.http.post(this.combosEndpoint, request)
			.map((res: Response) => res.json());
			// TODO: Handle Error
	}

}
