package graph

import "github.com/diegograssato/go-graphQL-api-mongo/mongo_api"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Account *mongo_api.Account
}
