package main

import (
	"log"
	"sync"

	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/inventory_service/transport"
	"github.com/hidayatullahap/evermos/pkg/db/mysql"
	"github.com/joho/godotenv"
)

type Inventory struct {
	ID        int64 `gorm:"id"`
	ProductID int64 `gorm:"product_id"`
	Qty       int64 `gorm:"qty"`
}

func (t Inventory) TableName() string {
	return "inventories"
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := entity.NewApp()

	db, err := mysql.Connect(app.Config.MysqlConnection)
	if err != nil {
		log.Fatalf("failed to connect database")
	}

	app.GormDb = db

	t := transport.NewTransport(app)

	wg := sync.WaitGroup{}
	wg.Add(1)

	t.GrpcServer.Start()

	wg.Wait()
}
