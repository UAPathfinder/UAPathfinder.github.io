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

	//ui junk
	//I really dont care
	Weekdays: Array<Weekday> = JSON.parse(JSON.stringify(WEEKDAYS));
	RawStartTime: string = "07:00";
	RawEndTime: string = "17:00";
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
