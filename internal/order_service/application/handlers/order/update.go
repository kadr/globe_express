package order_api_order_handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/order"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (oah *OrderAPIHandler) Update(c fiber.Ctx) error {
	oah.setLoggerFromReq(c)
	oah.logger.Info("updateing order")
	ctx := c.Context()
	var orderID uuid.UUID
	if id, err := uuid.Parse(c.Params("order_id")); err == nil {
		orderID = id
	} else {
		oah.logger.Error("api update order error: %w", err)
		return err
	}
	var updateDTO dto.UpdateDTO
	err := c.Bind().Body(&updateDTO)
	if err != nil {
		oah.logger.Error("api update order error: %w", err)
		return err
	}
	if err != nil {
		oah.logger.Error("api update order error: %w", err)
		return err
	}
	if id, ok := c.Value("userID").(uuid.UUID); ok {
		if !oah.orderService.IsOwner(ctx, id, orderID) {
			return fmt.Errorf("You can't edit not your order %w", api_errors.ErrorForbidden)
		}
	}
	order, err := oah.orderService.Update(ctx, orderID, dto.ToUpdateModel(updateDTO))
	if err != nil {
		oah.logger.Error("api update order error: %w", err)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTO(order))
}
