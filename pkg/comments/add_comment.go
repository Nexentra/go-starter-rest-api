package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/models"
	"github.com/nexentra/inteligpt/pkg/httputil"
)

// AddComment godoc
// @Summary      Add a comment
// @Description  add by json comment
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment  body      models.AddCommentRequestBody  true  "Add comment"
// @Success      201      {object}  models.Comment
// @Failure      400      {object}  httputil.HTTPError400
// @Failure      404      {object}  httputil.HTTPError404
// @Router       /comments [post]
func (h handler) AddComment(ctx *gin.Context) {

	body := models.AddCommentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		// ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var comment models.Comment

	// comment.ID = uuid.NewV4().String()
	comment.Title = body.Title
	comment.Author = body.Author
	comment.Slug = body.Slug

	if result := h.DB.Create(&comment); result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		// ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &comment)
}
