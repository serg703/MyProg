package main

// некоторые импорты нужны для проверки
import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<html>
	<head>
	  <title> 
		Тест приложения в контейнере 
	  </title>
	</head>
  <body>
  <h1>Тест</h1>
  Просто текст
  <p>Текст в параграфе</p>
  </body>
  </html>`))
}

func main() {
	// Регистрируем обработчик для пути "/"
	http.HandleFunc("/get", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
