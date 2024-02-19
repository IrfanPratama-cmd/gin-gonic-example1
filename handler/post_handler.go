package handler

import (
	"fmt"
	"gin-socmed/lib"
	"gin-socmed/model"
	"gin-socmed/service"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postHandler struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *postHandler {
	return &postHandler{
		service: service,
	}
}

func (h *postHandler) Create(c *gin.Context) {
	var post model.PostRequest

	if err := c.ShouldBind(&post); err != nil {
		lib.HandleError(c, &lib.BadRequestError{Message: err.Error()})
		return
	}

	if post.Picture != nil {
		if err := os.MkdirAll("/public/picture", 0755); err != nil {
			lib.HandleError(c, &lib.InternalServerError{Message: err.Error()})
			return
		}

		// Rename Picture
		ext := filepath.Ext(post.Picture.Filename)
		newFileName := uuid.New().String() + ext

		// Save Image to Directory
		dst := filepath.Join("public/picture", filepath.Base(newFileName))
		c.SaveUploadedFile(post.Picture, dst)

		post.Picture.Filename = fmt.Sprintf("%s/public/picture/%s", c.Request.Host, newFileName)
	}

	userID, _ := c.Get("userID")
	post.UserID = userID.(*uuid.UUID)

	if err := h.service.Create(&post); err != nil {
		lib.HandleError(c, err)
		return
	}

	res := lib.Response(model.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Success post your tweet",
		Data:       post,
	})

	c.JSON(http.StatusCreated, res)
}
