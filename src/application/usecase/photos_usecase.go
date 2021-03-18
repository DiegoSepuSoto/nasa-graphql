package usecase

import "github.com/diegosepusoto/nasa-graph-ql/src/domain/models"

type PhotosUseCase interface {
	GetNasaPhotos() ([]*models.Image, error)
}
