import { Pipe, PipeTransform } from '@angular/core';
import * as moment from 'moment';

// Displays times in UTC
@Pipe({
	name: 'time',
})
export class TimePipe implements PipeTransform {
	transform(
		input: string | moment.Moment,
		format: string,
	) {
		var time: moment.Moment;
		if (typeof input == "string") {
			time = moment(<string>input);
		} else if (input instanceof moment) {
			time = <moment.Moment>input;
		}

		return time.utc().format(format);
	}
}

