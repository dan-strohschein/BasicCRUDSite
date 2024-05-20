package main

import (
	"bcs/controllers"
	"bcs/data"
	"bcs/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// var c = models.Company{Id: uuid.New(), Name: "Test", Address: "123 happy lane", Phone: "", Email: "d@d.com"}

	// data.CreateCompany(c)

	dsn := "host=localhost user=dev password=letmein dbname=bcs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure UUID extension is enabled in PostgreSQL
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.AutoMigrate(&models.Company{}, &models.Contact{}, &models.Service{})

	companyRepo := data.NewCompanyRepository(db)
	contactRepo := data.NewContactRepository(db)
	serviceRepo := data.NewServiceRepository(db)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		controllers.RegisterCompanyRoutes(v1, companyRepo)
		controllers.RegisterContactRoutes(v1, contactRepo)
		controllers.RegisterServiceRoutes(v1, serviceRepo)
	}

	router.Run(":8080")
}
