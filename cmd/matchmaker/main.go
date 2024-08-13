package main

import (
	"log"
	"log/slog"
	"matchmaker/config"
	"matchmaker/internal/matchmaker"
	"matchmaker/internal/server"
	"matchmaker/pkg/logging"
)

func main() {
	cfg := config.Load("config.yaml")

	logfile := logging.MustLoad()
	defer logfile.Close()

	mm := matchmaker.NewMatchMaker(cfg)

	s := server.New(cfg, mm)
	log.Println("Running server")
	s.Run()
	slog.Debug("Service stopped")
}
