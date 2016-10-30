var webdriverio = require('webdriverio');
var options = {
	desiredCapabilities: {
		browserName: 'firefox'
	}
};

// var client = webdriverio.remote(options)
// 	.init()
// 	.url('https://google.com')
// 	.setValue("input[type='text']", "Hello World!")
// 	.saveScreenshot("google.png");

var url = "https://my.uakron.edu/psp/portprodg/EMPLOYEE/CS/c/COMMUNITY_ACCESS.CLASS_SEARCH.GBL";
 
// Fetch Page and Switch to Appropriate Frame
var client = webdriverio.remote(options)
	.init()
	.url(url)
	.frame("TargetContent");

// Get All the Subjects
client.execute(function() {
	return [].map.call(
		document.querySelectorAll("select[id^=SSR_CLSRCH_WRK_SUBJECT_] > option"),
		function(f) { return f.value }
	).splice(1);
}).then((response) => {
	for (var subject of response.value) {
		// Fetch Classes for Each Subject
		console.log(`Fetching for ${subject}`);

		// Get Value of Subject Combo Box
		client
			.getValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']")
			.then((v) => console.log(v));

		// Select Subject
		client.selectByValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']", subject)
			.then(() => client.saveScreenshot("./dump.png"));

		// Wait For Processing to Be Done
		client.waitForExist("#processing", 500, true)
			.getValue("select[id^='SSR_CLSRCH_WRK_SUBJECT']")
			.then((v) => console.log(v));

		// Click Search Button
		client.click("#CLASS_SRCH_WRK2_SSR_PB_CLASS_SRCH")
			// Wait For Processing to Be Done
			.waitForExist("#processing", 500, true)

		// client.finish();
		return;
	}
});

// var theFrameId = 1;
// var options = driver.wait(until.ableToSwitchToFrame(theFrameId))
// 	.then(() => {
// 		driver.switchTo().frame(theFrameId);
// 		return driver.findElements(By.css("select[id^=SSR_CLSRCH_WRK_SUBJECT_] > option"));
// 	})
// 	.then((elements) => elements.splice(1).map((e) => e.getAttribute('value')))

// var x = document.querySelectorAll("table#ACE_SSR_CLSRSLT_WRK_GROUPBOX1.PSGROUPBOX > tbody > tr:last-child > td:last-child > div > table.PABACKGROUNDINVISIBLEWBO > tbody > tr > td > table > tbody > tr")
