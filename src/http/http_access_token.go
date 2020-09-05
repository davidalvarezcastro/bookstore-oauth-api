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

// NewHandler handles oauth api requests
func NewHandler(service token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessToken, err := h.service.GetByID(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
	}

	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at tokenM.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}
