package main

import (
	"log"
	"sync"

	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/inventory_service/transport"
	"github.com/hidayatullahap/evermos/pkg/db/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// TestWithLock(db)
	// TestWithoutLock(db)

	t := transport.NewTransport(app)
	wg := sync.WaitGroup{}
	wg.Add(1)

	t.GrpcServer.Start()

	wg.Wait()
}

type Table struct {
	ID int64 `gorm:"id"`
}

func (t Table) TableName() string {
	return "table1"
}

func TestWithLock(db *gorm.DB) {
	wg := sync.WaitGroup{}
	UpdateCurrentProductQty(db, 1, 5)
	t := FindProductQty(db, 1)
	log.Printf("Current product qty before test %+v", t.Qty)

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(loop int) {
			UpdateInventoryWithLock(loop, db)
			wg.Done()
		}(i)
	}

	wg.Wait()

	t = FindProductQty(db, 1)
	log.Printf("Current product qty after test %+v", t.Qty)
}

func TestWithoutLock(db *gorm.DB) {
	wg := sync.WaitGroup{}
	UpdateCurrentProductQty(db, 1, 5)

	t := FindProductQty(db, 1)
	log.Printf("Current product qty before test %+v", t.Qty)

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(loop int) {
			UpdateInventoryWithoutLock(loop, db)
			wg.Done()
		}(i)
	}

	wg.Wait()

	t = FindProductQty(db, 1)
	log.Printf("Current product qty after test %+v", t.Qty)
}

func UpdateInventoryWithLock(i int, db *gorm.DB) error {
	var t Inventory
	productID := int64(1)

	tx := db.Begin()
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("product_id = ?", productID).
		Find(&t)

	log.Printf("update loop i: %v, current qty %v", i, t.Qty)
	if t.Qty > 0 {
		UpdateProductQty(tx, 1, 1)
	}

	tx.Commit()
	return nil
}

func UpdateInventoryWithoutLock(i int, db *gorm.DB) error {
	var t Inventory

	productID := int64(1)

	tx := db.Begin()
	tx.Where("product_id = ?", productID).
		Find(&t)

	log.Printf("update i %v, current qty %v", i, t.Qty)
	if t.Qty > 0 {
		UpdateProductQty(tx, productID, 1)
	}

	tx.Commit()
	log.Printf("done")
	return nil
}

func UpdateProductQty(tx *gorm.DB, productID int64, qty int64) {
	tx.Model(&Inventory{}).Where("product_id = ? ", productID).UpdateColumn("qty", gorm.Expr("qty - ?", qty))
}

func UpdateCurrentProductQty(tx *gorm.DB, productID int64, qty int64) {
	tx.Model(&Inventory{}).Where("product_id = ? ", productID).UpdateColumn("qty", qty)
}

func FindProductQty(db *gorm.DB, productID int64) (i Inventory) {
	db.Where("product_id = ?", productID).Find(&i)

	return
}
