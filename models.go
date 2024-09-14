// models.go
package main

import (
	"time"
)

type AccessRequest struct {
	ID          int       `json:"id"` // Добавляем поле ID
	UserID      string    `json:"user_id"`
	System      string    `json:"system"`
	Status      string    `json:"status"`
	RequestedAt time.Time `json:"requested_at"` // Добавляем поле RequestedAt
}

// Функция для создания нового запроса на доступ
func CreateAccessRequest(userID, system string) (int, error) {
	query := `INSERT INTO access_requests (user_id, system, status) VALUES ($1, $2, $3) RETURNING id`
	var requestID int
	err := db.QueryRow(query, userID, system, "pending").Scan(&requestID)
	if err != nil {
		return 0, err
	}
	return requestID, nil
}

// Функция для получения всех запросов на доступ
func GetAccessRequests() ([]AccessRequest, error) {
	rows, err := db.Query(`SELECT id, user_id, system, status, requested_at FROM access_requests`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []AccessRequest
	for rows.Next() {
		var request AccessRequest
		err := rows.Scan(&request.ID, &request.UserID, &request.System, &request.Status, &request.RequestedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}
