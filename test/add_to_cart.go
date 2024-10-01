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

	// Find the 'В корзину' button
	addToCartButton, err := wd.FindElement(selenium.ByXPATH, "//button[@class='item-card__button ui-button ui-button--colored item-card__button']//p[@class='item-card__button--text' and text()='В корзину']")
	if err != nil {
		log.Fatalf("Error finding the 'В корзину' button: %v", err)
	}

	// Click the 'Add to Cart' button
	if err := addToCartButton.Click(); err != nil {
		log.Fatalf("Error clicking on the 'В корзину' button: %v", err)
	}

	// Wait to ensure the product is added to the cart
	time.Sleep(10 * time.Second)
	fmt.Println("Product successfully added to cart!")

	fmt.Println("Test completed successfully")
}
