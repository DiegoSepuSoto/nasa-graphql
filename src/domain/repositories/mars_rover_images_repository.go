package repositories

import "github.com/diegosepusoto/nasa-graph-ql/src/domain/models"

type MarsRoverImagesRepository interface {
	GetMarsRoverImages() ([]*models.Image, error)
}
