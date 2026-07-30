package order_domain_product_service

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
)

type ProductModel struct {
	ID          *uuid.UUID
	OrderID     *uuid.UUID
	Name        string
	Description *string
	Price       float64
	Currency    string
	ShopUrl     string
	ShopName    string
	ImageUrls   *[]string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type ProductUpdateModel struct {
	Name        *string
	Description *string
	Price       *float64
	Currency    *string
	ShopUrl     *string
	ShopName    *string
	ImageUrls   *[]string
}

type ProductRepositoryIface interface {
	Create(ctx context.Context, schema domain_models.ProductModel) (domain_models.ProductModel, error)
	Update(ctx context.Context, tripID uuid.UUID, schema domain_models.ProductUpdateModel) (domain_models.ProductModel, error)
	Delete(ctx context.Context, tripID uuid.UUID) error
	GetDetail(ctx context.Context, tripID uuid.UUID) (domain_models.ProductModel, error)
	GetList(ctx context.Context, limit, offset int) ([]domain_models.ProductModel, error)
}

type ProductService struct {
	repo   ProductRepositoryIface
	logger *slog.Logger
}

func NewService(repo ProductRepositoryIface, logger *slog.Logger) *ProductService {
	return &ProductService{repo: repo, logger: logger}
}

func FromDomainModel(schema domain_models.ProductModel) ProductModel {
	tripModel := ProductModel{
		&schema.ID,
		schema.OrderID,
		schema.Name,
		schema.Description,
		schema.Price,
		schema.Currency,
		schema.ShopUrl,
		schema.ShopName,
		schema.ImageUrls,
		&schema.CreatedAt,
		schema.UpdatedAt,
	}

	return tripModel
}

func FromDomainModelList(result []domain_models.ProductModel) []ProductModel {
	var results []ProductModel
	for _, schema := range result {
		results = append(results, ProductModel{
			&schema.ID,
			schema.OrderID,
			schema.Name,
			schema.Description,
			schema.Price,
			schema.Currency,
			schema.ShopUrl,
			schema.ShopName,
			schema.ImageUrls,
			&schema.CreatedAt,
			schema.UpdatedAt,
		})
	}

	return results
}
