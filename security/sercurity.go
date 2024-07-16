package security

import (
	"baconator/api"
	"context"
)

type Security struct{}

func (s *Security) HandleOAuth2(ctx context.Context, operation string, t api.OAuth2) (context.Context, error) {
	/*p, err := idtoken.Validate(ctx, t.Token, "")
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", p)
	v, ok := p.Claims["email"]
	if !ok {
		return nil, fmt.Errorf("email missing")
	}

	ctx = context.WithValue(ctx, "email", v)

	*/
	return ctx, nil
}
