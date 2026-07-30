package order_api_product_service_interface

import (
	"context"

	"github.com/google/uuid"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
)

type ProductServiceIface interface {
	Create(ctx context.Context, schema product_service.ProductModel) (product_service.ProductModel, error)
	Update(ctx context.Context, tripID uuid.UUID, schema product_service.ProductUpdateModel) (product_service.ProductModel, error)
	Delete(ctx context.Context, tripID uuid.UUID) error
	GetDetail(ctx context.Context, tripID uuid.UUID) (product_service.ProductModel, error)
	GetList(ctx context.Context, limit, offset int) ([]product_service.ProductModel, error)
}
