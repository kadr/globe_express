package order_api_product_handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	product_service "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

type ProductAPIHandler struct {
	productService product_service.ProductServiceIface
	app            *fiber.App
	logger         *slog.Logger
}

func NewProductAPIHandler(productService product_service.ProductServiceIface, app *fiber.App, logger *slog.Logger) *ProductAPIHandler {
	return &ProductAPIHandler{productService: productService, app: app, logger: logger}
}

func (tah *ProductAPIHandler) RegisterHandlers() {
	routes := tah.app.Group("/api/v1/products")
	routes.Post("/", tah.Create)
	routes.Patch("/:product_id", tah.Update)
	routes.Get("/", tah.GetList)
	routes.Get("/:product_id", tah.GetDetail)
	routes.Delete("/:product_id", tah.Delete)
}

func (tah *ProductAPIHandler) setLoggerFromReq(c fiber.Ctx) {
	if logger, ok := c.Value("logger").(*slog.Logger); ok {
		tah.logger = logger
	}
}
