package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"matchmaker/internal/models"
	"net/http"
	"time"
)

func (s *Server) UsrHandler(c *gin.Context) {
	var usr models.User
	if err := c.BindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Debug(fmt.Sprintf("new user accepted: %+v", usr))
	usr.TimeOfJoin = time.Now()
	go s.mm.Process(&usr)
	c.JSON(http.StatusOK, usr)
}
