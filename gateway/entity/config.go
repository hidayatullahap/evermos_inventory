package entity

import (
	"os"
)

type Config struct {
	Services       Services
	HttpServerPort string
}

type Services struct {
	InventoryHost string
}

func NewConfig() Config {
	var c Config

	c.Services = Services{
		InventoryHost: os.Getenv("HOST_GRPC_INVENTORY_SERVICE"),
	}

	c.HttpServerPort = os.Getenv("HTTP_PORT")

	return c
}
