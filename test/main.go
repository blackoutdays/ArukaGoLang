package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	// Список файлов, которые нужно запустить по порядку
	files := []string{
		"app_links.go",
		"user_login.go",
		"go-top-button.go",
		// Добавь остальные файлы pls
	}

	for _, file := range files {
		// Конструируем полный путь к файлу
		filePath := filepath.Join("test", file)

		// Запускаем файл
		cmd := exec.Command("go", "run", filePath)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Ошибка выполнения %s: %v, Output: %s", file, err, string(output))
		}

		// Выводим результат выполнения
		fmt.Printf("Output of %s: %s\n", file, string(output))
	}
}
