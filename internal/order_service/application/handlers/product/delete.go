package order_api_product_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (pah *ProductAPIHandler) Delete(c fiber.Ctx) error {
	pah.setLoggerFromReq(c)
	pah.logger.Info("delete product")
	ctx := c.Context()
	var productID uuid.UUID
	if id, err := uuid.Parse(c.Params("product_id")); err == nil {
		productID = id
	} else {
		pah.logger.Error("api Delete product", "error:", err.Error())
		return err
	}
	err := pah.productService.Delete(ctx, productID)
	if err != nil {
		pah.logger.Error("api Delete product", "error:", err.Error())
		return err
	}
	c.Status(fiber.StatusNoContent)
	return nil
}
