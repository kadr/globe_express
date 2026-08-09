package order_api_services_interface

import (
	"context"
	"time"

	"github.com/google/uuid"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
)

type OrderModel struct {
	ID               *uuid.UUID
	CustomerID       uuid.UUID
	TravelerID       *uuid.UUID
	Status           *string
	PickupCity       string
	PickupCountry    string
	DeliveryCity     string
	DeliveryCountry  string
	RewardAmount     float64
	RewardCurrency   string
	DeliveryDeadline time.Time
	DeliveryDate     time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	Products         *[]product_service.ProductModel
}

type OrderUpdateModel struct {
	TravelerID       *uuid.UUID
	Status           *string
	PickupCity       *string
	PickupCountry    *string
	DeliveryCity     *string
	DeliveryCountry  *string
	RewardAmount     *float64
	RewardCurrency   *string
	DeliveryDeadline *time.Time
	DeliveryDate     *time.Time
}

type OrderServiceIface interface {
	Create(ctx context.Context, schema OrderModel) (OrderModel, error)
	Update(ctx context.Context, orderID uuid.UUID, schema OrderUpdateModel) (OrderModel, error)
	Cancel(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	GetDetail(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	GetList(ctx context.Context, limit, offset int) ([]OrderModel, error)
	GetCustomer(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]OrderModel, error)
	GetTravelerAvailable(ctx context.Context, limit, offset int) ([]OrderModel, error)
	Accept(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	Purchase(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	Transit(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	Deliver(ctx context.Context, orderID uuid.UUID) (OrderModel, error)
	IsOwner(ctx context.Context, customerID, orderID uuid.UUID) bool
}
