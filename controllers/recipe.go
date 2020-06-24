package controllers

import (
	"net/http"

	contextdb "github.com/gecapo/rest-api-golang/database"
	"github.com/gecapo/rest-api-golang/database/models"
	"github.com/gin-gonic/gin"
)

// Recipe api input model
type Recipe struct {
	Name         string `json:"name" binding:"required"`
	Instructions string `json:"instructions" binding:"required"`
}

// FindRecipes ..
// GET /recipe
// Find all recipes
func FindRecipes(c *gin.Context) {
	var recipes []models.RecipeDBO
	contextdb.DB.Find(&recipes)

	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

// CreateRecipe ...
// POST /recipe
// Create new recipe
func CreateRecipe(c *gin.Context) {
	var input Recipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipe := models.RecipeDBO{Name: input.Name, Instructions: input.Instructions}
	contextdb.DB.Create(&recipe)

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// FindRecipe ...
// GET /recipe/:id
// Find a recipe
func FindRecipe(c *gin.Context) { // Get model if exist
	var recipe models.RecipeDBO

	if err := contextdb.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// UpdateRecipe ...
// PUT /recipe/:id
// Update a recipe
func UpdateRecipe(c *gin.Context) {
	var recipe models.RecipeDBO
	if err := contextdb.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input Recipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contextdb.DB.Model(&recipe).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// PatchRecipe ...
// PATCH /recipe/:id
// Update a recipe
func PatchRecipe(c *gin.Context) {
	var recipe models.RecipeDBO
	if err := contextdb.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	type patchRecipe struct {
		Name         string `json:"name"`
		Instructions string `json:"instructions"`
	}

	// Validate input
	var input patchRecipe
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contextdb.DB.Model(&recipe).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// DeleteRecipe ...
// GET /recipe/:id
// Find a recipe
func DeleteRecipe(c *gin.Context) {
	//maybe soft delete
	var recipe models.RecipeDBO

	if err := contextdb.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	contextdb.DB.Delete(&recipe)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
