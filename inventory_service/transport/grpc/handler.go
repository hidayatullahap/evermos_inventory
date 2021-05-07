package grpc

import (
	"context"

	"github.com/hidayatullahap/evermos/inventory_service/action"
	"github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/pkg/proto/inventory"
)

type Handler struct {
	app             *entity.App
	inventoryAction action.IInventoryAction
}

func (h *Handler) UpdateInventoryQty(ctx context.Context, request *inventory.UpdateQtyRequest) (*inventory.NoResponse, error) {
	_, err := h.inventoryAction.UpdateInventoryQty(request.ProductId, request.Qty)

	return &inventory.NoResponse{}, err
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app:             app,
		inventoryAction: action.NewInventoryAction(app),
	}
}
