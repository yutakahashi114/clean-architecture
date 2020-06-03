package restaurant

import "github.com/yutakahashi114/clean-architecture/domain/model/client"

// Restaurant .
type Restaurant struct {
	ID        ID
	Name      Name
	Tags      Tags
	ClientUID client.ID
}

// NewRestaurant .
func NewRestaurant(
	id ID,
	name Name,
	tags Tags,
	clientUID client.ID,
) Restaurant {
	return Restaurant{
		ID:        id,
		Name:      name,
		Tags:      tags,
		ClientUID: clientUID,
	}
}
