// Package handlers implements the API handlers for the baconator API.
//
// It provides the implementation of the handler interface.
package handlers

import (
	"context"
	"fmt"

	"github.com/vamage/baconator/api"
	"github.com/vamage/baconator/config"
	"github.com/vamage/baconator/resources/terraform"
	"github.com/vamage/baconator/sql/user"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var initalSize = 10

// Handler is the struct that contains the implementation of the API handlers.
type Handler struct{}

// UsersGet implements the users.yaml GET method of the API
// This returns a list of users in the system.
func (t *Handler) UsersGet(ctx context.Context) ([]api.Entity, error) {
	c := config.GetConf()
	u := user.New(c.Pool)
	c.Logger.Info("listing users")
	list, err := u.ListUsers(ctx)
	if err != nil {
		c.Logger.Error("error listing users", "error", err)
		return nil, err404
	}
	var users []api.Entity
	for _, v := range list {
		meta := make(api.StringOrArray, initalSize)
		meta["name"] = api.NewStringStringOrArrayItem(v.Name)
		meta["namespace"] = api.NewStringStringOrArrayItem("baconator")
		profile := make(map[string]string, initalSize)
		profile["email"] = "baconator@bacon.com"
		profile["avatarUrl"] = "https://baconator.com/avatar.png"
		profile["displayName"] = "Baconator Son of Bacon"
		spec := make(api.StringOrArray, initalSize)
		meta["profile"] = api.NewStringOrArrayItem2StringOrArrayItem(profile)
		annotation := make(map[string]string, initalSize)
		annotation["backstage.io/managed-by-location"] = fmt.Sprintf("url:http://baconator:8081/user/%d", v.ID)
		annotation["backstage.io/managed-by"] = "baconator"
		annotation["backstage.io/managed-by-origin-location"] = fmt.Sprintf("url:http://baconator:8081/user/%d", v.ID)
		meta["annotations"] = api.NewStringOrArrayItem2StringOrArrayItem(annotation)
		spec["memberOf"] = api.NewStringArrayStringOrArrayItem([]string{"default/guests"})

		users = append(users, api.Entity{
			ApiVersion: api.NewOptString("backstage.io/v1alpha1"),
			Kind:       api.NewOptString("User"),
			Metadata:   api.NewOptStringOrArray(meta),
			Spec:       api.NewOptStringOrArray(spec),
		})
	}
	c.Logger.Info("listing users", "users", users)
	// return &users[0], nil
	return users, nil
}

var err404 = &api.ErrorStatusCode{
	StatusCode: 404,
	Response: api.Error{
		Code:    404,
		Message: "not found",
	},
}

// UserPatch implements the User Patch method of the API
// This is the implementation of the Patch method for the /user endpoint.
func (t *Handler) UserPatch(ctx context.Context, req *api.Entity) (*api.Entity, error) {
	// TODO implement me
	panic("implement me")
}

// UserPost implements the User POST method of the API
// This is the implementation of the POST method for the /user endpoint.
func (t *Handler) UserPost(ctx context.Context, req *api.Entity) (*api.Entity, error) {
	c := config.GetConf()
	u := user.New(c.Pool)
	m, ok := req.Metadata.Get()

	if !ok {
		c.Logger.Error("name missing")
		return nil, err404
	}
	n := m["name"]
	name, ok := n.GetString()
	if !ok {
		c.Logger.Error("name is array")
		return nil, err404
	}
	r, err := u.CreateUser(context.Background(), user.CreateUserParams{
		Name: name,
		Uuid: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
	})
	if err != nil {
		c.Logger.Error("error creating user", "error", err)
		return nil, err404
	}
	meta := make(api.StringOrArray, initalSize)
	meta["name"] = api.NewStringStringOrArrayItem(r.Name)

	entity := &api.Entity{
		ApiVersion: api.NewOptString("backstage.io/v1alpha1"),
		Kind:       api.NewOptString("User"),
		Metadata:   api.NewOptStringOrArray(meta),
		Spec:       api.NewOptStringOrArray(nil),
	}
	return entity, nil
}

// UsersUserIdGet implements the UsersUserId GET method of the API
// This is the implementation of the GET method for the /users/{userId} endpoint.
func (t *Handler) UsersUserIdGet(ctx context.Context, params api.UsersUserIdGetParams) (*api.Entity, error) {
	c := config.GetConf()
	u := user.New(c.Pool)
	f, err := u.GetUser(context.Background(), params.UserId)
	if err != nil {
		c.Logger.Error("error getting user", "errror", err)
		return nil, err404
	}
	meta := make(api.StringOrArray, initalSize)
	meta["name"] = api.NewStringStringOrArrayItem(f.Name)

	entity := &api.Entity{
		ApiVersion: api.NewOptString("backstage.io/v1alpha1"),
		Kind:       api.NewOptString("User"),
		Metadata:   api.NewOptStringOrArray(meta),
		Spec:       api.NewOptStringOrArray(nil),
	}

	return entity, nil
}

// ResourcesPost implements the Resources POST method of the API.
func (t *Handler) ResourcesPost(ctx context.Context, req *api.Resource) (*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

// ResourcesResourceIDGet implements the ResourcesResourceId GET method of the API.
func (t *Handler) ResourcesResourceIDGet(ctx context.Context,
	params api.ResourcesResourceIDGetParams,
) (*api.Resource, error) {
	if params.ResourceID == 0 {
		return nil, &api.ErrorStatusCode{
			StatusCode: 403,
			Response: api.Error{
				Code:    403,
				Message: "resource not found",
			},
		}
	}
	a := api.Resource{
		Name:           fmt.Sprintf("resource-%d", params.ResourceID),
		ResourceInputs: nil,
	}
	a.GetName()
	return &a, nil
}

// NewError creates a new error response if one is needed.
func (t *Handler) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	// TODO implement me
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    500,
			Message: "sure why not",
		},
	}
}

// ListResourceTypesGet implements the ListResourceTypes GET method of the API.
func (t *Handler) ListResourceTypesGet(ctx context.Context) (resp *api.Resource, err error) {
	// TODO implement me
	resp, err = terraform.ReadTF("https://github.com/vamage/baconator-modules", "testing-variables")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// WhoamiGet implements the Whoami GET method of the API.
func (t *Handler) WhoamiGet(ctx context.Context) (*api.Entity, error) {
	email, ok := ctx.Value("email").(string)
	if !ok {
		err := t.NewError(ctx, fmt.Errorf("email missing"))
		return nil, err
	}
	meta := make(api.StringOrArray, initalSize)
	meta["name"] = api.NewStringStringOrArrayItem(email)

	entity := &api.Entity{
		ApiVersion: api.NewOptString("backstage.io/v1alpha1"),
		Kind:       api.NewOptString("User"),
		Metadata:   api.NewOptStringOrArray(meta),
		Spec:       api.NewOptStringOrArray(nil),
	}

	return entity, nil
}
