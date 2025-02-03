package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/EredinHawk/redis/api/handlers"
	"github.com/EredinHawk/redis/cache"
)

func main() {
	server := constructServer()

	fmt.Println("Сервер localhost:8090 запущен и прослушивает входящие запросы...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ошибка при запуске сервера (%v)", err)
	}
}

// constructServer возвращает инициализированный сервер типа *http.Server
func constructServer() *http.Server {
	router := http.NewServeMux()
	redis := cache.NewRedisServer()
	router.HandleFunc("POST /", api.WrapHandler(redis))

	server := &http.Server{
		Addr:         "localhost:8090",
		Handler:      router, // HTTP мультиплексор, или по другому роутер
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server
}
