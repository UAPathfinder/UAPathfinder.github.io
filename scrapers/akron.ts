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
	.usingServer('http://192.168.1.217:4444/wd/hub')
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

	var department = await departmentSelector.findElement(By.css(`option:nth-child(${i})`))
		.then(option => option.getAttribute('value'));

	// Select Department
	await driver.executeScript((selector, value) => {
		selector.value = value;
		selector.onchange();
	}, departmentSelector, department);

	// TODO: Possibly store #processing
	// Wait for Processing
	await driver.findElement(By.id('processing'))
		.then(processing => driver.wait(until.elementIsNotVisible(processing)));

	// Press Search Button
	await driver.findElement(By.id('CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH'))
		.then(searchButton => searchButton.click())

	// Wait for Processing
	await driver.findElement(By.id('processing'))
		.then(processing => driver.wait(until.elementIsNotVisible(processing)))

	// TODO: Handle "Over 200 elements dialog" See Applied Music Dept for
	// example.

	// Extract Data
	// Get Courses
	var courses = await driver.findElements(By.css(coursesSelector))

	// TODO: Run this concurrently instead of sequentually.
	for (let course of courses) {
		// Course Title
		course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
			.then(title => title.getText())
			.then(titleText => console.log(titleText.trim()));

		// Course Sections
		var sections = await course.findElements(By.css(sectionsSelector));
		for (let section of sections) {
			// Class Number
			var numberPromise = section.findElement(By.css('a[id^=MTG_CLASS_NBR]'))
				.then(numberNode => numberNode.getText())
				.then(numberText => console.log(`Class Number: ${numberText}`));

			// Section Name
			var sectionNamePromise = section.findElement(By.css('a[id^=MTG_CLASSNAME]'))
				.then(sectionNode => sectionNode.getText())
				.then(section => console.log(`Section: ${section.replace('\n', ' ')}`))

			// Class Time + Date
			// Examples:
			// - Th 1:10PM - 2:50PM
			// - MoWe 1:10PM - 2:25PM
			// - TuTh 5:10PM - 6:00PM
			// - MoWeFr 8:50AM - 12:30PM
			//
			// TODO: Parse
			var daytimePromise = section.findElement(By.css('[id^=MTG_DAYTIME]'))
				.then(timeNode => timeNode.getText())
				.then(timeText => console.log(`Times: ${timeText}`));

			// Room
			var roomPromise = section.findElement(By.css('[id^=MTG_ROOM]'))
				.then(locNode => locNode.getText())
				.then(location => console.log(`Room: ${location}`));

			// Instructor
			var instructorPromise = section.findElement(By.css('[id^=MTG_INSTR]'))
				.then(instructorNode => instructorNode.getText())
				.then(instructor => console.log(`Instructor: ${instructor}`));

			// Units
			var unitsPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR2'))
				.then(unitsNode => unitsNode.getText())
				.then(units => console.log(`Units: ${units}`));

			// Enrolled
			var enrolledPromise = section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR3'))
				.then(enrolledNode => enrolledNode.getText())
				.then(enrolled => console.log(`Enrolled: ${enrolled}`))
				.then(_ => console.log("\n"));

			await Promise.all([numberPromise, sectionNamePromise, daytimePromise, roomPromise, instructorPromise, unitsPromise, enrolledPromise]);

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
		}
	}
}
