package action

import (
	"context"

	e "github.com/hidayatullahap/evermos/gateway/entity"
	"github.com/hidayatullahap/evermos/inventory_service/entity"
)

type IGatewayAction interface {
	Order(ctx context.Context, request entity.Inventory) error
	GetInventory(ctx context.Context, request entity.Inventory) (entity.Inventory, error)
}

type GatewayAction struct {
	app *e.App
}

func NewGatewayAction(app *e.App) *GatewayAction {
	return &GatewayAction{app}
}
