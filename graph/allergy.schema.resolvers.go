package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/echonabin/allergymanagement-api/graph/model"
)

// CreateAllergy is the resolver for the createAllergy field.
func (r *mutationResolver) CreateAllergy(ctx context.Context, input model.NewAllergy) (*model.Allergy, error) {
	panic(fmt.Errorf("not implemented: CreateAllergy - createAllergy"))
}

// Allergy is the resolver for the allergy field.
func (r *queryResolver) Allergy(ctx context.Context, id *string) ([]*model.Allergy, error) {
	panic(fmt.Errorf("not implemented: Allergy - allergy"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
