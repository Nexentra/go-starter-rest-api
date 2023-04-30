package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/models"
)

func (h handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	userName, ok := ctx.Get("user_name")

	var user models.User

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if result := h.DB.First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if userName != user.UserName {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	response := models.UserResponse{
        ID:       user.ID,
        UserName: user.UserName,
        Email:    user.Email,
    }

	ctx.JSON(http.StatusOK, response)
}
