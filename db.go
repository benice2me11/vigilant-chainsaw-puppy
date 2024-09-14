// db.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Глобальная переменная для подключения к базе данных
var db *sql.DB

// Функция для инициализации подключения к базе данных
func InitDB() {
	connStr := "user=admin password=password dbname=access_db sslmode=disable host=localhost"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке подключения к базе данных: %v", err)
	}
	fmt.Println("Успешно подключились к базе данных!")
}

// Функция для создания таблицы, если она не существует
func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS access_requests (
		id SERIAL PRIMARY KEY,
		user_id TEXT NOT NULL,
		system TEXT NOT NULL,
		status TEXT NOT NULL,
		requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	fmt.Println("Таблица access_requests успешно создана или уже существует.")
}
