package nasa

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/diegosepusoto/nasa-graph-ql/src/domain/models"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/entities"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/mappers"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa/utils"
	client "github.com/pzentenoe/httpclient-call-go"
)

type APIRepository struct {
	httpClient *client.HTTPClientCall
}

func NewNasaAPIRepository(httpClient *client.HTTPClientCall) *APIRepository {
	return &APIRepository{httpClient: httpClient}
}

func (r *APIRepository) GetMarsRoverPhotos() ([]*models.Photo, error) {
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

	if response.StatusCode != http.StatusOK {
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
