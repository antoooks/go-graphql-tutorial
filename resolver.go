package main

import (
	"context"

	"github.com/antoooks/go-graphql-tutorial/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Person(ctx context.Context) ([]*models.Person, error) {
	var result []*models.Person

	Adam := &models.Person{
		ID:   1,
		Name: "Adam",
		Pet:  models.Pet{
			ID: 1,
			Name: "Stew",
		},
	}

	result = append(result,Adam)
	return result,nil
}
func (r *queryResolver) Pet(ctx context.Context) ([]*models.Pet, error) {
	var result []*models.Pet

	Stew := &models.Pet{
			ID: 1,
			Name: "Stew",
		}

	result = append(result,Stew)
	return result,nil
}
