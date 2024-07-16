package handlers

import (
	"baconator/api"
	"baconator/config"
	"baconator/resources/terraform"
	"baconator/sql/user"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Handler struct{}

func (t *Handler) UsersYamlGet(ctx context.Context) (*api.User, error) {

	c := config.GetConf()
	u := user.New(c.Pool)
	c.Logger.Info("listing users")
	list, err := u.ListUsers(context.Background())
	if err != nil {
		c.Logger.Error("error listing users", err)
		return nil, err404
	}
	var users []api.User
	for _, v := range list {
		meta := make(api.UserMetadata)
		meta["name"] = v.Name
		spec := make(api.UserSpec)
		spec["memberOf"] = api.NewStringArrayUserSpecItem([]string{"guests"})
		users = append(users, api.User{
			ApiVersion: api.NewOptString("backstage.io/v1alpha1"),
			Kind:       api.NewOptString("User"),
			Metadata:   api.NewOptUserMetadata(meta),
			Spec:       api.NewOptUserSpec(spec),
		})
	}
	c.Logger.Info("listing users", "users", users)
	return &users[0], nil
	//return users, nil
}

var err404 = &api.ErrorStatusCode{
	StatusCode: 404,
	Response: api.Error{
		Code:    404,
		Message: "not found",
	},
}

func (t *Handler) UserPatch(ctx context.Context, req *api.User) (*api.User, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Handler) UserPost(ctx context.Context, req *api.User) (*api.User, error) {
	c := config.GetConf()
	u := user.New(c.Pool)
	r, err := u.CreateUser(context.Background(), user.CreateUserParams{
		Name: req.Name.Value,
		Uuid: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
	})
	if err != nil {
		c.Logger.Error("error creating user", err)
		return nil, err404
	}
	return &api.User{Name: api.NewOptString(r.Name),
		ID: api.OptInt64{
			Value: r.ID,
			Set:   true,
		},
	}, nil
}

func (t *Handler) UsersUserIdGet(ctx context.Context, params api.UsersUserIdGetParams) (*api.User, error) {
	c := config.GetConf()
	u := user.New(c.Pool)
	f, err := u.GetUser(context.Background(), params.UserId)
	if err != nil {
		c.Logger.Error("error getting user", err)
		return nil, err404
	}
	return &api.User{
		Name: api.NewOptString(f.Name),
		ID: api.OptInt64{
			Value: f.ID,
			Set:   false,
		},
	}, nil
}

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
		err := t.NewError(ctx, fmt.Errorf("email missing"))
		return nil, err
	}
	return &api.User{Name: api.NewOptString(email)}, nil
}
