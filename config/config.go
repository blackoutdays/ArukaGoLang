package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		fmt.Println("Error starting the ChromeDriver server:", err)
		return
	}
	defer service.Stop()

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{"--headless"},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println("Error creating new WebDriver session:", err)
		return
	}
	defer wd.Quit()

	err = wd.Get("https://smartdeal.kz/")
	if err != nil {
		fmt.Println("Error navigating to smartdeal.kz:", err)
		return
	}

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

	phoneNumberField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите номер телефона']")
	if err != nil {
		fmt.Println("Error finding the phone number field:", err)
		return
	}

	err = phoneNumberField.SendKeys("777 777 77 77")
	if err != nil {
		fmt.Println("Error entering phone number:", err)
		return
	}

	fmt.Println("Test completed successfully.")
}
