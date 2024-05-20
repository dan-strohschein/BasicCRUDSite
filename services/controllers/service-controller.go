package controllers

import (
	"net/http"

	"bcs/data"
	"bcs/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterServiceRoutes(router *gin.RouterGroup, repo *data.ServiceRepository) {
	serviceRouter := router.Group("/services")
	{
		serviceRouter.POST("/", func(c *gin.Context) {
			createService(c, repo)
		})
		serviceRouter.GET("/", func(c *gin.Context) {
			getServices(c, repo)
		})
		serviceRouter.GET("/:id", func(c *gin.Context) {
			getServiceByID(c, repo)
		})
		serviceRouter.PUT("/:id", func(c *gin.Context) {
			updateService(c, repo)
		})
		serviceRouter.DELETE("/:id", func(c *gin.Context) {
			deleteService(c, repo)
		})
	}
}

func createService(c *gin.Context, repo *data.ServiceRepository) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.Id = uuid.New()
	if err := repo.CreateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, service)
}

func getServices(c *gin.Context, repo *data.ServiceRepository) {
	services, err := repo.GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}

func getServiceByID(c *gin.Context, repo *data.ServiceRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	service, err := repo.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

func updateService(c *gin.Context, repo *data.ServiceRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.Id = id
	if err := repo.UpdateService(&service); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service)
}

func deleteService(c *gin.Context, repo *data.ServiceRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := repo.DeleteService(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
