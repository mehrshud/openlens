package services

import (
	"log"
	"./models"
	"./database"
)

// GetDashboard returns a dashboard by its ID.
func GetDashboard(config Config, dashboardID int) (Dashboard, error) {
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return Dashboard{}, err
	}

	dashboard, err := db.GetDashboard(dashboardID)
	if err != nil {
		log.Printf("Error getting dashboard: %v\n", err)
		return Dashboard{}, err
	}

	return dashboard, nil
}

// GetDashboards returns a list of all dashboards.
func GetDashboards(config Config) ([]Dashboard, error) {
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return []Dashboard{}, err
	}

	dashboards, err := db.GetDashboards()
	if err != nil {
		log.Printf("Error getting dashboards: %v\n", err)
		return []Dashboard{}, err
	}

	return dashboards, nil
}

// CreateDashboard creates a new dashboard.
func CreateDashboard(config Config, dashboard Dashboard) (Dashboard, error) {
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return Dashboard{}, err
	}

	createdDashboard, err := db.CreateDashboard(dashboard)
	if err != nil {
		log.Printf("Error creating dashboard: %v\n", err)
		return Dashboard{}, err
	}

	return createdDashboard, nil
}

// UpdateDashboard updates an existing dashboard.
func UpdateDashboard(config Config, dashboard Dashboard) (Dashboard, error) {
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return Dashboard{}, err
	}

	updatedDashboard, err := db.UpdateDashboard(dashboard)
	if err != nil {
		log.Printf("Error updating dashboard: %v\n", err)
		return Dashboard{}, err
	}

	return updatedDashboard, nil
}

// DeleteDashboard deletes a dashboard by its ID.
func DeleteDashboard(config Config, dashboardID int) error {
	db, err := database.Connect(config.DBURL)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return err
	}

	err = db.DeleteDashboard(dashboardID)
	if err != nil {
		log.Printf("Error deleting dashboard: %v\n", err)
		return err
	}

	return nil
}
