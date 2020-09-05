package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

// AccessToken stores neededndata for authentication
type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserID      int64  `json:"user_id"`
	ClientID    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
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
func GetNewAccessToken() AccessToken {
	return AccessToken{
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
