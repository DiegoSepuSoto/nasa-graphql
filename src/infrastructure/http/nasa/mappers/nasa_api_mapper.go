package mappers

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/entities"
)

func ImagesToDomain(apiResponse *entities.MarsRoverImages) []*models.Image {
	images := make([]*models.Image, 0)

	for _, photo := range apiResponse.LatestPhotos {
		images = append(images, &models.Image{
			Link: photo.ImgSrc,
			Camera: models.Camera{
				ID:   photo.Camera.ID,
				Name: photo.Camera.FullName,
			},
			Rover:         models.Rover{
				ID:   photo.Rover.ID,
				Name: photo.Rover.Name,
			},
			FormattedDate: photo.EarthDate,
		})
	}

	return images
}
