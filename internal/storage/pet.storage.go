package storage

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

func (h *Handler) AddPetMedia(petID, base64Image string) (string, error) {
	// Add base64 media to the storage
	// Decode the Base64-encoded image
	imageData, err := base64.StdEncoding.DecodeString(strings.Split(base64Image, ";base64,")[1])
	if err != nil {
		fmt.Println("Error decoding image data:", err)
		return "", fmt.Errorf("Error decoding image data")
	}

	imagePath := "pets/" + petID + ".jpg"

	bucket := "petpal-cloud-storage.appspot.com"
	wc := h.stg.Bucket(bucket).Object(imagePath).NewWriter(context.Background())
	_, err = wc.Write(imageData)
	if err != nil {
		fmt.Println("Error writing image data:", err)
		return "", fmt.Errorf("Error writing image data")
	}
	if err := wc.Close(); err != nil {
		fmt.Println("Error closing writer:", err)
		return "", fmt.Errorf("Error closing writer")
	}

	// Generate the direct URL for the uploaded image
	escapedImagePath := url.PathEscape(imagePath)
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucket, escapedImagePath)

	return url, nil
}
