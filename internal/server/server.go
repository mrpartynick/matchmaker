package server

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"matchmaker/config"
	"matchmaker/internal/matchmaker"
)

type Server struct {
	cfg *config.Config
	mm  *matchmaker.MatchMaker
	r   *gin.Engine
}

func New(cfg *config.Config, mm *matchmaker.MatchMaker) *Server {
	s := Server{
		cfg: cfg,
		mm:  mm,
	}
	s.r = gin.Default()
	s.r.POST("/users", s.UsrHandler)
	return &s
}

func (s *Server) Run() {
	slog.Debug("Running server...")
	s.r.Run(s.cfg.Host + ":" + s.cfg.Port)
}
