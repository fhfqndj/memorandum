# Используем официальный образ Go в качестве базового
FROM golang:1.23-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

COPY entrypoint.sh entrypoint.sh

# Копируем все файлы проекта в контейнер
COPY . .

ENTRYPOINT ["/app/backend/entrypoint.sh"]
# Запускаем приложение с использованием go run
CMD ["air", "-c", ".air.toml"]