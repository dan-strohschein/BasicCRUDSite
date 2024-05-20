package data

import (
	"bcs/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	DB *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

func (r *ServiceRepository) CreateService(service *models.Service) error {
	return r.DB.Create(service).Error
}

func (r *ServiceRepository) GetServices() ([]models.Service, error) {
	var services []models.Service
	err := r.DB.Find(&services).Error
	return services, err
}

func (r *ServiceRepository) GetServiceByID(id uuid.UUID) (*models.Service, error) {
	var service models.Service
	err := r.DB.First(&service, "id = ?", id).Error
	return &service, err
}

func (r *ServiceRepository) UpdateService(service *models.Service) error {
	return r.DB.Save(service).Error
}

func (r *ServiceRepository) DeleteService(id uuid.UUID) error {
	return r.DB.Delete(&models.Service{}, "id = ?", id).Error
}
