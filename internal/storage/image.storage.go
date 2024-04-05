package storage

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/2110366-2566-2/tortoise-app-backend/pkg/utils"
	"google.golang.org/api/iterator"
)

func (h *Handler) AddImage(ctx context.Context, name, folder, base64Image string) (string, error) {

	// split the base64 image string
	split_string, err := utils.ValidateBase64Image(base64Image)
	if err != nil {
		fmt.Println("Error validating image data:", err)
		return "", fmt.Errorf("invalid image data")
	}

	// Decode the Base64-encoded image
	imageData, err := base64.StdEncoding.DecodeString((*split_string)[1])
	if err != nil {
		fmt.Println("Error decoding image data:", err)
		return "", fmt.Errorf("error decoding image data")
	}

	imagePath := folder + "/" + name
	bucket := "petpal-cloud-storage.appspot.com"

	// Write the image data to the storage bucket
	wc := h.stg.Bucket(bucket).Object(imagePath).NewWriter(ctx)
	_, err = wc.Write(imageData)
	if err != nil {
		fmt.Println("Error writing image data:", err)
		return "", fmt.Errorf("error writing image data")
	}

	// Close the writer
	if err := wc.Close(); err != nil {
		fmt.Println("Error closing writer:", err)
		return "", fmt.Errorf("error closing writer")
	}

	// Generate the direct URL for the uploaded image
	escapedImagePath := url.PathEscape(imagePath)
	url := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", bucket, escapedImagePath)

	return url, nil
}

func (h *Handler) DeleteImage(ctx context.Context, name, folder string) error {
	bucket := "petpal-cloud-storage.appspot.com"
	imagePath := folder + "/" + name

	// Delete the object from the storage bucket
	if err := h.stg.Bucket(bucket).Object(imagePath).Delete(ctx); err != nil {
		fmt.Println("Error deleting image data:", err)
		return fmt.Errorf("error deleting image data")
	}

	return nil
}

func (h *Handler) DeleteFolder(ctx context.Context, folder string) error {
	// delete existing folder
	bucket := "petpal-cloud-storage.appspot.com"
	iter := h.stg.Bucket(bucket).Objects(ctx, &storage.Query{Prefix: folder + "/"})
	for {
		objAttrs, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("error deleting image data")
		}
		if err := h.stg.Bucket(bucket).Object(objAttrs.Name).Delete(ctx); err != nil {
			return fmt.Errorf("error deleting image data")
		}
	}
	return nil
}
