package models

import (
	"errors"
	"log"

	"github.com/graphql-go/graphql"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewUser creates a new user instance
func NewUser(id int, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

// Validate checks if the user is valid
func (u *User) Validate() error {
	if u.ID <= 0 {
		return errors.New("invalid user id")
	}
	if u.Name == "" {
		return errors.New("user name is required")
	}
	if u.Email == "" {
		return errors.New("user email is required")
	}
	return nil
}

// ToGraphQLType converts the user to a GraphQL object
func (u *User) ToGraphQLType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

// GetUser retrieves a user by id
func GetUser(id int) (*User, error) {
	// implement database query to retrieve user by id
	// for demonstration purposes, return a hardcoded user
	user := &User{
		ID:    1,
		Name:  "John Doe",
		Email: "john@example.com",
	}
	if id != user.ID {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func init() {
	// initialize user model
	log.Println("User model initialized")
}
