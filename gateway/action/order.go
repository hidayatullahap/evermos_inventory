package action

import (
	"context"

	"github.com/hidayatullahap/evermos/inventory_service/entity"

	"github.com/hidayatullahap/evermos/pkg/grpc"
	pb "github.com/hidayatullahap/evermos/pkg/proto/inventory"
)

func (a *GatewayAction) Order(ctx context.Context, request entity.Inventory) error {
	conn, err := grpc.Dial(a.app.Config.Services.InventoryHost)
	if err != nil {
		return err
	}

	defer conn.Close()

	pbReq := &pb.UpdateQtyRequest{
		ProductId: request.ProductID,
		Qty:       request.Quantity,
	}

	_, err = pb.NewInventoryClient(conn).UpdateInventoryQty(ctx, pbReq)
	if err != nil {
		return err
	}

	return nil
}
