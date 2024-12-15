package types

import (
	"context"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token string) (*oidc.IDToken, error) {

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}
	verifier := a.Verifier(oidcConfig)
	idToken, err := verifier.Verify(ctx, token)

	return idToken, err
}
