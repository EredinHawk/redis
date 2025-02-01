package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/EredinHawk/redis/cache"
)

// Запускает две итерации запроса к API: без кэша и с кэшем
func main() {
	cacheServer := cache.NewRedisServer()
	
	for i := 0; i < 2; i++ {

		result, err := cacheServer.CheckValue("key")
		ErrorCheck(err)

		if !result {

			httpResult := MustRequest()
			err := cacheServer.SetValue(cache.Cahe{Key: "key", Value: httpResult})
			ErrorCheck(err)

			fmt.Printf("Value from http - %v\n", httpResult)
		} else {

			value, err := cacheServer.GetValue()
			ErrorCheck(err)

			fmt.Printf("Value from cache - %v\n", value)
		}
	}
}

// MustRequest выполняет http запрос к API и выводит в консоль тело ответа
func MustRequest() string {
	resp, err := http.Get("http://127.0.0.1:8090/")
	ErrorCheck(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	ErrorCheck(err)
	
	return string(body)
}

// Error - обертка над проверкой ошибки
func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
