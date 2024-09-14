// handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func requestAccess(w http.ResponseWriter, r *http.Request) {
	var newRequest struct {
		UserID string `json:"user_id"`
		System string `json:"system"`
	}
	if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
		http.Error(w, "Неверные данные", http.StatusBadRequest)
		return
	}

	// Сохраняем запрос в базу данных
	requestID, err := CreateAccessRequest(newRequest.UserID, newRequest.System)
	if err != nil {
		http.Error(w, "Ошибка при создании запроса", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Запрос на доступ создан с ID: %d\n", requestID)
}

func getAccessRequests(w http.ResponseWriter, r *http.Request) {
	// Получаем запросы из базы данных
	requests, err := GetAccessRequests()
	if err != nil {
		http.Error(w, "Ошибка получения запросов", http.StatusInternalServerError)
		return
	}

	// Возвращаем запросы в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}
