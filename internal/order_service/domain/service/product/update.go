package order_domain_product_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
)

func (ts *ProductService) Update(ctx context.Context, productID uuid.UUID, schema ProductUpdateModel) (ProductModel, error) {
	newSchema, err := ToUpdateDomainModel(schema)
	if err != nil {
		return ProductModel{}, fmt.Errorf("product update service error: %w", err)
	}
	product, err := ts.repo.Update(ctx, productID, newSchema)
	if err != nil {
		return ProductModel{}, fmt.Errorf("product update service error: %w", err)
	}

	return FromDomainModel(product), nil
}

func ToUpdateDomainModel(schema ProductUpdateModel) (domain_models.ProductUpdateModel, error) {
	productModel, err := domain_models.NewUpdateProduct(
		schema.Name,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.Description,
		schema.Price,
		schema.ImageUrls,
	)
	if err != nil {
		return domain_models.ProductUpdateModel{}, err
	}

	return productModel, nil
}
