package action

import (
	"context"

	"github.com/hidayatullahap/evermos/inventory_service/entity"

	"github.com/hidayatullahap/evermos/pkg/grpc"
	pb "github.com/hidayatullahap/evermos/pkg/proto/inventory"
)

func (a *GatewayAction) GetInventory(ctx context.Context, request entity.Inventory) (entity.Inventory, error) {
	var inv entity.Inventory
	conn, err := grpc.Dial(a.app.Config.Services.InventoryHost)
	if err != nil {
		return inv, err
	}

	defer conn.Close()

	pbReq := &pb.GetInventoryRequest{
		ProductId: request.ProductID,
	}

	pbRes, err := pb.NewInventoryClient(conn).GetInventory(ctx, pbReq)
	if err != nil {
		return inv, err
	}

	inv.ID = pbRes.Id
	inv.Quantity = pbRes.Qty
	inv.ProductID = pbRes.ProductId

	return inv, nil
}
