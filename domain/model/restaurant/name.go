package restaurant

import "fmt"

// Name .
type Name string

// NewName .
func NewName(name string) (Name, error) {
	if name == "" {
		return "", fmt.Errorf("invalid name")
	}
	return Name(name), nil
}

// String .
func (name Name) String() string {
	return string(name)
}
