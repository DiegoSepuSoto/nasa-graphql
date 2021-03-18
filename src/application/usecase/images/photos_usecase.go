package images

import (
	"fmt"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/repositories"
	"github.com/diegosepusoto/nasa-graph-ql/src/utils"
)

type photosUseCase struct {
	nasaAPIRepository repositories.MarsRoverImagesRepository
}

func NewPhotosUseCase(nasaAPIRepository repositories.MarsRoverImagesRepository) *photosUseCase {
	return &photosUseCase{nasaAPIRepository: nasaAPIRepository}
}

func (u *photosUseCase) GetNasaPhotos() ([]*models.Image, error) {
	images, err := u.nasaAPIRepository.GetMarsRoverPhotos()
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		formattedDate, err := utils.FormatDate(image.Date, utils.DateYMDFormatHyphen, utils.DateDMYFormatSlash)
		if err != nil {
			fmt.Sprintf("could not format date: %s", image.Date)
			formattedDate = image.Date
		}

		image = &models.Image{
			Link:   image.Link,
			Camera: image.Camera,
			Rover:  image.Rover,
			Date:   formattedDate,
		}
	}

	return images, nil
}