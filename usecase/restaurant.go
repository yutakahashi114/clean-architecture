package usecase

import (
	"context"

	"github.com/yutakahashi114/clean-architecture/domain/model/restaurant"
)

// RestaurantUseCase .
type RestaurantUseCase interface {
	GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, error)
}

// NewRestaurantUseCase .
func NewRestaurantUseCase(restaurantRepository restaurant.Repository) RestaurantUseCase {
	return &restaurantUseCase{
		restaurantRepository: restaurantRepository,
	}
}

type restaurantUseCase struct {
	restaurantRepository restaurant.Repository
}

// GetByIDInput .
type GetByIDInput struct {
	ID uint64
}

func (in GetByIDInput) toDomain() (restaurant.ID, error) {
	return restaurant.NewID(in.ID)
}

// GetByIDOutput .
type GetByIDOutput struct {
	ID        uint64
	Name      string
	Tags      []string
	ClientUID uint64
}

func (u *restaurantUseCase) GetByID(ctx context.Context, in GetByIDInput) (*GetByIDOutput, error) {

	id, err := in.toDomain()
	if err != nil {
		return nil, err
	}

	r, err := u.restaurantRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetByIDOutput{
		ID:        r.ID.Uint64(),
		Name:      r.Name.String(),
		Tags:      r.Tags.Strings(),
		ClientUID: r.ClientUID.Uint64(),
	}, nil
}
