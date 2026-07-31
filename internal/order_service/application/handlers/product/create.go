package order_api_product_handlers

import (
	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
)

func (pah *ProductAPIHandler) Create(c fiber.Ctx) error {
	pah.setLoggerFromReq(c)
	pah.logger.Info("creating new product")
	ctx := c.Context()
	var createDTO dto.CreateDTO
	err := c.Bind().Body(&createDTO)
	if err != nil {
		pah.logger.Error("api Create product error:", err)
		return err
	}
	product, err := pah.productService.Create(ctx, dto.ToCreateModel(createDTO))
	if err != nil {
		pah.logger.Error("api Create product error:", err)
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToDTO(product))
}
