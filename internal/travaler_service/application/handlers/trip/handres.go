package travaler_api_trip_handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	trip_service "github.com/kadr/globe_express/internal/travaler_service/interfaces/api/handlers"
)

type TripAPIHandler struct {
	travalerService trip_service.TravalerServiceIface
	app             *fiber.App
	logger          *slog.Logger
}

func NewTripAPIHandler(travalerService trip_service.TravalerServiceIface, app *fiber.App, logger *slog.Logger) *TripAPIHandler {
	return &TripAPIHandler{travalerService: travalerService, app: app, logger: logger}
}

func (tah *TripAPIHandler) RegisterHandlers() {
	routes := tah.app.Group("/api/v1/trips")
	routes.Post("/", tah.Create)
	routes.Patch("/:trip_id", tah.Update)
	routes.Get("/", tah.GetList)
	routes.Get("/active", tah.GetActive)
	routes.Get("/completed", tah.GetCompleted)
	routes.Get("/:trip_id", tah.GetDetail)
	routes.Delete("/:trip_id/cancel", tah.Cancel)
}

func (tah *TripAPIHandler) setLoggerFromReq(c fiber.Ctx) {
	if logger, ok := c.Value("logger").(*slog.Logger); ok {
		tah.logger = logger
	}
}
