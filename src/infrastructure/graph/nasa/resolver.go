package nasa

import (
	"github.com/diegosepusoto/nasa-graph-ql/src/application/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	photosUseCase usecase.PhotosUseCase
}

func NewPhotosResolver(photosUseCase usecase.PhotosUseCase) *Resolver {
	return &Resolver{photosUseCase: photosUseCase}
}