// "Навыки":
// 1. Возвращает json ответ на POST запрос
// 2. Флаги. После компиляции
// go build -o skill
// (такая комманда компилирует приложение, но не устанавливает расширение exe, дописываю в ручную)
// после комманды
// $./skill
// возвращает в терминале
//Running server on :8080
// после комманды
// $./skill -a :8081
// возвращает в терминале
//Running server on :8081
//
// Update: 23.04.2025

package main

import (
	"fmt"
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	// обрабатываем аргументы командной строки
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	fmt.Println("Running server on", flagRunAddr)
	return http.ListenAndServe(flagRunAddr, http.HandlerFunc(webhook))
}

// //Такой функция run была до внедрения флагов
// func run() error {
// 	return http.ListenAndServe(`localhost:8080`, http.HandlerFunc(webhook))
// }

// функция webhook — обработчик HTTP-запроса
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// разрешаем только POST-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// установим правильный заголовок для типа данных
	w.Header().Set("Content-Type", "application/json")
	// пока установим ответ-заглушку, без проверки ошибок
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
}
