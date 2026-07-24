package travaler_api_trip_handlers

import (
	"github.com/gofiber/fiber/v3"
	trip_service "github.com/kadr/globe_express/internal/travaler_service/interfaces/api/handlers"
)

type TripAPIHandler struct {
	travalerService trip_service.TravalerServiceIface
	app             *fiber.App
}

func NewTripAPIHandler(travalerService trip_service.TravalerServiceIface, app *fiber.App) *TripAPIHandler {
	return &TripAPIHandler{travalerService: travalerService, app: app}
}

func (tah *TripAPIHandler) RegisterHandlers() {
	routes := tah.app.Group("/api/v1/trips")
	routes.Post("/", tah.Create)
	routes.Delete("/:id/cancel", tah.Cancel)
	routes.Post("/active", tah.GetActive)
	routes.Post("/completed", tah.GetCompleted)
	routes.Post("/:id", tah.GetDetail)
	routes.Post("/", tah.GetList)
	routes.Post("/:id", tah.Update)
}
