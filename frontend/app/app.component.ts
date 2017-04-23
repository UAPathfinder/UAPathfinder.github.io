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

	Weekdays: Weekday[] = WEEKDAYS;
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
		this.courses.forEach((course, idx) => {
			course.Priority = idx;
		});
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

          //this.schedules = response;
          //this.schedules = populateMeetingDays(response);
        },
				// TODO: Handle Properly
				err => {
          console.log("an error occured");
        }
			)
	}

  populateCourseDays(courses: Course[]){
    for (let course of courses){
      for (let thisClass of course.Classes){
        for (let weekday of thisClass.Weekdays){
          if (weekday.active){
            if (weekday.name = "S"){
              thisClass.Sunday = true;
            }else if (weekday.name = "M"){
              thisClass.Monday = true;
            }else if (weekday.name = "T"){
              thisClass.Tuesday = true;
            }else if (weekday.name = "W"){
              thisClass.Wednesday = true;
            }else if (weekday.name = "Th"){
              thisClass.Thursday = true;
            }else if (weekday.name = "F"){
              thisClass.Friday = true;
            }else if (weekday.name = "Su"){
              thisClass.Sunday = true;
            }
          }
        }
      }
    } 
  }


	// Helper method to get a course given a course id.
	 getCourse(id: string, courses: Array<Course>):Course {
		return this.courses
			.find((course) => course.Identifier == id);
	}

  getTime(input: number): string {
     return new Date(1000 * input).toISOString().substr(11, 8)
   }

   newTempClass(){
     var newClass:Class = new Class();
     this.tempCourse.Classes.push(newClass);
   }

   AddCourse(){
     this.courses.push(this.tempCourse);
     this.tempCourse = new Course();
     this.newTempClass();
   }

  populateMeetingDays(borks: Array<Schedule>): Array<Schedule>{
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

// Parses time from input elements into a json format.
function parseTime(input: string): string {
  alert("test")
alert(getTime(61800))
}
