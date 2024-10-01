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

	// Now find the input fields and the reset button
	priceInputFieldsXPath := "//div[@class='inputs']//input[@class='inputs__list--input']"
	resetButtonXPath := "//button[contains(@class, 'reset-button')]"

	// Wait for the input fields to appear
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		elements, err := wd.FindElements(selenium.ByXPATH, priceInputFieldsXPath)
		return err == nil && len(elements) == 2, nil
	}, 20*time.Second)
	if err != nil {
		fmt.Println("Error finding the price input fields:", err)
		return
	}

	// Find the input fields
	priceInputFields, err := wd.FindElements(selenium.ByXPATH, priceInputFieldsXPath)
	if err != nil {
		fmt.Println("Error finding the price input fields:", err)
		return
	}

	// Clear the input fields directly
	err = priceInputFields[0].Clear()
	if err != nil {
		fmt.Println("Error clearing the first price input field:", err)
		return
	}

	err = priceInputFields[1].Clear()
	if err != nil {
		fmt.Println("Error clearing the second price input field:", err)
		return
	}

	// Ensure both fields are cleared
	time.Sleep(1 * time.Second)

	// Enter values into the input fields
	err = priceInputFields[0].SendKeys("100")
	if err != nil {
		fmt.Println("Error entering value into the first price input field:", err)
		return
	}

	err = priceInputFields[1].SendKeys("50000")
	if err != nil {
		fmt.Println("Error entering value into the second price input field:", err)
		return
	}

	fmt.Println("Entered values into price input fields successfully")

	time.Sleep(5 * time.Second)
	// Wait for the reset button to appear
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, resetButtonXPath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		fmt.Println("Error finding the reset button:", err)
		return
	}

	// Find the reset button
	resetButton, err := wd.FindElement(selenium.ByXPATH, resetButtonXPath)
	if err != nil {
		fmt.Println("Error finding the reset button:", err)
		return
	}

	// Click the reset button
	err = resetButton.Click()
	if err != nil {
		fmt.Println("Error clicking the reset button:", err)
		return
	}

	fmt.Println("Clicked the reset button successfully")

	time.Sleep(5 * time.Second)
	fmt.Println("Test completed successfully")
}
