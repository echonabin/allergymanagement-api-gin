package graph

import "github.com/echonabin/allergymanagement-api/graph/model"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user []*model.User
}
