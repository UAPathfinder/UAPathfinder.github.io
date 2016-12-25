package scheduling

import (
	"time"
)

// A singular class. A class is something you could attend. There are often
// many classes for each course.
// Example: Data Structures starting at 3 PM in room 301 with professor x
type Class struct {
	// Class identifier string. Human readable.
	Identifier string

	// Course identifier string. Human readable.
	Course string

	// TODO: Omit Information About Validity In JSON
	Capacity   int64 `gorm:"column:capicity"`
	Registered int64
	Professor  string
	Location   string

	Times
}

func (class *Class) Events(props *EventProperties) []ClassEvent {
	events := []ClassEvent{}

	// TODO: This isn't elegant. Find a way to dump this data into a map, use
	// raw queries, or a different db schema.
	if class.Times.Sunday {
		events = append(events, ClassEvent{
			Day:             time.Sunday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Monday {
		events = append(events, ClassEvent{
			Day:             time.Monday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Tuesday {
		events = append(events, ClassEvent{
			Day:             time.Tuesday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Wednesday {
		events = append(events, ClassEvent{
			Day:             time.Wednesday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Thursday {
		events = append(events, ClassEvent{
			Day:             time.Thursday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Friday {
		events = append(events, ClassEvent{
			Day:             time.Friday,
			Class:           class,
			EventProperties: props,
		})
	}

	if class.Times.Saturday {
		events = append(events, ClassEvent{
			Day:             time.Saturday,
			Class:           class,
			EventProperties: props,
		})
	}

	return events
}

// An Event created from a Class.
type ClassEvent struct {
	Day time.Weekday
	*Class
	*EventProperties
}

func (evt ClassEvent) Properties() EventProperties {
	return *evt.EventProperties
}

func (evt *ClassEvent) handleTime(rawTime int64) *time.Time {
	if rawTime == 0 {
		// TODO: This will cause a panic if the database doesn't have time
		// information.
		return nil
	}

	t := time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)

	// Set Weekday
	diff := int(evt.Day) - int(t.Weekday())
	t.AddDate(0, 0, diff)

	// Add Raw Time
	t.Add(time.Duration(rawTime) * time.Second)
	return &t
}

func (evt ClassEvent) StartTime() time.Time {
	return *evt.handleTime(evt.Class.Times.RawStartTime)
}

func (evt ClassEvent) EndTime() time.Time {
	return *evt.handleTime(evt.Class.Times.RawEndTime)
}

type Times struct {
	Sunday    bool
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool

	RawStartTime int64 `gorm:"column:start_time"`
	RawEndTime   int64 `gorm:"column:end_time"`
}

// A group of classes which share some common characteristics. For example,
// 3960:401 Data Structures
type Course struct {
	Department string
	Identifier string

	Title       string
	Description string
	Units       int64 // TODO: This can be a decimal.
}

type Department struct {
	Identifier  string
	Title       string
	Description string
}

// Input to FindSchedules
type Criteria struct {
	MinimizeBreaks bool
	BreakWeight    int
}
