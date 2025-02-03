package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// TestHandlerSUM - юнит тест, который дважды выполняет http запрос к API.
// Успешным тест будет только тогда, когда результат первого запроса будет расчитан API, а второй получен из кэша.
// Для проверки ожидаемых данных проверяется значение заголовка http ответа 'cached'.
func TestHandlerSUM(t *testing.T) {
	// Задать ожидаемые параметры

	testCases := []struct {
		Excepted string
	}{
		{"false"}, {"true"},
	}

	// Дважды вызвать обработчик: без кэша и с кэшем
	for _, v := range testCases {
		// Выполннить http запрос с методом Get к API
		response, err := http.Post("http://localhost:8090/", "application/json", bytes.NewReader([]byte(`{"n1":2,"n2":2}`)))
		if err != nil {
			t.Fatalf("%v\n", err)
		}

		// Проверка ожидаемых данных
		cache_status := response.Header.Get("cached")
		if cache_status != v.Excepted {
			t.Errorf("excepted - %v, but recived - %v", v.Excepted, cache_status)
		}

		// Для информации вывести результат в консоль
		body, err := io.ReadAll(response.Body)
		if err != nil {
			t.Errorf("%v\n", err)
		}

		fmt.Println(string(body))
	}

}
