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

	// Find the 'Favourite' button
	favouriteBtn, err := wd.FindElement(selenium.ByXPATH, "//div[@class='item-card__actions--item']//span[@class='icon']")
	if err != nil {
		log.Fatalf("Error finding the 'Add to Favourite': %v", err)
	}

	// Click the 'Favourite' button
	if err := favouriteBtn.Click(); err != nil {
		log.Fatalf("Error clicking on the 'Add to favourite' button: %v", err)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Product successfully added to favourite!")

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
		// If "У вас пока нет товаров в избранном" is found, there are no favorite items
		fmt.Println("No items in favorites.")
	} else {
		// If favorite items are present, click on the first item card
		fmt.Println("Items found in favorites. Clicking on the first item.")
		firstItemXPath := "//div[@class='items-list__items']//div[@class='item-card'][1]//button[@class='item-card__button']"
		firstItem, err := wd.FindElement(selenium.ByXPATH, firstItemXPath)
		if err != nil {
			log.Fatalf("Error finding the first item: %v", err)
		}

		// Click on the first item card
		if err := firstItem.Click(); err != nil {
			log.Fatalf("Error clicking on the first item: %v", err)
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Test completed successfully")
}
