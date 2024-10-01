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

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Find the sorting option with value "6"
	sortOption6XPath := "//option[@class='select__select--option' and @value='6']"
	sortOption6, err := wd.FindElement(selenium.ByXPATH, sortOption6XPath)
	if err != nil {
		log.Fatalf("Error finding the sorting option with value '6': %v", err)
	}

	// Click on the sorting option "6"
	err = sortOption6.Click()
	if err != nil {
		log.Fatalf("Error clicking on the sorting option '6': %v", err)
	}

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Print the sorting option text
	text6, err := sortOption6.Text()
	if err != nil {
		log.Fatalf("Error getting the text of the sorting option '6': %v", err)
	}
	fmt.Printf("Sorting option selected: %s\n", text6)

	fmt.Println("Sort by 6 items per page completed successfully")

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Find the sorting option with value "9"
	sortOption9XPath := "//option[@class='select__select--option' and @value='9']"
	sortOption9, err := wd.FindElement(selenium.ByXPATH, sortOption9XPath)
	if err != nil {
		log.Fatalf("Error finding the sorting option with value '9': %v", err)
	}

	// Click on the sorting option "9"
	err = sortOption9.Click()
	if err != nil {
		log.Fatalf("Error clicking on the sorting option '9': %v", err)
	}

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Print the sorting option text
	text9, err := sortOption9.Text()
	if err != nil {
		log.Fatalf("Error getting the text of the sorting option '9': %v", err)
	}
	fmt.Printf("Sorting option selected: %s\n", text9)

	fmt.Println("Sort by 9 items per page completed successfully")

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Find the sorting option with value "9"
	sortOption12XPath := "//option[@class='select__select--option' and @value='12']"
	sortOption12, err := wd.FindElement(selenium.ByXPATH, sortOption12XPath)
	if err != nil {
		log.Fatalf("Error finding the sorting option with value '12': %v", err)
	}

	// Click on the sorting option "9"
	err = sortOption12.Click()
	if err != nil {
		log.Fatalf("Error clicking on the sorting option '12': %v", err)
	}

	// Wait for the sorting to be applied
	time.Sleep(5 * time.Second)

	// Print the sorting option text
	text12, err := sortOption12.Text()
	if err != nil {
		log.Fatalf("Error getting the text of the sorting option '12': %v", err)
	}
	fmt.Printf("Sorting option selected: %s\n", text12)

	fmt.Println("Sort by 12 items per page completed successfully")
	time.Sleep(5 * time.Second)
	fmt.Println("successful")
}
