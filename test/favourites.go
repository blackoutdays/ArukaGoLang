package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func main() {
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

	// Find the 'Favourite' button
	favouriteBtn, err := wd.FindElement(selenium.ByCSSSelector, ".item-card__actions--item")
	if err != nil {
		log.Fatalf("Error finding the 'Add to Favourite': %v", err)
	}

	// Click the 'Favourite' button
	if err := favouriteBtn.Click(); err != nil {
		log.Fatalf("Error clicking on the 'Add to Favourite' button: %v", err)
	}

	time.Sleep(20 * time.Second)
	fmt.Println("Product successfully added to favourite!")

	time.Sleep(10 * time.Second)
	fmt.Println("Test completed successfully")
}
