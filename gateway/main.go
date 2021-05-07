package main

import (
	"log"
	"sync"

	"github.com/hidayatullahap/evermos/gateway/entity"
	"github.com/hidayatullahap/evermos/gateway/transport"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := entity.NewApp()

	t := transport.NewTransport(app)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	t.HttpServer.Start()

	wg.Wait()
}
