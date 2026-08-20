package order_api_order_handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/order"
)

func (oah *OrderAPIHandler) GetTravelerAvailable(c fiber.Ctx) error {
	oah.setLoggerFromReq(c)
	oah.logger.Info("listing available to traveler orders")
	ctx := c.Context()
	limit := 100
	offset := 0
	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}
	orders, err := oah.orderService.GetTravelerAvailable(ctx, limit, offset)
	if err != nil {
		oah.logger.Error("api GetTravelerAvailable order", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTOList(orders))
}
