import { Course } from './course';
import * as moment from 'moment';
import { WEEKDAYS, Weekday } from './weekday';

// Request made from the client for requesting the possible combinations.
export class ScheduleRequest {

	constructor(startTime: string, endTime: string, courses: Course[], weekdays: Weekday[]){
		this.RawStartTime = startTime;
		this.RawEndTime = endTime;
		this.Courses = courses;
		this.populateWeekdays(weekdays);
	}

	populateWeekdays(weekdays: Weekday[]){
    for (let weekday of weekdays){
      if (weekday.active){
        if (weekday.name = "S"){
          this.Sunday = true;
        }else if (weekday.name = "M"){
          this.Monday = true;
        }else if (weekday.name = "T"){
          this.Tuesday = true;
        }else if (weekday.name = "W"){
          this.Wednesday = true;
        }else if (weekday.name = "Th"){
          this.Thursday = true;
        }else if (weekday.name = "F"){
          this.Friday = true;
        }else if (weekday.name = "Su"){
          this.Sunday = true;
        }
      }
    }
	}

	RawStartTime: string;
	RawEndTime: string;
	Courses: Array<Course>;

	Sunday: boolean = false;
	Monday: boolean = false;
	Tuesday: boolean = false;
	Wednesday: boolean = false;
	Thursday: boolean = false;
	Friday: boolean = false;
	Saturday: boolean = false;
}
