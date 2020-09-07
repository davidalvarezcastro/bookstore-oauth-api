package token

import (
	"strings"

	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/models/access_token"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/repository/db"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/repository/rest"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
)

// Service defining all the functions for access_token service
type Service interface {
	GetByID(string) (*token.AccessToken, *errors.RestErr)
	Create(token.AccessTokenRequest) (*token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.UserRepository
	dbRepo        db.Repository
}

// NewService return a Service interface
func NewService(usersRepo rest.UserRepository, dbRepo db.Repository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

// GetByID return an access token by id
func (s *service) GetByID(accessTokenID string) (*token.AccessToken, *errors.RestErr) {
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0 {
		return nil, errors.NewBadRequestError("invalid access token")
	}

	accessToken, err := s.dbRepo.GetByID(accessTokenID)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

// Create creates a new access token
func (s *service) Create(request token.AccessTokenRequest) (*token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := token.GetNewAccessToken(user.ID)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

// UpdateExpirationTime updated expiration time
func (s *service) UpdateExpirationTime(at token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}
