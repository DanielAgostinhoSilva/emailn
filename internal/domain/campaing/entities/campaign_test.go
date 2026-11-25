package entities

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type CampaignSuite struct {
	suite.Suite
}

func (s *CampaignSuite) TestNewCampaignIsValid() {
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(time.Hour)
	campaign, err := NewCampaign("Campaign A", []string{"a@example.com", "b@example.com"}, "Hello", createdAt, updatedAt)

	s.Require().NoError(err)
	s.Require().NotNil(campaign)
	_, uuidErr := uuid.Parse(campaign.id.Value())
	s.NoError(uuidErr)
	s.Equal("Campaign A", campaign.name.Value())
	s.Require().Len(campaign.contacts, 2)
	s.Equal("a@example.com", campaign.contacts[0].Value())
	s.Equal("b@example.com", campaign.contacts[1].Value())
	s.Equal("Hello", campaign.content.Value())
	s.Equal(createdAt, campaign.createdAt)
	s.Equal(updatedAt, campaign.updatedAt)
}

func (s *CampaignSuite) TestNewCampaignRejectsInvalidValues() {
	tests := []struct {
		name         string
		campaignName string
		contacts     []string
		content      string
		error        string
	}{
		{"name", "Ca", nil, "Hello", "campaign.name: name cannot be shorter than 3 characters"},
		{"contact", "Campaign A", []string{"a@example.com", "invalid"}, "Hello", "campaign.contacts[1]: contact must be a valid email address"},
		{"nil contacts", "Campaign A", nil, "Hello", "campaign.contacts: at least one contact is required"},
		{"empty contacts", "Campaign A", []string{}, "Hello", "campaign.contacts: at least one contact is required"},
		{"duplicate contact", "Campaign A", []string{"a@example.com", "b@example.com", "a@example.com"}, "Hello", "campaign.contacts[2]: duplicate contact"},
		{"duplicate contact ignoring case", "Campaign A", []string{"a@example.com", "A@Example.com"}, "Hello", "campaign.contacts[1]: duplicate contact"},
		{"content", "Campaign A", []string{"a@example.com"}, "", "campaign.content: email content is required"},
	}
	for _, test := range tests {
		s.Run(test.name, func() {
			campaign, err := NewCampaign(test.campaignName, test.contacts, test.content, time.Time{}, time.Time{})
			s.Nil(campaign)
			s.Require().EqualError(err, test.error)
		})
	}
}

func (s *CampaignSuite) TestNewCampaignRejectsUpdatedAtBeforeCreatedAt() {
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(-time.Second)

	campaign, err := NewCampaign("Campaign A", []string{"a@example.com"}, "Hello", createdAt, updatedAt)

	s.Nil(campaign)
	s.Require().EqualError(err, "campaign.updatedAt: cannot be before createdAt")
}

func (s *CampaignSuite) TestNewCampaignAcceptsUpdatedAtEqualToCreatedAt() {
	createdAt := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	campaign, err := NewCampaign("Campaign A", []string{"a@example.com"}, "Hello", createdAt, createdAt)

	s.Require().NoError(err)
	s.NotNil(campaign)
}

func TestCampaignSuite(t *testing.T) {
	suite.Run(t, new(CampaignSuite))
}
