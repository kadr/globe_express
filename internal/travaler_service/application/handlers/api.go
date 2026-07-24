package travaler_api_handlers

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	trip_handlers "github.com/kadr/globe_express/internal/travaler_service/application/handlers/trip"
	trip_service "github.com/kadr/globe_express/internal/travaler_service/interfaces/api/handlers"
)

type TravalerAPI struct {
	tripService trip_service.TravalerServiceIface
	app         *fiber.App
	logger      *slog.Logger
}

func NewTravalerAPI(tripService trip_service.TravalerServiceIface, logger *slog.Logger) *TravalerAPI {
	return &TravalerAPI{tripService: tripService, app: fiber.New(), logger: logger}
}

func (ta *TravalerAPI) RegisterMiddleware(middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		ta.app.Use(middleware)
	}
}

func (ta *TravalerAPI) RegisterHandlers() {
	trip_handlers.NewTripAPIHandler(ta.tripService, ta.app).RegisterHandlers()
}

func (ta *TravalerAPI) Start(address string) error {
	err := ta.app.Listen(address)
	if err != nil {
		return err
	}
	ta.logger.Info("Server start succeful in address: ", address)
	return nil
}

func (ta *TravalerAPI) Shutdown(ctx context.Context) {
	if err := ta.app.ShutdownWithContext(ctx); err != nil {
		ta.logger.Warn("can't shutdown server")
	}
}
