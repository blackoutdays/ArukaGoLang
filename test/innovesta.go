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

type FormData struct {
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	RoomType       string
	RoomDimensions string
	Comment        string
	SKU            string
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

	formData := FormData{
		FirstName:      "Ryan",
		LastName:       "Gosling",
		Email:          "test@example.com",
		Phone:          "+1234567890",
		RoomType:       "mall",
		RoomDimensions: "100",
		Comment:        "Sample comment",
		SKU:            "sample-sku",
	}

	// Connect to the WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error connecting to the WebDriver: %v", err)
	}
	defer func() {
		// Delay before quitting to ensure the page is not closed immediately after submission
		time.Sleep(20 * time.Second)

		if err := wd.Quit(); err != nil {
			log.Fatalf("Error quitting the WebDriver session: %v", err)
		}
	}()

	// Navigate to the website
	if err := wd.Get("https://innovesta.kz/"); err != nil {
		log.Fatalf("Error opening the page: %v", err)
	}

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Find the dropdown element by XPath
	dropdown, err := wd.FindElement(selenium.ByXPATH, "//select[@class='language']")
	if err != nil {
		log.Fatalf("Error finding the dropdown element: %v", err)
	}

	// Click on the dropdown to expand the options
	if err := dropdown.Click(); err != nil {
		log.Fatalf("Error clicking on the dropdown: %v", err)
	}

	// Find the desired option by XPath and click on it
	option, err := wd.FindElement(selenium.ByXPATH, "//option[@value='ru']")
	if err != nil {
		log.Fatalf("Error finding the option element: %v", err)
	}

	// Click on the option to select it
	if err := option.Click(); err != nil {
		log.Fatalf("Error clicking on the option: %v", err)
	}

	// Wait for the page to reload or for changes to take effect
	time.Sleep(5 * time.Second)

	fmt.Println("Language changed successfully")

	// Proceed to order page
	if err := wd.Get("https://innovesta.kz/bid"); err != nil {
		log.Fatalf("Error opening the order page: %v", err)
	}

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Fill the order form sequentially
	fields := map[string]string{
		"//input[@name='first_name']":      formData.FirstName,
		"//input[@name='last_name']":       formData.LastName,
		"//input[@name='email']":           formData.Email,
		"//input[@name='phone']":           formData.Phone,
		"//select[@name='room_type']":      formData.RoomType,
		"//input[@name='room_dimensions']": formData.RoomDimensions,
		"//textarea[@name='comment']":      formData.Comment,
		"//input[@name='sku']":             formData.SKU,
	}

	for selector, value := range fields {
		elem, err := wd.FindElement(selenium.ByXPATH, selector)
		if err != nil {
			log.Fatalf("Error finding form element: %v", err)
		}
		if err := elem.SendKeys(value); err != nil {
			log.Fatalf("Error sending keys to form element: %v", err)
		}

		// Introduce a delay between filling each field
		time.Sleep(2 * time.Second)
	}

	// Attach a file
	fileInput, err := wd.FindElement(selenium.ByXPATH, "//input[@type='file']")
	if err != nil {
		log.Fatalf("Error finding element for file upload: %v", err)
	}
	if err := fileInput.SendKeys("/Users/aruka/Downloads/5-mb-example-file.pdf"); err != nil {
		log.Fatalf("Error uploading file: %v", err)
	}

	// Submit the form using JavaScript
	submitButton, err := wd.FindElement(selenium.ByXPATH, "//button[@class='order article-page']//span[text()='Отправить']")
	if err != nil {
		log.Fatalf("Error finding submit button: %v", err)
	}

	// Use JavaScript to click the submit button
	_, err = wd.ExecuteScript("arguments[0].click();", []interface{}{submitButton})
	if err != nil {
		log.Fatalf("Error clicking submit button: %v", err)
	}

	// Wait for the success message or confirmation modal
	successMessageXPath := "//p[contains(text(),'Заявка успешно отправлена!')]"
	timeout := time.After(30 * time.Second)
	ticker := time.Tick(500 * time.Millisecond)
	var successFound bool

	for {
		select {
		case <-timeout:
			log.Fatalf("Timed out waiting for success message")
		case <-ticker:
			_, err := wd.FindElement(selenium.ByXPATH, successMessageXPath)
			if err == nil {
				successFound = true
				break
			}
		}
		if successFound {

		}
	}

	if successFound {
		fmt.Println("Form submitted successfully")
	} else {
		fmt.Println("Form submission failed")
	}

	// Additional delay before closing the browser
	time.Sleep(10 * time.Second)
	fmt.Println("Successfully submitted")
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"time"
//
//	"github.com/tebeka/selenium"
//)
//
//const (
//	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // replace with your driver path
//	port             = 8080
//)
//
//type FormData struct {
//	FirstName      string
//	LastName       string
//	Email          string
//	Phone          string
//	RoomType       string
//	RoomDimensions string
//	Comment        string
//	SKU            string
//}
//
//func main() {
//	// Set up ChromeDriver service
//	var opts []selenium.ServiceOption
//	selenium.SetDebug(false)
//	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
//	if err != nil {
//		log.Fatalf("Error starting the ChromeDriver server: %v", err)
//	}
//	defer func() {
//		if err := service.Stop(); err != nil {
//			log.Fatalf("Error stopping the ChromeDriver service: %v", err)
//		}
//	}()
//	formData := FormData{
//		FirstName:      "Ryan",
//		LastName:       "Gosling",
//		Email:          "test@example.com",
//		Phone:          "+1234567890",
//		RoomType:       "mall",
//		RoomDimensions: "100",
//		Comment:        "Sample comment",
//		SKU:            "sample-sku",
//	}
//	// Connect to the WebDriver
//	caps := selenium.Capabilities{"browserName": "chrome"}
//	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
//	if err != nil {
//		log.Fatalf("Error connecting to the WebDriver: %v", err)
//	}
//	defer func() {
//		// Delay before quitting to ensure the page is not closed immediately after submission
//		time.Sleep(10 * time.Second)
//
//		if err := wd.Quit(); err != nil {
//			log.Fatalf("Error quitting the WebDriver session: %v", err)
//		}
//	}()
//
//	// Navigate to the website
//	if err := wd.Get("https://innovesta.kz/"); err != nil {
//		log.Fatalf("Error opening the page: %v", err)
//	}
//
//	// Wait for the page to load
//	time.Sleep(5 * time.Second)
//
//	// Find the dropdown element by XPath
//	dropdown, err := wd.FindElement(selenium.ByXPATH, "//select[@class='language']")
//	if err != nil {
//		log.Fatalf("Error finding the dropdown element: %v", err)
//	}
//
//	// Click on the dropdown to expand the options
//	if err := dropdown.Click(); err != nil {
//		log.Fatalf("Error clicking on the dropdown: %v", err)
//	}
//
//	// Find the desired option by XPath and click on it
//	option, err := wd.FindElement(selenium.ByXPATH, "//option[@value='ru']")
//	if err != nil {
//		log.Fatalf("Error finding the option element: %v", err)
//	}
//
//	// Click on the option to select it
//	if err := option.Click(); err != nil {
//		log.Fatalf("Error clicking on the option: %v", err)
//	}
//
//	// Wait for the page to reload or for changes to take effect
//	time.Sleep(5 * time.Second)
//
//	fmt.Println("Language changed successfully")
//
//	// Proceed to order page
//	if err := wd.Get("https://innovesta.kz/bid"); err != nil {
//		log.Fatalf("Error opening the order page: %v", err)
//	}
//
//	// Wait for the page to load
//	time.Sleep(5 * time.Second)
//
//	// Fill the order form sequentially
//	fields := map[string]string{
//		"//input[@name='first_name']":      formData.FirstName,
//		"//input[@name='last_name']":       formData.LastName,
//		"//input[@name='email']":           formData.Email,
//		"//input[@name='phone']":           formData.Phone,
//		"//select[@name='room_type']":      formData.RoomType,
//		"//input[@name='room_dimensions']": formData.RoomDimensions,
//		"//textarea[@name='comment']":      formData.Comment,
//		"//input[@name='sku']":             formData.SKU,
//	}
//
//	for selector, value := range fields {
//		elem, err := wd.FindElement(selenium.ByXPATH, selector)
//		if err != nil {
//			log.Fatalf("Error finding form element: %v", err)
//		}
//		if err := elem.SendKeys(value); err != nil {
//			log.Fatalf("Error sending keys to form element: %v", err)
//		}
//
//		// Introduce a delay between filling each field
//		time.Sleep(2 * time.Second)
//	}
//
//	// Attach a file
//	fileInput, err := wd.FindElement(selenium.ByXPATH, "//input[@type='file']")
//	if err != nil {
//		log.Fatalf("Error finding element for file upload: %v", err)
//	}
//	if err := fileInput.SendKeys("/Users/aruka/Downloads/5-mb-example-file.pdf"); err != nil {
//		log.Fatalf("Error uploading file: %v", err)
//	}
//
//	//err = wd.WaitWithTimeoutAndInterval(func(wd selenium.WebDriver) (bool, error) {
//	//	element, err := wd.FindElement(selenium.ByXPATH, "//button[@type='submit']")
//	//	if err != nil {
//	//		return false, err
//	//	}
//	//	displayed, err := element.IsDisplayed()
//	//	if err != nil {
//	//		return false, err
//	//	}
//	//	enabled, err := element.IsEnabled()
//	//	if err != nil {
//	//		return false, err
//	//	}
//	//	if !displayed || !enabled {
//	//		return false, nil
//	//	}
//	//	return true, nil
//	//}, 10*time.Second, 1*time.Second)
//	//if err != nil {
//	//	log.Fatalf("Error waiting for submit button to be clickable: %v", err)
//	//}
//	//
//	//// Wait for the form to submit and for the new page to load
//	//time.Sleep(30 * time.Second)
//
//	// Submit the form
//	submitButton, err := wd.FindElement(selenium.ByXPATH, "//button[@type='submit']//span[text()='Отправить']")
//	if err != nil {
//		log.Fatalf("Error finding submit button: %v", err)
//	}
//	if err := submitButton.Click(); err != nil {
//		log.Fatalf("Error clicking submit button: %v", err)
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("successfully submitted")
//
//	time.Sleep(10 * time.Second)
//	fmt.Println("successfully submitted")
//}
