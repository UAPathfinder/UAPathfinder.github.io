Schema Design
=============

Things
------
* Departments:
Contain courses along with a description about what the department teaches.

* Courses:
A course is a type of class in a department. Courses may only be part of one
department. Courses have a title, a description, an identifier and a set of
classes which are teaching this course. For example, "Introduction to Writing" and "Java 1" are examples of courses. 

* Classes:
A class is an instance of a course. Classes have a time, instructor, location,
identifier and capacity. Classes are part of a department. For example:
"Beginning French" with Prof. Wilson at 1 P.M. in Bathford Hall.

Stuff We Need <small>ordered by importance</small>
-------------

* Classes for a Course, Time, Location or Instructor
This query is needed for the algorithm to build schedules and for browsing.

* Courses for a Department
This query will be used for browsing the course catalog.

Possible Databases
------------------

### Document Based (MongoDB, RethinkDB):
#### Departments Collection:
Used for browsing

    [
        {
            id: ..., // department identifier
            title: ..., // department title, displayed in UIs
            description: ...,
            courses: [
                {
                    // embed vital course information
                    title: ...,
                    description: ...,
                    id: ObjectId('course-object-id'),
                },
                // 100s-1000s of entries.
            ]
        },
        ...
    ]

#### Courses Collection:
Course detail view.

    [
        {
            id: ...,
            title: ...,
            description: ...,
            classes: [
                {
                    // embed class
                    professor: ...,
                    times: ...,
                    location: ...,
                    section_id: ...,
                    section_name: ...,
                },
                ...
            ],
        },
        ...
    ]

#### Classes Collection:

    [
        {
            id: ...,
            section_id: ...,
            professor: ...,
            times: ...,
            location: ...,
            section_id: ...,
            section_name: ...,
        },
        ....
    ]

Questions:
* Can the logic used the db?

