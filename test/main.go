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
		//"add_all_products_to_cart.go",
		//"add_all_products_to_favourite.go",
		//"add_to_cart.go",
		//"add_to_favourites.go",
		//"address_add_and_edit.go",
		//"address_edit.go",
		//"address_delete_negative.go",
		//"address_delete_positive.go",
		//"app_links.go",
		//"catalog.go",
		//"footer_policies.go",
		//"go_top_button.go",
		//"header_b2b.go",
		//"header_switch_city.go",
		//"header_switch_language_kz.go",
		//"products_pagination_linear_sorting.go",
		//"products_number_per_page.go",
		//"products_sort.go",
		//"support_chat.go",
		//"user_orders.go",
		"user_profile.go",
		"user_profile_edit.go",
		"user_profile_tab_cards.go",
		"user_exit.go",
		"user_login.go",
		// остальные файлы ;))
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
