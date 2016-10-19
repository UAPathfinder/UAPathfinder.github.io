import * as moment from 'moment';

export class Class {
	ClassId: number;
	CourseId: number;
	StartTime: moment.Moment;
	EndTime: moment.Moment;
	MeetingDays: string;
	ProfessorName: string;
	MeetingLocation: string;
}

