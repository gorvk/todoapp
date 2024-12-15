package initializers

import (
	"context"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorvk/todoapp/internal/types"
	"golang.org/x/oauth2"
)

var auth *types.Authenticator

func GetAuthInstance() *types.Authenticator {
	return auth
}

func ConnectAuth() {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+os.Getenv("AUTH0_DOMAIN")+"/",
	)
	if err != nil {
		panic(err)
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	auth = &types.Authenticator{
		Provider: provider,
		Config:   conf,
	}
}
