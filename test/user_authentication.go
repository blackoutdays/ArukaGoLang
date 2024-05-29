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

	// Find and click the login button
	loginButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Войти']")
	if err != nil {
		fmt.Println("Error finding the login button:", err)
		return
	}

	err = loginButton.Click()
	if err != nil {
		fmt.Println("Error clicking the login button:", err)
		return
	}

	time.Sleep(2 * time.Second)

	// Find the phone number input field and enter the phone number
	phoneNumberField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите номер телефона']")
	if err != nil {
		fmt.Println("Error finding the phone number field:", err)
		return
	}

	err = phoneNumberField.SendKeys("7 708 225 26 75")
	if err != nil {
		fmt.Println("Error entering phone number:", err)
		return
	}

	time.Sleep(20 * time.Second)

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

	// Wait for the SMS code input field to be present
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите код']")
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		fmt.Println("SMS code input field not found:", err)
		return
	}

	smsCodeField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите код']")
	if err != nil {
		fmt.Println("Error finding the SMS code input field:", err)
		return
	}

	// Here, you should replace "SMS_CODE" with the actual code received via SMS
	smsCode := "SMS_CODE" // You need to replace this with the actual code received
	err = smsCodeField.SendKeys(smsCode)
	if err != nil {
		fmt.Println("Error entering SMS code:", err)
		return
	}

	time.Sleep(20 * time.Second)
	fmt.Println("Test completed successfully")
}
