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

func (u *photosUseCase) GetNasaPhotos() ([]*models.Photos, error) {
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

		photo = &models.Photos{
			Link:   photo.Link,
			Camera: photo.Camera,
			Rover:  photo.Rover,
			Date:   formattedDate,
		}
	}

	return photos, nil
}