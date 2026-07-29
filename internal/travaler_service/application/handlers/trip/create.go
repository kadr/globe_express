package travaler_api_trip_handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	dto "github.com/kadr/globe_express/internal/travaler_service/application/dto/trip"
)

func (tah *TripAPIHandler) Create(c fiber.Ctx) error {
	tah.setLoggerFromReq(c)
	tah.logger.Info("creating new trip")
	ctx := c.Context()
	var createDTO dto.CreateDTO
	err := c.Bind().Body(&createDTO)
	if err != nil {
		tah.logger.Error("api Create trip error: %w", err)
		return err
	}
	userID := c.Value("userID").(uuid.UUID)
	schema, err := dto.ToCreateDomainModel(createDTO, userID)
	if err != nil {
		tah.logger.Error("api Create trip error: %w", err)
		return err
	}
	trip, err := tah.travalerService.Create(ctx, schema)
	if err != nil {
		tah.logger.Error("api Create trip error: %w", err)
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToDTO(trip))
}
