package handlers

import (
	"context"
	"fmt"

	"baconator/api"
	"baconator/resources/terraform"
)

type Handler struct{}

func (t *Handler) ResourcesPost(ctx context.Context, req *api.Resource) (*api.Resource, error) {
	// TODO implement me
	panic("implement me")
}

func (t *Handler) ResourcesResourceIdGet(ctx context.Context, params api.ResourcesResourceIdGetParams) (*api.Resource, error) {
	// TODO implement me
	if params.ResourceId == 0 {
		return nil, &api.ErrorStatusCode{
			StatusCode: 403,
			Response: api.Error{
				Code:    403,
				Message: "resource not found",
			},
		}
	}
	a := api.Resource{
		Name:           fmt.Sprintf("resource-%d", params.ResourceId),
		ResourceInputs: nil,
	}
	a.GetName()
	return &a, nil
}

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

func (t *Handler) ListResourceTypesGet(ctx context.Context) (resp *api.Resource, err error) {
	// TODO implement me
	resp, err = terraform.ReadTF("https://github.com/vamage/baconator-modules", "testing-variables")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type User struct {
	Name string `json:"name"`
}

func (t *Handler) WhoamiGet(ctx context.Context) (*api.User, error) {
	email, ok := ctx.Value("email").(string)
	if !ok {
		t.NewError(ctx, fmt.Errorf("email missing"))
	}
	return (*api.User)(&User{Name: email}), nil
}
