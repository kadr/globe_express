package order_api_order_handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type OrderAPIHandler struct {
	orderService services.OrderServiceIface
	app          *fiber.App
	logger       *slog.Logger
}

func NewOrderAPIHandler(orderService services.OrderServiceIface, app *fiber.App, logger *slog.Logger) *OrderAPIHandler {
	return &OrderAPIHandler{orderService: orderService, app: app, logger: logger}
}

func (oah *OrderAPIHandler) RegisterHandlers() {
	routes := oah.app.Group("/api/v1/orders")
	routes.Post("/", oah.Create)
	routes.Patch("/:order_id", oah.Update)
	routes.Get("/", oah.GetList)
	routes.Get("/customer", oah.GetCustomer)
	routes.Get("/available", oah.GetTravelerAvailable)
	routes.Post("/:order_id/accept", oah.Accept)
	routes.Post("/:order_id/purchase", oah.Purchase)
	routes.Post("/:order_id/transit", oah.Transit)
	routes.Post("/:order_id/deliver", oah.Deliver)
	routes.Get("/:order_id", oah.GetDetail)
	routes.Delete("/:order_id", oah.Cancel)
}

func (oah *OrderAPIHandler) setLoggerFromReq(c fiber.Ctx) {
	if logger, ok := c.Value("logger").(*slog.Logger); ok {
		oah.logger = logger
	}
}
