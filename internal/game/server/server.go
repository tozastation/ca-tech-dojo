package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tozastation/ca-tech-dojo/internal/game/user"
	"log"
	"os"
)

var (
	_serverPort = os.Getenv("PORT")
)

type Server struct {
	Router *gin.Engine
	UserCtrl user.IController
}

func NewServer(userCtrl user.IController) *Server {
	result := &Server{}
	result.Router = gin.Default()
	result.UserCtrl = userCtrl
	_ = result.SetUpRouting()
	return result
}

func (s *Server) SetUpRouting() error {
	//s.Router.Use(middleware.AuthMiddleware())
	v1 := s.Router.Group("/v1")
	{
		v1.GET("/user/get", s.UserCtrl.Get)
		v1.POST("/user/create", s.UserCtrl.Create)
		v1.PUT("/user/update", s.UserCtrl.Update)
	}
	return nil
}

func (s *Server) Run() error {
	log.Printf("info: listen on %s", _serverPort)
	return s.Router.Run(_serverPort)
}