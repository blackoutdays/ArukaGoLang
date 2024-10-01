// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver_mac_arm64-2/chromedriver" // replace with your driver
	port             = 8989
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

	time.Sleep(5 * time.Second)

	// Find the dropdown element by XPath
	dropdown, err := wd.FindElement(selenium.ByXPATH, "//div[@class='header-top__top--info']//a[@class='header-top__top--b2b']")
	if err != nil {
		log.Fatalf("Error finding the dropdown element: %v", err)
	}

	// Click on the dropdown to expand the options
	if err := dropdown.Click(); err != nil {
		log.Fatalf("Error clicking on the dropdown: %v", err)
	}

	fmt.Println("Clicked on the dropdown")

	// Wait for the page to reload or for changes to take effect
	time.Sleep(5 * time.Second)

	fmt.Println("Test completed successfully")
}
