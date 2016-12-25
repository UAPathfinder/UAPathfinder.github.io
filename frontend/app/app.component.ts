import { Component } from '@angular/core';
import { CourseService } from './course.service';

import { Course } from './models/course';
import { Combination } from './models/combination';
import { WEEKDAYS, Weekday } from './models/weekday';

import * as moment from 'moment';

import './rxjs-operators';

@Component({
    selector: 'my-app',
	templateUrl: 'app/app.component.html',
	providers: [ CourseService ],
})
export class AppComponent {
	constructor(private courseService: CourseService) { }

	weekdays: Array<Weekday> = WEEKDAYS;
	startTime: string = "07:00";
	endTime: string = "17:00";
	courses: Array<Course>;
	combinations: Array<Combination>;
	currentCombo: number = 0;

	ngOnInit() {
		this.courseService.getCourses()
			.subscribe(
				courses => this.courses = courses,
				// TODO: Handle Properly
				err => console.error(err)
			)
	}

	selectedCourses: Array<Course> = new Array();
	updateSelected(course: Course, evt: Event) {
		course.isSelected = !course.isSelected;
		if (course.isSelected) {
			// Just Became Selected
			this.selectedCourses.push(course);
		} else {
			// Just Became Un-Selected
			this.selectedCourses.splice(
				this.selectedCourses.indexOf(course), 1);
		}
	}

	// TODO: Proper form handling
	onSubmit() {
		this.courseService.getSchedules({
            Courses: this.selectedCourses.map(course => ({Course: course.Identifier})),
		})
			.subscribe(
				combos => this.combinations = combos,
				// TODO: Handle Properly
				err => console.error(err)
			)
	}

	// Helper method to get a course given a course id.
	// getCourse(id: number): Course {
	// 	return this.courses
	// 		.find((course) => course.CourseId == id)
	// }
}

// Parses time from input elements into a json format.
function parseTime(input: string): string {
	return moment.utc(input, "HH:mm")
		.set({'year': 0, 'month': 0, 'day': 0})
		// Bug in moment or go which puts a wired prefix when using year 0
		.toJSON()
		.replace(/^\+00/, '');
}

