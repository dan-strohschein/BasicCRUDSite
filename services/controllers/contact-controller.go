package controllers

import (
	"bcs/data"
	"bcs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterContactRoutes(router *gin.RouterGroup, repo *data.ContactRepository) {
	contactRouter := router.Group("/contacts")
	{
		contactRouter.POST("/", func(c *gin.Context) {
			createContact(c, repo)
		})
		contactRouter.GET("/", func(c *gin.Context) {
			getContacts(c, repo)
		})
		contactRouter.GET("/:id", func(c *gin.Context) {
			getContactByID(c, repo)
		})
		contactRouter.PUT("/:id", func(c *gin.Context) {
			updateContact(c, repo)
		})
		contactRouter.DELETE("/:id", func(c *gin.Context) {
			deleteContact(c, repo)
		})
	}
}

func createContact(c *gin.Context, repo *data.ContactRepository) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact.Id = uuid.New()
	if err := repo.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, contact)
}

func getContacts(c *gin.Context, repo *data.ContactRepository) {
	contacts, err := repo.GetContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}

func getContactByID(c *gin.Context, repo *data.ContactRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	contact, err := repo.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func updateContact(c *gin.Context, repo *data.ContactRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contact.Id = id
	if err := repo.UpdateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func deleteContact(c *gin.Context, repo *data.ContactRepository) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := repo.DeleteContact(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
