package main

import (
	"baconator/api"
	"cloud.google.com/go/auth/credentials/idtoken"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	service := &testService{}
	sec := &Security{}
	srv, err := api.NewServer(service, sec)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}

type testService struct {
}

type User struct {
	Name string `json:"name"`
}

func (t *testService) WhoamiGet(ctx context.Context) (api.WhoamiGetRes, error) {
	email, ok := ctx.Value("email").(string)
	fmt.Println("lalalalalalalal")
	if !ok {
		return &api.WhoamiGetForbidden{}, fmt.Errorf("was not set")
	}

	return (*api.User)(&User{Name: email}), nil
}

type Security struct{}

func (s *Security) HandleOAuth2(ctx context.Context, operation string, t api.OAuth2) (context.Context, error) {
	p, err := idtoken.Validate(ctx, t.Token, "")
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", p)
	v, ok := p.Claims["email"]
	if !ok {
		return nil, fmt.Errorf("email missing")
	}

	ctx = context.WithValue(ctx, "email", v)
	return ctx, nil
}
