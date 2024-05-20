package data

import (
	"bcs/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var DBConn *gorm.DB

//  func connectToDb() *gorm.DB {

// 	dsn := "host=localhost user=gorm password=letmein dbname=bcs port=5432 sslmode=disable TimeZone=eastern"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

//  	if err != nil {
//  		panic(err)
//  	}

//  	DBConn = db

//  	return db
//  }

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (r *CompanyRepository) CreateCompany(company *models.Company) error {
	return r.DB.Create(company).Error
}

func (r *CompanyRepository) GetCompanies() ([]models.Company, error) {
	var companies []models.Company
	err := r.DB.Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	err := r.DB.First(&company, "id = ?", id).Error
	return &company, err
}

func (r *CompanyRepository) UpdateCompany(company *models.Company) error {
	return r.DB.Save(company).Error
}

func (r *CompanyRepository) DeleteCompany(id uuid.UUID) error {
	return r.DB.Delete(&models.Company{}, "id = ?", id).Error
}
