package order_domain_order_service

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
)

// type OrderModel struct {
// 	ID               *uuid.UUID
// 	CustomerID       uuid.UUID
// 	TravelerID       *uuid.UUID
// 	Status           *string
// 	PickupCity       string
// 	PickupCountry    string
// 	DeliveryCity     string
// 	DeliveryCountry  string
// 	RewardAmount     float64
// 	RewardCurrency   string
// 	DeliveryDeadline time.Time
// 	DeliveryDate     time.Time
// 	CreatedAt        *time.Time
// 	UpdatedAt        *time.Time
// 	Product          *product_service.ProductModel
// }
//
// type OrderUpdateModel struct {
// 	TravelerID       *uuid.UUID
// 	Status           *string
// 	PickupCity       *string
// 	PickupCountry    *string
// 	DeliveryCity     *string
// 	DeliveryCountry  *string
// 	RewardAmount     *float64
// 	RewardCurrency   *string
// 	DeliveryDeadline *time.Time
// 	DeliveryDate     *time.Time
// }

type OrderRepositoryIface interface {
	Create(ctx context.Context, schema domain_models.OrderModel) (domain_models.OrderModel, error)
	Update(ctx context.Context, orderID uuid.UUID, schema domain_models.OrderUpdateModel) (domain_models.OrderModel, error)
	Cancel(ctx context.Context, orderID uuid.UUID) (domain_models.OrderModel, error)
	GetDetail(ctx context.Context, orderID uuid.UUID) (domain_models.OrderModel, error)
	GetList(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error)
	GetCustomer(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]domain_models.OrderModel, error)
	GetTravelerAvailable(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error)
}

type OrderService struct {
	repo           OrderRepositoryIface
	productService services.ProductServiceIface

	logger *slog.Logger
}

func NewService(repo OrderRepositoryIface, productService services.ProductServiceIface, logger *slog.Logger) *OrderService {
	return &OrderService{repo: repo, productService: productService, logger: logger}
}

func FromDomainModel(schema domain_models.OrderModel) services.OrderModel {
	status := string(schema.Status)
	orderModel := services.OrderModel{
		&schema.ID,
		schema.CustomerID,
		schema.TravelerID,
		&status,
		schema.PickupCity,
		schema.PickupCountry,
		schema.DeliveryCity,
		schema.DeliveryCountry,
		schema.RewardAmount,
		string(schema.RewardCurrency),
		schema.DeliveryDeadline,
		schema.DeliveryDate,
		&schema.CreatedAt,
		schema.UpdatedAt,
		nil,
	}
	if schema.Products != nil {
		products := make([]product_service.ProductModel, 0, len(*schema.Products))
		for _, product := range *schema.Products {
			item := product_service.ProductModel{
				&product.ID,
				product.OrderID,
				product.Name,
				product.Description,
				product.Price,
				string(product.Currency),
				product.ShopUrl,
				product.ShopName,
				product.ImageUrls,
				&product.CreatedAt,
				product.UpdatedAt,
			}

			products = append(products, item)
		}
		orderModel.Products = &products
	}

	return orderModel
}

func FromDomainModelList(result []domain_models.OrderModel) []services.OrderModel {
	var results []services.OrderModel
	for _, schema := range result {
		results = append(results, FromDomainModel(schema))
	}

	return results
}

func (os *OrderService) IsOwner(ctx context.Context, orderID, customerID uuid.UUID) bool {
	order, err := os.repo.GetDetail(ctx, orderID)
	if err != nil {
		return false
	}
	return order.CustomerID == customerID
}
