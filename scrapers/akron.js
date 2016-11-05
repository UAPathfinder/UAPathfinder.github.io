var webdriver = require('selenium-webdriver'),
    By = webdriver.By,
    until = webdriver.until;

var driver = new webdriver.Builder()
	.usingServer('http://192.168.1.217:4444/wd/hub')
    .forBrowser('firefox')
    .build();

// Get Cookies
var url1 = "https://my.uakron.edu/psp/portprodg/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL";

// Drop Directly Into Frame
var url2 = "https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL" +
	"?PortalActualURL=" + encodeURIComponent("https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL") +
	"&PortalContentURL=" + encodeURIComponent("https://campusss.uakron.edu/psc/csprodss/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL") +
	"&PortalContentProvider=" + "CS" +
	"&PortalCRefLabel=" + encodeURIComponent("Schedule of Classes") + 
	"&PortalRegistryName=" + "EMPLOYEE" +
	"&PortalServletURI=" + encodeURIComponent("https://my.uakron.edu/psp/portprodg/") +
	"&PortalURI=" + encodeURIComponent("https://my.uakron.edu/psp/portprodg/") +
	"&PortalHostNode=" + "EMPL" +
	"&NoCrumbs=" + "yes";

driver.get(url1)
	.then(() => driver.get(url2))

var departmentSelector = driver.findElement(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_]"));

var options = driver.findElements(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_] > option"))
 	.then(elements => elements.splice(160,1).map((e) => e.getAttribute('value')))
	.then(departments => departments.forEach((department) => {
		Promise.all([departmentSelector, department])
			.then(values => {
				console.log(`Department: ${values[1]}`);
				return driver.executeScript(
					(selector, value) => {
						selector.value = value;
						selector.onchange();
					},
					values[0], values[1]
				)
			})

			// Wait for Processing
			.then(() => driver.findElement(By.id('processing')))
			.then(processing => driver.wait(until.elementIsNotVisible(processing)))

			// Press Search Button
			.then(() => driver.findElement(By.id('CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH')))
			.then(searchButton => searchButton.click())

			// Wait for Processing
			.then(() => driver.findElement(By.id('processing')))
			.then(processing => driver.wait(until.elementIsNotVisible(processing)))

			// Extract Data
			// Get Courses
			.then(() => driver.findElements(By.css(coursesSelector)))
			.then(courses => courses.forEach(course => {
				// Course Title
				course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
					.then(title => title.getText())
					.then(titleText => console.log(titleText.trim()));

				// Course Sections
				course.findElements(By.css(sectionsSelector))
					.then(sections => sections.forEach(section => {
						// Class Number
						section.findElement(By.css('a[id^=MTG_CLASS_NBR]'))
							.then(numberNode => numberNode.getText())
							.then(numberText => console.log(`Class Number: ${numberText}`));

						// Section Name
						section.findElement(By.css('a[id^=MTG_CLASSNAME]'))
							.then(sectionNode => sectionNode.getText())
							.then(section => console.log(`Section: ${section.replace('\n', ' ')}`))

						// Class Time + Date
						// Examples:
						// - Th 1:10PM - 2:50PM
						// - MoWe 1:10PM - 2:25PM
						// - TuTh 5:10PM - 6:00PM
						// - MoWeFr 8:50AM - 12:30PM
						section.findElement(By.css('[id^=MTG_DAYTIME]'))
							.then(timeNode => timeNode.getText())
							.then(timeText => console.log(`Times: ${timeText}`));

						// Room
						section.findElement(By.css('[id^=MTG_ROOM]'))
							.then(locNode => locNode.getText())
							.then(location => console.log(`Room: ${location}`));

						// Instructor
						section.findElement(By.css('[id^=MTG_INSTR]'))
							.then(instructorNode => instructorNode.getText())
							.then(instructor => console.log(`Instructor: ${instructor}`));

						// Units
						section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR2'))
							.then(unitsNode => unitsNode.getText())
							.then(units => console.log(`Units: ${units}`));

						// Enrolled
						section.findElement(By.css('[id^=UA_DERIVED_SRCH_DESCR3'))
							.then(enrolledNode => enrolledNode.getText())
							.then(enrolled => console.log(`Enrolled: ${enrolled}`))
							.then(_ => console.log("\n"));

						// TODO: Textbook
						// When the textbook link is clicked, javascript submits
						// a form which responds with some html. The html is
						// inlined and script tags in the response executed. The
						// code which opens the window for textbook information
						// is somewhere in these script tags.

						// TODO: Status
						// Not sure what this means because even last semesters
						// classes are green. They can't be registered for.
						// Maybe this is for prereqs?
					}));

					// TODO: Commit to Rethink
					// TODO: Callback hell might have been better.
			}))
			.catch(err => console.error(err));
		})
	);

var coursesSelector = [
	"table#ACE_SSR_CLSRSLT_WRK_GROUPBOX1.PSGROUPBOX", "tbody", "tr:last-child",
	"td:last-child", "div", "table.PABACKGROUNDINVISIBLEWBO", "tbody", "tr",
	"td", "table", "tbody", "tr:not(:first-child)"
].join(' > ');

var sectionsSelector = [
	"[id^=ACE_SSR_CLSRSLT_WRK_GROUPBOX2]", "tbody", "tr:nth-child(2)",
	"td:last-child", "div", "table", "tbody", "tr", "td", "table", "tbody",
	"tr:nth-child(even)"
].join(' > ');

