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

func clickSubCatalog(wd selenium.WebDriver, xpath string) error {
	err := wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, xpath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		return fmt.Errorf("error finding the subcatalog button %q: %v", xpath, err)
	}

	subCatalogButton, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return fmt.Errorf("error finding the subcatalog button %q: %v", xpath, err)
	}

	err = subCatalogButton.Click()
	if err != nil {
		return fmt.Errorf("error clicking the subcatalog button %q: %v", xpath, err)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Click on subcatalog %q completed successfully\n", xpath)
	return nil
}

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

	// List of social media to click
	subCatalogs := []string{
		"//div[@class='contacts']//p[text()='Telegram']",
		"//div[@class='contacts']//p[text()='Instagram']",
		"//div[@class='contacts']//p[text()='TikTok']",
		"//div[@class='contacts']//p[text()='WhatsApp']",
		"//div[@class='contacts']//p[text()='Facebook']",
		"//div[@class='contacts']//p[text()='YouTube']",
	}

	// Click on each social media button
	for _, xpath := range subCatalogs {
		if err := clickSubCatalog(wd, xpath); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Test completed successfully")
}
