package orders_api_handlers

import (
	"context"
	"log/slog"

	"github.com/gofiber/fiber/v3"
	order_handlers "github.com/kadr/globe_express/internal/order_service/application/handlers/order"
	product_handlers "github.com/kadr/globe_express/internal/order_service/application/handlers/product"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type OrderAPI struct {
	orderService   services.OrderServiceIface
	productService services.ProductServiceIface
	app            *fiber.App
	logger         *slog.Logger
}

func NewProductAPI(orderService services.OrderServiceIface, productService services.ProductServiceIface, logger *slog.Logger) *OrderAPI {
	return &OrderAPI{orderService: orderService, productService: productService, app: fiber.New(), logger: logger}
}

func (ta *OrderAPI) RegisterMiddleware(middlewares ...fiber.Handler) {
	for _, middleware := range middlewares {
		ta.app.Use(middleware)
	}
}

func (ta *OrderAPI) RegisterHandlers() {
	product_handlers.NewProductAPIHandler(ta.productService, ta.app, ta.logger).RegisterHandlers()
	order_handlers.NewOrderAPIHandler(ta.orderService, ta.app, ta.logger).RegisterHandlers()
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
