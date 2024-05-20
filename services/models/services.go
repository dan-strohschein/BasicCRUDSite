package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Id          uuid.UUID
	Name        string
	Description string
	PriceRange  string
	Companies   []Company
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
