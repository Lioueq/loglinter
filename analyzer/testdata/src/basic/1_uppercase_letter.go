package basic

import (
	"log/slog"
)

func test1() {
	log := slog.New(slog.Default().Handler())

	log.Info("Starting server on port 8080")    // want "log message must start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log message must start with a lowercase letter"
}
