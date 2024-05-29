package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // replace with your driver
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

	// Use explicit wait to find the catalog button
	catalogButtonXPath := "//button[contains(@class, 'header__bottom--catalog')]"
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, catalogButtonXPath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		fmt.Println("Error finding the catalog button:", err)
		return
	}

	catalogButton, err := wd.FindElement(selenium.ByXPATH, catalogButtonXPath)
	if err != nil {
		fmt.Println("Error finding the catalog button:", err)
		return
	}

	err = catalogButton.Click()
	if err != nil {
		fmt.Println("Error clicking the catalog button:", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Click on catalog completed successfully")

	// Continue with finding a subcatalog button
	subCatalogButtonXPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Компьютеры']"
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButtonXPath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		fmt.Println("Error finding the subcatalog button:", err)
		return
	}

	subCatalogButton, err := wd.FindElement(selenium.ByXPATH, subCatalogButtonXPath)
	if err != nil {
		fmt.Println("Error finding the subcatalog button:", err)
		return
	}

	err = subCatalogButton.Click()
	if err != nil {
		fmt.Println("Error clicking the subcatalog button:", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Click on subcatalog completed successfully")

	// Continue with finding the nested subcatalog for "Комплектующие"
	nestedSubCatalogXPath := "//a[contains(@class, 'categories-menu__subcategories--title') and p[text()='Комплектующие']]"
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, nestedSubCatalogXPath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		fmt.Println("Error finding the nested subcatalog link:", err)
		return
	}

	nestedSubCatalogLink, err := wd.FindElement(selenium.ByXPATH, nestedSubCatalogXPath)
	if err != nil {
		fmt.Println("Error finding the nested subcatalog link:", err)
		return
	}

	err = nestedSubCatalogLink.Click()
	if err != nil {
		fmt.Println("Error clicking the nested subcatalog link:", err)
		return
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Click on nested subcatalog completed successfully")

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Find the sorting dropdown
	sortDropdownXPath := "//select[@class='select__select']"
	sortDropdown, err := wd.FindElement(selenium.ByXPATH, sortDropdownXPath)
	if err != nil {
		log.Fatalf("Error finding the sorting dropdown: %v", err)
	}

	// Click on the sorting dropdown to activate it
	err = sortDropdown.Click()
	if err != nil {
		log.Fatalf("Error clicking on the sorting dropdown: %v", err)
	}

	// Find the sorting options
	sortOptionsXPath := "//option[@class='select__select--option']"
	sortOptions, err := wd.FindElements(selenium.ByXPATH, sortOptionsXPath)
	if err != nil {
		log.Fatalf("Error finding the sorting options: %v", err)
	}

	// Iterate through each sorting option and apply the sorting
	for _, option := range sortOptions {
		text, err := option.Text()
		if err != nil {
			log.Fatalf("Error getting the text of the sorting option: %v", err)
		}

		// Check if the option is one of the desired criteria
		if text == "По популярности" || text == "По возрастанию цены" || text == "По убыванию цены" {
			// Click on the current sorting option
			err = option.Click()
			if err != nil {
				log.Fatalf("Error clicking on the sorting option '%s': %v", text, err)
			}

			// Wait for the sorting to be applied
			time.Sleep(5 * time.Second)

			fmt.Printf("Sorting by '%s' completed successfully\n", text)
		}
	}
	time.Sleep(5 * time.Second)

	fmt.Println("Test completed successfully")
}
