package basic

import (
	"log/slog"
)

var api_key string

func test5() {
	log := slog.New(slog.Default().Handler())

	log.Info("Starting server on port 8080🚀") // want "log message must start with a lowercase letter" "log message must not contain special characters or emoji"
	slog.Error("Ошибка" + api_key)            // want "log message must start with a lowercase letter" "log message must be in English" "log message may contain sensitive data"
}
