package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"./models"
	"./api/graphql"
	"./services/core"
)

// Main entry point of the application
func main() {
	// Initialize configuration
	var config models.Config
	if err := initConfig(&config); err != nil {
		log.Fatal(err)
	}

	// Initialize database connection
	if err := core.InitDatabase(config.DBURL); err != nil {
		log.Fatal(err)
	}

	// Initialize GraphQL schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type:        models.UserType,
					Description: "Get user by ID",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(int)
						if !ok {
							return nil, fmt.Errorf("invalid user ID")
						}
						user, err := core.GetUser(id)
						if err != nil {
							return nil, err
						}
						return user, nil
					},
				},
				"dashboard": &graphql.Field{
					Type:        graphql.NewList(models.DashboardType),
					Description: "Get all dashboards",
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						dashboards, err := core.GetDashboards()
						if err != nil {
							return nil, err
						}
						return dashboards, nil
					},
				},
			},
		}),
	})

	if err != nil {
		log.Fatal(err)
	}

	// Start GraphQL server
	if err := graphql.StartServer(schema); err != nil {
		log.Fatal(err)
	}
}

// Initialize configuration from environment variables or default values
func initConfig(config *models.Config) error {
	config.Debug = false
	config.DBURL = "localhost:5432"
	return nil
}
