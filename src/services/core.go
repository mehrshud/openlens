package services

import (
	"context"
	"fmt"
	"log"

	"src/database"
	"src/models"
)

// CoreService represents the core service of the application
type CoreService struct {
	db database.Database
}

// NewCoreService returns a new instance of the core service
func NewCoreService(db database.Database) *CoreService {
	return &CoreService{db: db}
}

// GetUser retrieves a user by ID
func (s *CoreService) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, err := s.db.GetUser(ctx, id)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

// GetDashboard retrieves a dashboard by ID
func (s *CoreService) GetDashboard(ctx context.Context, id int) (*models.Dashboard, error) {
	dashboard, err := s.db.GetDashboard(ctx, id)
	if err != nil {
		log.Printf("Error retrieving dashboard: %v", err)
		return nil, fmt.Errorf("failed to retrieve dashboard: %w", err)
	}
	return dashboard, nil
}

// GetVisualization retrieves a visualization by ID
func (s *CoreService) GetVisualization(ctx context.Context, id int) (*models.Visualization, error) {
	visualization, err := s.db.GetVisualization(ctx, id)
	if err != nil {
		log.Printf("Error retrieving visualization: %v", err)
		return nil, fmt.Errorf("failed to retrieve visualization: %w", err)
	}
	return visualization, nil
}

// UpdateConfig updates the application configuration
func (s *CoreService) UpdateConfig(ctx context.Context, config *models.Config) error {
	err := s.db.UpdateConfig(ctx, config)
	if err != nil {
		log.Printf("Error updating config: %v", err)
		return fmt.Errorf("failed to update config: %w", err)
	}
	return nil
}
