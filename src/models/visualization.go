// Package models provides data models for the OpenLens application.
package models

import (
	"errors"
	"log"
)

// Visualization represents a visualization in the system.
type Visualization struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewVisualization creates a new Visualization instance.
func NewVisualization(id int, name string, vizType string) (*Visualization, error) {
	if id <= 0 {
		return nil, errors.New("invalid visualization id")
	}
	if name == "" {
		return nil, errors.New("invalid visualization name")
	}
	if vizType == "" {
		return nil, errors.New("invalid visualization type")
	}
	viz := &Visualization{ID: id, Name: name, Type: vizType}
	log.Printf("New visualization created: %+v\n", viz)
	return viz, nil
}

// Validate checks if the visualization is valid.
func (v *Visualization) Validate() error {
	if v.ID <= 0 {
		return errors.New("invalid visualization id")
	}
	if v.Name == "" {
		return errors.New("invalid visualization name")
	}
	if v.Type == "" {
		return errors.New("invalid visualization type")
	}
	return nil
}
