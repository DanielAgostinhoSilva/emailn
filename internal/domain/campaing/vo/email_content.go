package vo

import "errors"

type EmailContent struct {
	value string
}

func (e EmailContent) Value() string {
	return e.value
}

func NewEmailContent(value string) (*EmailContent, error) {
	if value == "" {
		return nil, errors.New("email content is required")
	}

	return &EmailContent{value: value}, nil
}
