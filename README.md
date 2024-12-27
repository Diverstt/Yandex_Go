# Yandex_Go

## Описание
Yandex_Go — это микросервис для выполнения вычислений с использованием обратной польской нотации (RPN). Он предоставляет REST API для обработки математических операций.

## Технологии
- Go
- RPN (обратная польская нотация)

## Установка
```bash
git clone https://github.com/Diverstt/Yandex_Go.git
cd Yandex_Go
go run cmd/main.go 
```
## Тестирование
Тесты
```bash
go test ./...
```
## Тестирование с помощью curl для Win
```bash
1. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3+4\"}" http://localhost:8080/api/v1/calculate
```
- status code = 200
```bash
2. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3+\"}" http://localhost:8080/api/v1/calculate
```
- status code = 422
```bash
3. curl -X POST -H "Content-Type: application/json" -d "{\"expression\": \"3/0\"}" http://localhost:8080/api/v1/calculate
```
- status code = 500

## Структура проекта
```markdown
Yandex_Go/
├── cmd/
│   └── main.go
├── internal/
│   ├── application/
│   │   ├── handler.go
│   │   ├── handler_test.go
│   │   ├── middleware.go
│   │   └── service.go
│   └── infrastructure/
│       ├── config/
│       │   └── config.go
│       └── http/
│           ├── routes.go
│           └── server.go
├── pkg/
│   └── rpn/
│       ├── rpn.go
│       └── rpn_test.go
├── go.mod
└── README.md
```
