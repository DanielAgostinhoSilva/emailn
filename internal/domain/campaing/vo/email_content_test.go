package vo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type EmailContentSuite struct {
	suite.Suite
}

func (s *EmailContentSuite) TestEmailContentIsValid() {
	content, err := NewEmailContent("Hello, world!")
	s.Require().NoError(err)
	s.Require().NotNil(content)
	s.Equal("Hello, world!", content.Value())
}

func (s *EmailContentSuite) TestEmailContentIsEmpty() {
	content, err := NewEmailContent("")
	s.Require().Error(err)
	s.Nil(content)
	s.EqualError(err, "email content is required")
}

func TestEmailContentSuite(t *testing.T) {
	suite.Run(t, new(EmailContentSuite))
}