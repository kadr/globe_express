
# Этап сборки
FROM golang:1.26.5-alpine AS builder


ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64 

# Устанавливаем необходимые системные зависимости
RUN apk add --no-cache git ca-certificates gcc  \
  musl-dev \
  libc-dev \
  build-base \
  linux-headers && \
  update-ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go mod и sum файлы
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

RUN go build -ldflags="-w -s" -o app/bin/travaler-app ./cmd/travaler-service/main.go

# Финальный этап
FROM alpine:latest AS final

# Добавляем CA сертификаты для HTTPS запросов
RUN apk --no-cache add ca-certificates tzdata

RUN addgroup -g 1000 -S appuser && \
  adduser -u 1000 -S appuser -G appuser

WORKDIR /app

# Бинарник и config/users.json для логина (путь относительно cwd: config/users.json)
COPY --from=builder /app/bin/travaler-app /app/bin/travaler-app

# Смена владельца
RUN chown -R appuser:appuser /app

# Переключение на непривилегированного пользователя
USER appuser

# Открываем порт
EXPOSE 80

# Запускаем приложение
CMD ["/app/bin/travaler-app"]
