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
	if err := wd.Get("https://smartdeal.kz/"); err != nil {
		log.Fatalf("Ошибка открытия страницы: %v", err)
	}
	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Найти кнопку "main-button" по XPath
	mainButton, err := wd.FindElement(selenium.ByXPATH, "//button[@class='main-button']")
	if err != nil {
		log.Fatalf("Ошибка нахождения кнопки 'main-button': %v", err)
	}

	// Кликнуть на кнопку "main-button"
	if err := mainButton.Click(); err != nil {
		log.Fatalf("Ошибка клика на кнопку 'main-button': %v", err)
	}
	time.Sleep(5 * time.Second) // Ожидание загрузки страницы
	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Найти элементы чата
	whatsappButton, err := wd.FindElement(selenium.ByXPATH, "//button[@class='social-button whatsapp-button']")
	if err != nil {
		log.Fatalf("Ошибка нахождения кнопки WhatsApp: %v", err)
	}

	telegramButton, err := wd.FindElement(selenium.ByXPATH, "//button[@class='social-button telegram-button']")
	if err != nil {
		log.Fatalf("Ошибка нахождения кнопки Telegram: %v", err)
	}

	// Кликнуть на кнопки чата
	if err := whatsappButton.Click(); err != nil {
		log.Fatalf("Ошибка клика на кнопку WhatsApp: %v", err)
	}
	time.Sleep(5 * time.Second) // Ожидание загрузки страницы

	if err := telegramButton.Click(); err != nil {
		log.Fatalf("Ошибка клика на кнопку Telegram: %v", err)
	}
	time.Sleep(5 * time.Second) // Ожидание загрузки страницы

	fmt.Println("Тест завершен успешно")
}
