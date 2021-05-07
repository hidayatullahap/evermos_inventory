package repo

import (
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"gorm.io/gorm"
)

type InventoryRepo struct {
	db *gorm.DB
}

type IInventoryRepo interface {
	UpdateInventoryQty(productID int64, qty int64) error
}

func (r *InventoryRepo) UpdateInventoryQty(productID int64, qty int64) error {

	return nil
}

func NewInventoryRepo(app *entity.App) *InventoryRepo {
	return &InventoryRepo{
		db: app.GormDb,
	}
}
