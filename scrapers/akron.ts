import { By, until, Builder } from 'selenium-webdriver';

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

const coursesSelector = [
	"table#ACE_SSR_CLSRSLT_WRK_GROUPBOX1.PSGROUPBOX", "tbody", "tr:last-child",
	"td:last-child", "div", "table.PABACKGROUNDINVISIBLEWBO", "tbody", "tr",
	"td", "table", "tbody", "tr:not(:first-child)"
].join(' > ');

const sectionsSelector = [
	"[id^=ACE_SSR_CLSRSLT_WRK_GROUPBOX2]", "tbody", "tr:nth-child(2)",
	"td:last-child", "div", "table", "tbody", "tr", "td", "table", "tbody",
	"tr:nth-child(even)"
].join(' > ');

var driver = new Builder()
	// .usingServer('http://devbook.lan:4444/wd/hub')
	.usingServer('http://drone.lan:4444/wd/hub')
    .forBrowser('firefox')
    .build();

getMetadata();

async function getMetadata() {
	await driver.get(url1)
	await driver.get(url2)

	const departmentSelector = await driver.findElement(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_]"));
	const departments = (await departmentSelector.findElements(By.css('option'))).length;

	for (let i = 2; i < departments; i++) {
		await handleDepartment(i);
		await driver.get(url2)
	}
}

async function handleDepartment(i: number) {
	const departmentSelector = await driver.findElement(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_]"));
	const departmentOption = await departmentSelector.findElement(By.css(`option:nth-child(${i})`))
	const departmentName = await departmentOption.getText();

	console.log(departmentName);

	const department = await departmentOption.getAttribute('value');

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
	var courses = await driver.findElements(By.css(coursesSelector))
	var sectionPromises = [];

	for (let course of courses) {
		// Course Title
		const coursePromise = course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
			.then(title => title.getText())
			.then(titleText => titleText.trim());

		// Course Sections
		var sections = await course.findElements(By.css(sectionsSelector));

		for (let section of sections) {
			// Class Number
			var numberPromise = section.findElement(By.css('a[id^=MTG_CLASS_NBR]'))
				.then(numberNode => numberNode.getText());

			// Section Name
			var sectionNamePromise = section.findElement(By.css('a[id^=MTG_CLASSNAME]'))
				.then(sectionNode => sectionNode.getText())
				.then(section => section.replace('\n', ' '));

			// Class Time + Date
			// Examples:
			// - Th 1:10PM - 2:50PM
			// - MoWe 1:10PM - 2:25PM
			// - TuTh 5:10PM - 6:00PM
			// - MoWeFr 8:50AM - 12:30PM
			//
			// TODO: Parse
			var daytimePromise = section.findElement(By.css('[id^=MTG_DAYTIME]'))
				.then(timeNode => timeNode.getText());

			// Room
			var roomPromise = section.findElement(By.css('[id^=MTG_ROOM]'))
				.then(locNode => locNode.getText());

			// TODO: Possible newlines splitting instructors.
			// Instructor
			var instructorPromise = section.findElement(By.css('[id^=MTG_INSTR]'))
				.then(instructorNode => instructorNode.getText());

			// Units
			var unitsPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR2'))
				.then(unitsNode => unitsNode.getText());

			// Enrolled
			var enrolledPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR3'))
				.then(enrolledNode => enrolledNode.getText());

			var sectionPromise = Promise.all([
				coursePromise, numberPromise, sectionNamePromise, daytimePromise,
				roomPromise, instructorPromise, unitsPromise, enrolledPromise
			]).then(resolved => {
				let [course, sectionNumber, sectionName, daytime, room, instructor,
					units, enrolled] = resolved;

				return {
					id: sectionNumber,
					course: course,
					daytime: daytime, // TODO: Parse
					location: room,
					instructor: instructor,
					units: units,
					enrolled: enrolled,
				};
			})

			sectionPromises.push(sectionPromise);

			// TODO: Meeting Days
			// TODO: Textbook
			// When the textbook link is clicked, javascript submits a form
			// which responds with some html. The html is inlined and script
			// tags in the response executed. The code which opens the window
			// for textbook information is somewhere in these script tags. a

			// TODO: Status
			// Not sure what this means because even last semesters classes are
			// green. They can't be registered for. Maybe this is for prereqs?

			// TODO: Commit to Rethink

			// TODO: After 103 minutes of running, the page refreshes
			// spontaneously and the script breaks. Element references fail to
			// resolve and the script crashes. I believe the scraper got kicked
			// due to in-activity on the page (15 minutes?). Section after
			// writing.
		}
	}

	var sections = await Promise.all(sectionPromises)
	console.log(sections);
}
