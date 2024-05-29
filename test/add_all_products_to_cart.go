// comment success in english and names in ru pls ;<
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

	// Find the icon element by XPath
	sortBy, err := wd.FindElement(selenium.ByXPATH, "//span[@class='icon disable-icon']")
	if err != nil {
		log.Fatalf("Error finding the icon: %v", err)
	}

	// Click on the icon element
	err = sortBy.Click()
	if err != nil {
		log.Fatalf("Error clicking the icon: %v", err)
	}

	// Add a wait condition to ensure the click action is completed
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {

		// Check if the element is still visible after the click
		visible, err := sortBy.IsDisplayed()
		if err != nil {
			return false, err
		}
		// Return true if the element is visible
		return visible, nil
	}, 10*time.Second)
	if err != nil {
		log.Fatalf("Error waiting for the click action to complete: %v", err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Click on icon completed successfully")

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Find the "Select All" checkbox
	selectAllXPath := "//input[@type='checkbox' and @id='select-all']"
	selectAllCheckbox, err := wd.FindElement(selenium.ByXPATH, selectAllXPath)
	if err != nil {
		log.Fatalf("Error finding the Выбрать все' checkbox: %v", err)
	}

	// Click the "Select All" checkbox
	err = selectAllCheckbox.Click()
	if err != nil {
		log.Fatalf("Error clicking the 'Выбрать все' checkbox: %v", err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Clicked 'Select All' checkbox successfully")

	// Find the "Add to Cart" button
	addToCartButtonXPath := "//button[contains(@class, 'actions__button ui-button ui-button--colored actions__button')]"
	addToCartButton, err := wd.FindElement(selenium.ByXPATH, addToCartButtonXPath)
	if err != nil {
		log.Fatalf("Error finding the 'Add to Cart' button: %v", err)
	}

	// Click the "Add to Cart" button
	err = addToCartButton.Click()
	if err != nil {
		log.Fatalf("Error clicking the 'Add to Cart' button: %v", err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Clicked 'Add to Cart' button successfully")

	time.Sleep(15 * time.Second)
	fmt.Println("Test completed successfully")
}
