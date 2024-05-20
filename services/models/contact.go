package models

import (
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Id        uuid.UUID
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Companies []Company
	CreatedAt time.Time
	UpdatedAt time.Time
}
