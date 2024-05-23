package users

import (
	"baconator/api"
	"fmt"
	"golang.org/x/net/context"
)

type UserService struct {
}

type User struct {
	Name string `json:"name"`
}

func (t *UserService) WhoamiGet(ctx context.Context) (*api.User, error) {
	fmt.Sprintf("calledwhoami\n")
	return (*api.User)(&User{Name: "greg"}), nil
}
