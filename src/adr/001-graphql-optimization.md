// src/adr/001-graphql-optimization.go
package main

import (
	"log"

	"github.com/graphql-go/graphql"
	"./models"
	"./database"
)

// optimizeGraphQLQuery optimizes a GraphQL query by filtering out unnecessary fields
func optimizeGraphQLQuery(query string, config models.Config) (string, error) {
	// Try to parse the query
	ast, err := graphql.Parse(query)
	if err != nil {
		log.Printf("Error parsing GraphQL query: %v\n", err)
		return "", err
	}

	// Filter out unnecessary fields
	optimizedQuery := optimizeFields(ast, config)

	// Return the optimized query
	return optimizedQuery, nil
}

// optimizeFields optimizes the fields in a GraphQL query
func optimizeFields(ast *graphql.AST, config models.Config) string {
	// Initialize the optimized query
	optimizedQuery := ""

	// Iterate over the fields in the query
	for _, field := range ast.Fields {
		// Check if the field is necessary
		if isFieldNecessary(field, config) {
			// Add the field to the optimized query
			optimizedQuery += field.Name + " "
		}
	}

	// Return the optimized query
	return optimizedQuery
}

// isFieldNecessary checks if a field is necessary in a GraphQL query
func isFieldNecessary(field *graphql.Field, config models.Config) bool {
	// Check if the field is in the cache
	if isFieldCached(field, config) {
		return true
	}

	// Check if the field is required for the dashboard
	if isFieldRequiredForDashboard(field, config) {
		return true
	}

	// If the field is not necessary, return false
	return false
}

// isFieldCached checks if a field is in the cache
func isFieldCached(field *graphql.Field, config models.Config) bool {
	// Connect to the database
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return false
	}

	// Check if the field is in the cache
	cached, err := db.IsFieldCached(field.Name)
	if err != nil {
		log.Printf("Error checking if field is cached: %v\n", err)
		return false
	}

	// Return whether the field is cached
	return cached
}

// isFieldRequiredForDashboard checks if a field is required for the dashboard
func isFieldRequiredForDashboard(field *graphql.Field, config models.Config) bool {
	// Get the dashboard metadata
	metadata, err := getDashboardMetadata(config)
	if err != nil {
		log.Printf("Error getting dashboard metadata: %v\n", err)
		return false
	}

	// Check if the field is required for the dashboard
	for _, visualization := range metadata.Visualizations {
		if visualization.Type == field.Name {
			return true
		}
	}

	// If the field is not required, return false
	return false
}

// getDashboardMetadata gets the metadata for a dashboard
func getDashboardMetadata(config models.Config) (models.Dashboard, error) {
	// Connect to the database
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return models.Dashboard{}, err
	}

	// Get the dashboard metadata
	metadata, err := db.GetDashboardMetadata()
	if err != nil {
		log.Printf("Error getting dashboard metadata: %v\n", err)
		return models.Dashboard{}, err
	}

	// Return the dashboard metadata
	return metadata, nil
}
