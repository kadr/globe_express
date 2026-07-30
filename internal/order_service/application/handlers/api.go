package orders_api_handlers

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	product_handlers "github.com/kadr/globe_express/internal/order_service/application/handlers/product"
	product_service "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type OrderAPI struct {
	productService product_service.ProductServiceIface
	app            *fiber.App
	logger         *slog.Logger
}

func NewProductAPI(productService product_service.ProductServiceIface, logger *slog.Logger) *OrderAPI {
	return &OrderAPI{productService: productService, app: fiber.New(), logger: logger}
}

func (ta *OrderAPI) RegisterMiddleware(middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		ta.app.Use(middleware)
	}
}

func (ta *OrderAPI) RegisterHandlers() {
	product_handlers.NewProductAPIHandler(ta.productService, ta.app, ta.logger).RegisterHandlers()
}

func (ta *OrderAPI) Start(address string) error {
	err := ta.app.Listen(address)
	if err != nil {
		return err
	}
	ta.logger.Info("Order server start succeful in address: ", address)
	return nil
}

func (ta *OrderAPI) Shutdown(ctx context.Context) {
	if err := ta.app.ShutdownWithContext(ctx); err != nil {
		ta.logger.Warn("can't shutdown server")
	}
}
