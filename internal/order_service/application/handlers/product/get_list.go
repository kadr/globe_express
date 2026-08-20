package order_api_product_handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/order_service/application/dto/product"
)

func (pah *ProductAPIHandler) GetList(c fiber.Ctx) error {
	pah.setLoggerFromReq(c)
	pah.logger.Info("listing product")
	ctx := c.Context()
	limit := 100
	offset := 0
	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}
	products, err := pah.productService.GetList(ctx, limit, offset)
	if err != nil {
		pah.logger.Error("api GetList product", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTOList(products))
}
