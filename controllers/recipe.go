package controllers

import (
	"net/http"

	contextdb "github.com/gecapo/rest-api-golang/database"
	"github.com/gecapo/rest-api-golang/database/models"
	"github.com/gin-gonic/gin"
)

// GET /recipe
// Find all books
func FindRecipes(c *gin.Context) {
	var recipes []models.Recipe
	contextdb.DB.Find(&recipes)

	c.JSON(http.StatusOK, gin.H{"data": recipes})
}
