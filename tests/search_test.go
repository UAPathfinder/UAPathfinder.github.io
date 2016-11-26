package tests

import (
	"testing"
	"time"

	"github.com/mibzman/CourseCorrect/scheduling"
)

type Event struct {
	Start time.Time
	End   time.Time
	scheduling.EventProperties
}

func (evt Event) StartTime() time.Time {
	return evt.Start
}

func (evt Event) EndTime() time.Time {
	return evt.End
}

func (evt Event) Properties() scheduling.EventProperties {
	return evt.EventProperties
}

func TestDoesHaveConflictDuring(t *testing.T) {
	var cal scheduling.Calendar
	cal.Add(&Event{
		Start: scheduling.MustParseTime("3:00:00"),
		End:   scheduling.MustParseTime("7:00:00"),
	})

	conflicting := Event{
		Start: scheduling.MustParseTime("4:00:00"),
		End:   scheduling.MustParseTime("5:00:00"),
	}

	otherConflict := cal.DoesConflict(conflicting)
	if otherConflict != cal.Events[0] {
		t.Fatal("Returned wrong conflicting event.", conflicting, otherConflict)
	}
}

func TestDoesHaveConflictStartingDuring(t *testing.T) {
	var cal scheduling.Calendar
	cal.Add(&Event{
		Start: scheduling.MustParseTime("3:00:00"),
		End:   scheduling.MustParseTime("7:00:00"),
	})

	conflicting := Event{
		Start: scheduling.MustParseTime("6:00:00"),
		End:   scheduling.MustParseTime("9:00:00"),
	}

	otherConflict := cal.DoesConflict(conflicting)
	if otherConflict != cal.Events[0] {
		t.Fatal("Returned wrong conflicting event.", conflicting, otherConflict)
	}
}

func TestDoesHaveConflictEndingDuring(t *testing.T) {
	var cal scheduling.Calendar
	cal.Add(&Event{
		Start: scheduling.MustParseTime("3:00:00"),
		End:   scheduling.MustParseTime("7:00:00"),
	})

	conflicting := Event{
		Start: scheduling.MustParseTime("1:00:00"),
		End:   scheduling.MustParseTime("4:00:00"),
	}

	otherConflict := cal.DoesConflict(conflicting)
	if otherConflict != cal.Events[0] {
		t.Fatal("Returned wrong conflicting event.", conflicting, otherConflict)
	}
}
