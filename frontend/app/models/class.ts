import * as moment from 'moment';
import { WEEKDAYS, Weekday } from './weekday';

export class Class {
	Identifier: string = "";
	Course: string = "";
	Capicity: ValidInt64 = new ValidInt64();
	Registered: ValidInt64 = new ValidInt64();
	Professor: ValidString = new ValidString();
	Location: ValidString = new ValidString();
	Sunday: boolean = false;
	Monday: boolean = false;
	Tuesday: boolean = false;
	Wednesday: boolean = false;
	Thursday: boolean = false;
	Friday: boolean = false;
	Saturday: boolean = false;
	MeetingDays: string = "";
	RawStartTime: ValidInt64 = new ValidInt64();
	RawEndTime: ValidInt64 = new ValidInt64();

	//ui junk
	//I really dont care
	Weekdays: Array<Weekday> = JSON.parse(JSON.stringify(WEEKDAYS));
	StartTime: string = "07:00";
	EndTime: string = "17:00";
}

//TODO:make this one file.  I'm soooooo lazy
export class ValidInt64 {
	Int64: number = 0;
	Valid: boolean = true;
}

export class ValidString {
	String: string = "";
	Valid: boolean = true;
}
