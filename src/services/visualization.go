package services

import (
	"fmt"
	"log"

	"./models"
)

// GetVisualization fetches a visualization by ID
func GetVisualization(config models.Config, visualizationId int) (models.Visualization, error) {
	dbURL := config.DBURL
	if dbURL == "" {
		return models.Visualization{}, fmt.Errorf("database URL is not set")
	}

	// Connect to database and fetch visualization
	// For this example, we assume a simple database connection
	// and a 'visualizations' table with 'id' and 'name' columns
	// You should replace this with your actual database logic
	visualization, err := fetchVisualizationFromDB(dbURL, visualizationId)
	if err != nil {
		log.Printf("Error fetching visualization: %v", err)
		return models.Visualization{}, err
	}

	return visualization, nil
}

// fetchVisualizationFromDB is a helper function to fetch a visualization from the database
func fetchVisualizationFromDB(dbURL string, visualizationId int) (models.Visualization, error) {
	// For this example, we assume a simple database connection
	// and a 'visualizations' table with 'id' and 'name' columns
	// You should replace this with your actual database logic
	// We'll return a mock visualization for now
	return models.Visualization{
		ID:   visualizationId,
		Name: "Mock Visualization",
	}, nil
}

// GetVisualizationsForDashboard fetches all visualizations for a dashboard
func GetVisualizationsForDashboard(config models.Config, dashboardId int) ([]models.Visualization, error) {
	dbURL := config.DBURL
	if dbURL == "" {
		return []models.Visualization{}, fmt.Errorf("database URL is not set")
	}

	// Connect to database and fetch visualizations
	// For this example, we assume a simple database connection
	// and a 'visualizations' table with 'id' and 'name' columns
	// You should replace this with your actual database logic
	visualizations, err := fetchVisualizationsFromDB(dbURL, dashboardId)
	if err != nil {
		log.Printf("Error fetching visualizations: %v", err)
		return []models.Visualization{}, err
	}

	return visualizations, nil
}

// fetchVisualizationsFromDB is a helper function to fetch visualizations from the database
func fetchVisualizationsFromDB(dbURL string, dashboardId int) ([]models.Visualization, error) {
	// For this example, we assume a simple database connection
	// and a 'visualizations' table with 'id' and 'name' columns
	// You should replace this with your actual database logic
	// We'll return a mock visualization for now
	return []models.Visualization{
		{
			ID:   1,
			Name: "Mock Visualization 1",
		},
		{
			ID:   2,
			Name: "Mock Visualization 2",
		},
	}, nil
}