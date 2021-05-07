package http

import (
	netHttp "net/http"
	"strconv"

	"github.com/hidayatullahap/evermos/gateway/action"
	"github.com/hidayatullahap/evermos/gateway/entity"
	i "github.com/hidayatullahap/evermos/inventory_service/entity"
	"github.com/hidayatullahap/evermos/pkg/errors"
	"github.com/hidayatullahap/evermos/pkg/http"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	app    *entity.App
	action action.IGatewayAction
}

func (h *Handler) Order(c echo.Context) error {
	var req i.Inventory
	err := http.BindAndValidate(c, &req)
	if err != nil {
		return err
	}

	if req.ProductID == 0 {
		return errors.InvalidArgument("product id is required")
	}

	err = h.action.Order(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(netHttp.StatusCreated)
}
func (h *Handler) GetInventory(c echo.Context) error {
	var req i.Inventory
	id, _ := strconv.ParseInt(c.Param("productID"), 10, 64)
	req.ProductID = id

	pbRes, err := h.action.GetInventory(c.Request().Context(), req)
	if err != nil {
		return err
	}

	inv := i.Inventory{
		ID:        pbRes.ID,
		ProductID: pbRes.ProductID,
		Quantity:  pbRes.Quantity,
	}

	data := bson.M{"data": inv}
	return c.JSON(netHttp.StatusOK, data)
}

func NewHandler(app *entity.App) *Handler {
	return &Handler{
		app:    app,
		action: action.NewGatewayAction(app),
	}
}
