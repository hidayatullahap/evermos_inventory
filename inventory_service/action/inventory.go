package action

import (
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/inventory_service/repo"
)

type IInventoryAction interface {
	UpdateInventoryQty(productID int64, qty int64) (int64, error)
	GetInventory(productID int64) (entity.Inventory, error)
}

type InventoryAction struct {
	inventoryRepo repo.IInventoryRepo
}

func (r *InventoryAction) UpdateInventoryQty(productID int64, qty int64) (int64, error) {
	var availableQty int64
	err := r.inventoryRepo.UpdateInventoryQty(productID, qty)

	return availableQty, err
}

func (r *InventoryAction) GetInventory(productID int64) (entity.Inventory, error) {
	inv, err := r.inventoryRepo.FindProductInventory(productID)

	return inv, err
}

func NewInventoryAction(app *entity.App) IInventoryAction {
	return &InventoryAction{
		inventoryRepo: repo.NewInventoryRepo(app),
	}
}
