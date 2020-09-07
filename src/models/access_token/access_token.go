package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
	"github.com/davidalvarezcastro/bookstore-users-api/utils/crypto"
)

const (
	expirationTime       = 24
	grantTypePassword    = "password"
	grantTypeCredentials = "credentials"
)

// AccessToken stores needed data for authentication token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

// AccessTokenRequest stores needed data for auth request
type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`

	// User for password grant type
	Username string `json:"username"`
	Password string `json:"password"`

	// User for client_credential grant type
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Validate validates an access token request
func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case grantTypePassword:
		if at.Username == "" || at.Password == "" {
			return errors.NewBadRequestError("invalid password grant_type")
		}
		break
	case grantTypeCredentials:
		if at.ClientID == "" || at.ClientSecret == "" {
			return errors.NewBadRequestError("invalid credentials grant_type")
		}
		break
	default:
		return errors.NewBadRequestError("invalid grant_type parameter")
	}

	return nil
}

// Validate validates an access token
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token")
	}

	if at.UserID <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}

	if at.ClientID <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}

	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

// GetNewAccessToken return a new AccessToken struct
func GetNewAccessToken(userID int64) AccessToken {
	return AccessToken{
		UserID:  userID,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

// IsExpired returns if the AccessToken is expired
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)

	return expirationTime.Before(now)
}

// Generate generates a new access token from the user and the expiration date
func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserID, at.Expires))
}
