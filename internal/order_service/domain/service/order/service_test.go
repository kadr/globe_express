package order_domain_order_service

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/order_service/domain/models"
	api_errors "github.com/kadr/globe_express/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTripRepository struct {
	mock.Mock
}

func (m *MockTripRepository) Create(ctx context.Context, trip domain_models.ProductModel) (domain_models.ProductModel, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockTripRepository) GetDetail(ctx context.Context, id uuid.UUID) (domain_models.ProductModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain_models.ProductModel{}, args.Error(1)
	}
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockTripRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.ProductModel, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain_models.ProductModel), args.Error(1)
}

func (m *MockTripRepository) Update(ctx context.Context, id uuid.UUID, trip domain_models.ProductUpdateModel) (domain_models.ProductModel, error) {
	args := m.Called(ctx, id, trip)
	return args.Get(0).(domain_models.ProductModel), args.Error(1)
}

func (m *MockTripRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	now := time.Now()

	orderID := uuid.MustParse("9b825b1d-c744-4ca3-bac8-4aa1b13a98d3")
	links := []string{"https://some-link1.ru", "https://some-link2.ru"}
	incorrectImages := []string{"some-link1.ru", "some-link2.ru"}
	description := "Full desxcription"

	tests := []struct {
		name       string
		input      ProductModel
		mockResult domain_models.ProductModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful create without all fields",
			input: ProductModel{
				nil,
				nil,
				"Some product",
				nil,
				12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				nil,
				nil,
				nil,
			},
			mockResult: domain_models.ProductModel{
				uuid.New(),
				nil,
				"Some product",
				nil,
				12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				nil,
				now,
				nil,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "successful create with all fields",
			input: ProductModel{
				nil,
				&orderID,
				"Some product",
				&description,
				12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				&links,
				nil,
				nil,
			},
			mockResult: domain_models.ProductModel{
				uuid.New(),
				&orderID,
				"Some product",
				&description,
				12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				&links,
				now,
				nil,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "create with short name field",
			input: ProductModel{
				nil,
				nil,
				"So",
				nil,
				12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with negative price field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"RUR",
				"https://some-url.ru",
				"Shop name",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect currency field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"TR",
				"https://some-url.ru",
				"Shop name",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create without shop_url field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"TR",
				"",
				"Shop name",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect shop_url field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"TR",
				"shop-url.ru",
				"Shop name",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create without shop_name field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"TR",
				"https://some-url.ru",
				"",
				nil,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect image_urls field",
			input: ProductModel{
				nil,
				nil,
				"Some name",
				nil,
				-12.3,
				"TR",
				"https://some-url.ru",
				"",
				&incorrectImages,
				&now,
				nil,
			},
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			if !tt.wantErr {
				mockRepo.On("Create", ctx, mock.Anything).Return(tt.mockResult, tt.mockErr)
			}
			service := NewService(mockRepo, slog.Default())
			result, err := service.Create(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ProductModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.Equal(t, tt.input.Name, result.Name)
				if tt.input.OrderID != nil {
					assert.Equal(t, tt.input.OrderID, result.OrderID)
				}
				if tt.input.Description != nil {
					assert.Equal(t, tt.input.Description, result.Description)
				}
				assert.Equal(t, tt.input.Price, result.Price)
				assert.Equal(t, tt.input.Currency, result.Currency)
				assert.Equal(t, tt.input.ShopUrl, result.ShopUrl)
				assert.Equal(t, tt.input.ShopName, result.ShopName)
				if tt.input.ImageUrls != nil {
					assert.Equal(t, *tt.input.ImageUrls, *result.ImageUrls)
				}

				assert.NotEqual(t, uuid.Nil, result.ID, "ID should be generated")
				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
				assert.Nil(t, result.UpdatedAt, "UpdatedAt should be nil")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ID := uuid.New()
	newName := "New name"
	newDescription := "New description"
	newPrice := 21.36
	newCurrency := "USD"
	newShopUrl := "https://new-link.ru"
	newShopName := "Some new shop name"
	newImageUrls := []string{"https://link.ru"}
	incorrectName := "S"
	incorrectDescription := "M"
	incorrectPrice := -1.0
	incorrectCurrency := "TR"
	incorrectShopUrl := "shop-url.ru"
	incorrectShopName := "N"
	incorrectImageUrls := []string{"link.ru"}
	tests := []struct {
		name       string
		input      ProductUpdateModel
		successful domain_models.ProductModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful update",
			input: ProductUpdateModel{
				&newName,
				&newDescription,
				&newPrice,
				&newCurrency,
				&newShopUrl,
				&newShopName,
				&newImageUrls,
			},
			successful: domain_models.ProductModel{
				ID,
				nil,
				newName,
				&newDescription,
				newPrice,
				newCurrency,
				newShopUrl,
				newShopName,
				&newImageUrls,
				now,
				&now,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "update incorrect name",
			input: ProductUpdateModel{
				Name: &incorrectName,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect description",
			input: ProductUpdateModel{
				Description: &incorrectDescription,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect price",
			input: ProductUpdateModel{
				Price: &incorrectPrice,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect currency",
			input: ProductUpdateModel{
				Currency: &incorrectCurrency,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect shop_url",
			input: ProductUpdateModel{
				ShopUrl: &incorrectShopUrl,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect shop_name",
			input: ProductUpdateModel{
				ShopName: &incorrectShopName,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect image_urls",
			input: ProductUpdateModel{
				ImageUrls: &incorrectImageUrls,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			if !tt.wantErr {
				mockRepo.On("Update", ctx, ID, mock.Anything).Return(tt.successful, tt.mockErr)
			}
			service := NewService(mockRepo, slog.Default())
			result, err := service.Update(ctx, ID, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ProductModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.Nil(t, result.OrderID, "OrderID should be nil")
				assert.Equal(t, *tt.input.Name, result.Name)
				assert.Equal(t, *tt.input.Description, *result.Description)
				assert.Equal(t, *tt.input.Price, result.Price)
				assert.Equal(t, *tt.input.ShopUrl, result.ShopUrl)
				assert.Equal(t, *tt.input.ShopName, result.ShopName)
				assert.Equal(t, *tt.input.ImageUrls, *result.ImageUrls)

				assert.Equal(t, ID, *result.ID)
				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
				assert.NotNil(t, result.UpdatedAt, "UpdatedAt should be not nil")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetDetail(t *testing.T) {
	now := time.Now()
	ID := uuid.New()
	description := "desc"
	images := []string{"https://image.ru"}

	tests := []struct {
		name       string
		id         uuid.UUID
		mockResult domain_models.ProductModel
		mockErr    error
		wantErr    bool
		wantNil    bool
	}{
		{
			name: "successful get product",
			id:   ID,
			mockResult: domain_models.ProductModel{
				ID,
				nil,
				"Some",
				&description,
				2.3,
				"USD",
				"https://link.ru",
				"Name of shop",
				&images,
				now,
				&now,
			},
			mockErr: nil,
			wantErr: false,
			wantNil: false,
		},
		{
			name:       "product not found",
			id:         uuid.New(),
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorNotFound,
			wantErr:    true,
			wantNil:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockResult, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			result, err := service.GetDetail(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ProductModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.Nil(t, result.OrderID, "OrderID should be nil")
				assert.Equal(t, tt.mockResult.Name, result.Name)
				assert.Equal(t, *tt.mockResult.Description, *result.Description)
				assert.Equal(t, tt.mockResult.Price, result.Price)
				assert.Equal(t, tt.mockResult.Currency, result.Currency)
				assert.Equal(t, tt.mockResult.ShopUrl, result.ShopUrl)
				assert.Equal(t, tt.mockResult.ShopName, result.ShopName)
				assert.Equal(t, *tt.mockResult.ImageUrls, *result.ImageUrls)

				assert.Equal(t, ID, *result.ID)
				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
				assert.NotNil(t, result.UpdatedAt, "UpdatedAt should be not nil")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetList(t *testing.T) {
	now := time.Now()
	orderID := uuid.New()
	description := "descr"
	images := []string{"https://link.ru"}

	tests := []struct {
		name        string
		limit       int
		offset      int
		mockResults []domain_models.ProductModel
		mockErr     error
		wantCount   int
		wantErr     bool
	}{
		{
			name: "successful get all",
			mockResults: []domain_models.ProductModel{
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
			},
			mockErr:   nil,
			wantCount: 3,
			wantErr:   false,
		},
		{
			name:  "successful get one record",
			limit: 1,
			mockResults: []domain_models.ProductModel{
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
			},
			mockErr:   nil,
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:   "successful get with offset",
			offset: 1,
			mockResults: []domain_models.ProductModel{
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
				{
					uuid.New(),
					&orderID,
					"Addidas",
					&description,
					125.36,
					"USD",
					"https://addidas.com",
					"addidas",
					&images,
					now,
					nil,
				},
			},
			mockErr:   nil,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:        "get list with empty filters",
			mockResults: []domain_models.ProductModel{},
			mockErr:     nil,
			wantCount:   0,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetList", ctx, tt.limit, tt.offset).Return(tt.mockResults, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			results, err := service.GetList(ctx, tt.limit, tt.offset)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ProductModel{}, results)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCount, len(results))
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDelete(t *testing.T) {
	ID := uuid.New()
	ID2 := uuid.New()
	now := time.Now()
	description := "desc"
	tests := []struct {
		name       string
		id         uuid.UUID
		mockResult domain_models.ProductModel
		mockErr    error
		wantErr    bool
		callDetete bool
	}{
		{
			name: "successful delete",
			id:   ID,
			mockResult: domain_models.ProductModel{
				ID,
				nil,
				"UK",
				&description,
				23.1,
				"USD",
				"link",
				"name",
				nil,
				now,
				nil,
			},
			mockErr:    nil,
			wantErr:    false,
			callDetete: true,
		},
		{
			name:       "delete non exists product",
			id:         uuid.New(),
			mockResult: domain_models.ProductModel{},
			mockErr:    api_errors.ErrorNotFound,
			wantErr:    true,
			callDetete: true,
		},
		{
			name: "delete product which in order",
			id:   ID2,
			mockResult: domain_models.ProductModel{
				ID2,
				&ID2,
				"UK",
				&description,
				23.1,
				"USD",
				"link",
				"name",
				nil,
				now,
				nil,
			},
			mockErr:    api_errors.ErrorCantDeleteProduct,
			wantErr:    true,
			callDetete: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockResult, nil)
			if tt.callDetete {
				mockRepo.On("Delete", ctx, tt.id).Return(tt.mockErr)
			}

			service := NewService(mockRepo, slog.Default())
			err := service.Delete(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
