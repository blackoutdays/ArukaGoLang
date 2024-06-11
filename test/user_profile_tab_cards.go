//comment success in english and names in ru pls ;<

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	chromeDriverPath = "/Users/aruka/Downloads/chromedriver-mac-arm64-2/chromedriver" // замените на путь к вашему драйверу
	port             = 8080
)

func main() {
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

	// Навигация на веб-сайт
	if err := wd.Get("https://smartdeal.kz/"); err != nil {
		log.Fatalf("Ошибка открытия страницы: %v", err)
	}

	// Ожидание загрузки страницы
	time.Sleep(5 * time.Second)

	// Нахождение кнопки "Войти" и клик по ней
	loginButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Войти']")
	if err != nil {
		log.Fatalf("Ошибка нахождения кнопки входа: %v", err)
	}

	err = loginButton.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по кнопке входа: %v", err)
	}
	fmt.Println("Click on 'Войти' completed successfully")
	time.Sleep(2 * time.Second)

	// Нахождение поля для ввода номера телефона и ввод номера
	phoneNumberField, err := wd.FindElement(selenium.ByXPATH, "//input[@placeholder='Введите номер телефона']")
	if err != nil {
		log.Fatalf("Ошибка нахождения поля для ввода номера телефона: %v", err)
	}

	err = phoneNumberField.SendKeys("7 777 832 63 35")
	if err != nil {
		log.Fatalf("Ошибка ввода номера телефона: %v", err)
	}
	fmt.Println("Click on Favourite card completed successfully")

	time.Sleep(2 * time.Second)
	// Нахождение и клик кнопки "Получить код"
	smsButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Получить код']")
	if err != nil {
		fmt.Println("Ошибка нахождения кнопки SMS:", err)
		return
	}
	err = smsButton.Click()
	if err != nil {
		fmt.Println("Ошибка клика по кнопке SMS:", err)
		return
	}

	// Ожидание появления полей для ввода кода SMS
	time.Sleep(5 * time.Second)
	smsCodeBlocks, err := wd.FindElements(selenium.ByCSSSelector, ".auth__code-form--otp .otp-input-container")
	if err != nil {
		log.Fatalf("Ошибка нахождения блоков для ввода кода SMS: %v", err)
	}

	// Ввод кода SMS
	smsCode := "0000" // Замените на ваш реальный код SMS
	for i, block := range smsCodeBlocks {
		inputField, err := block.FindElement(selenium.ByCSSSelector, "input")
		if err != nil {
			log.Fatalf("Ошибка нахождения поля ввода в блоке кода SMS %d: %v", i+1, err)
		}

		blockCode := smsCode[i*4 : (i+1)*4]
		err = inputField.SendKeys(blockCode)
		if err != nil {
			log.Fatalf("Ошибка ввода блока кода SMS %d: %v", i+1, err)
		}

		time.Sleep(5 * time.Second)

		// Нахождение и клик по кнопке профиля
		profileButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='+77778326335']")
		if err != nil {
			log.Fatalf("Ошибка нахождения кнопки профиля: %v", err)
		}

		err = profileButton.Click()
		if err != nil {
			log.Fatalf("Ошибка клика по кнопке профиля: %v", err)
		}

		time.Sleep(5 * time.Second)
		// Нахождение и клик по карточке "Мои заказы"
		orderCard, err := wd.FindElement(selenium.ByXPATH, "//div[@class='main-tab__cards']//p[text()='Мои заказы']")
		if err != nil {
			log.Fatalf("Ошибка при поиске карточки 'Мои заказы': %v", err)
		}

		if err = orderCard.Click(); err != nil {
			log.Fatalf("Ошибка клика по карточке 'Мои заказы': %v", err)
		}

		time.Sleep(5 * time.Second)
		// Нахождение и клик по кнопке "Профиль"
		personalAccountButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Профиль']")
		if err != nil {
			log.Fatalf("Ошибка нахождения кнопки 'Профиль': %v", err)
		}

		err = personalAccountButton.Click()
		if err != nil {
			log.Fatalf("Ошибка клика по кнопке 'Профиль': %v", err)
		}
		time.Sleep(5 * time.Second)

		// Нахождение и клик по карточке "Мои адреса"
		addressesCard, err := wd.FindElement(selenium.ByXPATH, "//div[@class='main-tab__cards']//div[@class='tab-card__header--title']//p[text()='Мои адреса']")
		if err != nil {
			log.Fatalf("Ошибка при поиске карточки 'Мои адреса': %v", err)
		}

		if err = addressesCard.Click(); err != nil {
			log.Fatalf("Ошибка клика по карточке 'Мои адреса': %v", err)
		}

		time.Sleep(5 * time.Second)
		// Нахождение и клик по кнопке "Профиль"
		personalAccountButton2, err := wd.FindElement(selenium.ByXPATH, "//p[text()='Профиль']")
		if err != nil {
			log.Fatalf("Ошибка нахождения кнопки 'Профиль': %v", err)
		}

		err = personalAccountButton2.Click()
		if err != nil {
			log.Fatalf("Ошибка клика по кнопке 'Профиль': %v", err)
		}
		time.Sleep(5 * time.Second)

		// Нахождение и клик по карточке "Избранное"
		favouriteCard, err := wd.FindElement(selenium.ByXPATH, "//div[@class='main-tab__cards']//div[@class='tab-card__header--title']//p[text()='Избранное']")
		if err != nil {
			log.Fatalf("Ошибка клика по карточке: %v", err)
		}
		if err = favouriteCard.Click(); err != nil {
			log.Fatalf("Ошибка клика по карточке 'Избранное': %v", err)
		}
		fmt.Println("Click on Favourite card completed successfully")
		time.Sleep(5 * time.Second)
	}

	time.Sleep(5 * time.Second)

	// Нахождение и клик по кнопке профиля
	profileButton, err := wd.FindElement(selenium.ByXPATH, "//p[text()='+77778326335']")
	if err != nil {
		log.Fatalf("Ошибка нахождения кнопки профиля: %v", err)
	}

	err = profileButton.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по кнопке профиля: %v", err)
	}

	time.Sleep(5 * time.Second)

	// Нахождение кнопки Telegram и клик по ней
	telegramButton, err := wd.FindElement(selenium.ByCSSSelector, "a.button[href='https://t.me/smartdealkz']")
	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки Telegram: %v", err)
	}
	err = telegramButton.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по кнопке Telegram: %v", err)
	}

	// Нахождение кнопки WhatsApp и клик по ней
	whatsappButton, err := wd.FindElement(selenium.ByCSSSelector, "a.button[href='https://api.whatsapp.com/send?phone=77089450729']")
	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки WhatsApp: %v", err)
	}
	err = whatsappButton.Click()
	if err != nil {
		log.Fatalf("Ошибка клика по кнопке WhatsApp: %v", err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Тест успешно завершен")

}
