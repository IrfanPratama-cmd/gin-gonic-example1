package handler

import (
	"gin-socmed/lib"
	"gin-socmed/model"
	"gin-socmed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var register model.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		lib.HandleError(c, &lib.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		lib.HandleError(c, err)
		return
	}

	res := lib.Response(model.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register Succesfully",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var login model.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		lib.HandleError(c, &lib.BadRequestError{Message: err.Error()})
	}

	result, err := h.service.Login(&login)
	if err != nil {
		lib.HandleError(c, err)
		return
	}

	res := lib.Response(model.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login Succesfull",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}
