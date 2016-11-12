Queries
-------

A collection of useful SQL queries.

## Course with Highest Credit

    SELECT * FROM "courses"
    ORDER BY units DESC
    LIMIT 1

## Number of Credits offered by Each Department

    SELECT SUM(courses.units) AS units, departments.title FROM "classes"
    JOIN "courses" ON classes.course = courses.identifier
    JOIN "departments" on courses.department = departments.identifier
    GROUP BY courses.department
    ORDER BY units DESC

## Find All Actual Classes for a Course

    SELECT * FROM "classes"
    JOIN "courses" ON classes.course = courses.identifier
    WHERE department = "7520" AND classes.start_time IS NOT NULL

## Find Classes on Saturday or Sunday

    SELECT * FROM "classes"
    WHERE sunday IS 1 OR saturday IS 1

## Find Largest Class

    SELECT * FROM "classes"
    ORDER BY registered DESC

## Find All Professors

    SELECT DISTINCT professor FROM "classes"

## Which Professors Teach the Most Classes

    SELECT professor, COUNT(professor) AS classes FROM "classes"
    GROUP BY professor
    ORDER BY classes DESC

## What classes do these professors teach

    SELECT title, COUNT(title) AS sections FROM "classes"
    JOIN "courses" ON classes.course = courses.identifier
    WHERE professor IS "Brian Davis, Charlotte LaBelle"
    GROUP BY title
    ORDER BY sections DESC

