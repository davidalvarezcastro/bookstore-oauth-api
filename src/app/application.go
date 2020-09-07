package app

import (
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/http"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/repository/db"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/repository/rest"
	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts the oauth api
func StartApp() {
	atHandler := http.NewAccessTokenHandler(
		token.NewService(rest.NewRestUsersRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
