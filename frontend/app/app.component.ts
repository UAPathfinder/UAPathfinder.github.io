import { Component } from '@angular/core';
import { CourseService } from './course.service';

import { Course } from './models/course';
import { Class } from './models/class';
import { Schedule } from './models/schedule';
import { ScheduleRequest } from './models/schedule-request';
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

	Weekdays: Weekday[] = JSON.parse(JSON.stringify(WEEKDAYS));
	startTime: string = "07:00";
	endTime: string = "17:00";
	courses: Array<Course> = new Array<Course>();

  schedules: Array<Schedule>;
	currentSchedule: number = 0;

  classes: Array<Class>;
  currentClass: number = 0;

  filterText: string = "";

  tempCourse: Course;

	ngOnInit() {
    this.tempCourse = new Course();
    this.newTempClass();
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
		for (var thisCourse of this.courses){
      for (var thisClass of thisCourse.Classes){
        thisClass.Manditory = true;
        thisClass.Priority = 1;
      }
    }
    var requests:any[] = [];
    this.courses.forEach(function(course){
      requests.push({
        Course: course.Identifier,
        Weight: 10,
        Optional: false,
      })
    })

    this.populateCourseDays(this.courses);

    var scheduleRequest = new ScheduleRequest(this.startTime, this.endTime, this.courses, this.Weekdays);

    console.log(scheduleRequest);
  
		this.courseService.getSchedules(scheduleRequest)
			.subscribe(
				response => {

          this.schedules = response;
        },
				// TODO: Handle Properly
				err => {
          console.log("an error occured");
        }
			)
	}

  populateCourseDays(courses: Course[]){
    for (var course of courses){
      for (var thisClass of course.Classes){
        thisClass.parseWeekdays();
      }
    } 
  }


	// Helper method to get a course given a course id.
	 getCourse(id: string, courses: Array<Course>):Course {
		return this.courses
			.find((course) => course.Identifier == id);
	}

  getTime(input: number): string {
      console.log(input);
      return new Date(1000 * input).toISOString().substr(11, 8);
   }

   newTempClass(){
     var newClass:Class = new Class();
     this.tempCourse.Classes.push(newClass);
   }

   AddCourse(){
     for (var thisClass of this.tempCourse.Classes){
       thisClass.Course = this.tempCourse.Title.String;
     }
     this.courses.push(this.tempCourse);
     this.tempCourse = new Course();
     this.newTempClass();
   }

  getMeetingDays(thisClass: Class) : string{
    var output: string = "";
    if (thisClass.Sunday){
      output += "S";
    }
    if (thisClass.Monday){
      output += "M";
    }
    if (thisClass.Tuesday){
      output += "T";
    }
    if (thisClass.Wednesday){
      output += "W";
    }
    if (thisClass.Thursday){
      output += "Th";
    }
    if (thisClass.Friday){
      output += "F";
    }
    if (thisClass.Saturday){
      output += "Su";
    }
    return output;
  }


}
