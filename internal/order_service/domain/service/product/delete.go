package order_domain_product_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

func (ps *ProductService) Delete(ctx context.Context, productID uuid.UUID) error {
	if product, err := ps.repo.GetDetail(ctx, productID); err == nil {
		if product.OrderID != nil {
			return api_errors.ErrorCantDeleteProduct
		}
	}
	err := ps.repo.Delete(ctx, productID)
	if err != nil {
		return fmt.Errorf("product Delete service error: %w", err)
	}

	return nil
}
