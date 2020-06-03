package grpc

import (
	"context"

	"github.com/yutakahashi114/clean-architecture/domain/model/client"
	"github.com/yutakahashi114/clean-architecture/domain/model/restaurant"
	"github.com/yutakahashi114/clean-architecture/infrastructure/grpc/proto"
	"google.golang.org/grpc"
)

type repository struct {
	client proto.OtherRestaurantServiceClient
}

// NewRestaurantRepository .
func NewRestaurantRepository(conn *grpc.ClientConn) restaurant.Repository {
	return &repository{
		client: proto.NewOtherRestaurantServiceClient(conn),
	}
}

func toEntity(r *proto.Restaurant) (*restaurant.Restaurant, error) {
	id, err := restaurant.NewID(r.GetId())
	if err != nil {
		return nil, err
	}
	name, err := restaurant.NewName(r.GetName())
	if err != nil {
		return nil, err
	}
	clientUID, err := client.NewID(r.GetClientUid())
	if err != nil {
		return nil, err
	}
	tags, err := restaurant.NewTags(r.GetTags())
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
	res, err := r.client.GetOtherRestaurantByID(ctx, &proto.GetOtherRestaurantByIDRequest{
		Id: id.Uint64(),
	})
	if err != nil {
		return nil, err
	}

	return toEntity(res.GetRestaurant())
}
