package murid

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MuridHandler struct {
	service IMuridService
}

func NewMuridHandler(service IMuridService) *MuridHandler {
	return &MuridHandler{service: service}
}

func (h *MuridHandler) CreateMurid(c *gin.Context) {
	var murid Murid
	if err := c.ShouldBindJSON(&murid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateMurid(context.Background(), murid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, murid)
}

func (h *MuridHandler) GetMurid(c *gin.Context) {
	var murid Murid
	if err := c.ShouldBindJSON(&murid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.service.GetMurid(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

//func (h *MuridHandler) UpdateMurid(c *gin.Context) {
//	var murid Murid
//	if err := c.ShouldBindJSON(&murid); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	murid, err := h.service.UpdateMurid(context.Background(), murid)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, murid)
//}
