package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/generated"
	"github.com/diegosepusoto/nasa-graph-ql/src/infrastructure/graph/nasa/model"
)

func (r *queryResolver) Photos(ctx context.Context) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
