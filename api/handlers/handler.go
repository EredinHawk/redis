package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/EredinHawk/redis/cache"
)

// result - содержит в себе результат вычислений.
// Выведен в отдельную переменную для отображения в двух местах в коде
var result int

// Numbers структура, которая содержит в себе числа для суммирования
type Numbers struct {
	N1 int `json:"n1"`
	N2 int `json:"n2"`
}

// WrapHandler это обертка над обработчиком, которая прнимает на входе указатель на экземпляр RedisServer типа и
// возвращает обработчик.
//
// Обработчик, суммирует два числа из тела запроса и возвращает результат вычисления клиенту.
// При этом, если в кэше уже записан результат, то вычисление не будет производится. Обработчик вернет это значение из кэша.
func WrapHandler(redis *cache.RedisServer) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Считывание двух чисел из тела запроса
		numbers, err := scanNumbers(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Формирование из этих двух чисел ключ в кэше
		key := keyGen(numbers)

		// Проверка по ключу наличие данных в кэше
		cache_check, err := redis.CheckValue(key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Если данных в кэше нет, произвести вычисление
		if !cache_check {
			result = sum(numbers)

			// Результат вычисления записать в кэш
			err := redis.SetValue(cache.Cahe{Key: key, Value: result})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Add("cached", "false")
			fmt.Fprintf(w, "%v + %v = %v - данные получены из расчета API", numbers.N1, numbers.N2, result)

		// Если данные в кэше присутствуют, получить их
		} else {
			result, err = redis.GetValue(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Add("cached", "true")
			fmt.Fprintf(w, "%v + %v = %v - данные получены из кэша", numbers.N1, numbers.N2, result)
		}
	}
}

// scanNumbers сканирует из тела HTTP запроса два числа (JSON декодирование во внутреннюю структуру Go)
func scanNumbers(r *http.Request) (*Numbers, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var numbers *Numbers
	err = json.Unmarshal(body, &numbers)
	if err != nil {
		return nil, err
	}

	return numbers, nil
}

// keyGen возвращает строку - ключ, полученная их двух чисел, разделенные нижним подчеркиванием (пример '2_5')
func keyGen(numbers *Numbers) string {
	return fmt.Sprintf("%v_%v", numbers.N1, numbers.N2)
}

// sum суммирует два числа и возвращает результат.
// Вообще sum - это иммитация каких-то дорогостоящих, по части памяти, вычислений.
func sum(numbers *Numbers) int {
	return numbers.N1 + numbers.N2
}
