// Package models define the structure of data models used in OpenLens.
package models

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

// Dashboard represents a collection of visualizations.
type Dashboard struct {
	ID             int                 `json:"id"`
	Name           string             `json:"name"`
	Visualizations []Visualization     `json:"visualizations"`
}

// Visualization represents a single visualization in a dashboard.
type Visualization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewDashboard creates a new dashboard with the given name and visualizations.
func NewDashboard(name string, visualizations []Visualization) *Dashboard {
	if name == "" {
		log.Fatal("Dashboard name cannot be empty")
	}
	return &Dashboard{
		Name:           name,
		Visualizations: visualizations,
	}
}

// Validate checks if the dashboard is valid.
func (d *Dashboard) Validate() error {
	if d.Name == "" {
		return fmt.Errorf("dashboard name cannot be empty")
	}
	if len(d.Visualizations) == 0 {
		return fmt.Errorf("dashboard must have at least one visualization")
	}
	return nil
}
