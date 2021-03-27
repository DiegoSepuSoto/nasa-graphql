package photos

import (
	"fmt"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/repositories"
	"github.com/diegosepusoto/nasa-graph-ql/src/utils"
)

type photosUseCase struct {
	nasaAPIRepository repositories.MarsRoverPhotosRepository
}

func NewPhotosUseCase(nasaAPIRepository repositories.MarsRoverPhotosRepository) *photosUseCase {
	return &photosUseCase{nasaAPIRepository: nasaAPIRepository}
}

func (u *photosUseCase) GetNasaPhotos() ([]*models.Photo, error) {
	photos, err := u.nasaAPIRepository.GetMarsRoverPhotos()
	if err != nil {
		return nil, err
	}

	for _, photo := range photos {
		formattedDate, err := utils.FormatDate(photo.Date, utils.DateYMDFormatHyphen, utils.DateDMYFormatSlash)
		if err != nil {
			fmt.Sprintf("could not format date: %s", photo.Date)
			formattedDate = photo.Date
		}

		photo.Date = formattedDate
	}

	return photos, nil
}