package handler

import (
	"context"

	"github.com/yutakahashi114/clean-architecture/controller/grpc/handler/proto"
	"github.com/yutakahashi114/clean-architecture/usecase"
)

type restaurantHandler struct {
	restaurantUseCase usecase.RestaurantUseCase
}

// NewRestaurantHandler .
func NewRestaurantHandler(restaurantUseCase usecase.RestaurantUseCase) proto.RestaurantServiceServer {
	return &restaurantHandler{
		restaurantUseCase: restaurantUseCase,
	}
}

func (h *restaurantHandler) GetRestaurantByID(ctx context.Context, req *proto.GetRestaurantByIDRequest) (*proto.GetRestaurantByIDResponse, error) {

	out, err := h.restaurantUseCase.GetByID(ctx, usecase.GetByIDInput{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &proto.GetRestaurantByIDResponse{
		Restaurant: &proto.Restaurant{
			Id:        out.ID,
			Name:      out.Name,
			Tags:      out.Tags,
			ClientUid: out.ClientUID,
		},
	}, nil
}
