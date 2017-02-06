import * as moment from 'moment';

export class Class {
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

//TODO:make this one file.  I'm soooooo lazy
export class ValidInt64 {
	Int64: number;
	Valid: boolean;
}

export class ValidString {
	String: string;
	Valid: boolean;
}
