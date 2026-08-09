package order_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
	shared_money "github.com/kadr/globe_express/internal/shared/money"
	api_errors "github.com/kadr/globe_express/pkg/errors"
	_ "github.com/lib/pq"
)

type OrderResult struct {
	ID               uuid.UUID       `db:"id"`
	CustomerID       uuid.UUID       `db:"customer_id"`
	TravelerID       *uuid.UUID      `db:"traveler_id"`
	Status           string          `db:"status"`
	PickupCity       string          `db:"pickup_city"`
	PickupCountry    string          `db:"pickup_country"`
	DeliveryCity     string          `db:"delivery_city"`
	DeliveryCountry  string          `db:"delivery_country"`
	RewardAmount     float64         `db:"reward_amount"`
	RewardCurrency   string          `db:"reward_currency"`
	DeliveryDeadline time.Time       `db:"delivery_deadline"`
	DeliveryDate     time.Time       `db:"delivery_date"`
	CreatedAt        time.Time       `db:"created_at"`
	UpdatedAt        *time.Time      `db:"updated_at"`
	Products         json.RawMessage `db:"products"`
}

type OrderRepository struct {
	db      *sqlx.DB
	timeout int
	logger  *slog.Logger
}

func NewOrderRepository(db *sqlx.DB, timeout int, logger *slog.Logger) *OrderRepository {
	return &OrderRepository{db: db, timeout: timeout, logger: logger}
}

func (or *OrderRepository) Create(ctx context.Context, schema domain_models.OrderModel) (domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	insertFields := []string{"customer_id", "status", "pickup_city", "pickup_country", "delivery_city", "delivery_country", "reward_amount", "reward_currency", "delivery_deadline", "delivery_date"}
	args := map[string]any{
		"customer_id":       schema.CustomerID,
		"status":            schema.Status,
		"pickup_city":       schema.PickupCity,
		"pickup_country":    schema.PickupCountry,
		"delivery_city":     schema.DeliveryCity,
		"delivery_country":  schema.DeliveryCountry,
		"reward_amount":     schema.RewardAmount,
		"reward_currency":   schema.RewardCurrency,
		"delivery_deadline": schema.DeliveryDeadline,
		"delivery_date":     schema.DeliveryDate,
	}
	if schema.TravelerID != nil {
		insertFields = append(insertFields, "traveler_id")
		args["traveler_id"] = *schema.TravelerID
	}

	insertValueFields := []string{}
	for _, field := range insertFields {
		insertValueFields = append(insertValueFields, fmt.Sprintf(":%s", field))
	}
	stmt := fmt.Sprintf(`
	INSERT INTO orders(%s)
	VALUES (%s) 
	RETURNING *
	`, strings.Join(insertFields, ","), strings.Join(insertValueFields, ","))
	rows, err := or.db.NamedQueryContext(ctx, stmt, args)
	if err != nil {
		or.logger.Error(err.Error())
		return domain_models.OrderModel{}, fmt.Errorf("order  repository, create err: %w", err)
	}
	defer rows.Close()

	var result OrderResult
	if rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			or.logger.Error(err.Error())
			return domain_models.OrderModel{}, fmt.Errorf("order  repository, create err: %w", err)
		}
	}
	return or.toOrderDomainModel(result), nil
}

