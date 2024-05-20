package models

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	Id        uuid.UUID
	Name      string
	Address   string
	Phone     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
