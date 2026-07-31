package order_domain_product_service

import (
	"context"
	"fmt"

	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
)

func (ps *ProductService) Create(ctx context.Context, schema ProductModel) (ProductModel, error) {
	domainSchema, err := ToDomainModel(schema)
	if err != nil {
		return ProductModel{}, fmt.Errorf("product service error: %w", err)
	}
	product, err := ps.repo.Create(ctx, domainSchema)
	if err != nil {
		return ProductModel{}, fmt.Errorf("product service error: %w", err)
	}

	return FromDomainModel(product), nil
}

func ToDomainModel(schema ProductModel) (domain_models.ProductModel, error) {
	productModel, err := domain_models.NewProduct(
		schema.Name,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.Price,
		schema.Description,
		schema.OrderID,
		schema.ImageUrls,
	)
	if err != nil {
		return domain_models.ProductModel{}, err
	}

	return productModel, nil
}
