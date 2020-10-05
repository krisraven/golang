package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPizzasHandler(t *testing.T) {
	tt := []struct {
		name				string
		method			string
		input 			*Pizzas
		want				string
		statusCode	int
	}{
		{
			name:				"without pizzas",
			method:			http.MethodGet,
			input:			&Pizzas{},
			want:				"Error: No pizzas found",
			statusCode:	http.StatusNotFound,
		},
		{
			name:		"with pizzas",
			method:	http.MethodGet,
			input: &Pizzas{
				Pizza{
					ID:			1,
					Name:		"Foo",
					Price:	10,
				},
			},
			want:				`[{"id":1,"name":"Foo","price":10}]`,
			statusCode:	http.StatusOK,
		},
		{
			name:				"with bad method",
			method:			http.MethodPost,
			input:			&Pizzas{},
			want:				"Method not allowed",
			statusCode:	http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(tc.method, "/orders", nil)
			responseRecorder := httptest.NewRecorder()

			pizzasHandler{tc.input}.ServeHTTP(responseRecorder, request)

			if responseRecorder.Code != tc.statusCode {
				t.Errorf("Want status '%d', got '%d'", tc.statusCode, responseRecorder.Code)
			}

			if strings.TrimSpace(responseRecorder.Body.String()) != tc.want {
				t.Errorf("Want '%s', got '%s'", tc.want, responseRecorder.Body)
			}
		})
	}
}