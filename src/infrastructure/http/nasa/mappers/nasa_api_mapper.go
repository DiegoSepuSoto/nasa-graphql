package mappers

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/entities"
)

func PhotosToDomain(apiResponse *entities.MarsRoverPhotos) []*models.Photo {
	photos := make([]*models.Photo, 0)

	for i := range apiResponse.LatestPhotos {
		photos = append(photos, &models.Photo{
			Link: apiResponse.LatestPhotos[i].ImgSrc,
			Camera: models.Camera{
				ID:   apiResponse.LatestPhotos[i].Camera.ID,
				Name: apiResponse.LatestPhotos[i].Camera.FullName,
			},
			Rover: models.Rover{
				ID:   apiResponse.LatestPhotos[i].Rover.ID,
				Name: apiResponse.LatestPhotos[i].Rover.Name,
			},
			Date: apiResponse.LatestPhotos[i].EarthDate,
		})
	}

	return photos
}
