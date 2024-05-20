package data

import (
	"bcs/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactRepository struct {
	DB *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{DB: db}
}

func (r *ContactRepository) CreateContact(contact *models.Contact) error {
	return r.DB.Create(contact).Error
}

func (r *ContactRepository) GetContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	err := r.DB.Find(&contacts).Error
	return contacts, err
}

func (r *ContactRepository) GetContactByID(id uuid.UUID) (*models.Contact, error) {
	var contact models.Contact
	err := r.DB.First(&contact, "id = ?", id).Error
	return &contact, err
}

func (r *ContactRepository) UpdateContact(contact *models.Contact) error {
	return r.DB.Save(contact).Error
}

func (r *ContactRepository) DeleteContact(id uuid.UUID) error {
	return r.DB.Delete(&models.Contact{}, "id = ?", id).Error
}
