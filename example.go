package main

import (
	"log/slog"
	"os"
)

var password, apiKey, token string

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	//1. Лог-сообщения должны начинаться со строчной буквы
	//❌Неправильно
	log.Info("Starting server on port 8080")
	slog.Error("Failed to connect to database")

	//✅Правильно
	log.Info("starting server on port 8080")
	slog.Error("failed to connect to database")

	//2. Лог-сообщения должны быть только на английском языке
	//❌Неправильно
	log.Info("запуск сервера")
	log.Error("ошибка подключения к базе данных")

	//✅Правильно
	log.Info("starting server")
	log.Error("failed to connect to database")

	//3. Лог-сообщения не должны содержать спецсимволы или эмодзи
	//❌Неправильно
	log.Info("server started!🚀")
	log.Error("connection failed!!!")
	log.Warn("warning: something went wrong...")

	//✅Правильно
	log.Info("server started")
	log.Error("connection failed")
	log.Warn("something went wrong")

	//4. Лог-сообщения не должны содержать потенциально чувствительные данные
	//❌Неправильно
	log.Info("user password: " + password)
	log.Debug("api_key=" + apiKey)
	log.Info("token: " + token)

	//✅Правильно
	log.Info("user authenticated successfully")
	log.Debug("api request completed")
	log.Info("token validated")

}
