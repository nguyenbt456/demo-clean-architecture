package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenbt456/demo-clean-architecture/model"
	"github.com/nguyenbt456/demo-clean-architecture/util"
)

type BookHandler struct {
	bookService model.BookService
}

func NewBookHandler(bService model.BookService) *BookHandler {
	return &BookHandler{bService}
}

func (h *BookHandler) GetByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	books, err := h.bookService.GetByUserID(userID)
	if err != nil {
		util.HandleError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, books)
}
