package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // replace with your driver path
	port             = 8080
)

func main() {
	// Set up ChromeDriver service
	var opts []selenium.ServiceOption
	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		log.Fatalf("Error starting the ChromeDriver server: %v", err)
	}
	defer func() {
		if err := service.Stop(); err != nil {
			log.Fatalf("Error stopping the ChromeDriver service: %v", err)
		}
	}()

	// Connect to the WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver: %v", err)
	}
	defer func() {
		if err := wd.Quit(); err != nil {
			log.Fatalf("Error quitting the WebDriver session: %v", err)
		}
	}()

	// Navigate to the website
	if err := wd.Get("https://smartdeal.kz/"); err != nil {
		log.Fatalf("Error opening the page: %v", err)
	}

	// Find the dropdown element by XPath
	dropdown, err := wd.FindElement(selenium.ByXPATH, "//select[@class='header-top__top--locale']")
	if err != nil {
		log.Fatalf("Error finding the dropdown element: %v", err)
	}

	// Click on the dropdown to expand the options
	if err := dropdown.Click(); err != nil {
		log.Fatalf("Error clicking on the dropdown: %v", err)
	}

	// Find the desired option by XPath and click on it
	option, err := wd.FindElement(selenium.ByXPATH, "//option[@value='kk']")
	if err != nil {
		log.Fatalf("Error finding the option element: %v", err)
	}

	// Click on the option to select it
	if err := option.Click(); err != nil {
		log.Fatalf("Error clicking on the option: %v", err)
	}

	// Wait for the page to reload or for changes to take effect
	time.Sleep(5 * time.Second)

	fmt.Println("Language changed successfully")
}
