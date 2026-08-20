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

func (pah *ProductAPIHandler) RegisterHandlers() {
	routes := pah.app.Group("/api/v1/products")
	routes.Post("/", pah.Create)
	routes.Patch("/:product_id", pah.Update)
	routes.Get("/", pah.GetList)
	routes.Get("/:product_id", pah.GetDetail)
	routes.Delete("/:product_id", pah.Delete)
}

func (tah *ProductAPIHandler) setLoggerFromReq(c fiber.Ctx) {
	if logger, ok := c.Value("logger").(*slog.Logger); ok {
		tah.logger = logger
	}
}
