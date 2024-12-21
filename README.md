# Yandex_Go

## Описание
Yandex_Go — это микросервис для выполнения вычислений с использованием обратной пользовательской нотации (RPN). Он предоставляет REST API для обработки математических операций.

## Технологии
- Go
- RPN (обратная пользовательская нотация)

## Установка
1. Ручная вариация:
   ```bash
   git clone https://github.com/Diverstt/Yandex_Go.git
   cd Yandex_Go
   go mod download
   go run cmd/main.go 

2. Автоматическая вариация:
    ```bash
    ./start.sh

## Тестирование
go test ./...

## Тестирование с помощью curl для Win(Отличие от Unix систем в отсуствии поддержи переносов строк)
1. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3+4\"}" http://localhost:8080/api/v1/calculate
- status code = 200
2. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3+\"}" http://localhost:8080/api/v1/calculate
- status code = 422
3. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3/0\"}" http://localhost:8080/api/v1/calculate
- status code = 500

## Структура проекта
Yandex_Go/
├── cmd/
│ └── main.go # Точка входа в приложение
├── internal/
│ ├── application/
│ │ ├── handler.go # Обработчики HTTP-запросов
│ │ ├── handler_test.go # Тестирование обработчиков
│ │ ├── middleware.go # Middleware для обработки запросов
│ │ └── service.go # Бизнес-логика
│ └── infrastructure/
│ ├── config/
│ │ └── config.go # Конфигурации приложения
│ └── http/
│ ├── routes.go # Определение маршрутов
│ └── server.go # Запуск HTTP-сервера
├── pkg/
│ └── rpn/
│ └── rpn.go # Логика для работы с RPN
│ └── rpn_test.go # Тестирование логики RPN
├── calc_service
├── go.mod # Управление зависимостями
└── README.md # Документация проекта