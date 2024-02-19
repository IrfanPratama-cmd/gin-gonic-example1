package model

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        *uuid.UUID     `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Before Create Data
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if b.ID != nil {
		return nil
	}
	id, err := uuid.NewRandom()
	b.ID = &id

	return err
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	if nil != b.ID {
		return nil
	}
	now := strfmt.DateTime(time.Now().UTC())
	b.UpdatedAt = time.Time(now)
	return nil
}
