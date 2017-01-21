-- A SQL database for a school's class information.

CREATE TABLE departments (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,

	-- A short identifier
	identifier  TEXT NOCASE,

	-- A long identifier
	title       TEXT,

	description TEXT,

	UNIQUE(identifier)
);

-- Courses are a type of class. An example is Art-1040: An Introduction
-- to Modern Art.
CREATE TABLE courses (
	id          INTEGER PRIMARY KEY AUTOINCREMENT,

	-- The department this course is part of. For example Art, History or
	-- Science.
	department  TEXT NOCASE,

	-- A human readable identifier for this course.
	identifier  TEXT NOCASE,

	-- A human readable title describing this course.
	title       TEXT,

	-- A human readable description.
	description TEXT,

	-- The units this class counts for. This is used to calculate cost and
	-- for graduation requirements.
	units       INTEGER,

	FOREIGN KEY(department) REFERENCES departments(identifier),
	UNIQUE(
		department, identifier
	)
);

CREATE TABLE classes (
	id INTEGER PRIMARY KEY AUTOINCREMENT,

	-- A unique, human readable string used to reference classes.
	identifier TEXT NOCASE,

	-- The ID of the course this class corresponds to.
	course TEXT NOCASE,

	-- Maximum number of students that can attend a class.
	capicity INTEGER,
	
	-- Number of students currently registered for this class.
	registered INTEGER,

	-- The name of the professor.
	professor TEXT,

	-- The start time of the class measured as the number of seconds since 12 AM
	-- (the one when it's dark out).
	start_time INTEGER,

	-- The end time of the class. Measured the same way as `start_time`.
	end_time INTEGER,

	-- Days of the week the class is held on.
	sunday BOOLEAN,
	monday BOOLEAN,
	tuesday BOOLEAN,
	wednesday BOOLEAN,
	thursday BOOLEAN,
	friday BOOLEAN,
	saturday BOOLEAN,

	-- The starting day of this class.
	start_date INTEGER,

	-- The ending day of this class.
	end_date INTEGER,

	-- The location this class is held.
	location TEXT,

	FOREIGN KEY(course) REFERENCES courses(identifier)
);

-- TODO: Generalize the units structure. Some courses have a range of credit
-- depending.

/*
-- A mapping of courses to their dependencies.
CREATE TABLE course_dependencies (
	id INTEGER PRIMARY KEY AUTOINCREMENT,

	-- The course which we are specifying dependencies for.
	course INTEGER,
	FOREIGN KEY(course) REFERENCES courses(id),

	-- A course which is a dependency to `course`.
	dependency INTEGER
	FOREIGN KEY(dependency) REFERENCES courses(id),

	UNIQUE(
		course, dependency
	)
)
*/

