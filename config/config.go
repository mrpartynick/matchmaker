package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}

type Env struct {
	Env string `yaml:"env" env-default:"debug"`
}

type Matchmaker struct {
	GroupSize        int `yaml:"group_size" env-default:"5"`
	SkillTolerance   int `yaml:"skill_tolerance" env-default:"2"`
	LatencyTolerance int `yaml:"latency_tolerance" env-default:"100"`
}

type Config struct {
	Server     `yaml:"server"`
	Env        `yaml:"env"`
	Matchmaker `yaml:"matchmaker"`
}

func Load(path string) *Config {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		slog.Error(fmt.Sprintf("config load error: %v. Loading default", err))
		setDefault(&cfg)
	}
	return &cfg
}

func setDefault(cfg *Config) {
	s := Server{
		Host: "localhost",
		Port: "8080",
	}
	e := Env{Env: "debug"}
	m := Matchmaker{
		GroupSize:        5,
		SkillTolerance:   2,
		LatencyTolerance: 100,
	}
	cfg.Server = s
	cfg.Env = e
	cfg.Matchmaker = m
}
