package mappers

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/entities"
)

func PhotosToDomain(apiResponse *entities.MarsRoverPhotos) []*models.Photos {
	photos := make([]*models.Photos, 0)

	for _, photo := range apiResponse.LatestPhotos {
		photos = append(photos, &models.Photos{
			Link: photo.ImgSrc,
			Camera: models.Camera{
				ID:   photo.Camera.ID,
				Name: photo.Camera.FullName,
			},
			Rover:         models.Rover{
				ID:   photo.Rover.ID,
				Name: photo.Rover.Name,
			},
			Date: photo.EarthDate,
		})
	}

	return photos
}
