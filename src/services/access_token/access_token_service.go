package token

import (
	"strings"

	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/models/access_token"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
)

// Repository temp
type Repository interface {
	GetByID(accessToken string) (*token.AccessToken, *errors.RestErr)
	Create(token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token.AccessToken) *errors.RestErr
}

// Service defining all the functions for access_token service
type Service interface {
	GetByID(string) (*token.AccessToken, *errors.RestErr)
	Create(token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo Repository
}

// NewService return a Service interface
func NewService(repo Repository) Service {
	return &service{
		restUsersRepo: repo,
	}
}

// GetByID return an access token by id
func (s *service) GetByID(accessTokenID string) (*token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}

	accessToken, err := s.restUsersRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// Create creates a new access token
func (s *service) Create(at token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.restUsersRepo.Create(at)
}

// UpdateExpirationTime updated expiration time
func (s *service) UpdateExpirationTime(at token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.restUsersRepo.UpdateExpirationTime(at)
}
