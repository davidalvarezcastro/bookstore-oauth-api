package db

import (
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/clients/cassandra"
	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/models/access_token"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

// Repository interface storing db repository functions
type Repository interface {
	GetByID(string) (*token.AccessToken, *errors.RestErr)
	Create(token.AccessToken) *errors.RestErr
	UpdateExpirationTime(token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

// NewRepository return a new Repository
func NewRepository() Repository {
	return &dbRepository{}
}

// GetByID returns an access token from database by id
func (r *dbRepository) GetByID(id string) (*token.AccessToken, *errors.RestErr) {
	var result token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserID,
		&result.ClientID,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found with given id")
		}
		return nil, errors.NewInternalServerError("error when trying to get current id")
	}

	return &result, nil
}

// Create creates a new token in database
func (r *dbRepository) Create(at token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserID,
		at.ClientID,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error when trying to save access token in database")
	}

	return nil
}

// UpdateExpirationTime updates the token's expiration time in database
func (r *dbRepository) UpdateExpirationTime(at token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError("error when trying to update current resource")
	}

	return nil
}
