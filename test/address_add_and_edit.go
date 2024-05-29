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

		// Find the login button and click it
		myOrdersButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Мои заказы']")
		if err != nil {
			log.Fatalf("Error finding the myOrders button: %v", err)
		}

		err = myOrdersButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the myOrders button: %v", err)
		}

		//Wait for the page to load
		time.Sleep(5 * time.Second)

		// Find the 'My addresses' button and click it
		myAddressesButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Мои адреса']")
		if err != nil {
			log.Fatalf("Error finding the 'My addresses' button: %v", err)
		}

		err = myAddressesButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the 'My addresses' button: %v", err)
		}
		time.Sleep(5 * time.Second)

		// 'My addresses' edit
		time.Sleep(5 * time.Second)
		myAddressesAddButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Добавить новый адрес']")
		if err != nil {
			log.Fatalf("Error finding the 'Add new address' button: %v", err)
		}

		err = myAddressesAddButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the 'My addresses' button: %v", err)
		}

		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		addressStreetField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Улица/микрорайон, № дома']")
		if err != nil {
			log.Fatalf("Error finding the address street field: %v", err)
		}

		err = addressStreetField.SendKeys("Назарбаева 235")
		if err != nil {
			log.Fatalf("Error entering address street: %v", err)
		}

		//add a floor in the address
		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		addressFloorField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Этаж']")
		if err != nil {
			log.Fatalf("Error finding the address floor field: %v", err)
		}

		err = addressFloorField.SendKeys("12")
		if err != nil {
			log.Fatalf("Error entering floor in the address: %v", err)
		}

		//add an apartment/office building in the address
		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		addressBldgField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Кв./офис']")
		if err != nil {
			log.Fatalf("Error finding the address building (apartmnet/office field: %v", err)
		}

		err = addressBldgField.SendKeys("24")
		if err != nil {
			log.Fatalf("Error entering building (apartmnet/office) in the address: %v", err)
		}

		//click button to 'Add'
		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		addAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Добавить']")
		if err != nil {
			log.Fatalf("Error finding the 'Add' button: %v", err)
		}

		err = addAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Add' address: %v", err)
		}

		//click button to 'Edit'
		//Wait for the page to load
		time.Sleep(5 * time.Second)
		//add street in address
		editAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Изменить']")
		if err != nil {
			log.Fatalf("Error finding the 'Edit' button: %v", err)
		}

		err = editAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Edit' address: %v", err)
		}

		//edit the saved address
		time.Sleep(5 * time.Second)

		// Add street in address
		addressSavedStreetField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Улица/микрорайон, № дома']")
		if err != nil {
			log.Fatalf("Error finding the address street field: %v", err)
		}

		err = addressSavedStreetField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address street field: %v", err)
		}

		err = addressSavedStreetField.SendKeys("New Назарбаева 235")
		if err != nil {
			log.Fatalf("Error entering address street: %v", err)
		}

		// Add floor in address
		time.Sleep(5 * time.Second)

		addressSavedFloorField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Этаж']")
		if err != nil {
			log.Fatalf("Error finding the address floor field: %v", err)
		}

		err = addressSavedFloorField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address floor field: %v", err)
		}

		err = addressSavedFloorField.SendKeys("1")
		if err != nil {
			log.Fatalf("Error entering floor in the address: %v", err)
		}

		// Add an apartment/office building in the address
		time.Sleep(5 * time.Second)

		addressSavedBldgField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Кв./офис']")
		if err != nil {
			log.Fatalf("Error finding the address building (apartment/office) field: %v", err)
		}

		err = addressSavedBldgField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address building (apartment/office) field: %v", err)
		}

		err = addressSavedBldgField.SendKeys("24")
		if err != nil {
			log.Fatalf("Error entering building (apartment/office) in the address: %v", err)
		}

		// Click the 'Edit' button
		time.Sleep(5 * time.Second)

		addEditedAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Изменить']")
		if err != nil {
			log.Fatalf("Error finding the 'Edit' button: %v", err)
		}

		err = addEditedAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Edit' address: %v", err)
		}

		time.Sleep(5 * time.Second)
		fmt.Println("Test completed successfully")
	}
}
