package transport

import (
	"github.com/hidayatullahap/evermos/gateway/entity"
	"github.com/hidayatullahap/evermos/gateway/transport/http"
)

type Transport struct {
	HttpServer *http.Server
}

func NewTransport(app *entity.App) *Transport {
	return &Transport{
		http.NewHttpServer(app),
	}
}
