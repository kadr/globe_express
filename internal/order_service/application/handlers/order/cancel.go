package order_api_order_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/order"
)

func (oah *OrderAPIHandler) Cancel(c fiber.Ctx) error {
	oah.setLoggerFromReq(c)
	oah.logger.Info("cancel order")
	ctx := c.Context()
	var orderID uuid.UUID
	if id, err := uuid.Parse(c.Params("order_id")); err == nil {
		orderID = id
	} else {
		oah.logger.Error("api Cancel order error: %w", err)
		return err
	}
	order, err := oah.orderService.Cancel(ctx, orderID)
	if err != nil {
		oah.logger.Error("api Cancel order error: %w", err)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTO(order))
}
