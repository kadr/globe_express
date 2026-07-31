package order_product_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	api_errors "github.com/kadr/globe_express/pkg/errors"
	_ "github.com/lib/pq"
)

type ProductResult struct {
	ID          uuid.UUID  `db:"id"`
	OrderID     *uuid.UUID `db:"order_id"`
	Name        string     `db:"name"`
	Description *string    `db:"description"`
	Price       float64    `db:"price"`
	Currency    string     `db:"currency"`
	ShopUrl     string     `db:"shop_url"`
	ShopName    string     `db:"shop_name"`
	ImageUrls   *string    `db:"image_urls"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

type ProductRepository struct {
	db      *sqlx.DB
	timeout int
	logger  *slog.Logger
}

func NewRepository(db *sqlx.DB, timeout int, logger *slog.Logger) *ProductRepository {
	return &ProductRepository{db: db, timeout: timeout, logger: logger}
}

func (pr *ProductRepository) Create(ctx context.Context, schema domain_models.ProductModel) (domain_models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(pr.timeout))
	defer cancel()
	insertFields := []string{"name", "price", "currency", "shop_url", "shop_name"}
	args := map[string]any{"name": schema.Name, "price": schema.Price, "currency": schema.Currency, "shop_url": schema.ShopUrl, "shop_name": schema.ShopName}
	if schema.OrderID != nil {
		insertFields = append(insertFields, "order_id")
		args["order_id"] = *schema.OrderID
	}
	if schema.Description != nil {
		insertFields = append(insertFields, "description")
		args["description"] = *schema.Description
	}
	if schema.ImageUrls != nil {
		insertFields = append(insertFields, "image_urls")
		args["image_urls"] = strings.Join(*schema.ImageUrls, ",")
	}

	insertValueFields := []string{}
	for _, field := range insertFields {
		insertValueFields = append(insertValueFields, fmt.Sprintf(":%s", field))
	}
	stmt := fmt.Sprintf(`
	INSERT INTO products(%s)
	VALUES (%s) 
	RETURNING id,order_id,name,description,price,currency,shop_url,shop_name,image_urls,created_at,updated_at
	`, strings.Join(insertFields, ","), strings.Join(insertValueFields, ","))
	rows, err := pr.db.NamedQueryContext(ctx, stmt, args)
	if err != nil {
		pr.logger.Error(err.Error())
		return domain_models.ProductModel{}, fmt.Errorf("product  repository, create err: %w", err)
	}
	defer rows.Close()

	var result ProductResult
	if rows.Next() {
		err := rows.StructScan(&result)
		if err != nil {
			pr.logger.Error(err.Error())
			return domain_models.ProductModel{}, fmt.Errorf("product  repository, create err: %w", err)
		}
	}
	return toDomainModel(result), nil
}

func (pr *ProductRepository) Update(ctx context.Context, productID uuid.UUID, schema domain_models.ProductUpdateModel) (domain_models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(pr.timeout))
	defer cancel()
	var setParts []string
	args := map[string]interface{}{"id": productID}
	if schema.Name != nil {
		setParts = append(setParts, "name=:name")
		args["name"] = *schema.Name
	}
	if schema.Price != nil {
		setParts = append(setParts, "price=:price")
		args["price"] = *schema.Price
	}
	if schema.Currency != nil {
		setParts = append(setParts, "currency=:currency")
		args["currency"] = *schema.Currency
	}
	if schema.ShopUrl != nil {
		setParts = append(setParts, "shop_url=:shop_url")
		args["shop_url"] = *schema.ShopUrl
	}
	if schema.ShopName != nil {
		setParts = append(setParts, "shop_name=:shop_name")
		args["shop_name"] = *schema.ShopName
	}
	if schema.Description != nil {
		setParts = append(setParts, "description=:description")
		args["description"] = *schema.Description
	}
	if schema.ImageUrls != nil {
		setParts = append(setParts, "image_urls=:image_urls")
		args["image_urls"] = strings.Join(*schema.ImageUrls, ",")
	}

	stmt := fmt.Sprintf(
		`
		UPDATE products SET %s, updated_at=NOW() WHERE id=:id 
		RETURNING id,order_id,name,description,price,currency,shop_url,shop_name,image_urls,created_at,updated_at`,
		strings.Join(setParts, ","),
	)
	pr.logger.Debug("stmt", stmt)
	rows, err := pr.db.NamedQueryContext(ctx, stmt, args)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Warn(err.Error())
			return domain_models.ProductModel{}, fmt.Errorf("product  repository, update record no found: %w", api_errors.ErrorNotFound)
		}
		pr.logger.Error(err.Error())
		return domain_models.ProductModel{}, fmt.Errorf("product  repository, update err: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		var result ProductResult
		err := rows.StructScan(&result)
		if err != nil {
			pr.logger.Error(err.Error())
			return domain_models.ProductModel{}, fmt.Errorf("product  repository, update err: %w", err)
		}

		return toDomainModel(result), nil
	}
	return domain_models.ProductModel{}, fmt.Errorf("product  repository, update err: can't update record")
}

func (pr *ProductRepository) Delete(ctx context.Context, productID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(pr.timeout))
	defer cancel()
	stmt := "DELETE FROM products WHERE id = $1"

	res, err := pr.db.ExecContext(ctx, stmt, productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Warn(err.Error())
			return fmt.Errorf("product  repository, record no found: %w", api_errors.ErrorNotFound)
		}
		pr.logger.Error(err.Error())
		return fmt.Errorf("product  repository, Cancel err: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("product  repository, Cancel err: %w", err)
	}
	if count == 0 {
		pr.logger.Warn("no rows found to delete")
	}

	return nil
}

func (pr *ProductRepository) GetDetail(ctx context.Context, productID uuid.UUID) (domain_models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(pr.timeout))
	defer cancel()
	query := `
	SELECT id,order_id,name,description,price,currency,shop_url,shop_name,image_urls,created_at,updated_at
	FROM products WHERE id = $1 ORDER BY created_at DESC`

	var result ProductResult
	err := pr.db.GetContext(ctx, &result, query, productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Warn(err.Error())
			return domain_models.ProductModel{}, fmt.Errorf("product  repository, record no found: %w", api_errors.ErrorNotFound)
		}
		pr.logger.Error(err.Error())
		return domain_models.ProductModel{}, fmt.Errorf("product  repository, GetDetail err: %w", err)
	}

	return toDomainModel(result), nil
}

func (pr *ProductRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.ProductModel, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(pr.timeout))
	defer cancel()
	query := `
	SELECT id,order_id,name,description,price,currency,shop_url,shop_name,image_urls,created_at,updated_at
	FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	var results []ProductResult
	err := pr.db.SelectContext(ctx, &results, query, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			pr.logger.Warn(err.Error())
			return []domain_models.ProductModel{}, nil
		}
		pr.logger.Error(err.Error())
		return []domain_models.ProductModel{}, fmt.Errorf("product  repository, GetList err: %w", err)
	}
	return toDomainModelList(results), nil
}

func toDomainModel(result ProductResult) domain_models.ProductModel {
	var images []string
	if result.ImageUrls != nil {
		for _, image := range strings.Split(*result.ImageUrls, ",") {
			images = append(images, image)
		}
	}
	return domain_models.ProductModel{
		result.ID,
		result.OrderID,
		result.Name,
		result.Description,
		result.Price,
		result.Currency,
		result.ShopUrl,
		result.ShopName,
		&images,
		result.CreatedAt,
		result.UpdatedAt,
	}
}

func toDomainModelList(products []ProductResult) []domain_models.ProductModel {
	var results []domain_models.ProductModel
	for _, res := range products {
		results = append(results, toDomainModel(res))
	}
	return results
}
