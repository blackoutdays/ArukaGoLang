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

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Scroll to the bottom of the page to ensure elements are loaded
	_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
	if err != nil {
		log.Fatalf("Error scrolling the page: %v", err)
	}

	// Wait for the scroll to complete
	time.Sleep(5 * time.Second)

	// Locate the go-top element
	goTopXPath := "//span[@class='icon go-top']"
	goTopElement, err := wd.FindElement(selenium.ByXPATH, goTopXPath)
	if err != nil {
		log.Fatalf("Error finding the go-top element: %v", err)
	}

	// Scroll the element into view
	_, err = wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{goTopElement})
	if err != nil {
		log.Fatalf("Error scrolling the go-top element into view: %v", err)
	}

	// Wait a bit for the scroll to complete
	time.Sleep(1 * time.Second)

	// Click the go-top element using JavaScript to bypass click interception issues
	_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{goTopElement})
	if err != nil {
		log.Fatalf("Error clicking on the go-top element: %v", err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Test completed successfully")
}
