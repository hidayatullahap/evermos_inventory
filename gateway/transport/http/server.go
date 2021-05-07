package http

import (
	"fmt"
	"log"

	"github.com/hidayatullahap/evermos/gateway/entity"
	"github.com/labstack/echo/v4"
)

const defaultPort = "1313"

type Server struct {
	app *entity.App
	e   *echo.Echo
}

func (s *Server) Start() {
	setupDefaultMiddleware(s.e)
	setupRoutes(s.e, s.app)
	ErrorHandler(s.e)

	port := s.app.Config.HttpServerPort
	if port == "" {
		port = defaultPort
	}

	s.e.Server.Addr = fmt.Sprintf(":%v", port)
	s.e.HideBanner = false

	log.Printf("Http server start at port %s", port)
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%v", port)))
}

func NewHttpServer(app *entity.App) *Server {
	return &Server{
		app: app,
		e:   echo.New(),
	}
}
