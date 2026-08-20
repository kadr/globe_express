package travaler_api_trip_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/travaler_service/application/dto/trip"
)

func (tah *TripAPIHandler) Update(c fiber.Ctx) error {
	tah.setLoggerFromReq(c)
	tah.logger.Info("updateing trip")
	ctx := c.Context()
	var tripID uuid.UUID
	if id, err := uuid.Parse(c.Params("trip_id")); err == nil {
		tripID = id
	} else {
		tah.logger.Error("api update trip", "error:", err.Error())
		return err
	}
	var updateDTO dto.UpdateDTO
	err := c.Bind().Body(&updateDTO)
	if err != nil {
		tah.logger.Error("api update trip", "error:", err.Error())
		return err
	}
	if err != nil {
		tah.logger.Error("api update trip", "error:", err.Error())
		return err
	}
	trip, err := tah.travalerService.Update(ctx, tripID, dto.ToUpdateModel(updateDTO))
	if err != nil {
		tah.logger.Error("api update trip", "error:", err.Error())
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ToDTO(trip))
}
