package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/yutakahashi114/clean-architecture/domain/model/client"
	"github.com/yutakahashi114/clean-architecture/domain/model/restaurant"
)

type restaurantRepository struct {
	db DB
}

// NewRestaurantRepository .
func NewRestaurantRepository(db DB) restaurant.Repository {
	return &restaurantRepository{db}
}

func (r *restaurantRepository) GetByID(ctx context.Context, id restaurant.ID) (*restaurant.Restaurant, error) {
	rst := restaurantDTO{}
	has, err := r.db.engine.Where("id = ?", id).Get(&rst)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("not found")
	}

	rts := restaurantTagDTOs{}
	err = r.db.engine.Where("restaurant_id = ?", id).Find(&rts)
	if err != nil {
		return nil, err
	}

	return rst.toEntity(rts)
}

type restaurantDTO struct {
	ID        uint64 `xorm:"autoincr"`
	Name      string
	ClientUID uint64
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (r *restaurantDTO) toEntity(rts restaurantTagDTOs) (*restaurant.Restaurant, error) {
	id, err := restaurant.NewID(r.ID)
	if err != nil {
		return nil, err
	}
	name, err := restaurant.NewName(r.Name)
	if err != nil {
		return nil, err
	}
	clientUID, err := client.NewID(r.ClientUID)
	if err != nil {
		return nil, err
	}
	tags, err := rts.toTags()
	if err != nil {
		return nil, err
	}
	rst := restaurant.NewRestaurant(
		id,
		name,
		tags,
		clientUID,
	)
	return &rst, nil
}

type restaurantTagDTO struct {
	ID           uint64 `xorm:"autoincr"`
	restaurantID uint64
	Name         string
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
	DeletedAt    time.Time `xorm:"deleted"`
}

type restaurantTagDTOs []restaurantTagDTO

func (rts restaurantTagDTOs) toTags() (restaurant.Tags, error) {
	tags := make(restaurant.Tags, len(rts))
	for i, rt := range rts {
		tag, err := restaurant.NewTag(rt.Name)
		if err != nil {
			return nil, err
		}
		tags[i] = tag
	}
	return tags, nil
}
