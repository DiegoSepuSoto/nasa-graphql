package photos

import (
	"errors"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var mockPhotos = []*models.Photo{
	{
		Link:          "https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG",
		Camera:        models.Camera{
			ID:   23,
			Name: "Chemistry and Camera Complex",
		},
		Rover:         models.Rover{
			ID:   5,
			Name: "Curiosity",
		},
		Date: "2021-03-17",
	},
}

var mockPhotosWithErrorInDate = []*models.Photo{
	{
		Link:          "https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG",
		Camera:        models.Camera{
			ID:   23,
			Name: "Chemistry and Camera Complex",
		},
		Rover:         models.Rover{
			ID:   5,
			Name: "Curiosity",
		},
		Date: "2021-0&-17",
	},
}

var expectedPhotos = []*models.Photo{
	{
		Link:          "https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG",
		Camera:        models.Camera{
			ID:   23,
			Name: "Chemistry and Camera Complex",
		},
		Rover:         models.Rover{
			ID:   5,
			Name: "Curiosity",
		},
		Date: "17/03/2021",
	},
}

type mockNASAAPIRepository struct {
	mock.Mock
}

func (r *mockNASAAPIRepository) GetMarsRoverPhotos() ([]*models.Photo, error) {
	m := r.Called()

	if m.Get(0) == nil {
		return nil, m.Error(1)
	}
	return m.Get(0).([]*models.Photo), m.Error(1)
}

func Test_photosUseCase_GetNasaPhotos(t *testing.T) {
	t.Parallel()

	t.Run("when GetNasaPhotos is successful", func(t *testing.T) {
		nasaAPIRepository := new(mockNASAAPIRepository)
		nasaAPIRepository.On("GetMarsRoverPhotos").Return(mockPhotos, nil)

		nasaPhotosUseCase := NewPhotosUseCase(nasaAPIRepository)

		photos, err := nasaPhotosUseCase.GetNasaPhotos()

		assert.NoError(t, err)
		assert.NotNil(t, photos)
		assert.Equal(t, expectedPhotos, photos)
	})

	t.Run("when nasaAPIRepository returns an error", func(t *testing.T) {
		nasaAPIRepository := new(mockNASAAPIRepository)
		nasaAPIRepository.On("GetMarsRoverPhotos").Return(nil, errors.New("error in repository"))

		nasaPhotosUseCase := NewPhotosUseCase(nasaAPIRepository)

		photos, err := nasaPhotosUseCase.GetNasaPhotos()

		assert.Error(t, err)
		assert.Nil(t, photos)
	})

	t.Run("when nasaAPIRepository returns one photo with an error in date", func(t *testing.T) {
		nasaAPIRepository := new(mockNASAAPIRepository)
		nasaAPIRepository.On("GetMarsRoverPhotos").Return(mockPhotosWithErrorInDate, nil)

		nasaPhotosUseCase := NewPhotosUseCase(nasaAPIRepository)

		photos, err := nasaPhotosUseCase.GetNasaPhotos()

		assert.NoError(t, err)
		assert.Equal(t, mockPhotosWithErrorInDate, photos)
	})
}
