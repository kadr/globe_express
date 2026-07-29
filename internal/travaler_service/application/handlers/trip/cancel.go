package travaler_api_trip_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func (tah *TripAPIHandler) Cancel(c fiber.Ctx) error {
	tah.setLoggerFromReq(c)
	tah.logger.Info("canceling trip")
	ctx := c.Context()
	var tripID uuid.UUID
	if id, err := uuid.Parse(c.Params("trip_id")); err == nil {
		tripID = id
	} else {
		tah.logger.Error("api Cancel trip error: %w", err)
		return err
	}
	err := tah.travalerService.Cancel(ctx, tripID)
	if err != nil {
		tah.logger.Error("api Cancel trip error: %w", err)
		return err
	}
	c.Status(fiber.StatusNoContent)
	return nil
}
