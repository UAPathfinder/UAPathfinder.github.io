import { Class } from './class';
import * as moment from 'moment';

export class Schedule {
	// The classes included in this combination.
	Events: Event[];	
	Classes: Class[];
	Score: number;
}

export class Event{
	StartTime: moment.Moment;
	EndTime: moment.Moment;
	Weight:   number;
	Optional: boolean;
}
