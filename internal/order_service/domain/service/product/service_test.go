package order_domain_product_service

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	domain_models "github.com/kadr/globe_express/internal/travaler_service/domain/models"
	api_errors "github.com/kadr/globe_express/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTripRepository struct {
	mock.Mock
}

func (m *MockTripRepository) Create(ctx context.Context, trip domain_models.TripModel) (domain_models.TripModel, error) {
	args := m.Called(ctx, trip)
	return args.Get(0).(domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) GetDetail(ctx context.Context, id uuid.UUID) (domain_models.TripModel, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return domain_models.TripModel{}, args.Error(1)
	}
	return args.Get(0).(domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) GetActive(ctx context.Context) ([]domain_models.TripModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) GetCompleted(ctx context.Context) ([]domain_models.TripModel, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) GetList(ctx context.Context, limit, offset int) ([]domain_models.TripModel, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) Update(ctx context.Context, id uuid.UUID, trip domain_models.TripUpdateModel) (domain_models.TripModel, error) {
	args := m.Called(ctx, id, trip)
	return args.Get(0).(domain_models.TripModel), args.Error(1)
}

func (m *MockTripRepository) Cancel(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	now := time.Now()

	travelerID := uuid.MustParse("9b825b1d-c744-4ca3-bac8-4aa1b13a98d3")
	tripID := uuid.MustParse("12345678-1234-1234-1234-123456789012")
	incorrectStatus := "active"

	tests := []struct {
		name       string
		inputTrip  TripModel
		mockResult domain_models.TripModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful create",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
			},
			mockResult: domain_models.TripModel{
				ID:            tripID,
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "London",
				Status:        "planned",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
				CreatedAt:     now,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "create with short from_country field",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "U",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
			},
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with short from_city field",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "Ne",
				ToCountry:     "UK",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
			},
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with short to_country field",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "U",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
			},
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with short to_city field",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "L",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
			},
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorFieldValidation,
			wantErr:    true,
		},
		{
			name: "create with incorrect status",
			inputTrip: TripModel{
				TravelerID:    travelerID,
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
				Status:        &incorrectStatus,
			},
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorIncorrectStatus,
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
			result, err := service.Create(ctx, tt.inputTrip)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.Equal(t, tt.inputTrip.TravelerID, result.TravelerID)
				assert.Equal(t, tt.inputTrip.FromCountry, result.FromCountry)
				assert.Equal(t, tt.inputTrip.FromCity, result.FromCity)
				assert.Equal(t, tt.inputTrip.ToCountry, result.ToCountry)
				assert.Equal(t, tt.inputTrip.ToCity, result.ToCity)
				assert.Equal(t, tt.mockResult.Status, *result.Status)
				assert.True(t, result.DepartureDate.Equal(tt.inputTrip.DepartureDate))
				assert.True(t, result.ArrivalDate.Equal(tt.inputTrip.ArrivalDate))

				assert.NotEqual(t, uuid.Nil, result.ID, "ID should be generated")
				assert.Equal(t, "planned", *result.Status, "Status should be 'planned'")
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
	newFromCountry := "Spain"
	newFromCity := "Madrid"
	newToCountry := "Germany"
	newToCity := "Born"
	newDepartureDate := now
	newArrivalDate := now.Add(56 * time.Hour)
	newWeight := 5.0
	newSize := 5.0
	newStatus := "ongoing"
	incorrectFromCountry := "S"
	incorrectFromCity := "M"
	incorrectToCountry := "G"
	incorrectToCity := "Bo"
	incorrectWeight := 105.0
	incorrectSize := 75.0
	incorrectStatus := "draft"
	tests := []struct {
		name       string
		inputTrip  TripUpdateModel
		successful domain_models.TripModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful update",
			inputTrip: TripUpdateModel{
				FromCountry:   &newFromCountry,
				FromCity:      &newFromCity,
				ToCountry:     &newToCountry,
				ToCity:        &newToCity,
				DepartureDate: &newDepartureDate,
				ArrivalDate:   &newArrivalDate,
				MaxWeightKG:   &newWeight,
				MaxSizeCM3:    &newSize,
				Status:        &newStatus,
			},
			successful: domain_models.TripModel{
				ID:            ID,
				TravelerID:    uuid.New(),
				FromCountry:   newFromCountry,
				FromCity:      newFromCity,
				ToCountry:     newToCountry,
				ToCity:        newToCity,
				DepartureDate: newDepartureDate,
				ArrivalDate:   newArrivalDate,
				MaxWeightKG:   newWeight,
				MaxSizeCM3:    newSize,
				Status:        newStatus,
				CreatedAt:     now,
				UpdatedAt:     &now,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "update incorrect from_country",
			inputTrip: TripUpdateModel{
				FromCountry: &incorrectFromCountry,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect from_city",
			inputTrip: TripUpdateModel{
				FromCity: &incorrectFromCity,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect to_country",
			inputTrip: TripUpdateModel{
				ToCountry: &incorrectToCountry,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect to_city",
			inputTrip: TripUpdateModel{
				ToCity: &incorrectToCity,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect max_weight",
			inputTrip: TripUpdateModel{
				MaxWeightKG: &incorrectWeight,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect max_size",
			inputTrip: TripUpdateModel{
				MaxSizeCM3: &incorrectSize,
			},
			mockErr: api_errors.ErrorFieldValidation,
			wantErr: true,
		},
		{
			name: "update incorrect status",
			inputTrip: TripUpdateModel{
				Status: &incorrectStatus,
			},
			mockErr: api_errors.ErrorIncorrectStatus,
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
			result, err := service.Update(ctx, ID, tt.inputTrip)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.NotEqual(t, uuid.Nil, result.TravelerID, "TravalerID should be not nil")
				assert.Equal(t, *tt.inputTrip.FromCountry, result.FromCountry)
				assert.Equal(t, *tt.inputTrip.FromCity, result.FromCity)
				assert.Equal(t, *tt.inputTrip.ToCountry, result.ToCountry)
				assert.Equal(t, *tt.inputTrip.ToCity, result.ToCity)
				assert.Equal(t, *tt.inputTrip.Status, *result.Status)
				assert.True(t, result.DepartureDate.Equal(*tt.inputTrip.DepartureDate))
				assert.True(t, result.ArrivalDate.Equal(*tt.inputTrip.ArrivalDate))

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

	tests := []struct {
		name     string
		id       uuid.UUID
		mockTrip domain_models.TripModel
		mockErr  error
		wantErr  bool
		wantNil  bool
	}{
		{
			name: "successful get detail",
			id:   ID,
			mockTrip: domain_models.TripModel{
				ID:            ID,
				TravelerID:    uuid.New(),
				FromCountry:   "USA",
				FromCity:      "New York",
				ToCountry:     "UK",
				ToCity:        "London",
				DepartureDate: now,
				ArrivalDate:   now.Add(24 * time.Hour),
				Status:        "active",
				CreatedAt:     now,
				UpdatedAt:     &now,
			},
			mockErr: nil,
			wantErr: false,
			wantNil: false,
		},
		{
			name:     "trip not found",
			id:       uuid.New(),
			mockTrip: domain_models.TripModel{},
			mockErr:  api_errors.ErrorNotFound,
			wantErr:  true,
			wantNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockTrip, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			result, err := service.GetDetail(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, result)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)

				assert.NotEqual(t, uuid.Nil, result.TravelerID, "TravalerID should be not nil")
				assert.Equal(t, tt.mockTrip.FromCountry, result.FromCountry)
				assert.Equal(t, tt.mockTrip.FromCity, result.FromCity)
				assert.Equal(t, tt.mockTrip.ToCountry, result.ToCountry)
				assert.Equal(t, tt.mockTrip.ToCity, result.ToCity)
				assert.Equal(t, tt.mockTrip.Status, *result.Status)
				assert.True(t, result.DepartureDate.Equal(tt.mockTrip.DepartureDate))
				assert.True(t, result.ArrivalDate.Equal(tt.mockTrip.ArrivalDate))

				assert.Equal(t, ID, *result.ID)
				assert.NotZero(t, result.CreatedAt, "CreatedAt should be set")
				assert.NotNil(t, result.UpdatedAt, "UpdatedAt should be not nil")
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetActive(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		mockTrips []domain_models.TripModel
		mockErr   error
		wantCount int
		wantErr   bool
	}{
		{
			name: "successful get active trips",
			mockTrips: []domain_models.TripModel{
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "planned",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     nil,
				},
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "UK",
					FromCity:      "London",
					Status:        "ongoing",
					DepartureDate: now,
					ArrivalDate:   now.Add(48 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     &now,
				},
			},
			mockErr:   nil,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "no active trips",
			mockTrips: []domain_models.TripModel{},
			mockErr:   nil,
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetActive", ctx).Return(tt.mockTrips, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			results, err := service.GetActive(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, results)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tt.wantCount, len(results))
				for _, res := range results {
					assert.NotEqual(t, "completed", *res.Status)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetCompleted(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		mockTrips []domain_models.TripModel
		mockErr   error
		wantCount int
		wantErr   bool
	}{
		{
			name: "successful get active trips",
			mockTrips: []domain_models.TripModel{
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "completed",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     nil,
				},
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "UK",
					FromCity:      "London",
					Status:        "completed",
					DepartureDate: now,
					ArrivalDate:   now.Add(48 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     &now,
				},
			},
			mockErr:   nil,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "no active trips",
			mockTrips: []domain_models.TripModel{},
			mockErr:   nil,
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetCompleted", ctx).Return(tt.mockTrips, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			results, err := service.GetCompleted(ctx)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, results)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)

				assert.Equal(t, tt.wantCount, len(results))
				for _, res := range results {
					assert.Equal(t, "completed", *res.Status)
				}
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetList(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name      string
		limit     int
		offset    int
		mockTrips []domain_models.TripModel
		mockErr   error
		wantCount int
		wantErr   bool
	}{
		{
			name: "successful get all",
			mockTrips: []domain_models.TripModel{
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "planned",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     nil,
				},
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "ongoing",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     &now,
				},
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "completed",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     &now,
				},
			},
			mockErr:   nil,
			wantCount: 3,
			wantErr:   false,
		},
		{
			name:  "successful get one record",
			limit: 1,
			mockTrips: []domain_models.TripModel{
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "planned",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     nil,
				},
			},
			mockErr:   nil,
			wantCount: 1,
			wantErr:   false,
		},
		{
			name:   "successful get with offset",
			offset: 1,
			mockTrips: []domain_models.TripModel{
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "planned",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     nil,
				},
				{
					ID:            uuid.New(),
					TravelerID:    uuid.New(),
					FromCountry:   "USA",
					FromCity:      "New York",
					Status:        "completed",
					DepartureDate: now,
					ArrivalDate:   now.Add(24 * time.Hour),
					CreatedAt:     now,
					UpdatedAt:     &now,
				},
			},
			mockErr:   nil,
			wantCount: 2,
			wantErr:   false,
		},
		{
			name:      "get list with empty filters",
			mockTrips: []domain_models.TripModel{},
			mockErr:   nil,
			wantCount: 0,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetList", ctx, tt.limit, tt.offset).Return(tt.mockTrips, tt.mockErr)

			service := NewService(mockRepo, slog.Default())
			results, err := service.GetList(ctx, tt.limit, tt.offset)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, TripModel{}, results)
				assert.ErrorIs(t, err, tt.mockErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantCount, len(results))
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCancel(t *testing.T) {
	ID := uuid.New()
	ID2 := uuid.New()
	now := time.Now()
	tests := []struct {
		name       string
		id         uuid.UUID
		mockResult domain_models.TripModel
		mockErr    error
		wantErr    bool
	}{
		{
			name: "successful cancel",
			id:   ID,
			mockResult: domain_models.TripModel{
				ID:            ID,
				TravelerID:    uuid.New(),
				FromCountry:   "UK",
				FromCity:      "London",
				ToCountry:     "USA",
				ToCity:        "New York",
				DepartureDate: time.Now(),
				ArrivalDate:   time.Now().Add(54 * time.Hour),
				CreatedAt:     now,
				Status:        "planned",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:       "cancel non-existent trip",
			id:         uuid.New(),
			mockResult: domain_models.TripModel{},
			mockErr:    api_errors.ErrorNotFound,
			wantErr:    true,
		},
		{
			name: "cancel not planned trip",
			id:   ID2,
			mockResult: domain_models.TripModel{
				ID:            ID2,
				TravelerID:    uuid.New(),
				FromCountry:   "UK",
				FromCity:      "London",
				ToCountry:     "USA",
				ToCity:        "New York",
				DepartureDate: time.Now(),
				ArrivalDate:   time.Now().Add(54 * time.Hour),
				CreatedAt:     now,
				Status:        "ongoing",
			},
			mockErr: api_errors.ErrorCanNotCancel,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mockRepo := new(MockTripRepository)

			mockRepo.On("GetDetail", ctx, tt.id).Return(tt.mockResult, nil)
			if !tt.wantErr {
				mockRepo.On("Cancel", ctx, tt.id).Return(tt.mockErr)
			}

			service := NewService(mockRepo, slog.Default())
			err := service.Cancel(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
