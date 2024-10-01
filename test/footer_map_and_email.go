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

	// Прокрутка страницы вниз до футера
	_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
	if err != nil {
		log.Fatalf("Ошибка прокрутки страницы: %v", err)
	}

	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Нахождение ссылки на Google Карты и клик по ней
	googleMapsLink, err := wd.FindElement(selenium.ByCSSSelector, "a.contacts__item[href^='https://www.google.com/maps']")
	if err != nil {
		log.Fatalf("Ошибка при поиске ссылки на Google Карты: %v", err)
	}
	err = googleMapsLink.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по ссылке на Google Карты: %v", err)
	}

	// Переключение на новое окно (Google Карты)
	err = wd.SwitchWindow("")
	if err != nil {
		log.Fatalf("Ошибка переключения на окно Google Карты: %v", err)
	}

	// Ждем, пока загрузится карта
	time.Sleep(5 * time.Second)

	// Переключение обратно на основное окно
	err = wd.SwitchWindow("")
	if err != nil {
		log.Fatalf("Ошибка переключения на основное окно: %v", err)
	}

	// Нахождение ссылки на почту и клик по ней
	mailLink, err := wd.FindElement(selenium.ByCSSSelector, "a.contacts__item[href^='mailto']")
	if err != nil {
		log.Fatalf("Ошибка при поиске ссылки на почту: %v", err)
	}
	err = mailLink.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по ссылке на почту: %v", err)
	}

	// Переключение на новое окно (почта)
	err = wd.SwitchWindow("")
	if err != nil {
		log.Fatalf("Ошибка переключения на окно почты: %v", err)
	}

	// Ждем, пока загрузится страница почты
	time.Sleep(5 * time.Second)

	// Переключение обратно на основное окно
	err = wd.SwitchWindow("")
	if err != nil {
		log.Fatalf("Ошибка переключения на основное окно: %v", err)
	}

	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Вывод текущего URL страницы
	currentURL, err := wd.CurrentURL()
	if err != nil {
		log.Printf("Ошибка получения текущего URL: %v", err)
	} else {
		fmt.Printf("Текущий URL: %s\n", currentURL)
	}
}
