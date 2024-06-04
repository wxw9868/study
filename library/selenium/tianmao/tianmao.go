package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tebeka/selenium"
)

var (
	sigs chan os.Signal
)

func main() {
	const (
		// These paths will be different on your system.
		seleniumPath    = "E:\\selenium\\selenium-server-standalone-3.8.1.jar"
		geckoDriverPath = "E:\\selenium\\geckodriver.exe"
		port            = 8080
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
	if err := wd.Get("https://login.taobao.com/"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// Get a reference to the text box containing code.
	name, err := wd.FindElement(selenium.ByID, "fm-login-id")
	if err != nil {
		panic(err)
	}
	// Enter some new code in text box.
	if err = name.SendKeys("18201108862"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	passs, err := wd.FindElement(selenium.ByID, "fm-login-password")
	if err != nil {
		panic(err)
	}
	// Enter some new code in text box.
	if err = passs.SendKeys("1994wei.igi"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, ".password-login")
	if err != nil {
		panic(err)
	}
	if err = btn.Click(); err != nil {
		fmt.Println(err)
	}

	// Navigate to the simple playground interface.
	if err := wd.Get("https://cart.taobao.com/cart.htm?spm=a220o.7406545.a2226mz.12.86cb7eec42fyt8&from=btop"); err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)

	// J_CheckBoxShop 选中购物车中某个店铺的所有商品
	// J_CheckBoxItem 选中购物车中的一个商品
	boxs,err := wd.FindElements(selenium.ByClassName,"J_CheckBoxItem")
	if err != nil {
		panic(err)
	}
	fmt.Println("boxs: ",boxs)
    // https://m.tb.cn/h.45U0RPu?sm=400695


	<-sigs
}

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}
