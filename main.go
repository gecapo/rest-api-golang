package main

import (
	contextdb "github.com/gecapo/rest-api-golang/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	contextdb.ConnectDataBase()

	r.Run()
}
