package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/models"
	"github.com/nexentra/inteligpt/pkg/common/utils"
	"github.com/nexentra/inteligpt/pkg/httputil"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) LoginUser(ctx *gin.Context) {

	body := models.LoginUserBody{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var storedUser models.User
	result := h.DB.Where("user_name = ?", body.UserName).First(&storedUser)
	if result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(body.Password)); err != nil {
		httputil.NewError(ctx, http.StatusUnauthorized, err)
		return
	}

	user := storedUser

	if result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		return
	}
	if token, err := utils.CreateToken(user.UserName); err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"token": token})
	}
}
