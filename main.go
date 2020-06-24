package main

import (
	"github.com/gecapo/rest-api-golang/controllers"
	contextdb "github.com/gecapo/rest-api-golang/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	contextdb.ConnectDataBase()

	r.GET("/recipe", controllers.FindRecipes)
	r.POST("/recipe", controllers.CreateRecipe)
	r.GET("/recipe/:id", controllers.FindRecipe)
	r.PUT("/recipe/:id", controllers.UpdateRecipe)
	r.PATCH("/recipe/:id", controllers.PatchRecipe)

	r.Run()
}
