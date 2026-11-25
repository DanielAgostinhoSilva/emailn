package vo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NameSuite struct {
	suite.Suite
}

func (s *NameSuite) TestNameIsValid() {
	name, err := NewName("Campaign A")
	s.Require().NoError(err)
	s.Require().NotNil(name)
	s.Equal("Campaign A", name.Value())
}

func (s *NameSuite) TestNameIsInvalid() {
	name, err := NewName("")
	s.Require().Error(err)
	s.Require().Nil(name)
	s.Require().Equal("name is required", err.Error())
}

func (s *NameSuite) TestNameIsTooShort() {
	name, err := NewName("Ca")
	s.Require().Error(err)
	s.Require().Nil(name)
	s.Require().Equal("name cannot be shorter than 3 characters", err.Error())
}

func TestNameSuite(t *testing.T) {
	suite.Run(t, new(NameSuite))
}
