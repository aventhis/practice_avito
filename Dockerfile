# Этап 1: билд
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum, устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код и собираем бинарник
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server ./cmd/pvz_service/main


# Этап 2: минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем только бинарник
COPY --from=builder /app/server .

EXPOSE 8080

# Запуск
CMD ["./server"]


#Ты разделяешь Dockerfile на этапы (stage):
#
#Stage 1 — сборка
#Используешь образ с Go (golang:1.21-alpine)
#Ставишь зависимости, собираешь приложение
#
#Stage 2 — финальный образ
#Используешь минимальный образ (alpine или scratch)
#Копируешь только готовый бинарник