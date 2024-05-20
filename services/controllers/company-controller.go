package controllers

import (
	"bcs/data"
	"bcs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterCompanyRoutes(router *gin.RouterGroup, repo *data.CompanyRepository) {
	companyRouter := router.Group("/companies")
	{
		companyRouter.POST("/", func(c *gin.Context) {
			createCompany(c, repo)
		})
		companyRouter.GET("/", func(c *gin.Context) {
			getCompanies(c, repo)
		})
		companyRouter.GET("/:id", func(c *gin.Context) {
			getCompanyByID(c, repo)
		})
		companyRouter.PUT("/:id", func(c *gin.Context) {
			updateCompany(c, repo)
		})
		companyRouter.DELETE("/:id", func(c *gin.Context) {
			deleteCompany(c, repo)
		})
	}
}

func createCompany(c *gin.Context, repo *data.CompanyRepository) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.Id = uuid.New()
	if err := repo.CreateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, company)
}

func getCompanies(c *gin.Context, repo *data.CompanyRepository) {
	companies, err := repo.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func getCompanyByID(c *gin.Context, repo *data.CompanyRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	company, err := repo.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func updateCompany(c *gin.Context, repo *data.CompanyRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company.Id = id
	if err := repo.UpdateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func deleteCompany(c *gin.Context, repo *data.CompanyRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := repo.DeleteCompany(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
