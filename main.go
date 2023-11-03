package main

import (
	"restapi/database"
	"restapi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.InitDB()

	router.InitRouter(r, db)

	r.Run(":8080")

}
