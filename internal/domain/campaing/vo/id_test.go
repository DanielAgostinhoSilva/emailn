package vo

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type IDSuite struct {
	suite.Suite
}

func (s *IDSuite) TestNewIDValido() {
	id, err := NewID("123")
	s.Require().NoError(err)
	s.Require().NotNil(id)
	s.Equal("123", id.Value())
}

func (s *IDSuite) TestQuandoIdForVazioDeveLancarUmErro() {
	_, err := NewID("")
	s.Require().Error(err)
	s.Require().Equal("id is required", err.Error())
}

func TestIDSuite(t *testing.T) {
	suite.Run(t, new(IDSuite))
}
