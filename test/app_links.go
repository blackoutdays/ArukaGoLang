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

	// Locate and click the Google Play link
	googlePlayLinkXPath := "//div[@class='download-app__icons']//a[@href='https://play.google.com/store/apps/details?id=app.smartdeal.kz']"
	googlePlayLink, err := wd.FindElement(selenium.ByXPATH, googlePlayLinkXPath)
	if err != nil {
		log.Fatalf("Error finding the Google Play link: %v", err)
	}

	// Scroll the element into view
	_, err = wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{googlePlayLink})
	if err != nil {
		log.Fatalf("Error scrolling the Google Play link into view: %v", err)
	}

	time.Sleep(1 * time.Second) // Wait a bit for the scroll to complete

	// Click the Google Play link using JavaScript to bypass click interception issues
	_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{googlePlayLink})
	if err != nil {
		log.Fatalf("Error clicking on the Google Play link: %v", err)
	}

	// Wait for the new tab to open
	time.Sleep(5 * time.Second)

	// Get all window handles
	windowHandles, err := wd.WindowHandles()
	if err != nil {
		log.Fatalf("Error getting window handles: %v", err)
	}

	// Switch to the new window
	wd.SwitchWindow(windowHandles[len(windowHandles)-1])
	time.Sleep(5 * time.Second)

	// Close the new window
	err = wd.Close()
	if err != nil {
		log.Fatalf("Error closing the new window: %v", err)
	}

	// Switch back to the main window
	wd.SwitchWindow(windowHandles[0])
	time.Sleep(5 * time.Second)

	// Scroll to the bottom of the page again
	_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
	if err != nil {
		log.Fatalf("Error scrolling the page: %v", err)
	}

	// Wait for the scroll to complete
	time.Sleep(5 * time.Second)

	// Locate and click the App Store link
	appStoreLinkXPath := "//div[@class='download-app__icons']//a[@href='https://apps.apple.com/kz/app/smart-deal/id6479311597']"
	appStoreLink, err := wd.FindElement(selenium.ByXPATH, appStoreLinkXPath)
	if err != nil {
		log.Fatalf("Error finding the App Store link: %v", err)
	}

	// Scroll the element into view
	_, err = wd.ExecuteScript("arguments[0].scrollIntoView(true);", []interface{}{appStoreLink})
	if err != nil {
		log.Fatalf("Error scrolling the App Store link into view: %v", err)
	}

	time.Sleep(1 * time.Second) // Wait a bit for the scroll to complete

	// Click the App Store link using JavaScript to bypass click interception issues
	_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{appStoreLink})
	if err != nil {
		log.Fatalf("Error clicking on the App Store link: %v", err)
	}

	// Wait for the new tab to open
	time.Sleep(5 * time.Second)

	// Get all window handles
	windowHandles, err = wd.WindowHandles()
	if err != nil {
		log.Fatalf("Error getting window handles: %v", err)
	}

	// Switch to the new window
	wd.SwitchWindow(windowHandles[len(windowHandles)-1])
	time.Sleep(5 * time.Second)

	// Close the new window
	err = wd.Close()
	if err != nil {
		log.Fatalf("Error closing the new window: %v", err)
	}

	// Switch back to the main window
	wd.SwitchWindow(windowHandles[0])
	time.Sleep(5 * time.Second)

	fmt.Println("Test completed successfully")
}
