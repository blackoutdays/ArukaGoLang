// main.go
package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	// Список файлов для выполнения по порядку
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
		//"catalog_click_all.go",
		//"favourite_cart_without_favourite_item.go",
		//"favouriteCart_with_favouriteItem.go",
		//"footer_map_and_email.go",
		//"footer_social_media.go",
		//"footer_policies.go",
		//"go_top_button.go",
		//"header_b2b.go",
		//"header_switch_city.go", //dev fix
		//"header_switch_language_kz.go",
		//"products_filter.go",
		//"products_pagination.go", //why
		"products_pagination_linear_sorting.go",
		"products_number_per_page.go",
		"products_sort.go",
		"support_chat.go",
		"user_login.go",
		"user_orders.go",
		"user_profile.go",
		"user_profile_edit.go",
		"user_profile_tab_cards.go",
		"user_cart_payment.go",
		"user_exit.go",
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
