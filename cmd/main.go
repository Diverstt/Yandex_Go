package main

import (
	"github.com/Diverstt/Yandex_Go/internal/application"
	"github.com/Diverstt/Yandex_Go/internal/infrastructure/http"
)

func main() {
	// Инициализация сервиса
	service := &application.CalcService{}

	// Инициализация обработчика
	handler := application.NewCalcHandler(service)

	// Запуск сервера
	http.StartServer(*handler)
}
