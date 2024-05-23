package main

import (
	"baconator/api"
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	ctx := context.Background()
	gts, err := google.DefaultTokenSource(ctx)
	if err != nil {
		fmt.Printf("d%+v", err)

	}
	client, err := api.NewClient("http://127.0.0.1:8080/", &tokenSource{t: gts})
	if err != nil {
		fmt.Printf("sss%+v", err)
	}
	r, err := client.WhoamiGet(ctx)
	if err != nil {
		fmt.Printf("dd %+v", err)
	}
	fmt.Printf("zz%+v", r)

}

type tokenSource struct {
	t oauth2.TokenSource
}

func (t *tokenSource) OAuth2(ctx context.Context, opperationName string) (api.OAuth2, error) {
	// Construct the GoogleCredentials object which obtains the default configuration from your
	// working environment.
	credentials, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/cloud-platform.read-only")
	if err != nil {
		fmt.Printf("failed to generate default credentials: %w", err)
	}

	// Get the ID token.
	// Once you've obtained the ID token, you can use it to make an authenticated call
	// to the target audience.

	token, err := credentials.TokenSource.Token()
	if err != nil {
		return api.OAuth2{}, err
	}
	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		return api.OAuth2{}, fmt.Errorf("token did not contain an id_token")
	}
	return api.OAuth2{Token: idToken}, nil
}
