// Package users provides the implementation of the users API.
package users

import (
	"github.com/vamage/baconator/api"

	"golang.org/x/net/context"
)

// UserService is the struct that contains the implementation of the API handlers.
type UserService struct{}

// WhoamiGet returns the user information.
func (t *UserService) WhoamiGet(ctx context.Context) (*api.User, error) {
	return &api.User{Name: api.NewOptString("greg")}, nil
}
