package comments

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
	// Simple group: v1
	v1 := router.Group("/api/v1/")
	{
		comments := v1.Group("/comments")
		{
			comments.POST("", h.AddComment)
			comments.GET("", h.GetComments)
			comments.GET("/:id", h.GetComment)
			comments.PUT("/:id", h.UpdateComment)
			comments.DELETE("/:id", h.DeleteComment)
		}
	}

	// exactly the same as:
	dashboard := router.Group("/dashboard/")
{
	fineTunes := dashboard.Group("/fine-tunes")
	fineTunes.Use(jwt.AuthRequired())
	{
		fineTunes.POST("", h.AddComment)
		fineTunes.GET("", h.GetComments)
		fineTunes.GET("/:id", h.GetComment)
		fineTunes.PUT("/:id", h.UpdateComment)
		fineTunes.DELETE("/:id", h.DeleteComment)
	}
}	
}