package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type statusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCod int, message string) {
	if gin.Mode() != gin.TestMode {
		logrus.Error(message)
	}
	c.AbortWithStatusJSON(statusCod, ErrorResponse{Message: message})
}
