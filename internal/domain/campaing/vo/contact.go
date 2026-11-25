package vo

import (
	"errors"
	"net/mail"
)

type Contact struct {
	value string
}

func (c Contact) Value() string {
	return c.value
}

func NewContact(value string) (*Contact, error) {
	address, err := mail.ParseAddress(value)
	if err != nil || address.Address != value {
		return nil, errors.New("contact must be a valid email address")
	}

	return &Contact{value: value}, nil
}
