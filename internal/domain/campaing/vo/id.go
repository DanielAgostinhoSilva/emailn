package vo

import "errors"

type ID struct {
	value string
}

func (I ID) Value() string {
	return I.value
}

func NewID(value string) (*ID, error) {
	if value == "" {
		return nil, errors.New("id is required")
	}
	return &ID{value: value}, nil
}
