package order_api_order_handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/order"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (oah *OrderAPIHandler) Create(c fiber.Ctx) error {
	oah.setLoggerFromReq(c)
	oah.logger.Info("creating new order")
	ctx := c.Context()
	var createDTO dto.CreateDTO
	err := c.Bind().Body(&createDTO)
	if err != nil {
		oah.logger.Error("api Create order", "error:", err.Error())
		return err
	}
	if id, ok := c.Value("userID").(uuid.UUID); ok {
		createDTO.CustomerID = id
	} else {
		return fmt.Errorf("user id not found. %w", api_errors.ErrorBadRequest)
	}
	order, err := oah.orderService.Create(ctx, dto.ToCreateModel(createDTO))
	if err != nil {
		oah.logger.Error("api Create order", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToDTO(order))
}
