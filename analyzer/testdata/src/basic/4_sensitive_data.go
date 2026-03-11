package basic

import (
	"log/slog"
)

var password, apiKey, token, creds string

func test4() {
	log := slog.New(slog.Default().Handler())

	log.Info("user password: " + password) // want "log message may contain sensitive data"
	log.Debug("api_key=" + apiKey)         // want "log message may contain sensitive data"
	log.Info("token: " + token)            // want "log message may contain sensitive data"
	log.Info("token: " + creds)            // want "log message may contain sensitive data"
}
