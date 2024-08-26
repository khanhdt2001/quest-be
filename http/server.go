package http

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewHTTP(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}
func (a *Server) Run(address string) error {
	return a.router.Run(address)
}
