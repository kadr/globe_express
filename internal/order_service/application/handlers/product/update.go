package order_api_product_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
)

func (pah *ProductAPIHandler) Update(c fiber.Ctx) error {
	pah.setLoggerFromReq(c)
	pah.logger.Info("updateing product")
	ctx := c.Context()
	var productID uuid.UUID
	if id, err := uuid.Parse(c.Params("product_id")); err == nil {
		productID = id
	} else {
		pah.logger.Error("api update product error: %w", err)
		return err
	}
	var updateDTO dto.UpdateDTO
	err := c.Bind().Body(&updateDTO)
	if err != nil {
		pah.logger.Error("api update product error: %w", err)
		return err
	}
	if err != nil {
		pah.logger.Error("api update product error: %w", err)
		return err
	}
	product, err := pah.productService.Update(ctx, productID, dto.ToUpdateModel(updateDTO))
	if err != nil {
		pah.logger.Error("api update product error: %w", err)
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTO(product))
}
