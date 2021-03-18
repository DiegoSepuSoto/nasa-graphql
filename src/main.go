package main

import (
	"fmt"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa"
	client "github.com/pzentenoe/httpclient-call-go"
	"net/http"
	"os"
)

func main() {
	httpClient := client.NewHTTPClientCall(&http.Client{}).Host(os.Getenv("NASA_API_HOST"))

	photosRepo := nasa.NewNasaAPIRepository(httpClient)

	photos, err := photosRepo.GetMarsRoverPhotos()

	if err != nil {
		fmt.Sprintln("something wrong happened calling the API")
	}

	fmt.Println(photos[0].Link)
}
