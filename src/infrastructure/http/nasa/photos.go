package nasa

import (
	"encoding/json"
	"errors"
	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/entities"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/mappers"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/utils"
	client "github.com/pzentenoe/httpclient-call-go"
	"io/ioutil"
	"net/http"
)

type nasaAPIRepository struct {
	httpClient *client.HTTPClientCall
}

func NewNasaAPIRepository(httpClient *client.HTTPClientCall) *nasaAPIRepository {
	return &nasaAPIRepository{httpClient: httpClient}
}

func (r *nasaAPIRepository) GetMarsRoverPhotos() ([]*models.Photos, error) {

	response, err := r.httpClient.
		Headers(http.Header{"Content-Type": {"application/json; charset=UTF-8"}}).
		Path("/rovers/curiosity/latest_photos").
		Method("GET").
		Params(utils.APIKeyParam()).
		Do()

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("the server answered with a wrong http status code")
	}

	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return nil, err2
	}

	var nasaAPIResponse *entities.MarsRoverPhotos
	if err3 := json.Unmarshal(data, &nasaAPIResponse); err3 != nil {
		return nil, err3
	}

	return mappers.PhotosToDomain(nasaAPIResponse), nil
}