func (or *OrderRepository) Update(ctx context.Context, orderID uuid.UUID, schema domain_models.OrderUpdateModel) (domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	var setParts []string
	args := map[string]interface{}{"id": orderID}
	if schema.TravelerID != nil {
		setParts = append(setParts, "traveler_id=:traveler_id")
		args["traveler_id"] = *schema.TravelerID
	}
	if schema.Status != nil {
		setParts = append(setParts, "status=:status")
		args["status"] = *schema.Status
	}
	if schema.PickupCity != nil {
		setParts = append(setParts, "pickup_city=:pickup_city")
		args["pickup_city"] = *schema.PickupCity
	}
	if schema.PickupCountry != nil {
		setParts = append(setParts, "pickup_country=:pickup_country")
		args["pickup_country"] = *schema.PickupCountry
	}
	if schema.DeliveryCity != nil {
		setParts = append(setParts, "delivery_city=:delivery_city")
		args["delivery_city"] = *schema.DeliveryCity
	}
	if schema.DeliveryCountry != nil {
		setParts = append(setParts, "delivery_country=:delivery_country")
		args["delivery_country"] = *schema.DeliveryCountry
	}
	if schema.RewardAmount != nil {
		setParts = append(setParts, "reward_amount=:reward_amount")
		args["reward_amount"] = *schema.RewardAmount
	}
	if schema.RewardCurrency != nil {
		setParts = append(setParts, "reward_currency=:reward_currency")
		args["reward_currency"] = *schema.RewardCurrency
	}
	if schema.DeliveryDeadline != nil {
		setParts = append(setParts, "delivery_deadline=:delivery_deadline")
		args["delivery_deadline"] = *schema.DeliveryDeadline
	}
	if schema.DeliveryDate != nil {
		setParts = append(setParts, "delivery_date=:delivery_date")
		args["delivery_date"] = *schema.DeliveryDate
	}

	stmt := fmt.Sprintf(
		`
		UPDATE orders SET %s, updated_at=NOW() WHERE id=:id 
		RETURNING *
		`,
		strings.Join(setParts, ","),
	)
	rows, err := or.db.NamedQueryContext(ctx, stmt, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return domain_models.OrderModel{}, fmt.Errorf("order  repository, update record no found: %w", api_errors.ErrorNotFound)
		}
		or.logger.Error(err.Error())
		return domain_models.OrderModel{}, fmt.Errorf("order  repository, update err: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var result OrderResult
		err := rows.StructScan(&result)
		if err != nil {
			or.logger.Error(err.Error())
			return domain_models.OrderModel{}, fmt.Errorf("order  repository, update err: %w", err)
		}

		return or.toOrderDomainModel(result), nil
	}
	return domain_models.OrderModel{}, fmt.Errorf("order  repository, update err: can't update record")
}

func (or *OrderRepository) Cancel(ctx context.Context, orderID uuid.UUID) (domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	stmt := "UPDATE orders SET status=$1 WHERE id = $2"

	row := or.db.QueryRowxContext(ctx, stmt, string(vo.Cancelled), orderID)
	var result OrderResult
	err := row.StructScan(&result)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return domain_models.OrderModel{}, fmt.Errorf("order  repository, record no found: %w", api_errors.ErrorNotFound)
		}
		or.logger.Error(err.Error())
		return domain_models.OrderModel{}, fmt.Errorf("order  repository, Cancel err: %w", err)
	}

	return or.toOrderDomainModel(result), nil
}

func (or *OrderRepository) GetDetail(ctx context.Context, orderID uuid.UUID) (domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	query := `
	SELECT o.*, COALESCE(
			(SELECT JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', p.id,
					'order_id', p.order_id,
					'name', p.name,
					'description', p.description,
					'price', p.price,
					'currency', p.currency,
					'shop_url', p.shop_url,
					'shop_name', p.shop_name,
					'image_urls', p.image_urls,
					'created_at', TO_CHAR(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
					'updated_at', TO_CHAR(p.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				)
			) FROM products p WHERE p.order_id = o.id),
			'[]'::json
		) as products
	FROM orders o 
	WHERE o.id = $1 
	ORDER BY o.created_at DESC
	`

	var result OrderResult
	err := or.db.GetContext(ctx, &result, query, orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return domain_models.OrderModel{}, fmt.Errorf("order  repository, record no found: %w", api_errors.ErrorNotFound)
		}
		or.logger.Error(err.Error())
		return domain_models.OrderModel{}, fmt.Errorf("order  repository, GetDetail err: %w", err)
	}

	return or.toOrderDomainModel(result), nil
}

func (or *OrderRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	query := `
	SELECT o.*, COALESCE(
			(SELECT JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', p.id,
					'order_id', p.order_id,
					'name', p.name,
					'description', p.description,
					'price', p.price,
					'currency', p.currency,
					'shop_url', p.shop_url,
					'shop_name', p.shop_name,
					'image_urls', p.image_urls,
					'created_at', TO_CHAR(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
					'updated_at', TO_CHAR(p.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				)
			) FROM products p WHERE p.order_id = o.id),
			'[]'::json
		) as products
	FROM orders o 
	ORDER BY o.created_at DESC LIMIT $1 OFFSET $2`

	var results []OrderResult
	err := or.db.SelectContext(ctx, &results, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return []domain_models.OrderModel{}, nil
		}
		or.logger.Error(err.Error())
		return []domain_models.OrderModel{}, fmt.Errorf("order  repository, GetList err: %w", err)
	}
	return or.toOrderDomainModelList(results), nil
}

