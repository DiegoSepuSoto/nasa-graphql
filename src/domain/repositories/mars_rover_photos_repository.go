package repositories

import "github.com/diegosepusoto/nasa-graph-ql/src/domain/models"

type MarsRoverPhotosRepository interface {
	GetMarsRoverPhotos() ([]*models.Photos, error)
}
