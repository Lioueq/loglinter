package basic

import (
	"log/slog"
)

func test6() {
	log := slog.New(slog.Default().Handler())

	log.Info("starting server on port 8080")    // ok
	slog.Error("failed to connect to database") // ok
	log.Info("starting server")                 // ok
	log.Error("failed to connect to database")  // ok
	log.Info("server started")                  // ok
	log.Error("connection failed")              // ok
	log.Warn("something went wrong")            // ok
	log.Info("user authenticated successfully") // ok
	log.Debug("api request completed")          // ok
	log.Info("token validated")                 // ok
}
