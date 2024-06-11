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

	// Wait for the page to load
	time.Sleep(5 * time.Second)

	// Find the 'В корзину' button
	addToCartButton, err := wd.FindElement(selenium.ByCSSSelector, ".item-card__button.ui-button.ui-button--colored.item-card__button p.item-card__button--text")
	if err != nil {
		log.Fatalf("Error finding the 'В корзину' button: %v", err)
	}

	// Click the 'Add to Cart' button
	if err := addToCartButton.Click(); err != nil {
		log.Fatalf("Error clicking on the 'В корзину' button: %v", err)
	}

	// Wait to ensure the product is added to the cart
	time.Sleep(5 * time.Second)
	fmt.Println("Product successfully added to cart!")

	time.Sleep(5 * time.Second)

	// Find and click the "basket" icon
	basketIconXPath := "//a[@href='/basket']//span[@class='icon']"
	basketIcon, err := wd.FindElement(selenium.ByXPATH, basketIconXPath)
	if err != nil {
		log.Fatalf("Error finding the basket icon: %v", err)
	}
	if err := basketIcon.Click(); err != nil {
		log.Fatalf("Error clicking the basket icon: %v", err)
	}

	fmt.Println("Basket icon clicked successfully")

	// Wait for the basket page to load
	time.Sleep(5 * time.Second)

	// Wait for the loader to disappear
	loaderXPath := "//div[@class='global-loader']"
	for {
		loader, err := wd.FindElement(selenium.ByXPATH, loaderXPath)
		if err != nil {
			break // Loader not found, proceed
		}
		isDisplayed, err := loader.IsDisplayed()
		if err != nil {
			log.Fatalf("Error checking loader visibility: %v", err)
		}
		if !isDisplayed {
			break // Loader not visible, proceed
		}
		time.Sleep(1 * time.Second)
	}

	// Click the "Подтвердить" button
	checkoutButtonXPath := "//p[text()='Подтвердить']"
	checkoutButton, err := wd.FindElement(selenium.ByXPATH, checkoutButtonXPath)
	if err != nil {
		log.Fatalf("Error finding the 'Подтвердить' button: %v", err)
	}
	if err := checkoutButton.Click(); err != nil {
		log.Fatalf("Error clicking the 'Подтвердить' button: %v", err)
	}

	// Wait for the delivery options to load
	time.Sleep(5 * time.Second)

	// Select the first delivery address checkbox
	firstCheckbox, err := wd.FindElement(selenium.ByXPATH, "//div[@class='item-checkboxes']//div[@class='checkbox'][1]//input[@type='radio']")
	if err != nil {
		log.Fatalf("Error finding the first delivery address checkbox: %v", err)
	}
	if err := firstCheckbox.Click(); err != nil {
		log.Fatalf("Error clicking the first delivery address checkbox: %v", err)
	}

	// Wait for the address selection to process
	time.Sleep(2 * time.Second)

	// Confirm the delivery address
	confirmDeliveryButtonXPath := "//button[@class='button ui-button ui-button--colored button']//p[text()='Подтвердить']"
	confirmDeliveryButton, err := wd.FindElement(selenium.ByXPATH, confirmDeliveryButtonXPath)
	if err != nil {
		log.Fatalf("Error finding the 'Подтвердить' button: %v", err)
	}
	if err := confirmDeliveryButton.Click(); err != nil {
		log.Fatalf("Error clicking the 'Подтвердить' button: %v", err)
	}

	// Select the card payment checkbox
	cardPaymentCheckbox, err := wd.FindElement(selenium.ByXPATH, "//input[@type='radio' and @id='epay']")
	if err != nil {
		log.Fatalf("Error finding the card payment checkbox: %v", err)
	}
	if err := cardPaymentCheckbox.Click(); err != nil {
		log.Fatalf("Error clicking the card payment checkbox: %v", err)
	}

	time.Sleep(5 * time.Second)

	// Click on the slider
	sliderElement, err := wd.FindElement(selenium.ByCSSSelector, ".slider.round")
	if err != nil {
		log.Fatalf("Error finding the slider element: %v", err)
	}
	if err := sliderElement.Click(); err != nil {
		log.Fatalf("Error clicking the slider element: %v", err)
	}

	// Find and click the "Оформить заказ" button
	submitOrderButton, err := wd.FindElement(selenium.ByCSSSelector, ".payment-type__button.ui-button.ui-button--colored.payment-type__button p.payment-type__button--text")
	if err != nil {
		log.Fatalf("Error finding the 'Оформить заказ' button: %v", err)
	}
	if err := submitOrderButton.Click(); err != nil {
		log.Fatalf("Error clicking the 'Оформить заказ' button: %v", err)
	}

	// Wait for the address confirmation to process
	time.Sleep(5 * time.Second)

	fmt.Println("Test completed successfully")
}
