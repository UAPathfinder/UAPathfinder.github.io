var webdriver = require('selenium-webdriver'),
    By = webdriver.By,
    until = webdriver.until;

var driver = new webdriver.Builder()
	.usingServer('http://localhost:4444/wd/hub')
    .forBrowser('firefox')
    .build();

// Fetch Page and Switch to Appropriate Frame
var url = "https://my.uakron.edu/psp/portprodg/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL";
driver.get(url);

var theFrameId = 1;
driver.wait(until.ableToSwitchToFrame(theFrameId))
	.then(() => {
		driver.switchTo().frame(theFrameId);
		return driver.findElements(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_] > option"));
	})
	.then((elements) => elements.splice(1).forEach((element) => {
		driver.findElement(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_]"))
			.then((combobox) => combobox.click())
			.then(() => element.click());
		})
	)


// driver.findElements(By.css("frame")).then((tables) => console.log(tables));
	// .then((subjects) => subjects.map((element) => element.getAttribute("value")))
	// .then((stuff) => console.log(stuff))
	// .then(() => driver.quit());
 
// Get All the Subjects
// client.execute(function() {
// 	return [].map.call(
// 		document.querySelectorAll("select[id^=SSR_CLSRCH_WRK_SUBJECT_] > option"),
// 		function(f) { return f.value }
// 	).splice(1);
// }).then((response) => {
// 	for (var subject of response.value) {
// 		// Fetch Classes for Each Subject
// 		console.log(`Fetching for ${subject}`);
// 
// 		// Get Value of Subject Combo Box
// 		client
// 			.getValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']")
// 			.then((v) => console.log(v));
// 
// 		// Select Subject
// 		client.selectByValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']", subject)
// 			.then(() => client.saveScreenshot("./dump.png"));
// 
// 		// Wait For Processing to Be Done
// 		client.waitForExist("#processing", 500, true)
// 			.getValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']")
// 			.then((v) => console.log(v));
// 
// 		// Click Search Button
// 		client.click("#CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH")
// 			// Wait For Processing to Be Done
// 			.waitForExist("#processing", 500, true)
// 
// 		// client.finish();
// 		return;
// 	}
// });
