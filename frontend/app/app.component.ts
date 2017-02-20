import { Component } from '@angular/core';
import { CourseService } from './course.service';

import { Course } from './models/course';
import { Schedule } from './models/schedule';
import { CombinationsRequest } from './models/combinations-request';
import { CoursesRequest } from './models/combinations-request';
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
	//combinations: Array<Combination>;
  schedules: Array<Schedule>;
	currentSchedule: number = 0;

  classes: Array<Class>;
  currentClass: number = 0;

  filterText: string = "";

	ngOnInit() {
		this.courseService.getCourses()
			.subscribe(
				courses => {
          this.courses = courses;
          //getCourse("3460 210", this.courses);
        },
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
		this.selectedCourses.forEach((course, idx) => {
			course.Priority = idx;
		});
    var requests:any[] = [];
    this.selectedCourses.forEach(function(course){
      requests.push({
        Course: course.Identifier,
        Weight: 10,
        Optional: false,
      })
    })

		this.courseService.getSchedules({
			Courses: requests,
		})
			.subscribe(
				response => {
          //this.schedules = response;
          this.schedules = populateMeetingDays(response);
        },
				// TODO: Handle Properly
				err => console.error(err)
			)
	}


	// Helper method to get a course given a course id.
	 getCourse(id: string, courses: Array<Course>):Course {
		return this.courses
			.find((course) => course.Identifier == id);
	}

  getTime(input: number): string {
     return new Date(1000 * input).toISOString().substr(11, 8)
 }

  function populateMeetingDays(borks: Array<Schedule>): Array<Schedule>{
    console.log("wow, I hit the method");
    for (let schedule of borks){
      for (event of schedule.Classes){
        event.MeetingDays = "";
        if (event.Sunday){
          event.MeetingDays += "S";
        }
        if (event.Monday){
          event.MeetingDays += "M";
        }
        if (event.Tuesday){
          event.MeetingDays += "T";
        }
        if (event.Wednesday){
          event.MeetingDays += "W";
        }
        if (event.Thursday){
          event.MeetingDays += "Th";
        }
        if (event.Friday){
          event.MeetingDays += "F";
        }
        if (event.Saturday){
          event.MeetingDays += "Su";
        }
      }
    }
    return borks;
  }
}

/*
function populateMeetingDays(borks: Array<Class>): Array<Class>{
  //console.log(borks);
  for (let schedule of borks){
    for (event of schedule.Events){
      event.MeetingDays = "test";
      /*if (event.Sunday){
        event.MeetingDays += "S";
      }
      if (event.Monday){
        event.MeetingDays += "M";
      }
      if (event.Tuesday){
        event.MeetingDays += "T";
      }
      if (event.Wednesday){
        event.MeetingDays += "W";
      }
      if (event.Thursday){
        event.MeetingDays += "Th";
      }
      if (event.Friday){
        event.MeetingDays += "F";
      }
      if (event.Saturday){
        event.MeetingDays += "Su";
      }
    }
  }
  return borks;
}

}*/

// Parses time from input elements into a json format.
function parseTime(input: string): string {
  alert("test")
alert(getTime(61800))
}
