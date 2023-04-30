package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/models"
	"github.com/nexentra/inteligpt/pkg/common/utils"
	"github.com/nexentra/inteligpt/pkg/httputil"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) RegisterUser(ctx *gin.Context) {

	body := models.User{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Email = body.Email
	user.UserName = body.UserName

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hashedPassword)

	result := h.DB.Create(&user)

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
