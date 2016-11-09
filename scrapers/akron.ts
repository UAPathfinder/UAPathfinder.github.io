import "reflect-metadata";

import { By, until, Builder } from 'selenium-webdriver';
import { createConnection } from "typeorm";

import { Course } from './entity/course';
import { Class } from './entity/class';
import { Department } from './entity/department';

import { Time } from './time';

// Get Cookies
const url1 = "https://my.uakron.edu/psp/portprodg/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL";

// Drop Directly Into Frame
const url2 = "https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL" +
	"?PortalActualURL=" + encodeURIComponent("https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL") +
	"&PortalContentURL=" + encodeURIComponent("https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL") +
	"&PortalContentProvider=" + "CS" +
	"&PortalCRefLabel=" + encodeURIComponent("Schedule of Classes") + 
	"&PortalRegistryName=" + "EMPLOYEE" +
	"&PortalServletURI=" + encodeURIComponent("https://my.uakron.edu/psp/portprodg/") +
	"&PortalURI=" + encodeURIComponent("https://my.uakron.edu/psp/portprodg/") +
	"&PortalHostNode=" + "EMPL" +
	"&NoCrumbs=" + "yes";

var db;
const dbPromise = createConnection({
	driver: {
		type: "sqlite",
		storage: "../data/akron.db",
	},
	logging: {
		logQueries: true,
		logFailedQueryError: true,
		logOnlyFailedQueries: true,
		logSchemaCreation: true,
	},
	entities: [
		Course, Class, Department,
	],
});

var driver = new Builder()
	// .usingServer('http://devbook.lan:4444/wd/hub')
	.usingServer('http://drone.lan:4444/wd/hub')
	.forBrowser('firefox')
	.build();

getMetadata();

async function getMetadata() {
	db = await dbPromise;
	await driver.get(url1)
	await driver.get(url2)

	// Initialize Page
	const departmentSelector = await driver.findElement(By.id("SSR_CLSRCH_WRK_SUBJECT_SRCH$2"));
	const departments = (await departmentSelector.findElements(By.css('option'))).length;

	// Disable Timeout
	await driver.executeScript(_ => {
		// Patch Timeout Code
		window.getLastAccessTime = function() {
			return -1;
		};
	});

	const processing = await driver.findElement(By.id('processing'));

	for (let i = 2; i < departments; i++) {
		console.log(`Progress: ${i}/${departments}`);
		await handleDepartment(i, processing);

		// Start New Search
		await driver.findElement(By.id("CLASS_SRCH_WRK2_SSR_PB_NEW_SEARCH"))
			.then(button => button.click());

		// TODO: Handle Empty Department

		await driver.wait(until.elementIsNotVisible(processing));
	}
}

