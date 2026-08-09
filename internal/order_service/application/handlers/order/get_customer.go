package order_api_order_handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/order"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (oah *OrderAPIHandler) GetCustomer(c fiber.Ctx) error {
	oah.setLoggerFromReq(c)
	oah.logger.Info("listing customer orders")
	ctx := c.Context()
	limit := 100
	offset := 0
	var userID uuid.UUID
	if id, ok := c.Value("userID").(uuid.UUID); ok {
		userID = id
	} else {
		return fmt.Errorf("user id not found. %w", api_errors.ErrorBadRequest)
	}
	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}
	orders, err := oah.orderService.GetCustomer(ctx, userID, limit, offset)
	if err != nil {
		oah.logger.Error("api GetCustomer order error: %w", err)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTOList(orders))
}
