package main

import (
	"matchmaker/config"
	"matchmaker/internal/matchmaker"
	"matchmaker/internal/server"
	"matchmaker/pkg/logging"
)

// TODO: Добавить таймер понижения толерантности
// TODO: Разобраться с логированием

func main() {
	cfg := config.Load("config.yaml")

	logfile := logging.MustLoad()
	defer logfile.Close()

	mm := matchmaker.NewMatchMaker(cfg)

	s := server.New(cfg, mm)
	s.Run()
}
