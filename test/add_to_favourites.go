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

	//<div data-v-e2ced900="" class="item-card__actions--item"><span data-v-e2ced900="" class="icon"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
	//<path d="M12 8.19431C10 3.49988 3 3.99988 3 9.99991C3 15.9999 12 20.9999 12 20.9999C12 20.9999 21 15.9999 21 9.99991C21 3.99988 14 3.49988 12 8.19431Z" stroke="#6F6F6F" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
	//</svg>
	//</span></div>
	//// Find the 'Favourite' button
	//favouriteBtn, err := wd.FindElement(selenium.ByCSSSelector, ".item-card__actions--item")
	//if err != nil {
	//	log.Fatalf("Error finding the 'Add to Favourite': %v", err)
	//}

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

	time.Sleep(10 * time.Second)
	fmt.Println("Test completed successfully")
}
