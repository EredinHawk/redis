package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
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
		req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(`{"n1":2,"n2":2}`)))
		if err != nil {
			t.Fatalf("%v\n", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HandlerSUM)
		handler.ServeHTTP(rr, req)

		result := rr.Result().Header.Get("cached")
		
		// Проверить соответствие ожидаемым результатам
		if result != v.Excepted {
			t.Fatalf("Excepted - %v, but received - %v", v.Excepted, result)
		}

		fmt.Println(rr.Body.String())
	}

}
