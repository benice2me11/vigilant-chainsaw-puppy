// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Инициализация базы данных
	InitDB()

	// Создание таблицы, если она не существует
	CreateTable()

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/request-access", requestAccess).Methods("POST")
	r.HandleFunc("/get-access-requests", getAccessRequests).Methods("GET")

	// Запуск сервера
	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
