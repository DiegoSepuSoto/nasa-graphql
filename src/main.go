package main

import (
	"fmt"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa"
	client "github.com/pzentenoe/httpclient-call-go"
	"net/http"
)

func main() {
	httpClient := client.NewHTTPClientCall(&http.Client{})

	imagesRepo := nasa.NewNasaAPIRepository(httpClient)

	images, err := imagesRepo.GetMarsRoverImages()

	if err != nil {
		fmt.Sprintln("something wrong happened calling the API")
	}

	fmt.Println(images[0].Link)
}
