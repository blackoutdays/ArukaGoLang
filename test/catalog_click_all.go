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

func clickSubCatalog(wd selenium.WebDriver, xpath string) error {
	err := wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByXPATH, xpath)
		return err == nil, nil
	}, 10*time.Second)
	if err != nil {
		return fmt.Errorf("error finding the subcatalog button %q: %v", xpath, err)
	}

	subCatalogButton, err := wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		return fmt.Errorf("error finding the subcatalog button %q: %v", xpath, err)
	}

	err = subCatalogButton.Click()
	if err != nil {
		return fmt.Errorf("error clicking the subcatalog button %q: %v", xpath, err)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Click on subcatalog %q completed successfully\n", xpath)
	return nil
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

	// List of subcatalogs to click
	subCatalogs := []string{
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Телефоны и гаджеты']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Компьютеры']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Бытовая техника']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='ТВ, Аудио, Видео']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Красота и здоровье']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Автотовары']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Сетевое и серверное оборудование']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Строительство, ремонт']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Товары для дома и дачи']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Детские товары']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Мебель']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Спорт, туризм']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Аксессуары']",
		"//div[contains(@class, 'categories-menu__list--item')]//p[text()='Канцелярские товары']",
	}

	// Click on each subcatalog button
	for _, xpath := range subCatalogs {
		if err := clickSubCatalog(wd, xpath); err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Test completed successfully")
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
//	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // replace with your driver
//	port             = 8080
//)
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
//
//	// Connect to the WebDriver
//	caps := selenium.Capabilities{"browserName": "chrome"}
//	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
//	if err != nil {
//		log.Fatalf("Error connecting to the WebDriver: %v", err)
//	}
//	defer func() {
//		if err := wd.Quit(); err != nil {
//			log.Fatalf("Error quitting the WebDriver session: %v", err)
//		}
//	}()
//
//	// Navigate to the website
//	if err := wd.Get("https://smartdeal.kz/"); err != nil {
//		log.Fatalf("Error opening the page: %v", err)
//	}
//
//	// Use explicit wait to find the catalog button
//	catalogButtonXPath := "//button[contains(@class, 'header__bottom--catalog')]"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, catalogButtonXPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the catalog button:", err)
//		return
//	}
//
//	catalogButton, err := wd.FindElement(selenium.ByXPATH, catalogButtonXPath)
//	if err != nil {
//		fmt.Println("Error finding the catalog button:", err)
//		return
//	}
//
//	err = catalogButton.Click()
//	if err != nil {
//		fmt.Println("Error clicking the catalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on catalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton1XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Телефоны и гаджеты']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton1XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Телефоны и гаджеты' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton1, err := wd.FindElement(selenium.ByXPATH, subCatalogButton1XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Телефоны и гаджеты' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton1.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Телефоны и гаджеты' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Телефоны и гаджеты' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton2XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Компьютеры']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton2XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Компьютеры' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton2, err := wd.FindElement(selenium.ByXPATH, subCatalogButton2XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Компьютеры' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton2.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Компьютеры' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Компьютеры' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton3XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Бытовая техника']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton3XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Бытовая техника' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton3, err := wd.FindElement(selenium.ByXPATH, subCatalogButton3XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Бытовая техника' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton3.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Бытовая техника' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Бытовая техника' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton4XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='ТВ, Аудио, Видео']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton4XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'ТВ, Аудио, Видео' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton4, err := wd.FindElement(selenium.ByXPATH, subCatalogButton2XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'ТВ, Аудио, Видео' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton4.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'ТВ, Аудио, Видео' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'ТВ, Аудио, Видео' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton5XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Красота и здоровье']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton5XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Красота и здоровье' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton5, err := wd.FindElement(selenium.ByXPATH, subCatalogButton5XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Красота и здоровье' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton5.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Красота и здоровье' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Красота и здоровье' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton6XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Автотовары']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton6XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Автотовары' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton6, err := wd.FindElement(selenium.ByXPATH, subCatalogButton6XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Автотовары' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton6.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Автотовары' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Автотовары' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton7XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Сетевое и серверное оборудование']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton7XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Сетевое и серверное оборудование' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton7, err := wd.FindElement(selenium.ByXPATH, subCatalogButton7XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Сетевое и серверное оборудование' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton7.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Сетевое и серверное оборудование' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Сетевое и серверное оборудование' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton8XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Строительство, ремонт']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton8XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Строительство, ремонт' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton8, err := wd.FindElement(selenium.ByXPATH, subCatalogButton8XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Строительство, ремонт' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton8.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Строительство, ремонт' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Строительство, ремонт' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton9XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Товары для дома и дачи']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton9XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Товары для дома и дачи' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton9, err := wd.FindElement(selenium.ByXPATH, subCatalogButton9XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Товары для дома и дачи' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton9.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Товары для дома и дачи' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Товары для дома и дачи' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton10XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Детские товары']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton10XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Детские товары' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton10, err := wd.FindElement(selenium.ByXPATH, subCatalogButton10XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Детские товары' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton10.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Детские товары' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Детские товары' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton11XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Мебель']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton11XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Мебель' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton11, err := wd.FindElement(selenium.ByXPATH, subCatalogButton2XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Мебель' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton11.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Мебель' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Мебель' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton12XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Спорт, туризм']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton12XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Спорт, туризм' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton12, err := wd.FindElement(selenium.ByXPATH, subCatalogButton2XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Спорт, туризм' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton12.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Спорт, туризм' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Спорт, туризм' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton13XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Аксессуары']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton13XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Аксессуары' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton13, err := wd.FindElement(selenium.ByXPATH, subCatalogButton13XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Аксессуары' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton13.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Аксессуары' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Аксессуары' subcatalog completed successfully")
//
//	// Continue with finding a subcatalog button
//	subCatalogButton14XPath := "//div[contains(@class, 'categories-menu__list--item')]//p[text()='Канцелярские товары']"
//	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
//		_, err := wd.FindElement(selenium.ByXPATH, subCatalogButton14XPath)
//		return err == nil, nil
//	}, 10*time.Second)
//	if err != nil {
//		fmt.Println("Error finding the 'Канцелярские товары' subcatalog button:", err)
//		return
//	}
//
//	subCatalogButton14, err := wd.FindElement(selenium.ByXPATH, subCatalogButton14XPath)
//	if err != nil {
//		fmt.Println("Error finding the 'Канцелярские товары' subcatalog button:", err)
//		return
//	}
//
//	err = subCatalogButton14.Click()
//	if err != nil {
//		fmt.Println("Error clicking the 'Канцелярские товары' subcatalog button:", err)
//		return
//	}
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Click on 'Канцелярские товары' subcatalog completed successfully")
//
//	time.Sleep(5 * time.Second)
//	fmt.Println("Test completed successfully")
//}
