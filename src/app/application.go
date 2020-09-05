package app

import (
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/http"
	"github.com/davidalvarezcastro/bookstore-oauth-api/src/repository/db"
	token "github.com/davidalvarezcastro/bookstore-oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp starts the oauth api
func StartApp() {
	atService := token.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
