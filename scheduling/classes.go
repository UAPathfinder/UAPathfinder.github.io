package scheduling

import (
	"database/sql"
	// "time"
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
	Capicity   sql.NullInt64
	Registered sql.NullInt64
	Professor  sql.NullString
	Location   sql.NullString

	Times
}

type Times struct {
	// TODO: How are null values handled without nullable?
	Sunday    bool
	Monday    bool
	Tuesday   bool
	Wednesday bool
	Thursday  bool
	Friday    bool
	Saturday  bool

	RawStartTime sql.NullInt64 `gorm:"column:start_time"`
	RawEndTime   sql.NullInt64 `gorm:"column:end_time"`
}

// A group of classes which share some common characteristics. For example,
// 3960:401 Data Structures
type Course struct {
	Classes []Class

	Title     sql.NullString
	Priority  int
	Manditory bool
}
