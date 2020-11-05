package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// eCode represent for error codes
type eCode struct {
	StatusCode   int `json:"status_code"`
	InternalCode int `json:"internal_code"`
}

// errorMap is map which contains errors
var errorMap = map[string]eCode{
	"record not found":                  eCode{http.StatusBadRequest, 1001},
	"UserID is null":                    eCode{http.StatusBadRequest, 2001},
	"UserID is invalid":                 eCode{http.StatusBadRequest, 2002},
	"BookID is invalid":                 eCode{http.StatusBadRequest, 2003},
	"Can't get user from database":      eCode{http.StatusServiceUnavailable, 9001},
	"Can't get book from database":      eCode{http.StatusServiceUnavailable, 9002},
	"Can't get user_book from database": eCode{http.StatusServiceUnavailable, 9003},
}

// HandleError handle difference error types
func HandleError(c *gin.Context, message string) {
	if errorMap[message] == (eCode{}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "9999",
			"message": message,
		})
		return
	}

	c.JSON(errorMap[message].StatusCode, gin.H{
		"code":    errorMap[message].InternalCode,
		"message": message,
	})
}
