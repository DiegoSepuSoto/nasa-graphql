package nasa

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/model"
)

func PhotosToSchema(domainPhotos []*models.Photo) []*model.Photo {
	schemaPhotos := make([]*model.Photo, 0)

	for _, photo := range domainPhotos {
		schemaPhotos = append(schemaPhotos, &model.Photo{
			Link:   photo.Link,
			Camera: cameraToSchema(photo.Camera),
			Rover:  roverToSchema(photo.Rover),
			Date:   photo.Date,
		})
	}

	return schemaPhotos
}

func cameraToSchema(domainCamera models.Camera) *model.Camera {
	return &model.Camera{
		ID:   domainCamera.ID,
		Name: domainCamera.Name,
	}
}

func roverToSchema(domainRover models.Rover) *model.Rover {
	return &model.Rover{
		ID:   domainRover.ID,
		Name: domainRover.Name,
	}
}
