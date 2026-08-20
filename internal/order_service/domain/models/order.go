package order_domain_models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
	shared_money "github.com/kadr/globe_express/internal/shared/money"
	api_errors "github.com/kadr/globe_express/pkg/errors"
)

type OrderModel struct {
	ID               uuid.UUID
	CustomerID       uuid.UUID
	TravelerID       *uuid.UUID
	Status           vo.OrderStatus
	PickupCity       string
	PickupCountry    string
	DeliveryCity     string
	DeliveryCountry  string
	RewardAmount     float64
	RewardCurrency   shared_money.Currency
	DeliveryDeadline time.Time
	DeliveryDate     time.Time
	CreatedAt        time.Time
	UpdatedAt        *time.Time
	Products         *[]ProductModel
}
type OrderUpdateModel struct {
	TravelerID       *uuid.UUID
	Status           *vo.OrderStatus
	PickupCity       *string
	PickupCountry    *string
	DeliveryCity     *string
	DeliveryCountry  *string
	RewardAmount     *float64
	RewardCurrency   *shared_money.Currency
	DeliveryDeadline *time.Time
	DeliveryDate     *time.Time
}

func NewOrder(
	customerID uuid.UUID,
	pickupCity, pickupCountry, deliveryCity, deliveryCountry, rewardCurrency string,
	rewardAmount float64,
	deliveryDate, deliveryDeadline time.Time,
	status *string,
	travelerID *uuid.UUID,
) (OrderModel, error) {
	newStatus := vo.Pending
	var err error
	if status != nil {
		s, err := vo.NewStatus(*status)
		if err != nil {
			return OrderModel{}, err
		}
		newStatus = s
	}
	newRewardCurrency, err := vo.NewCurrency(rewardCurrency)
	if err != nil {
		return OrderModel{}, err
	}
	if pickupCity, err = vo.NewLocation(pickupCity); err != nil {
		return OrderModel{}, err
	}
	if pickupCountry, err = vo.NewLocation(pickupCountry); err != nil {
		return OrderModel{}, err
	}
	if deliveryCity, err = vo.NewLocation(deliveryCity); err != nil {
		return OrderModel{}, err
	}
	if deliveryCountry, err = vo.NewLocation(deliveryCountry); err != nil {
		return OrderModel{}, err
	}

	if rewardAmount, err = vo.NewPrice(rewardAmount); err != nil {
		return OrderModel{}, err
	}
	if deliveryDeadline.Before(deliveryDate) {
		return OrderModel{}, fmt.Errorf("incorrect delivery dates, delivery_deadline can not be less than delivery_date. %w", api_errors.ErrorFieldValidation)
	}
	mount := time.Duration(24*30) * time.Hour
	if deliveryDate.After(time.Now().Add(mount * 3)) {
		return OrderModel{}, fmt.Errorf("incorrect delivery date, can not be mare than 3 month. %w", api_errors.ErrorFieldValidation)
	}
	return OrderModel{
		uuid.Nil,
		customerID,
		travelerID,
		newStatus,
		pickupCity,
		pickupCountry,
		deliveryCity,
		deliveryCountry,
		rewardAmount,
		newRewardCurrency,
		deliveryDeadline,
		deliveryDate,
		time.Now(),
		nil,
		nil,
	}, nil
}

func NewUpdateOrder(
	pickupCity, pickupCountry, deliveryCity, deliveryCountry, rewardCurrency *string,
	rewardAmount *float64,
	deliveryDate, deliveryDeadline *time.Time,
	status *string,
	travelerID *uuid.UUID,
) (OrderUpdateModel, error) {
	var newStatus *vo.OrderStatus
	var err error
	if status != nil {
		s, err := vo.NewStatus(*status)
		if err != nil {
			return OrderUpdateModel{}, err
		}
		newStatus = &s
	}
	var newRewardCurrency shared_money.Currency
	if rewardCurrency != nil {
		newRewardCurrency, err = vo.NewCurrency(*rewardCurrency)
		if err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if pickupCity != nil {
		if *pickupCity, err = vo.NewLocation(*pickupCity); err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if pickupCountry != nil {
		if *pickupCountry, err = vo.NewLocation(*pickupCountry); err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if deliveryCity != nil {
		if *deliveryCity, err = vo.NewLocation(*deliveryCity); err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if deliveryCountry != nil {
		if *deliveryCountry, err = vo.NewLocation(*deliveryCountry); err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if rewardAmount != nil {
		if *rewardAmount, err = vo.NewPrice(*rewardAmount); err != nil {
			return OrderUpdateModel{}, err
		}
	}
	if deliveryDeadline != nil && deliveryDate != nil {
		if deliveryDeadline.Before(*deliveryDate) {
			return OrderUpdateModel{}, fmt.Errorf("incorrect delivery dates, delivery_deadline can not be less than delivery_date. %w", api_errors.ErrorFieldValidation)
		}
	}
	if deliveryDate != nil {
		mount := time.Duration(24*30) * time.Hour

		if deliveryDate.After(time.Now().Add(mount * 3)) {
			return OrderUpdateModel{}, fmt.Errorf("incorrect delivery date, can not be mare than 3 month. %w", api_errors.ErrorFieldValidation)
		}
	}
	return OrderUpdateModel{
		travelerID,
		newStatus,
		pickupCity,
		pickupCountry,
		deliveryCity,
		deliveryCountry,
		rewardAmount,
		&newRewardCurrency,
		deliveryDeadline,
		deliveryDate,
	}, nil
}
