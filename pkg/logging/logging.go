package logging

import (
	"log"
	"log/slog"
	"os"
	"time"
)

func MustLoad() *os.File {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		slog.Debug("Logs dir are not exists. Creating new...")
		if err := os.Mkdir("logs", 0755); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
		slog.Debug("Logs dir created")
	}

	date := time.Now().Format("2006-01-02")

	if _, err := os.Stat("logs/" + date + ".log"); os.IsNotExist(err) {
		slog.Debug("Creating new log file")
		f, err := os.Create("logs/" + date + ".log")
		if err != nil {
			log.Fatalf("Failed to create log file: %v", err)
		}
		log.SetOutput(f)
		return f
	}

	f, err := os.Open("logs/" + date + ".log")
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(f)
	return f
}
