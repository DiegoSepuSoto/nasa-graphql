package repositories

import "github.com/diegosepusoto/nasa-graph-ql/src/domain/models"

type MarsRoverImagesRepository interface {
	GetMarsRoverPhotos() ([]*models.Image, error)
}
