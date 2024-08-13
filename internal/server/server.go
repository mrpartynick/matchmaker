package server

import (
	"github.com/gin-gonic/gin"
	"matchmaker/config"
)

type Server struct {
	cfg *config.Config
	r   *gin.Engine
}

func New(cfg *config.Config) *Server {
	s := Server{cfg: cfg}
	s.r = gin.Default()
	return &s
}

func (s *Server) Run() {
	s.r.Run(s.cfg.Port)
}
