package usecase

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/yutakahashi114/clean-architecture/domain/model/restaurant"
	mock "github.com/yutakahashi114/clean-architecture/domain/model/restaurant/mock"
)

func Test_restaurantUsecase_GetByID(t *testing.T) {

	type args struct {
		in GetByIDInput
	}

	cases := []struct {
		name   string
		args   args
		mock   func(*mock.MockRepository)
		want   *GetByIDOutput
		errMsg string
	}{
		{
			name:   "failed to domain",
			args:   args{in: GetByIDInput{ID: 0}},
			errMsg: "invalid id",
		},
		{
			name: "failed to get by id",
			args: args{in: GetByIDInput{ID: 1}},
			mock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(gomock.Any(), restaurant.ID(1)).Return(nil, fmt.Errorf("something"))
			},
			errMsg: "something",
		},
		{
			name: "success",
			args: args{in: GetByIDInput{ID: 1}},
			mock: func(mock *mock.MockRepository) {
				mock.EXPECT().GetByID(gomock.Any(), restaurant.ID(1)).Return(&restaurant.Restaurant{
					ID:        1,
					Name:      "name",
					Tags:      restaurant.Tags{"tag1", "Tag2"},
					ClientUID: 2,
				}, nil)
			},
			want: &GetByIDOutput{
				ID:        1,
				Name:      "name",
				Tags:      []string{"tag1", "Tag2"},
				ClientUID: 2,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			repository := mock.NewMockRepository(ctrl)

			if c.mock != nil {
				c.mock(repository)
			}
			u := restaurantUseCase{
				restaurantRepository: repository,
			}

			got, err := u.GetByID(ctx, c.args.in)

			assert.Equal(t, c.want, got)
			if c.errMsg != "" {
				assert.EqualError(t, err, c.errMsg)
			}
		})
	}
}
