package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver"
	port             = 8080
)

func main() {
	// Set up ChromeDriver service
	var opts []selenium.ServiceOption
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

	// Find and click the "favorites" icon
	favoritesIconXPath := "//a[@href='/favorites']//span[@class='icon']"
	favoritesIcon, err := wd.FindElement(selenium.ByXPATH, favoritesIconXPath)
	if err != nil {
		log.Fatalf("Error finding the favorites icon: %v", err)
	}
	if err := favoritesIcon.Click(); err != nil {
		log.Fatalf("Error clicking the favorites icon: %v", err)
	}

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Check if the "У вас пока нет товаров в избранном" message is present
	noOrdersXPath := "//p[contains(text(),'У вас пока нет товаров в избранном')]"
	noOrdersElement, err := wd.FindElement(selenium.ByXPATH, noOrdersXPath)

	if err == nil && noOrdersElement != nil {
		// If "У вас пока нет товаров в избранном" is found, click "Начать покупки" button
		fmt.Println("No items in favorites. Clicking 'Начать покупки'.")
		shopButtonXPath := "//p[contains(text(),'Начать покупки')]/ancestor::button"
		shopButton, err := wd.FindElement(selenium.ByXPATH, shopButtonXPath)
		if err != nil {
			log.Fatalf("Error finding 'Начать покупки' button: %v", err)
		}
		_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{shopButton})
		if err != nil {
			log.Fatalf("Error clicking on 'Начать покупки' button: %v", err)
		}
	} else {
		// If items are present, click on the first item
		fmt.Println("Items found in favorites. Clicking on the first item.")
		firstItemXPath := "//div[contains(@class, 'favorites-item')][1]"
		firstItem, err := wd.FindElement(selenium.ByXPATH, firstItemXPath)
		if err != nil {
			log.Fatalf("Error finding the first item: %v", err)
		}
		_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{firstItem})
		if err != nil {
			log.Fatalf("Error clicking on the first item: %v", err)
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Test completed successfully")
}
