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

	// Find the login button and click it
	loginButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Войти']")
	if err != nil {
		log.Fatalf("Error finding the login button: %v", err)
	}

	err = loginButton.Click()
	if err != nil {
		log.Fatalf("Error clicking the login button: %v", err)
	}

	time.Sleep(2 * time.Second)

	// Find the phone number input field and enter the phone number
	phoneNumberField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите номер телефона']")
	if err != nil {
		log.Fatalf("Error finding the phone number field: %v", err)
	}

	err = phoneNumberField.SendKeys("7 777 832 63 35")
	if err != nil {
		log.Fatalf("Error entering phone number: %v", err)
	}

	time.Sleep(2 * time.Second)
	// Find and click the SMS button
	smsButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Получить код']")
	if err != nil {
		fmt.Println("Error finding the SMS button:", err)
		return
	}
	err = smsButton.Click()
	if err != nil {
		fmt.Println("Error clicking the SMS button:", err)
		return
	}

	// Wait for the SMS code input fields to be present
	time.Sleep(5 * time.Second)
	smsCodeBlocks, err := wd.FindElements(selenium.ByCSSSelector, ".auth__code-form--otp .otp-input-container")
	if err != nil {
		log.Fatalf("Error finding SMS code blocks: %v", err)
	}

	// Input the SMS code blocks
	smsCode := "0000" // Replace with your actual SMS code
	for i, block := range smsCodeBlocks {
		// Find the input field within the current block
		inputField, err := block.FindElement(selenium.ByCSSSelector, "input")
		if err != nil {
			log.Fatalf("Error finding input field in SMS code block %d: %v", i+1, err)
		}

		// Input the corresponding block of the SMS code into the input field
		blockCode := smsCode[i*4 : (i+1)*4]
		err = inputField.SendKeys(blockCode)
		if err != nil {
			log.Fatalf("Error entering SMS code block %d: %v", i+1, err)
		}
	}

	time.Sleep(34 * time.Second)
	fmt.Println("Test completed successfully")
}
