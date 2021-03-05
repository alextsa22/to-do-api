package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCod int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCod, Error{Message: message})
}
