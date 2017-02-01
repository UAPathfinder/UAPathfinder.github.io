import { Class } from './class';
import * as moment from 'moment';

export class Schedule {
	// The classes included in this combination.
	Events: Event[];
	Classes: Class[];
	Score: number;
}

export class Event{
	Day: number;
	Identifier: string;
	Course: string;
	Capicity: ValidInt64;
	Registered: ValidInt64;
	Professor: ValidString;
	Location: ValidString;
	Sunday: boolean;
	Monday: boolean;
	Tuesday: boolean;
	Wednesday: boolean;
	Thursday: boolean;
	Friday: boolean;
	Saturday: boolean;
	MeetingDays: string;
	RawStartTime: ValidInt64;
	RawEndTime: ValidInt64;
}

export class ValidInt64 {
	Int64: number;
	Valid: boolean;
}

export class ValidString {
	String: string;
	Valid: boolean;
}
