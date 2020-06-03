package restaurant

import "context"

// Repository .
type Repository interface {
	GetByID(context.Context, ID) (*Restaurant, error)
}
