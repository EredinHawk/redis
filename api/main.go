package main

import (
	"fmt"
	"log"
	"net/http"
)

// Запускает локально http сервер собработчиком, который возвращает просто строку
func main() {
	
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Data from API")
	})

	log.Println("Сервер запущен по адресу 127.0.0.1:8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
