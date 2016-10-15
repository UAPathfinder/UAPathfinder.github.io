import { Component } from '@angular/core';

import { Course } from './course';

@Component({
    selector: 'my-app',
	templateUrl: 'app/app.component.html',
})
export class AppComponent {
	// TODO: Source from backend
	courses: Array<Courses> = [
		{id: "MATH-1620", name: "Calculus II"},
		{id: "CHEM-1010", name: "Introduction to Inorganic Chemistry"},
	];

	combos: Array<Array<Courses>> = [
		[this.courses[0], this.courses[1]],
		[this.courses[1], this.courses[0]],
	];

	currentCombo = this.combos[0];
}
