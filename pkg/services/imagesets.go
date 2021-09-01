package services

import (
	"context"

	"github.com/redhatinsights/edge-api/pkg/clients/imagebuilder"
	"github.com/redhatinsights/edge-api/pkg/models"
)

// ImageSetsServiceInterface defines the interface that helps handle
// the business logic of ImageSets
type ImageSetsServiceInterface interface {
	GetAllImageSets(image *models.Image, account string) error
}

// NewImageSetsService gives a instance of the main implementation of a ImageSetsServiceInterface
func NewImageSetsService(ctx context.Context) ImageSetsServiceInterface {
	return &ImageSetsService{imageBuilder: imagebuilder.InitClient(ctx)}
}

// ImageSetsService is the main implementation of a ImageSetsServiceInterface
type ImageSetsService struct {
	ctx          context.Context
	imageBuilder imagebuilder.ClientInterface
}

// GetAllImageSets to org group of images into one
func (s *ImageSetsService) GetAllImageSets(image *models.Image, account string) error {

	return nil
}
