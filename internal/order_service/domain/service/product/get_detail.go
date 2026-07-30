package order_domain_product_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (ts *ProductService) GetDetail(ctx context.Context, productID uuid.UUID) (ProductModel, error) {
	product, err := ts.repo.GetDetail(ctx, productID)
	if err != nil {
		return ProductModel{}, fmt.Errorf("product GetDetail service error: %w", err)
	}

	return FromDomainModel(product), nil
}
