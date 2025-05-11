package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	// описываем ожидаемое тело ответа при успешном запросе
	successBody := `{
        "response": {
            "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
    }`

	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: successBody},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			//func httptest.NewRequest(method string, target string, body io.Reader) *http.Request
			//NewRequest возвращает новый входящий запрос сервера (*http.Request), подходящий для передачи в [http.Handler] для тестирования.
			//method пустой означает "GET"
			//target - запрашиваемый URL, "эндпойнт" (?)
			//body (текст запроса) может быть равно nil.
			//Если body имеет тип *bytes.Reader, *strings.Reader или *bytes.Buffer, свойство Request.ContentLength считается установленным
			r := httptest.NewRequest(tc.method, "/", nil)

			//func httptest.NewRecorder() *httptest.ResponseRecorder
			//NewRecorder возвращает инициализированный [ResponseRecorder] - регистратор(?) ответов
			//Получается ResponseRecorder это "создатель" интерфейса http.ResponseWriter, "таблицы" ответа сервера
			w := httptest.NewRecorder()

			// вызовем хендлер как обычную функцию, без запуска самого сервера и передадим в нее
			// готовые http.ResponseWriter и *http.Request
			webhook(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			// проверим корректность полученного тела ответа, если мы его ожидаем
			if tc.expectedBody != "" {
				// assert.JSONEq помогает сравнить две JSON-строки
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
