package repo

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/pkg/db/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func testFailed(t *testing.T, err error) {
	t.Errorf("test failed, error occured %v", err)
}

func loadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitInventoryRepo() IInventoryRepo {
	loadEnv()
	app := entity.NewApp()
	db, err := mysql.Connect(app.Config.MysqlConnection)
	if err != nil {
		panic(err)
	}

	app.GormDb = db

	return NewInventoryRepo(app)
}

// setupInventoryRow will create or update inventory with random product id
func setupInventoryRow(repo IInventoryRepo) int64 {
	rand.Seed(time.Now().UnixNano())
	min := 10
	max := 500
	id := int64(rand.Intn(max-min+1) + min)
	qty := int64(5)

	inv := entity.Inventory{
		ProductID: id,
		Quantity:  qty,
	}

	cInv, err := repo.FindProductInventory(id)
	if err != nil {
		panic(err)
	}

	if cInv.ID == 0 {
		errCreate := repo.CreateProductInventory(inv)
		if errCreate != nil {
			panic(errCreate)
		}
	} else {
		errUpdate := repo.UpdateProductQty(id, qty)
		if errUpdate != nil {
			panic(errUpdate)
		}
	}

	return id
}

// TestInventoryRepo_UpdateInventoryQty will test race condition on reduce inventory qty
// test will spawn 50 go routine to simulate 50 person buy at same product simultaneously
// product id will only available with 5 qty
// the expected result is product qty will be 0 with no negative value
//
// cmd to run:
// go test ./inventory_service/repo -tags integration -count 1 -v -run=^TestInventoryRepo_UpdateInventoryQty
func TestInventoryRepo_UpdateInventoryQty(t *testing.T) {
	repo := InitInventoryRepo()
	id := setupInventoryRow(repo)

	log.Printf("testing product id: %v", id)
	log.Printf("fetching inventory row of id %v ...", id)
	inv, err := repo.FindProductInventory(id)
	if err != nil {
		testFailed(t, err)
	}

	log.Printf("success, qty of row id %v = %v", id, inv.Quantity)

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(loop int) {
			repo.UpdateInventoryQty(id, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	inv, err = repo.FindProductInventory(id)
	if err != nil {
		testFailed(t, err)
	}

	log.Printf("test finished, qty of row id %v = %v", id, inv.Quantity)
	assert.Equal(t, int64(0), inv.Quantity, "Qty should 0")
}

// TestInventoryRepo_UpdateInventoryQty will test race condition on reduce inventory qty
// test will spawn 50 go routine to simulate 50 person buy at same product simultaneously
// product id will only available with 5 qty
// the expected result is product qty will be negative value
//
// cmd to run:
// go test ./inventory_service/repo -tags integration -count 1 -v -run=^TestInventoryRepo_UpdateInventoryQtyRaceCondition
func TestInventoryRepo_UpdateInventoryQtyRaceCondition(t *testing.T) {
	repo := InitInventoryRepo()
	id := setupInventoryRow(repo)

	log.Printf("testing product id: %v", id)
	log.Printf("fetching inventory row of id %v ...", id)
	inv, err := repo.FindProductInventory(id)
	if err != nil {
		testFailed(t, err)
	}

	log.Printf("success, qty of row id %v = %v", id, inv.Quantity)

	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func(loop int) {
			repo.UpdateInventoryQtyRaceCondition(id, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	inv, err = repo.FindProductInventory(id)
	if err != nil {
		testFailed(t, err)
	}

	log.Printf("test finished, qty of row id %v = %v", id, inv.Quantity)

	if inv.Quantity > 0 {
		t.Errorf("expected result is negative value")
	}
}
