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

		// Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the login button and click it
		profileButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='+77778326335']")
		if err != nil {
			log.Fatalf("Error finding the login button: %v", err)
		}

		err = profileButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the login button: %v", err)
		}

		//Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the login button and click it
		personalAccountButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Профиль']")
		if err != nil {
			log.Fatalf("Error finding the login button: %v", err)
		}

		err = personalAccountButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the login button: %v", err)
		}

		//Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the
		sortBySquare, err := wd.FindElement(selenium.ByXPATH, "//span[@class='icon main-tab__info--edit active-icon']")
		if err != nil {
			log.Fatalf("Error finding the edit: %v", err)
		}

		// Click on the edit element
		err = sortBySquare.Click()
		if err != nil {
			log.Fatalf("Error clicking the edit: %v", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Println("Click on edit completed successfully")

		time.Sleep(5 * time.Second)
		//add street in address
		myEmail, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Email']")
		if err != nil {
			log.Fatalf("Error finding the name field: %v", err)
		}

		err = myEmail.SendKeys("arukalarkins@icloud.com")
		if err != nil {
			log.Fatalf("Error entering in the name: %v", err)
		}

		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		myName, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Имя']")
		if err != nil {
			log.Fatalf("Error finding the name field: %v", err)
		}

		err = myName.SendKeys("Aruka")
		if err != nil {
			log.Fatalf("Error entering in the name: %v", err)
		}
		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		mySurname, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Фамилия']")
		if err != nil {
			log.Fatalf("Error finding the name field: %v", err)
		}

		err = mySurname.SendKeys("Larkins")
		if err != nil {
			log.Fatalf("Error entering in the name: %v", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Println("Click on edit completed successfully")
		//Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the login button and click it
		saveButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Сохранить']")
		if err != nil {
			log.Fatalf("Error finding the login button: %v", err)
		}

		err = saveButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the login button: %v", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Println("Click on edit completed successfully")

		//Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the login button and click it
		backButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Назад']")
		if err != nil {
			log.Fatalf("Error finding the login button: %v", err)
		}

		err = backButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the login button: %v", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Println("Test completed successfully")
	}
}
