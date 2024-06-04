package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tebeka/selenium"
	//"github.com/go-vgo/robotgo"
)

var (
	sigs chan os.Signal
)

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

const (
	// These paths will be different on your system.
	seleniumPath    = "E:\\selenium\\selenium-server-standalone-3.8.1.jar"
	geckoDriverPath = "E:\\selenium\\geckodriver.exe"
	port            = 8080
)

func main() {
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
	caps := selenium.Capabilities{
		"browserName":  "firefox",
		"Content-Type": "application/json",
	}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("https://passport.bilibili.com/login"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// Get a reference to the text box containing code.
	name, err := wd.FindElement(selenium.ByCSSSelector, "#login-username")
	if err != nil {
		panic(err)
	}
	// Enter some new code in text box.
	if err = name.SendKeys("18201108862"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	passs, err := wd.FindElement(selenium.ByCSSSelector, "#login-passwd")
	if err != nil {
		panic(err)
	}
	// Enter some new code in text box.
	if err = passs.SendKeys("986845663wxw"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, ".btn-login")
	if err != nil {
		panic(err)
	}
	btn.Click()

	time.Sleep(2 * time.Second)

	slider, err := wd.FindElement(selenium.ByCSSSelector, ".geetest_slider_button")
	if err != nil {
		panic(err)
	}
	if err = slider.MoveTo(100, 0); err != nil {
		fmt.Println(err)
	}

	/*
		// Navigate to the simple playground interface.
		if err := wd.Get("https://www.baidu.com/"); err != nil {
			panic(err)
		}
		wd.SetImplicitWaitTimeout(4)
		input,err:=wd.FindElement(selenium.ByCSSSelector,"#kw")
		if err!=nil{
			panic(err)
		}
		input.SendKeys("亿动华源")
		search,err:=wd.FindElement(selenium.ByCSSSelector,"#su")
		if err!=nil{
			panic(err)
		}
		search.Click()
	*/

	<-sigs
}
