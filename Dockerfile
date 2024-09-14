# Dockerfile

# Используем официальное минималистичное изображение Golang
FROM golang:1.20-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Копируем все файлы проекта
COPY . .

# Собираем приложение
RUN go build -o main .

# Открываем порт
EXPOSE 8080

# Команда для запуска приложения
CMD ["./main"]
