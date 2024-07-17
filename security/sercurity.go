// Package security provides the implementation of the security interface.
package security

import (
	"context"

	"github.com/vamage/baconator/api"
)

// Security is the struct that contains the implementation of the security interface.
type Security struct{}

// HandleOAuth2 implements the HandleOAuth2 method of the security interface.
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
