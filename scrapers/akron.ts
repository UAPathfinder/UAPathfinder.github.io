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
	await driver.get(url1)
	await driver.get(url2)

	const departmentSelector = await driver.findElement(By.id("SSR_CLSRCH_WRK_SUBJECT_SRCH$2"));
	const departments = (await departmentSelector.findElements(By.css('option'))).length;
	db = await dbPromise;

	for (let i = 2; i < departments; i++) {
		console.log(`Progress: ${i}/${departments}`);
		await handleDepartment(i);
		await driver.get(url2)
	}
}

async function handleDepartment(i: number) {
	const departmentSelector = await driver.findElement(By.id("SSR_CLSRCH_WRK_SUBJECT_SRCH$2"));
	const departmentOption = await departmentSelector.findElement(By.css(`option:nth-child(${i})`))
	const departmentName = await departmentOption.getText()
		.then(text => text.split(/\s{4}/)[0])
		// Example:
		//     Divorce Mediation    (1800)
		//     ->
		//     Divorce Mediation

	const department = await departmentOption.getAttribute('value');

	console.log(departmentName);

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
	const processing = await driver.findElement(By.id('processing'));
	await driver.wait(until.elementIsNotVisible(processing));

	// Press Search Button
	await driver.findElement(By.id('CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH'))
		.then(searchButton => searchButton.click())

	// Wait for Processing
	await driver.wait(until.elementIsNotVisible(processing));

	// Handle "Over 200 elements dialog" See Applied Music Dept. for example.
	await driver.findElements(By.id('#ICSave'))
		.then(elements => {
			let confirmButton = elements[0];
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

	for (let course of courses) {
		// Course Title
		// Example:
		//     " 7400 123 - Fundamentals of Construction " ->
		//     ["7400 123", "Fundamentals of Construction"]

		// TODO: Batch these.
		const coursePromise = course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
			.then(title => title.getText())
			.then(titleText => titleText.trim().replace(/\s+/g, ' ').split(' - '))
			// Persist
			.then(async information => {
				await deptModelPromise;
				let [identifier, title] = information;
				let course = new Course();
				course.department = department;
				course.identifier = identifier;
				course.title = title;
				// TODO: Description
				// TODO: Units

				db.entityManager.persist(course);
				return information;
			});

		// Course Sections
		var sections = await course.findElements(By.css('[id^=trSSR_CLSRCH_MTG1]'));

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

			// TODO: Possible newlines splitting instructors.
			// Instructor
			var instructorPromise = section.findElement(By.css('[id^=MTG_INSTR]'))
				.then(instructorNode => instructorNode.getText());

			// Units
			var unitsPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR2]'))
				.then(unitsNode => unitsNode.getText());

			// Enrolled
			var enrolledPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR3]'))
				.then(enrolledNode => enrolledNode.getText());

			var sectionPromise = Promise.all([
				coursePromise, numberPromise, sectionNamePromise, daytimePromise,
				roomPromise, instructorPromise, unitsPromise, enrolledPromise
			]).then(resolved => {
				let [[courseIdentifier, courseTitle], sectionNumber, sectionName,
					daysTimes, room, instructor, units, enrolled] = resolved;

				// TODO: Move units to course.

				let klass = new Class();
				klass.identifier = sectionNumber;
				klass.course = courseIdentifier;
				klass.location = room;
				// TODO: professor -> instructor?
				klass.professor = instructor;
				klass.registered = enrolled;

				if (daysTimes) {
					let [days, times] = daysTimes;
					// TODO: wtf is this case?

					if (times) {
						[klass.start_time, klass.end_time] = times
					}

					// Copy Days (sunday-saturday)
					Object.assign(klass, days);
				}

				// TODO: start_date, end_date
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

	var classes = await Promise.all(sectionPromises)
	if (classes.length != 0)
		db.entityManager.persist(classes);

	console.log(classes);
}
