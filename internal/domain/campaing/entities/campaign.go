package entities

import (
	"emailn/internal/domain/campaing/vo"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	id        vo.ID
	name      vo.Name
	contacts  []vo.Contact
	content   vo.EmailContent
	createdAt time.Time
	updatedAt time.Time
}

func (c Campaign) Id() vo.ID {
	return c.id
}

func (c Campaign) Name() vo.Name {
	return c.name
}

func (c Campaign) Contacts() []vo.Contact {
	return c.contacts
}

func (c Campaign) Content() vo.EmailContent {
	return c.content
}

func (c Campaign) CreatedAt() time.Time {
	return c.createdAt
}

func (c Campaign) UpdatedAt() time.Time {
	return c.updatedAt
}

func NewCampaign(
	name string,
	contacts []string,
	content string,
	createdAt time.Time,
	updatedAt time.Time,
) (*Campaign, error) {
	idVo, err := vo.NewID(uuid.NewString())
	if err != nil {
		return nil, fmt.Errorf("campaign.id: %w", err)
	}
	nameVo, err := vo.NewName(name)
	if err != nil {
		return nil, fmt.Errorf("campaign.name: %w", err)
	}
	if len(contacts) == 0 {
		return nil, errors.New("campaign.contacts: at least one contact is required")
	}
	contactsVo := make([]vo.Contact, 0, len(contacts))
	seen := make(map[string]struct{}, len(contacts))
	for i, contact := range contacts {
		contactVo, err := vo.NewContact(contact)
		if err != nil {
			return nil, fmt.Errorf("campaign.contacts[%d]: %w", i, err)
		}
		key := strings.ToLower(contactVo.Value())
		if _, exists := seen[key]; exists {
			return nil, fmt.Errorf("campaign.contacts[%d]: duplicate contact", i)
		}
		seen[key] = struct{}{}
		contactsVo = append(contactsVo, *contactVo)
	}
	contentVo, err := vo.NewEmailContent(content)
	if err != nil {
		return nil, fmt.Errorf("campaign.content: %w", err)
	}
	if updatedAt.Before(createdAt) {
		return nil, errors.New("campaign.updatedAt: cannot be before createdAt")
	}
	return &Campaign{
		id:        *idVo,
		name:      *nameVo,
		contacts:  contactsVo,
		content:   *contentVo,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
