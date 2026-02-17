package api

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"src/api/models"
	"log"
)

// GraphQL API schema
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"user": &graphql.Field{
					Type: models.UserType,
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id, ok := p.Args["id"].(int)
						if !ok {
							return nil, gqlerrors.NewFormattedError("id must be an integer")
						}
						user, err := models.GetUser(id)
						if err != nil {
							log.Println(err)
							return nil, gqlerrors.NewFormattedError("failed to retrieve user")
						}
						return user, nil
					},
				},
				"dashboards": &graphql.Field{
					Type: graphql.NewList(models.DashboardType),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						dashboards, err := models.GetDashboards()
						if err != nil {
							log.Println(err)
							return nil, gqlerrors.NewFormattedError("failed to retrieve dashboards")
						}
						return dashboards, nil
					},
				},
			},
		}),
	},
)

func GetSchema() *graphql.Schema {
	return &schema
}
