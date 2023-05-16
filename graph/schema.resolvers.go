package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/savvie-com/subgraph-shoplist/graph/generated"
	"github.com/savvie-com/subgraph-shoplist/graph/model"
)

// Foo is the resolver for the foo field.
func (r *queryResolver) Foo(ctx context.Context, id string) (*model.Foo, error) {
	return FindFoo(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
