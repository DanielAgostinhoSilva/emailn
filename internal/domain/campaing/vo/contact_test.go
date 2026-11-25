package vo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ContactSuite struct {
	suite.Suite
}

func (s *ContactSuite) TestContactIsValid() {
	for _, value := range []string{"user@example.com", "first.last+tag@example.co.uk"} {
		s.Run(value, func() {
			contact, err := NewContact(value)
			s.Require().NoError(err)
			s.Require().NotNil(contact)
			s.Equal(value, contact.Value())
		})
	}
}

func (s *ContactSuite) TestContactIsInvalid() {
	for _, value := range []string{"", "invalid", "user@", "@example.com", "User <user@example.com>", " user@example.com "} {
		s.Run(value, func() {
			contact, err := NewContact(value)
			s.Require().Error(err)
			s.Nil(contact)
		})
	}
}

func TestContactSuite(t *testing.T) {
	suite.Run(t, new(ContactSuite))
}