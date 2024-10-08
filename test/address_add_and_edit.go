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

	if err := wd.Get("https://smartdeal.kz/"); err != nil {
		log.Fatalf("Error opening the page: %v", err)
	}

	time.Sleep(5 * time.Second)

	loginButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Войти']")
	if err != nil {
		log.Fatalf("Error finding the login button: %v", err)
	}

	err = loginButton.Click()
	if err != nil {
		log.Fatalf("Error clicking the login button: %v", err)
	}
	fmt.Println("Login button clicked successfully")

	time.Sleep(2 * time.Second)

	phoneNumberField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите номер телефона']")
	if err != nil {
		log.Fatalf("Error finding the phone number field: %v", err)
	}

	err = phoneNumberField.SendKeys("7 777 832 63 35")
	if err != nil {
		log.Fatalf("Error entering phone number: %v", err)
	}

	fmt.Println("Phone number entered successfully")

	time.Sleep(2 * time.Second)

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

	fmt.Println("SMS button clicked successfully")

	time.Sleep(5 * time.Second)
	smsCodeBlocks, err := wd.FindElements(selenium.ByCSSSelector, ".auth__code-form--otp .otp-input-container")
	if err != nil {
		log.Fatalf("Error finding SMS code blocks: %v", err)
	}

	smsCode := "0000" // Replace with your actual SMS code
	for i, block := range smsCodeBlocks {
		inputField, err := block.FindElement(selenium.ByCSSSelector, "input")
		if err != nil {
			log.Fatalf("Error finding input field in SMS code block %d: %v", i+1, err)
		}

		blockCode := smsCode[i*4 : (i+1)*4]
		err = inputField.SendKeys(blockCode)
		if err != nil {
			log.Fatalf("Error entering SMS code block %d: %v", i+1, err)
		}

		fmt.Println("SMS code entered successfully")

		time.Sleep(5 * time.Second)

		profileButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='+77778326335']")
		if err != nil {
			log.Fatalf("Error finding the profile button: %v", err)
		}

		err = profileButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the profile button: %v", err)
		}

		fmt.Println("Profile button clicked successfully")

		time.Sleep(5 * time.Second)

		personalAccountButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Профиль']")
		if err != nil {
			log.Fatalf("Error finding the login button: %v", err)
		}

		err = personalAccountButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the login button: %v", err)
		}

		fmt.Println("Profile information button clicked successfully")

		time.Sleep(5 * time.Second)

		myAddressesButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Мои адреса']")
		if err != nil {
			log.Fatalf("Error finding the 'Мои адреса' button: %v", err)
		}

		err = myAddressesButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the 'Мои адреса' button: %v", err)
		}
		time.Sleep(5 * time.Second)

		fmt.Println("Button 'Мои адреса' clicked successfully")

		time.Sleep(5 * time.Second)
		myAddressesAddButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Добавить новый адрес']")
		if err != nil {
			log.Fatalf("Error finding the 'Добавить новый адрес' button: %v", err)
		}

		err = myAddressesAddButton.Click()
		if err != nil {
			log.Fatalf("Error clicking the 'Добавить новый адрес' button: %v", err)
		}

		fmt.Println("Button 'Добавить новый адрес' clicked successfully")

		time.Sleep(5 * time.Second)
		//add street in address
		addressStreetField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Улица/микрорайон, № дома']")
		if err != nil {
			log.Fatalf("Error finding the address street field 'Улица/микрорайон, № дома': %v", err)
		}

		err = addressStreetField.SendKeys("Назарбаева 235")
		if err != nil {
			log.Fatalf("Error entering address street: %v", err)
		}

		fmt.Println("Field 'Улица/микрорайон, № дома' completed successfully")

		time.Sleep(5 * time.Second)

		addressFloorField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Этаж']")
		if err != nil {
			log.Fatalf("Error finding the address floor field: %v", err)
		}

		err = addressFloorField.SendKeys("12")
		if err != nil {
			log.Fatalf("Error entering floor in the address: %v", err)
		}

		fmt.Println("Field 'Этаж' completed successfully")

		time.Sleep(5 * time.Second)

		addressBldgField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Кв./офис']")
		if err != nil {
			log.Fatalf("Error finding the address building (apartmnet/office field: %v", err)
		}

		err = addressBldgField.SendKeys("24")
		if err != nil {
			log.Fatalf("Error entering building (apartmnet/office) in the address: %v", err)
		}

		fmt.Println("Field 'Кв./офис' completed successfully")

		time.Sleep(5 * time.Second)

		addAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Добавить']")
		if err != nil {
			log.Fatalf("Error finding the 'Добавить' button: %v", err)
		}

		err = addAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Добавить' address: %v", err)
		}

		fmt.Println("Button 'Добавить' clicked successfully")

		time.Sleep(5 * time.Second)

		editAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Изменить']")
		if err != nil {
			log.Fatalf("Error finding the 'Изменить' button: %v", err)
		}

		err = editAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Изменить' address: %v", err)
		}

		fmt.Println("Button 'Изменить' clicked successfully")

		time.Sleep(5 * time.Second)

		addressSavedStreetField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Улица/микрорайон, № дома']")
		if err != nil {
			log.Fatalf("Error finding the address street field: %v", err)
		}

		err = addressSavedStreetField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address street field: %v", err)
		}

		fmt.Println("Existing field in 'Улица/микрорайон, № дома' cleared successfully")

		err = addressSavedStreetField.SendKeys("New Назарбаева 235")
		if err != nil {
			log.Fatalf("Error entering address street: %v", err)
		}

		fmt.Println("NEW field 'Улица/микрорайон, № дома' completed successfully")

		time.Sleep(5 * time.Second)

		addressSavedFloorField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Этаж']")
		if err != nil {
			log.Fatalf("Error finding the address floor field: %v", err)
		}

		err = addressSavedFloorField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address 'Этаж' field: %v", err)
		}

		fmt.Println("Existing field in 'Этаж' cleared successfully")

		err = addressSavedFloorField.SendKeys("1")
		if err != nil {
			log.Fatalf("Error entering 'Этаж' in the address: %v", err)
		}

		fmt.Println("NEW field 'Этаж' completed successfully")

		time.Sleep(5 * time.Second)

		addressSavedBldgField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Кв./офис']")
		if err != nil {
			log.Fatalf("Error finding the address 'Кв./офис' field: %v", err)
		}

		err = addressSavedBldgField.Clear() // Clear existing text
		if err != nil {
			log.Fatalf("Error clearing address 'Кв./офис' field: %v", err)
		}

		fmt.Println("Existing field in 'Кв./офис' cleared successfully")

		err = addressSavedBldgField.SendKeys("24")
		if err != nil {
			log.Fatalf("Error entering 'Кв./офис' in the address: %v", err)
		}

		fmt.Println("NEW field 'Кв./офис' completed successfully")

		time.Sleep(5 * time.Second)

		addEditedAddress, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Изменить']")
		if err != nil {
			log.Fatalf("Error finding the 'Изменить' button: %v", err)
		}

		err = addEditedAddress.Click()
		if err != nil {
			log.Fatalf("Error clicking 'Изменить' address: %v", err)
		}

		fmt.Println("Button 'Изменить' clicked successfully")

		time.Sleep(10 * time.Second)
		fmt.Println("Test completed successfully")
	}
}
