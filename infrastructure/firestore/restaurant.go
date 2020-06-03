package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/yutakahashi114/clean-architecture/domain/model/client"
	"github.com/yutakahashi114/clean-architecture/domain/model/restaurant"
)

type repository struct {
	client *firestore.Client
}

// NewRestaurantRepository .
func NewRestaurantRepository(client *firestore.Client) restaurant.Repository {
	return &repository{client}
}

type restaurantDTO struct {
	ID        uint64   `firestore:"-"`
	Name      string   `firestore:"name"`
	Tags      []string `firestore:"tags"`
	ClientUID uint64   `firestore:"client_uid"`
}

func (r *restaurantDTO) toEntity() (*restaurant.Restaurant, error) {
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
	tags, err := restaurant.NewTags(r.Tags)
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

func (r *repository) GetByID(ctx context.Context, id restaurant.ID) (*restaurant.Restaurant, error) {
	doc, err := r.client.Doc("restaurants/" + id.String()).Get(ctx)
	if err != nil {
		return nil, err
	}

	rst := restaurantDTO{}

	err = doc.DataTo(&rst)
	if err != nil {
		return nil, err
	}

	return rst.toEntity()
}
