package http

import (
	"net/http"
	"strings"

	tokenM "github.com/davidalvarezcastro/bookstore-oauth-api/src/models/access_token"
	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/services/access_token"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler defines AccessTokenHandler functions
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
}

type accessTokenHandler struct {
	service token.Service
}

// NewAccessTokenHandler handles oauth api requests
func NewAccessTokenHandler(service token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var request tokenM.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	accessToken, err := h.service.Create(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, accessToken)
}
