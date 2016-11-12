// TODO: Move into library. This is useful code for the frontend.

export class Time {
	// Time in 24 hour time.
	hour: number = 0;
	minute: number = 0;
	second: number = 0;

	// Time in the hour:minute (AM/PM) format. For example: 9:55AM, 11:35AM
	constructor(time: string) {
		let [hour, minute] = time.split(":")
			.map(text => Number.parseInt(text.substr(0, 2)));

		let meridiem: string = time.substr(-2, 2);

		if (meridiem.toUpperCase() == "PM") {
			hour += 12;
		}

		this.hour = hour;
		this.minute = minute;
	}

	toSeconds(): number {
		return this.hour * 3600 + this.minute * 60 + this.second;
	}
}

