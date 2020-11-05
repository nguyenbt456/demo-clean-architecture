package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/demo-clean-architecture/model"
	"github.com/nguyenbt456/demo-clean-architecture/util"
)

type UserHandler struct {
	userService model.UserService
}

func NewUserHandler(userService model.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetByID(c *gin.Context) {
	userID := c.Param("user_id")

	user, err := h.userService.GetByID(userID)
	if err != nil {
		util.HandleError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
