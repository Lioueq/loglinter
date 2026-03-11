package basic

import (
	"log/slog"
)

func test2() {
	log := slog.New(slog.Default().Handler())

	log.Info("запуск сервера")                    // want "log message must be in English"
	log.Error("ошибка подключения к базе данных") // want "log message must be in English"
}
