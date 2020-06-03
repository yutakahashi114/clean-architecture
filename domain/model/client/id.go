package client

import "fmt"

// ID .
type ID uint64

// NewID .
func NewID(id uint64) (ID, error) {
	if id <= 0 {
		return 0, fmt.Errorf("invalid id")
	}
	return ID(id), nil
}

// Uint64 .
func (id ID) Uint64() uint64 {
	return uint64(id)
}
