package services

import (
	"github.com/2110366-2566-2/tortoise-app-backend/internal/models"
	"github.com/2110366-2566-2/tortoise-app-backend/internal/storage"
	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	handler *storage.Handler
}

func NewStorageHandler(handler *storage.Handler) *StorageHandler {
	return &StorageHandler{handler: handler}
}

// UploadFile godoc
// @Method POST
// @Summary Upload file
// @Description Upload file to storage
// @Endpoint /api/v1/upload
func (h *StorageHandler) UploadFile(c *gin.Context) {
	// Get base64 file from request
	var petMedia models.PetMedia
	if err := c.ShouldBindJSON(&petMedia); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Upload file to storage
	url, err := h.handler.AddPetMedia(petMedia.ID, petMedia.Media)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "url": url})
}
