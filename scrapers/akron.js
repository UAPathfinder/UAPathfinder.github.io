var webdriver = require('selenium-webdriver'),
    By = webdriver.By,
    until = webdriver.until;

var driver = new webdriver.Builder()
	.usingServer('http://localhost:4444/wd/hub')
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
				console.log(values[1]);
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
				course.findElement(By.css('.PAGROUPBOXLABELLEVEL1'))
					.then(title => title.getText())
					.then(titleText => console.log(titleText));

					// TODO: Extract Sections
					// TODO: EXtract Data
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

