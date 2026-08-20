package order_domain_order_service

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	product_service "github.com/kadr/globe_express/internal/order_service/domain/service/product"
	vo "github.com/kadr/globe_express/internal/order_service/domain/value_objects"
	services "github.com/kadr/globe_express/internal/order_service/interfaces/api/handlers"
	api_errors "github.com/kadr/globe_express/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, trip domain_models.ProductModel) (domain_models.ProductModel, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockProductRepository) GetDetail(ctx context.Context, id uuid.UUID) (domain_models.ProductModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain_models.ProductModel{}, args.Error(1)
	}
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockProductRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.ProductModel, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain_models.ProductModel), args.Error(1)
}

func (m *MockProductRepository) Update(ctx context.Context, id uuid.UUID, trip domain_models.ProductUpdateModel) (domain_models.ProductModel, error) {
	args := m.Called(ctx, id, trip)
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockProductRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(ctx context.Context, trip domain_models.OrderModel) (domain_models.OrderModel, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) GetDetail(ctx context.Context, id uuid.UUID) (domain_models.OrderModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain_models.OrderModel{}, args.Error(1)
	}
	return args.Get(0).(domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) Update(ctx context.Context, id uuid.UUID, trip domain_models.OrderUpdateModel) (domain_models.OrderModel, error) {
	args := m.Called(ctx, id, trip)
	return args.Get(0).(domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) Cancel(ctx context.Context, id uuid.UUID) (domain_models.OrderModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) GetCustomer(ctx context.Context, id uuid.UUID, limit, offset int) ([]domain_models.OrderModel, error) {
	args := m.Called(ctx, id, limit, offset)
	return args.Get(0).([]domain_models.OrderModel), args.Error(1)
}

func (m *MockOrderRepository) GetTravelerAvailable(ctx context.Context, limit, offset int) ([]domain_models.OrderModel, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain_models.OrderModel), args.Error(1)
}

func TestCreate(t *testing.T) {
	now := time.Now()
	customerID := uuid.New()
	travelerID := uuid.New()
	pending := string(vo.Pending)

	tests := []struct {
		name       string
		input      services.OrderModel
		mockResult domain_models.OrderModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful create without all fields",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       nil,
				Status:           nil,
				PickupCity:       "Los Angeles",
				PickupCountry:    "USA",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{
				ID:               uuid.New(),
				CustomerID:       customerID,
				TravelerID:       nil,
				Status:           vo.Pending,
				PickupCity:       "Los Angeles",
				PickupCountry:    "USA",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        now,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "successful create with all fields",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "Los Angeles",
				PickupCountry:    "USA",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{
				ID:               uuid.New(),
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           vo.Pending,
				PickupCity:       "Los Angeles",
				PickupCountry:    "USA",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        now,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "create with incorrect pickup city name",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "L",
				PickupCountry:    "USA",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect pickup country name",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "Los Angeles",
				PickupCountry:    "U",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect delivery city name",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "Los Ageles",
				PickupCountry:    "USA",
				DeliveryCity:     "L",
				DeliveryCountry:  "UK",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect delivery country name",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "USA",
				PickupCountry:    "Los Angeles",
				DeliveryCity:     "London",
				DeliveryCountry:  "U",
				RewardAmount:     12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with negative amount",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "USA",
				PickupCountry:    "Los Angeles",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     -12.4,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with zero amount",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "USA",
				PickupCountry:    "Los Angeles",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     0,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect currency",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "USA",
				PickupCountry:    "Los Angeles",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     0,
				RewardCurrency:   "TR",
				DeliveryDeadline: now.Add(time.Hour * 152),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect delivery date and delivery deadline",
			input: services.OrderModel{
				ID:               nil,
				CustomerID:       customerID,
				TravelerID:       &travelerID,
				Status:           &pending,
				PickupCity:       "USA",
				PickupCountry:    "Los Angeles",
				DeliveryCity:     "London",
				DeliveryCountry:  "UK",
				RewardAmount:     0,
				RewardCurrency:   "USD",
				DeliveryDeadline: now.Add(time.Hour * 15),
				DeliveryDate:     now.Add(time.Hour * 24),
				CreatedAt:        nil,
				UpdatedAt:        nil,
				Products:         nil,
			},
			mockResult: domain_models.OrderModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockOrderRepository)
			mockProductRepo := new(MockProductRepository)

			if !tt.wantErr {
				mockRepo.On("Create", ctx, mock.Anything).Return(tt.mockResult, tt.mockErr)
			}
			productService := product_service.NewService(mockProductRepo, slog.Default())
			service := NewService(mockRepo, productService, slog.Default())
			result, err := service.Create(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, services.OrderModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.Equal(t, tt.input.CustomerID, result.CustomerID)
				if tt.input.TravelerID != nil {
					assert.Equal(t, *tt.input.TravelerID, *result.TravelerID)
				}
				assert.Equal(t, pending, *result.Status)
				assert.Equal(t, tt.input.PickupCity, result.PickupCity)
				assert.Equal(t, tt.input.PickupCountry, result.PickupCountry)
				assert.Equal(t, tt.input.DeliveryCity, result.DeliveryCity)
				assert.Equal(t, tt.input.DeliveryCountry, result.DeliveryCountry)
				assert.Equal(t, tt.input.RewardAmount, result.RewardAmount)
				assert.Equal(t, tt.input.RewardCurrency, result.RewardCurrency)
				assert.True(t, tt.input.DeliveryDate.Equal(result.DeliveryDate))
				assert.True(t, tt.input.DeliveryDeadline.Equal(result.DeliveryDeadline))

				assert.NotEqual(t, uuid.Nil, result.ID, "ID should be generated")
				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
				assert.Nil(t, result.UpdatedAt, "UpdatedAt should be nil")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

// func TestUpdate(t *testing.T) {
// 	now := time.Now()
// 	ID := uuid.New()
// 	newName := "New name"
// 	newDescription := "New description"
// 	newPrice := 21.36
// 	newCurrency := "USD"
// 	newShopUrl := "https://new-link.ru"
// 	newShopName := "Some new shop name"
// 	newImageUrls := []string{"https://link.ru"}
// 	incorrectName := "S"
// 	incorrectDescription := "M"
// 	incorrectPrice := -1.0
// 	incorrectCurrency := "TR"
// 	incorrectShopUrl := "shop-url.ru"
// 	incorrectShopName := "N"
// 	incorrectImageUrls := []string{"link.ru"}
// 	tests := []struct {
// 		name       string
// 		input      OrderUpdateModel
// 		successful domain_models.OrderModel
// 		mockErr    error
// 		wantErr    bool
// 	}{
// 		{
// 			name: "successful update",
// 			input: OrderUpdateModel{
// 				&newName,
// 				&newDescription,
// 				&newPrice,
// 				&newCurrency,
// 				&newShopUrl,
// 				&newShopName,
// 				&newImageUrls,
// 			},
// 			successful: domain_models.OrderModel{
// 				ID,
// 				nil,
// 				newName,
// 				&newDescription,
// 				newPrice,
// 				newCurrency,
// 				newShopUrl,
// 				newShopName,
// 				&newImageUrls,
// 				now,
// 				&now,
// 			},
// 			mockErr: nil,
// 			wantErr: false,
// 		},
// 		{
// 			name: "update incorrect name",
// 			input: OrderUpdateModel{
// 				Name: &incorrectName,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect description",
// 			input: OrderUpdateModel{
// 				Description: &incorrectDescription,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect price",
// 			input: OrderUpdateModel{
// 				Price: &incorrectPrice,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect currency",
// 			input: OrderUpdateModel{
// 				Currency: &incorrectCurrency,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect shop_url",
// 			input: OrderUpdateModel{
// 				ShopUrl: &incorrectShopUrl,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect shop_name",
// 			input: OrderUpdateModel{
// 				ShopName: &incorrectShopName,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 		{
// 			name: "update incorrect image_urls",
// 			input: OrderUpdateModel{
// 				ImageUrls: &incorrectImageUrls,
// 			},
// 			mockErr: api_errors.ErrorFieldValidation,
// 			wantErr: true,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			mockRepo := new(MockOrderRepository)
//
// 			if !tt.wantErr {
// 				mockRepo.On("Update", ctx, ID, mock.Anything).Return(tt.successful, tt.mockErr)
// 			}
// 			service := NewService(mockRepo, slog.Default())
// 			result, err := service.Update(ctx, ID, tt.input)
//
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Equal(t, OrderModel{}, result)
// 				assert.ErrorIs(t, err, tt.mockErr)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, result)
//
// 				assert.Nil(t, result.OrderID, "OrderID should be nil")
// 				assert.Equal(t, *tt.input.Name, result.Name)
// 				assert.Equal(t, *tt.input.Description, *result.Description)
// 				assert.Equal(t, *tt.input.Price, result.Price)
// 				assert.Equal(t, *tt.input.ShopUrl, result.ShopUrl)
// 				assert.Equal(t, *tt.input.ShopName, result.ShopName)
// 				assert.Equal(t, *tt.input.ImageUrls, *result.ImageUrls)
//
// 				assert.Equal(t, ID, *result.ID)
// 				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
// 				assert.NotNil(t, result.UpdatedAt, "UpdatedAt should be not nil")
// 			}
//
// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }
//
// func TestGetDetail(t *testing.T) {
// 	now := time.Now()
// 	ID := uuid.New()
// 	description := "desc"
// 	images := []string{"https://image.ru"}
//
// 	tests := []struct {
// 		name       string
// 		id         uuid.UUID
// 		mockResult domain_models.OrderModel
// 		mockErr    error
// 		wantErr    bool
// 		wantNil    bool
// 	}{
// 		{
// 			name: "successful get product",
// 			id:   ID,
// 			mockResult: domain_models.OrderModel{
// 				ID,
// 				nil,
// 				"Some",
// 				&description,
// 				2.3,
// 				"USD",
// 				"https://link.ru",
// 				"Name of shop",
// 				&images,
// 				now,
// 				&now,
// 			},
// 			mockErr: nil,
// 			wantErr: false,
// 			wantNil: false,
// 		},
// 		{
// 			name:       "product not found",
// 			id:         uuid.New(),
// 			mockResult: domain_models.OrderModel{},
// 			mockErr:    api_errors.ErrorNotFound,
// 			wantErr:    true,
// 			wantNil:    true,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			mockRepo := new(MockOrderRepository)
//
// 			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockResult, tt.mockErr)
//
// 			service := NewService(mockRepo, slog.Default())
// 			result, err := service.GetDetail(ctx, tt.id)
//
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Equal(t, OrderModel{}, result)
// 				assert.ErrorIs(t, err, tt.mockErr)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.NotNil(t, result)
//
// 				assert.Nil(t, result.OrderID, "OrderID should be nil")
// 				assert.Equal(t, tt.mockResult.Name, result.Name)
// 				assert.Equal(t, *tt.mockResult.Description, *result.Description)
// 				assert.Equal(t, tt.mockResult.Price, result.Price)
// 				assert.Equal(t, tt.mockResult.Currency, result.Currency)
// 				assert.Equal(t, tt.mockResult.ShopUrl, result.ShopUrl)
// 				assert.Equal(t, tt.mockResult.ShopName, result.ShopName)
// 				assert.Equal(t, *tt.mockResult.ImageUrls, *result.ImageUrls)
//
// 				assert.Equal(t, ID, *result.ID)
// 				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
// 				assert.NotNil(t, result.UpdatedAt, "UpdatedAt should be not nil")
// 			}
//
// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }
//
// func TestGetList(t *testing.T) {
// 	now := time.Now()
// 	orderID := uuid.New()
// 	description := "descr"
// 	images := []string{"https://link.ru"}
//
// 	tests := []struct {
// 		name        string
// 		limit       int
// 		offset      int
// 		mockResults []domain_models.OrderModel
// 		mockErr     error
// 		wantCount   int
// 		wantErr     bool
// 	}{
// 		{
// 			name: "successful get all",
// 			mockResults: []domain_models.OrderModel{
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 			},
// 			mockErr:   nil,
// 			wantCount: 3,
// 			wantErr:   false,
// 		},
// 		{
// 			name:  "successful get one record",
// 			limit: 1,
// 			mockResults: []domain_models.OrderModel{
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 			},
// 			mockErr:   nil,
// 			wantCount: 1,
// 			wantErr:   false,
// 		},
// 		{
// 			name:   "successful get with offset",
// 			offset: 1,
// 			mockResults: []domain_models.OrderModel{
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 				{
// 					uuid.New(),
// 					&orderID,
// 					"Addidas",
// 					&description,
// 					125.36,
// 					"USD",
// 					"https://addidas.com",
// 					"addidas",
// 					&images,
// 					now,
// 					nil,
// 				},
// 			},
// 			mockErr:   nil,
// 			wantCount: 2,
// 			wantErr:   false,
// 		},
// 		{
// 			name:        "get list with empty filters",
// 			mockResults: []domain_models.OrderModel{},
// 			mockErr:     nil,
// 			wantCount:   0,
// 			wantErr:     false,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			mockRepo := new(MockOrderRepository)
//
// 			mockRepo.On("GetList", ctx, tt.limit, tt.offset).Return(tt.mockResults, tt.mockErr)
//
// 			service := NewService(mockRepo, slog.Default())
// 			results, err := service.GetList(ctx, tt.limit, tt.offset)
//
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 				assert.Equal(t, OrderModel{}, results)
// 				assert.ErrorIs(t, err, tt.mockErr)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.wantCount, len(results))
// 			}
//
// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }
//
// func TestDelete(t *testing.T) {
// 	ID := uuid.New()
// 	ID2 := uuid.New()
// 	now := time.Now()
// 	description := "desc"
// 	tests := []struct {
// 		name       string
// 		id         uuid.UUID
// 		mockResult domain_models.OrderModel
// 		mockErr    error
// 		wantErr    bool
// 		callDetete bool
// 	}{
// 		{
// 			name: "successful delete",
// 			id:   ID,
// 			mockResult: domain_models.OrderModel{
// 				ID,
// 				nil,
// 				"UK",
// 				&description,
// 				23.1,
// 				"USD",
// 				"link",
// 				"name",
// 				nil,
// 				now,
// 				nil,
// 			},
// 			mockErr:    nil,
// 			wantErr:    false,
// 			callDetete: true,
// 		},
// 		{
// 			name:       "delete non exists product",
// 			id:         uuid.New(),
// 			mockResult: domain_models.OrderModel{},
// 			mockErr:    api_errors.ErrorNotFound,
// 			wantErr:    true,
// 			callDetete: true,
// 		},
// 		{
// 			name: "delete product which in order",
// 			id:   ID2,
// 			mockResult: domain_models.OrderModel{
// 				ID2,
// 				&ID2,
// 				"UK",
// 				&description,
// 				23.1,
// 				"USD",
// 				"link",
// 				"name",
// 				nil,
// 				now,
// 				nil,
// 			},
// 			mockErr:    api_errors.ErrorCantDeleteOrder,
// 			wantErr:    true,
// 			callDetete: false,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctx := context.Background()
// 			mockRepo := new(MockOrderRepository)
//
// 			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockResult, nil)
// 			if tt.callDetete {
// 				mockRepo.On("Delete", ctx, tt.id).Return(tt.mockErr)
// 			}
//
// 			service := NewService(mockRepo, slog.Default())
// 			err := service.Delete(ctx, tt.id)
//
// 			if tt.wantErr {
// 				assert.Error(t, err)
// 			} else {
// 				assert.NoError(t, err)
// 			}
//
// 			mockRepo.AssertExpectations(t)
// 		})
// 	}
// }
