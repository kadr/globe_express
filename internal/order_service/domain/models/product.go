package order_domain_models

import (
	"time"

	"github.com/google/uuid"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
)

type ProductModel struct {
	ID          uuid.UUID
	OrderID     *uuid.UUID
	Name        string
	Description *string
	Price       float64
	Currency    string
	ShopUrl     string
	ShopName    string
	ImageUrls   *[]string
	CreatedAt   time.Time
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

func NewProduct(name, currency, shopUrl, shopName string, price float64, description *string, orderID *uuid.UUID, imageUrls *[]string) (ProductModel, error) {
	var err error
	name, err = vo.NewName(name)
	if err != nil {
		return ProductModel{}, err
	}
	price, err = vo.NewPrice(price)
	if err != nil {
		return ProductModel{}, err
	}
	currency, err = vo.NewCurrency(currency)
	if err != nil {
		return ProductModel{}, err
	}
	shopUrl, err = vo.NewUrl(shopUrl)
	if err != nil {
		return ProductModel{}, err
	}
	shopName, err = vo.NewName(shopName)
	if err != nil {
		return ProductModel{}, err
	}
	if description != nil {
		*description, err = vo.NewDescription(*description)
		if err != nil {
			return ProductModel{}, err
		}
	}
	if imageUrls != nil {
		for _, url := range *imageUrls {
			_, err = vo.NewUrl(url)
			if err != nil {
				return ProductModel{}, err
			}
		}
	}
	return ProductModel{
		uuid.Nil,
		orderID,
		name,
		description,
		price,
		currency,
		shopUrl,
		shopName,
		imageUrls,
		time.Now(),
		nil,
	}, nil
}

func NewUpdateProduct(name, currency, shopUrl, shopName, description *string, price *float64, imageUrls *[]string) (ProductUpdateModel, error) {
	var err error
	if name != nil {
		*name, err = vo.NewName(*name)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if price != nil {
		*price, err = vo.NewPrice(*price)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if currency != nil {
		*currency, err = vo.NewCurrency(*currency)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if shopUrl != nil {
		*shopUrl, err = vo.NewUrl(*shopUrl)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if shopName != nil {
		*shopName, err = vo.NewName(*shopName)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if description != nil {
		*description, err = vo.NewDescription(*description)
		if err != nil {
			return ProductUpdateModel{}, err
		}
	}
	if imageUrls != nil {
		for _, url := range *imageUrls {
			_, err = vo.NewUrl(url)
			if err != nil {
				return ProductUpdateModel{}, err
			}
		}
	}
	return ProductUpdateModel{
		name,
		description,
		price,
		currency,
		shopUrl,
		shopName,
		imageUrls,
	}, nil
}
