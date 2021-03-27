package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	nasa2 "github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/generated"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa"
	client "github.com/pzentenoe/httpclient-call-go"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := client.NewHTTPClientCall(&http.Client{}).Host(os.Getenv("NASA_API_HOST"))

	photosRepo := nasa.NewNasaAPIRepository(httpClient)

	photos, _ := photosRepo.GetMarsRoverPhotos()

	fmt.Println(photos[0].Link)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &nasa2.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
