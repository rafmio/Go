# Используем официальный многоплатформенный образ Go
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем исходный код
COPY primitive-http-server.go .

# Собираем бинарник
RUN go build -o server primitive-http-server.go

# Второй этап: минимальный образ
FROM alpine:latest AS final

# Установим tzdata для корректного отображения времени
RUN apk --no-cache add ca-certificates tzdata

# Установим рабочую директорию
WORKDIR /root/

# Копируем бинарник из builder-образа
COPY --from=builder /app/server .

# Открываем порт
EXPOSE 8080

# Запускаем сервер
CMD ["./server"]
