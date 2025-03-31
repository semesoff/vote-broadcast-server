package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"vote-broadcast-server/services/gateway/pkg/api/handlers"
	"vote-broadcast-server/services/gateway/pkg/models"
)

func InitRoutes(mux *gin.Engine, services map[string]models.ServiceConfig, handlers handlers.Handlers) {
	authService, ok := services["auth"]
	if !ok {
		log.Fatalln(errors.New("auth service not found"))
		return
	}
	mux.Handle(authService.Routes[0].Method, authService.Routes[0].Path, handlers.RegisterUser)
	mux.Handle(authService.Routes[1].Method, authService.Routes[1].Path, handlers.LoginUser)

	pollService, ok := services["poll"]
	if !ok {
		log.Fatalln(errors.New("poll service not found"))
		return
	}
	mux.Handle(pollService.Routes[0].Method, pollService.Routes[0].Path, handlers.GetPolls)
	mux.Handle(pollService.Routes[1].Method, pollService.Routes[1].Path, handlers.CreatePoll)
	mux.Handle(pollService.Routes[2].Method, pollService.Routes[2].Path, handlers.GetPoll)

	voteService, ok := services["vote"]
	if !ok {
		log.Fatalln(errors.New("poll service not found"))
		return
	}
	mux.Handle(voteService.Routes[0].Method, voteService.Routes[0].Path, handlers.CreateVote)
	mux.Handle(voteService.Routes[1].Method, voteService.Routes[1].Path, handlers.GetVotes)
}
