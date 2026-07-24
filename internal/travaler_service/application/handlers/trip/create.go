package travaler_api_trip_handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	dto "github.com/kadr/globe_express/internal/travaler_service/application/dto/trip"
)

func (tah *TripAPIHandler) Create(c fiber.Ctx) error {
	ctx := c.Context()
	var createDTO dto.CreateDTO
	err := c.Bind().Body(&createDTO)
	if err != nil {
		slog.Error("api create trip error: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	schema, err := dto.ToDomainModel(createDTO)
	if err != nil {
		slog.Error("api create trip error: %w", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	trip, err := tah.travalerService.Create(ctx, schema)
	if err != nil {
		slog.Error("api create trip error: %w", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ToDTO(trip))
}