func (or *OrderRepository) GetCustomer(ctx context.Context, customerID uuid.UUID, limit, offset int) ([]domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	query := `
	SELECT o.*, COALESCE(
			(SELECT JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', p.id,
					'order_id', p.order_id,
					'name', p.name,
					'description', p.description,
					'price', p.price,
					'currency', p.currency,
					'shop_url', p.shop_url,
					'shop_name', p.shop_name,
					'image_urls', p.image_urls,
					'created_at', TO_CHAR(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
					'updated_at', TO_CHAR(p.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				)
			) FROM products p WHERE p.order_id = o.id),
			'[]'::json
		) as products
	FROM orders o 
	WHERE customer_id=$1 
	ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	var results []OrderResult
	err := or.db.SelectContext(ctx, &results, query, customerID, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return []domain_models.OrderModel{}, nil
		}
		or.logger.Error(err.Error())
		return []domain_models.OrderModel{}, fmt.Errorf("order  repository, GetList err: %w", err)
	}
	return or.toOrderDomainModelList(results), nil
}

func (or *OrderRepository) GetTravelerAvailable(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(or.timeout))
	defer cancel()
	query := `
	SELECT o.*, COALESCE(
			(SELECT JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', p.id,
					'order_id', p.order_id,
					'name', p.name,
					'description', p.description,
					'price', p.price,
					'currency', p.currency,
					'shop_url', p.shop_url,
					'shop_name', p.shop_name,
					'image_urls', p.image_urls,
					'created_at', TO_CHAR(p.created_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"'),
					'updated_at', TO_CHAR(p.updated_at, 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"')
				)
			) FROM products p WHERE p.order_id = o.id),
			'[]'::json
		) as products
	FROM orders o 
	WHERE traveler_id IS NULL AND status = $1 
	ORDER BY created_at DESC LIMIT $2 OFFSET $3`

	var results []OrderResult
	err := or.db.SelectContext(ctx, &results, query, vo.Pending, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			or.logger.Warn(err.Error())
			return []domain_models.OrderModel{}, nil
		}
		or.logger.Error(err.Error())
		return []domain_models.OrderModel{}, fmt.Errorf("order  repository, GetList err: %w", err)
	}
	return or.toOrderDomainModelList(results), nil
}

func (or *OrderRepository) toOrderDomainModel(result OrderResult) domain_models.OrderModel {
	res := domain_models.OrderModel{
		ID:               result.ID,
		CustomerID:       result.CustomerID,
		TravelerID:       result.TravelerID,
		Status:           vo.OrderStatus(result.Status),
		PickupCity:       result.PickupCity,
		PickupCountry:    result.PickupCountry,
		DeliveryCity:     result.DeliveryCity,
		DeliveryCountry:  result.DeliveryCountry,
		RewardAmount:     result.RewardAmount,
		RewardCurrency:   shared_money.Currency(result.RewardCurrency),
		DeliveryDeadline: result.DeliveryDeadline,
		DeliveryDate:     result.DeliveryDate,
		CreatedAt:        result.CreatedAt,
		UpdatedAt:        result.UpdatedAt,
		Products:         nil,
	}
	if len(result.Products) > 0 && string(result.Products) != "null" {
		products := make([]domain_models.ProductModel, 0, len(result.Products))
		var dbProducts []ProductResult
		err := json.Unmarshal(result.Products, &dbProducts)
		if err == nil {
			for _, product := range dbProducts {
				item := domain_models.ProductModel{
					ID:          product.ID,
					OrderID:     product.OrderID,
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
					Currency:    shared_money.Currency(product.Currency),
					ShopUrl:     product.ShopUrl,
					ShopName:    product.ShopName,
					ImageUrls:   nil,
					CreatedAt:   product.CreatedAt,
					UpdatedAt:   product.UpdatedAt,
				}
				if product.ImageUrls != nil {
					images := strings.Split(*product.ImageUrls, ",")
					item.ImageUrls = &images
				}

				products = append(products, item)
			}
			res.Products = &products
		} else {
			or.logger.Warn(err.Error())
		}

	}

	return res
}

func (or *OrderRepository) toOrderDomainModelList(orders []OrderResult) []domain_models.OrderModel {
	var results []domain_models.OrderModel
	for _, res := range orders {
		results = append(results, or.toOrderDomainModel(res))
	}
	return results
}