async function handleDepartment(i: number, processing: any) {
	const departmentSelector = await driver.findElement(By.id("SSR_CLSRCH_WRK_SUBJECT_SRCH$2"));
	const departmentOption = await departmentSelector.findElement(By.css(`option:nth-child(${i})`))
	const departmentName = await departmentOption.getText()
		.then(text => text.split(/\s{4}/)[0])
		// Example:
		//     Divorce Mediation    (1800)
		//     ->
		//     Divorce Mediation
	console.log(departmentName);

	const department = await departmentOption.getAttribute('value');

	let deptModel = new Department();
	deptModel.title = departmentName;
	deptModel.identifier = department;

	const deptModelPromise = db.entityManager.persist(deptModel);

	// Select Department
	await driver.executeScript((selector, value) => {
		selector.value = value;
		selector.onchange();
	}, departmentSelector, department);

	// Wait for Processing
	await driver.wait(until.elementIsNotVisible(processing));

	// Press Search Button
	await driver.findElement(By.id('CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH'))
		.then(searchButton => searchButton.click())

	// Wait for Processing
	await driver.wait(until.elementIsNotVisible(processing));

	// Handle "Over 200 elements dialog" See Applied Music Dept. for example.
	await driver.findElements(By.id('#ICSave'))
		.then(elements => {
			let [confirmButton] = elements;
			if (confirmButton)
				return confirmButton.click()
		});

	await driver.wait(until.elementIsNotVisible(processing));

	// Extract Data
	// Get Courses
	var courses = await driver.findElements(
		By.css([
			'#ACE_$ICField$12$$0'.replace(/\$/g, '\\$'),
			'tbody', 'tr:not(:first-child)'
		].join(' > '))
	);

	var sectionPromises = [];
	var coursePromises = [];

	for (let course of courses) {
		// Course Sections
		var sections = await course.findElements(By.css('[id^=trSSR_CLSRCH_MTG1]'));

		// Units
		var unitsPromise;
		if (sections.length == 0) {
			// Can't determine units.
			unitsPromise = Promise.resolve(null);
		} else {
			unitsPromise = sections[0].findElements(By.css('[id^=UA_DERIVED_SRCH_DESCR2]'))
				.then(nodes => {
					let [unitsNode] = nodes;
					if (unitsNode)
						return unitsNode.getText();
				});
		}

		// Course Title
		// Example:
		//     " 7400 123 - Fundamentals of Construction " ->
		//     ["7400 123", "Fundamentals of Construction"]
		const courseTextPromise = course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
			.then(title => title.getText())
			.then(titleText => titleText.trim().replace(/\s+/g, ' ').split(' - '));

		const coursePromise = Promise.all([courseTextPromise, unitsPromise])
			.then(courseInformation => {
				let [[identifier, title], units] = courseInformation;

				let course = new Course();
				course.department = department;
				course.identifier = identifier;
				course.title = title;
				course.units = units;

				// TODO: Description. The massive JSON file might have some
				// information.
				// https://www.uakron.edu/academics_majors/class-search/data/courses.json
				return course;
			});

		coursePromises.push(coursePromise);

		for (let section of sections) {
			// Class Number
			var numberPromise = section.findElement(By.css('[id^=MTG_CLASS_NBR]'))
				.then(numberNode => numberNode.getText());

			// Section Name
			var sectionNamePromise = section.findElement(By.css('[id^=MTG_CLASSNAME]'))
				.then(sectionNode => sectionNode.getText())
				.then(section => section.replace('\n', ' '));

			// Class Time + Date
			// Examples:
			// - Th 1:10PM - 2:50PM
			// - MoWe 1:10PM - 2:25PM
			// - TuTh 5:10PM - 6:00PM
			// - MoWeFr 8:50AM - 12:30PM
			var daytimePromise = section.findElement(By.css('[id^=MTG_DAYTIME]'))
				.then(timeNode => timeNode.getText())
				.then(text => {
					if (text == "TBA")
						return null;

					// Handle Days
					var days = {
						"monday": false,
						"tuesday": false,
						"wednesday": false,
						"thursday": false,
						"friday": false,
						"saturday": false,
						"sunday": false,
					};

					const mapping = {
						"Mo": "monday",
						"Tu": "tuesday",
						"We": "wednesday",
						"Th": "thursday",
						"Fr": "friday",

						// Can't confirm if these last few are right. Haven't seen
						// them in the wild.
						"Sa": "saturday",
						"Su": "sunday",
					};

					var i;
					for (i = 0; i < text.length && text[i] != ' '; i += 2) {
						let rawDay = text.substr(i, 2);
						let rawToNice = mapping[rawDay];
						if (rawToNice) {
							days[rawToNice] = true;
							delete mapping[rawDay];
						} else {
							// Duplicate days or some wired string.
							return null;
						}
					}

					let times = text.substr(i + 1)
						.split(' - ')
						.map(time => new Time(time).toSeconds());

					return [days, times];
				});

			// Room
			var roomPromise = section.findElement(By.css('[id^=MTG_ROOM]'))
				.then(locNode => locNode.getText());

			// Instructor
			var instructorPromise = section.findElement(By.css('[id^=MTG_INSTR]'))
				.then(instructorNode => instructorNode.getText())
				.then(text => text.replace(/\s+/g, ' '));

			var datesPromise = section.findElement(By.css('[id^=MTG_TOPIC]'))
				.then(elem => elem.getText())
				.then(dateText => {
					if (dateText == "TBA")
						return null;

					// Example: 08/29/2016 - 12/11/2016
					return dateText.split(' - ')
						.map(date => {
							let [month, day, year] = date.split('/');
							return (new Date(
								parseInt(year), parseInt(month) + 1, parseInt(day)
							)).valueOf();
						});
				});

			// Enrolled
			var enrolledPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR3]'))
				.then(enrolledNode => enrolledNode.getText());

			var sectionPromise = Promise.all([
				coursePromise, numberPromise, sectionNamePromise, daytimePromise,
				roomPromise, instructorPromise, datesPromise, enrolledPromise,
			]).then(resolved => {
				let [course, sectionNumber, sectionName, daysTimes, room,
					instructor, dates, enrolled] = resolved;

				let klass = new Class();
				klass.identifier = sectionNumber;
				klass.course = course.identifier;
				klass.location = room;
				// TODO: professor -> instructor?
				klass.professor = instructor;
				klass.registered = enrolled;

				if (dates) {
					[klass.start_date, klass.end_date] = dates;
				}

				if (daysTimes) {
					let [days, times] = daysTimes;
					// TODO: wtf is this case?

					if (times) {
						[klass.start_time, klass.end_time] = times
					}

					// Copy Days (sunday-saturday)
					Object.assign(klass, days);
				}

				return klass;
			});

			sectionPromises.push(sectionPromise);

			// TODO: Textbook
			// When the textbook link is clicked, javascript submits a form
			// which responds with some html. The html is inlined and script
			// tags in the response executed. The code which opens the window
			// for textbook information is somewhere in these script tags. a

			// TODO: After 103 minutes of running, the page refreshes
			// spontaneously and the script breaks. Element references fail to
			// resolve and the script crashes. I believe the scraper got kicked
			// due to in-activity on the page (15 minutes?). Section after
			// writing.

			// TODO: Do Something About Database Conflicts
		}
	}

	await deptModelPromise;
	var courses = await Promise.all(coursePromises);
	console.log(courses);
	if (courses.length > 0)
		db.entityManager.persist(courses);

	var classes = await Promise.all(sectionPromises)
	console.log(classes);
	if (classes.length > 0)
		db.entityManager.persist(classes);
}
