package order_domain_product_service

import (
	"context"
	"fmt"
)

func (ts *ProductService) GetList(ctx context.Context, limit, offset int) ([]ProductModel, error) {
	products, err := ts.repo.GetList(ctx, limit, offset)
	if err != nil {
		return []ProductModel{}, fmt.Errorf("product GetList service error: %w", err)
	}

	return FromDomainModelList(products), nil
}
