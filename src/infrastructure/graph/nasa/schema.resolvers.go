package nasa

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/generated"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/model"
)

func (r *queryResolver) Photos(ctx context.Context) ([]*model.Photo, error) {
	mockPhotos := []*model.Photo{{
		Link:   "link to mock photos[0]",
		Camera: nil,
		Rover:  nil,
		Date:   "date to mock photos[0]",
	}}
	return mockPhotos, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
