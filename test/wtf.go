package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // Путь к chromedriver
	port             = 8080
)

func main() {
	// Установка сервиса ChromeDriver
	var opts []selenium.ServiceOption
	selenium.SetDebug(false)
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера ChromeDriver: %v", err)
	}
	defer func() {
		if err := service.Stop(); err != nil {
			log.Fatalf("Ошибка остановки сервиса ChromeDriver: %v", err)
		}
	}()

	// Подключение к WebDriver
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Ошибка подключения к WebDriver: %v", err)
	}
	defer func() {
		if err := wd.Quit(); err != nil {
			log.Fatalf("Ошибка завершения сеанса WebDriver: %v", err)
		}
	}()

	// Переход на веб-сайт
	if err := wd.Get("https://smartdeal.kz/categories/processory"); err != nil {
		log.Fatalf("Ошибка открытия страницы: %v", err)
	}

	// Find the icon element by XPath
	sortBy, err := wd.FindElement(selenium.ByXPATH, "//span[@class='icon disable-icon']")
	if err != nil {
		log.Fatalf("Error finding the icon: %v", err)
	}

	// Click on the icon element
	err = sortBy.Click()
	if err != nil {
		log.Fatalf("Error clicking the icon: %v", err)
	}

	// Add a wait condition to ensure the click action is completed
	err = wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {

		// Check if the element is still visible after the click
		visible, err := sortBy.IsDisplayed()
		if err != nil {
			return false, err
		}
		// Return true if the element is visible
		return visible, nil
	}, 10*time.Second)
	if err != nil {
		log.Fatalf("Error waiting for the click action to complete: %v", err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Click on icon completed successfully")

	// Wait for the page to load
	time.Sleep(5 * time.Second)
	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)
	// Find the "Select All" checkbox
	selectAllXPath := "//input[@type='checkbox' and @id='select-all']"
	selectAllCheckbox, err := wd.FindElement(selenium.ByXPATH, selectAllXPath)
	if err != nil {
		log.Fatalf("Error finding the 'Выбрать все' checkbox: %v", err)
	}

	// Click the "Select All" checkbox
	err = selectAllCheckbox.Click()
	if err != nil {
		log.Fatalf("Error clicking the 'Выбрать все' checkbox: %v", err)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Clicked 'Select All' checkbox successfully")

	// Добавить задержку, чтобы дать время для появления кнопки
	time.Sleep(5 * time.Second)

	// Найти кнопку "Добавить все товары в избранное"
	addToFavoritesAllButton, err := wd.FindElement(selenium.ByCSSSelector, ".products__list--select-all .actions__list--item .icon")
	if err != nil {
		log.Fatalf("Ошибка поиска кнопки 'Добавить все товары в избранное': %v", err)
	}

	// Нажать кнопку "Добавить все товары в избранное"
	err = addToFavoritesAllButton.Click()
	if err != nil {
		log.Fatalf("Ошибка нажатия кнопки 'Добавить все товары в избранное': %v", err)
	}
	time.Sleep(15 * time.Second)

	fmt.Println("Все товары добавлены в избранное успешно")

}
