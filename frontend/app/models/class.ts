// See scheduling/classes.go
export interface Class {
	Identifier: string;
	Course: string;

    Sunday: boolean;
	Monday: boolean;
	Tuesday: boolean;
	Wednesday: boolean;
	Thursday: boolean;
	Friday: boolean;
	Saturday: boolean;

	RawStartTime: number;
	RawEndTime: number;

    Capacity: number;
    Registered: number;
	Professor: string;
	Location: string;
}

