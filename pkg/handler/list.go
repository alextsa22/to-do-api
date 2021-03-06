package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	// check userIdentity handler
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
