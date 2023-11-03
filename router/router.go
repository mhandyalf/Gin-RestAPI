package router

import (
	"restapi/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB) {
	AuthHandler := handler.NewAuthHandler(db)

	r.POST("/register", AuthHandler.RegisterUser)
	r.POST("/login", AuthHandler.LoginUser)
}
