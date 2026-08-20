package travaler_api_trip_handlers

import (
	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/travaler_service/application/dto/trip"
)

func (tah *TripAPIHandler) GetCompleted(c fiber.Ctx) error {
	tah.setLoggerFromReq(c)
	tah.logger.Info("geting active trips")
	ctx := c.Context()
	trips, err := tah.travalerService.GetCompleted(ctx)
	if err != nil {
		tah.logger.Error("api active trip", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTOList(trips))
}
