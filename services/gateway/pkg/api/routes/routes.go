package routes

import (
	"github.com/gin-gonic/gin"
	"vote-broadcast-server/services/gateway/pkg/api/handlers"
	"vote-broadcast-server/services/gateway/pkg/models"
)

func InitRoutes(mux *gin.Engine, services map[string]models.ServiceConfig, handlers handlers.Handlers) {
	authService := services["auth"]
	mux.Handle(authService.Routes[0].Method, authService.Routes[0].Path, handlers.RegisterUser)
	mux.Handle(authService.Routes[1].Method, authService.Routes[1].Path, handlers.LoginUser)
}
