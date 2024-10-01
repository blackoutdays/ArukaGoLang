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

	// Названия ссылок, которые нужно открыть
	desiredLinks := []string{
		"/policy/1",
		"/policy/2",
		"/policy/3",
		"/policy/4",
		"/policy/5",
		"/policy/6",
	}

	// Обход всех ссылок и переход по нужным
	for _, link := range desiredLinks {
		// Прокрутка страницы вниз до футера
		_, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
		if err != nil {
			log.Fatalf("Ошибка прокрутки страницы: %v", err)
		}

		// Ожидание загрузки страницы
		time.Sleep(5 * time.Second)

		// Формируем полный URL ссылки
		fullURL := "https://smartdeal.kz" + link

		// Открываем ссылку
		if err := wd.Get(fullURL); err != nil {
			log.Printf("Ошибка открытия страницы %s: %v", fullURL, err)
			continue
		}

		// Ожидание загрузки страницы
		time.Sleep(5 * time.Second)

		// Вывод текущего URL страницы
		currentURL, err := wd.CurrentURL()
		if err != nil {
			log.Printf("Ошибка получения текущего URL для ссылки %s: %v", fullURL, err)
		} else {
			fmt.Printf("Ссылка %s ведет на %s\n", fullURL, currentURL)
		}
	}
}
