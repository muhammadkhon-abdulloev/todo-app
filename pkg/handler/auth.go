package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammadkhon-abdulloev/todo-app"
	// "github.com/sirupsen/logrus"
)



func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":id,
	})
}

type SingInInput struct {
	Userame string `json:"username" blinding:"required"`
	Password string `json:"password" blinding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	token, err := h.services.Authorization.GenerateToken(input.Userame, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":token,
	})
}