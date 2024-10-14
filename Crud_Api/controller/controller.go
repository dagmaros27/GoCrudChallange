package controller

import (
	"example/crud_api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPersonsController handles GET /person
func GetAllPersonsController(c *gin.Context) {
	persons := model.GetAllPersons()
	c.JSON(http.StatusOK, persons)
}

// GetPersonByIDController handles GET /person/:id
func GetPersonByIDController(c *gin.Context) {
	id := c.Param("id")
	person, exists := model.GetPersonByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// CreatePersonController handles POST /person
func CreatePersonController(c *gin.Context) {
	var newPerson model.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdPerson := model.CreatePerson(newPerson)
	c.JSON(http.StatusCreated, createdPerson)
}

// UpdatePersonController handles PUT /person/:id
func UpdatePersonController(c *gin.Context) {
	id := c.Param("id")
	var updatedPerson model.Person
	if err := c.ShouldBindJSON(&updatedPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	person, exists := model.UpdatePerson(id, updatedPerson)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// DeletePersonController handles DELETE /person/:id
func DeletePersonController(c *gin.Context) {
	id := c.Param("id")
	if !model.DeletePerson(id) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
