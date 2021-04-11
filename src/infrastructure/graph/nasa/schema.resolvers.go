package nasa

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/generated"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/model"
)

type queryResolver struct {
	*Resolver
}

func (r *queryResolver) Photos(ctx context.Context) ([]*model.Photo, error) {
	photos, err := r.photosUseCase.GetNasaPhotos()
	if err != nil {
		return nil, err
	}

	return PhotosToSchema(photos), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
