package handlers

import (
	"fmt"
	"net/http"
)

type Handlers struct{}

func NeHandlers() *Handlers {
	return &Handlers{}
}

// Регистрируем обработчик для всех запросов
func (h Handlers) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad reqest!")
		return
	}
}

func (h Handlers) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "By, Go!")
}
