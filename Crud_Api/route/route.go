package route

import (
	"example/crud_api/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/person", controller.GetAllPersonsController)
	r.GET("/person/:id", controller.GetPersonByIDController)
	r.POST("/person", controller.CreatePersonController)
	r.PUT("/person/:id", controller.UpdatePersonController)
	r.DELETE("/person/:id", controller.DeletePersonController)

	// Handle non-existing endpoints
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
	})

	return r
}
