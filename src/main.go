package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	useCase "github.com/diegosepusoto/nasa-graph-ql/src/application/usecase/photos"
	graphqlNasa "github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/generated"
	httpNasa "github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/http/nasa"
	client "github.com/pzentenoe/httpclient-call-go"
	"log"
	"net/http"
	"os"
)

func main() {
	httpClient := client.NewHTTPClientCall(&http.Client{}).Host(os.Getenv("NASA_API_HOST"))

	photosRepo := httpNasa.NewNasaAPIRepository(httpClient)

	photosUseCase := useCase.NewPhotosUseCase(photosRepo)

	photosResolver := graphqlNasa.NewPhotosResolver(photosUseCase)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: photosResolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
}