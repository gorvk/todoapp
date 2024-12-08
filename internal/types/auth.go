package types

import (
	"context"
	"errors"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Profile struct {
	Aud        string `json:"aud"`
	Exp        int    `json:"exp"`
	Iat        int    `json:"iat"`
	Iss        string `json:"iss"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	Picture    string `json:"picture"`
	Sid        string `json:"sid"`
	Sub        string `json:"sub"`
	Updated_at string `json:"updated_at"`
}

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)

	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
