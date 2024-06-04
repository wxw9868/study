package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"io/ioutil"
)

func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "E:/Users/Administrator/Desktop/python/selenium-server-standalone-3.8.1.jar"
		geckoDriverPath = "E:/Users/Administrator/Desktop/python/geckodriver.exe"
		port            = 8081
	)
	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(ioutil.Discard),       // Output debug information to STDERR.
	}
	selenium.SetDebug(false)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://www.baidu.com/"); err != nil {
		panic(err)
	}
	wd.SetImplicitWaitTimeout(4)
	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#kw")
	if err != nil {
		panic(err)
	}
	//// Remove the boilerplate code already in the text box.
	//if err := elem.Clear(); err != nil {
	//	panic(err)
	//}

	// Enter some new code in text box.
	err = elem.SendKeys("中国移动")
	if err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#su")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	//// Wait for the program to finish running and get the output.
	//outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#content_left")
	//if err != nil {
	//	panic(err)
	//}
	//
	//var output string
	//output, err = outputDiv.Text()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("output: ",output)

	// Program exited.
}
