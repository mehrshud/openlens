package utils

import (
	"errors"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"../models"
)

// CustomError is a custom error type that includes a message and a code
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// HandleError handles and logs the error
func HandleError(err error, code int) *CustomError {
	if err != nil {
		log.Printf("error: %v", err)
		return &CustomError{Code: code, Message: err.Error()}
	}
	return nil
}

// ValidateUser validates a given User object
func ValidateUser(user *models.User) error {
	if user == nil {
		return HandleError(errors.New("user is nil"), 400).Error
	}
	if user.ID <= 0 {
		return HandleError(errors.New("invalid user id"), 400).Error
	}
	if user.Name == "" || user.Email == "" {
		return HandleError(errors.New("name and email are required"), 400).Error
	}
	return nil
}

// ValidateConfig validates a given Config object
func ValidateConfig(config *models.Config) error {
	if config == nil {
		return HandleError(errors.New("config is nil"), 400).Error
	}
	if !config.Debug && config.DBURL == "" {
		return HandleError(errors.New("db url is required when not in debug mode"), 400).Error
	}
	return nil
}

// ValidateDashboard validates a given Dashboard object
func ValidateDashboard(dashboard *models.Dashboard) error {
	if dashboard == nil {
		return HandleError(errors.New("dashboard is nil"), 400).Error
	}
	if dashboard.ID <= 0 {
		return HandleError(errors.New("invalid dashboard id"), 400).Error
	}
	if dashboard.Name == "" {
		return HandleError(errors.New("name is required"), 400).Error
	}
	for _, visualization := range dashboard.Visualizations {
		if visualization.ID <= 0 || visualization.Name == "" || visualization.Type == "" {
			return HandleError(errors.New("invalid visualization"), 400).Error
		}
	}
	return nil
}