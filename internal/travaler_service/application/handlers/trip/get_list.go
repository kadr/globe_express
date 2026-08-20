package travaler_api_trip_handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/travaler_service/application/dto/trip"
)

func (tah *TripAPIHandler) GetList(c fiber.Ctx) error {
	tah.setLoggerFromReq(c)
	tah.logger.Info("listing trip")
	ctx := c.Context()
	limit := 100
	offset := 0
	if l, err := strconv.Atoi(c.Query("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(c.Query("offset")); err == nil {
		offset = o
	}
	trips, err := tah.travalerService.GetList(ctx, limit, offset)
	if err != nil {
		tah.logger.Error("api GetList trip", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTOList(trips))
}
