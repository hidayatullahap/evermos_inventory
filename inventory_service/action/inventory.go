package action

import (
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/inventory_service/repo"
)

type IInventoryAction interface {
	UpdateInventoryQty(productID int64, qty int64) error
}

type InventoryAction struct {
	userRepo repo.IInventoryRepo
}

func (r *InventoryAction) UpdateInventoryQty(productID int64, qty int64) error {

	return nil
}

func NewInventoryAction(app *entity.App) IInventoryAction {
	return &InventoryAction{
		userRepo: repo.NewInventoryRepo(app),
	}
}
