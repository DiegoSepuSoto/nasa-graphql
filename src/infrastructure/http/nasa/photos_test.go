package nasa

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	client "github.com/pzentenoe/httpclient-call-go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectedPhotos = []*models.Photos{
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

func Test_nasaAPIRepository_GetMarsRoverPhotos(t *testing.T) {
	t.Parallel()

	t.Run("when getMarsRoverPhotos executes successfully", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{\"latest_photos\":[{\"id\":817017,\"sol\":3061,\"camera\":{\"id\":23,\"name\":\"CHEMCAM\",\"rover_id\":5,\"full_name\":\"Chemistry and Camera Complex\"},\"img_src\":\"https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG\",\"earth_date\":\"2021-03-17\",\"rover\":{\"id\":5,\"name\":\"Curiosity\",\"landing_date\":\"2012-08-06\",\"launch_date\":\"2011-11-26\",\"status\":\"active\"}}]}"))
		}))

		defer server.Close()

		httpClient := client.NewHTTPClientCall(server.Client()).Host(server.URL)
		nasaAPIRepository := NewNasaAPIRepository(httpClient)

		photos, err := nasaAPIRepository.GetMarsRoverPhotos()

		assert.NoError(t, err)
		assert.Equal(t, expectedPhotos, photos)
	})

	t.Run("when getMarsRoverPhotos executes with error in server", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{\"latest_photos\":[{\"id\":817017,\"sol\":3061,\"camera\":{\"id\":23,\"name\":\"CHEMCAM\",\"rover_id\":5,\"full_name\":\"Chemistry and Camera Complex\"},\"img_src\":\"https://mars.nasa.gov/msl-raw-photos/proj/msl/redops/ods/surface/sol/03061/opgs/edr/ccam/CR0_669228842EDR_F0870792CCAM01061M_.JPG\",\"earth_date\":\"2021-03-17\",\"rover\":{\"id\":5,\"name\":\"Curiosity\",\"landing_date\":\"2012-08-06\",\"launch_date\":\"2011-11-26\",\"status\":\"active\"}}]}"))
		}))

		defer server.Close()

		httpClient := client.NewHTTPClientCall(server.Client()).Host("")
		nasaAPIRepository := NewNasaAPIRepository(httpClient)

		photos, err := nasaAPIRepository.GetMarsRoverPhotos()

		assert.Error(t, err)
		assert.Nil(t, photos)
		assert.Equal(t, "empty host", err.Error())
	})

	t.Run("when getMarsRoverPhotos executes with wrong HTTP code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadGateway)
		}))

		defer server.Close()

		httpClient := client.NewHTTPClientCall(server.Client()).Host(server.URL)
		nasaAPIRepository := NewNasaAPIRepository(httpClient)

		photos, err := nasaAPIRepository.GetMarsRoverPhotos()

		assert.Error(t, err)
		assert.Nil(t, photos)
		assert.Equal(t, "the server answered with a wrong http status code", err.Error())
	})

	t.Run("when getMarsRoverPhotos executes with error reading the body", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("&"))
		}))

		defer server.Close()

		httpClient := client.NewHTTPClientCall(server.Client()).Host(server.URL)
		nasaAPIRepository := NewNasaAPIRepository(httpClient)

		photos, err := nasaAPIRepository.GetMarsRoverPhotos()

		assert.Error(t, err)
		assert.Nil(t, photos)
		assert.Equal(t, "invalid character '&' looking for beginning of value", err.Error())
	})

	t.Run("when getMarsRoverPhotos executes with error unmarshalling to json", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("{\"version\":\"1\",\"alive\":true?}"))
		}))

		defer server.Close()

		httpClient := client.NewHTTPClientCall(server.Client()).Host(server.URL)
		nasaAPIRepository := NewNasaAPIRepository(httpClient)

		photos, err := nasaAPIRepository.GetMarsRoverPhotos()

		assert.Error(t, err)
		assert.Nil(t, photos)
		assert.Equal(t, "invalid character '?' after object key:value pair", err.Error())
	})
}
