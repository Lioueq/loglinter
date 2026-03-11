package basic

import (
	"log/slog"
)

func test3() {
	log := slog.New(slog.Default().Handler())

	log.Info("server started!🚀")                 // want "log message must not contain special characters or emoji"
	log.Error("connection failed!!!")            // want "log message must not contain special characters or emoji"
	log.Warn("warning: something went wrong...") // want "log message must not contain special characters or emoji"
}
