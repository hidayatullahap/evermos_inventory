package repo

import (
	"strconv"

	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InventoryRepo struct {
	db *gorm.DB
}

type IInventoryRepo interface {
	UpdateInventoryQty(productID int64, qty int64) error
}

func (r *InventoryRepo) UpdateInventoryQty(productID int64, qty int64) error {
	var inv entity.Inventory
	var err error

	tx := r.db.Begin()
	tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("product_id = ?", productID).
		Find(&inv)

	if inv.Quantity >= qty {
		err = tx.Model(&inv).Where("product_id = ? ", productID).UpdateColumn("quantity", gorm.Expr("quantity - ?", qty)).Error
	} else {
		err = errors.InvalidArgument("can't buy product, available quantity is " + strconv.FormatInt(inv.Quantity, 10))
	}

	tx.Commit()

	return err
}

func NewInventoryRepo(app *entity.App) *InventoryRepo {
	return &InventoryRepo{
		db: app.GormDb,
	}
}
