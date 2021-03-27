package nasa

import (
	"context"
	"errors"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var mockPhotos = []*models.Photo{
	{
		Link: "https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG",
		Camera: models.Camera{
			ID:   23,
			Name: "Chemistry and Camera Complex",
		},
		Rover: models.Rover{
			ID:   5,
			Name: "Curiosity",
		},
		Date: "17/03/2021",
	},
}

var expectedPhotos = []*model.Photo{
	{
		Link: "https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG",
		Camera: &model.Camera{
			ID:   23,
			Name: "Chemistry and Camera Complex",
		},
		Rover: &model.Rover{
			ID:   5,
			Name: "Curiosity",
		},
		Date: "17/03/2021",
	},
}

type mockPhotosUseCase struct {
	mock.Mock
}

func (u *mockPhotosUseCase) GetNasaPhotos() ([]*models.Photo, error) {
	m := u.Called()

	if m.Get(0) == nil {
		return nil, m.Error(1)
	}

	return m.Get(0).([]*models.Photo), nil
}

func Test_queryResolver_Photos(t *testing.T) {
	t.Parallel()

	t.Run("when 'Photos' query resolver executes successfully", func(t *testing.T) {
		nasaPhotosUseCase := new(mockPhotosUseCase)
		nasaPhotosUseCase.On("GetNasaPhotos").Return(mockPhotos, nil)
		var ctx context.Context

		photosResolver := NewPhotosResolver(nasaPhotosUseCase)
		photos, err := photosResolver.Query().Photos(ctx)

		assert.NoError(t, err)
		assert.NotNil(t, photos)
		assert.Equal(t, expectedPhotos, photos)
	})

	t.Run("when 'Photos' query resolver executes with an error in useCase", func(t *testing.T) {
		nasaPhotosUseCase := new(mockPhotosUseCase)
		nasaPhotosUseCase.On("GetNasaPhotos").Return(nil, errors.New("error in useCase"))
		var ctx context.Context

		photosResolver := NewPhotosResolver(nasaPhotosUseCase)
		photos, err := photosResolver.Query().Photos(ctx)

		assert.Error(t, err)
		assert.Nil(t, photos)
	})
}
