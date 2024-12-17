package types

import (
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

type UserClaim struct {
	UserId string `json:"sub" db:"user_id"`
}

func (a *Authenticator) VerifyIDToken(r *http.Request) (UserClaim, error) {
	var userClaim UserClaim = UserClaim{}
	authToken := r.Header.Get("Authorization")
	verifier := a.Verifier(&oidc.Config{
		ClientID: a.ClientID,
	})

	idToken, err := verifier.Verify(r.Context(), authToken)
	if err != nil {
		return userClaim, err
	}

	err = idToken.Claims(&userClaim)
	fmt.Println(userClaim)
	return userClaim, err
}
