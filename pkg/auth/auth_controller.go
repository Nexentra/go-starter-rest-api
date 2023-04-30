package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/middlware/jwt"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	auth := router.Group("/auth/")
	{
        auth.POST("register", h.RegisterUser)
        auth.POST("login", h.LoginUser)
		user := auth.Group("/user/")
		user.Use(jwt.AuthRequired())
        user.GET(":id", h.GetUser)
	}
}