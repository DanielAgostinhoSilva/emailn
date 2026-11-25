package vo

import "errors"

type Name struct {
	value string
}

func (n Name) Value() string {
	return n.value
}

func NewName(value string) (*Name, error) {
	if value == "" {
		return nil, errors.New("name is required")
	}
	if len(value) < 3 {
		return nil, errors.New("name cannot be shorter than 3 characters")
	}

	return &Name{value: value}, nil
}
