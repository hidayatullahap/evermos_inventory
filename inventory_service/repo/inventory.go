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
	UpdateInventoryQtyRaceCondition(productID int64, qty int64) error
	UpdateProductQty(productID int64, qty int64) error
	FindProductInventory(productID int64) (entity.Inventory, error)
	CreateProductInventory(inventory entity.Inventory) error
}

// UpdateInventoryQty will check the available quantity before reduce to current inventory row
// to prevent race condition select will locked until update is executed
// update only occur when qty is available
// transaction will always committed to prevent deadlock
func (r *InventoryRepo) UpdateInventoryQty(productID int64, qty int64) error {
	var inv entity.Inventory
	var err error

	if qty < 0 {
		return errors.InvalidArgument("positive number is required")
	}

	tx := r.db.Begin()
	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("product_id = ?", productID).
		Find(&inv).
		Error

	if inv.ProductID == 0 {
		return errors.NotFound("product not found")
	}

	if err != nil {
		return err
	}

	if inv.Quantity >= qty {
		err = tx.Model(&inv).Where("product_id = ? ", productID).UpdateColumn("quantity", gorm.Expr("quantity - ?", qty)).Error
	} else {
		err = errors.InvalidArgument("can't buy product, available quantity is " + strconv.FormatInt(inv.Quantity, 10))
	}

	tx.Commit()

	return err
}

// UpdateInventoryQtyRaceCondition will check the available quantity before reduce to current inventory row
// but it doesnt have row lock this function will result on negative inventory number because
// inventory qty check will overlap with other user
func (r *InventoryRepo) UpdateInventoryQtyRaceCondition(productID int64, qty int64) error {
	var inv entity.Inventory
	var err error

	if qty < 0 {
		return errors.InvalidArgument("positive number is required")
	}

	err = r.db.Where("product_id = ?", productID).Find(&inv).Error
	if err != nil {
		return err
	}

	if inv.Quantity >= qty {
		err = r.db.Model(&inv).Where("product_id = ? ", productID).UpdateColumn("quantity", gorm.Expr("quantity - ?", qty)).Error
	} else {
		err = errors.InvalidArgument("can't buy product, available quantity is " + strconv.FormatInt(inv.Quantity, 10))
	}

	return err
}

// UpdateProductQty this function is to update exact qty
func (r *InventoryRepo) UpdateProductQty(productID int64, qty int64) error {
	err := r.db.Model(&entity.Inventory{}).Where("product_id = ? ", productID).UpdateColumn("quantity", qty).Error
	return err
}

func (r *InventoryRepo) FindProductInventory(productID int64) (i entity.Inventory, err error) {
	err = r.db.Where("product_id = ?", productID).Find(&i).Error

	if productID == 0 {
		err = errors.NotFound("product not found")
	}

	return
}

func (r *InventoryRepo) CreateProductInventory(inventory entity.Inventory) error {
	err := r.db.Create(&inventory).Error
	return err
}

func NewInventoryRepo(app *entity.App) *InventoryRepo {
	return &InventoryRepo{
		db: app.GormDb,
	}
}
