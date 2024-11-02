package model

import (
	"time"

	"github.com/google/uuid"
	
)

type Base struct {
	
	Id        uuid.UUID `gorm:"primary_key: unique;type:uuid;default:uuid_generate_v4()" json:"id"` 
	CreatedAt time.Time `gorm:"type:timestamp(6)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp(6)" json:"updated_at"`
}

