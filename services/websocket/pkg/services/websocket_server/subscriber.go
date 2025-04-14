package websocket_server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	utils2 "vote-broadcast-server/services/websocket/pkg/utils"
)

var upgrader = websocket.Upgrader{}

func (s *ServerManager) subscribeClientToMethod(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[1:]
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	switch method {
	case "getPolls":
		if err := s.addClientToGetPolls(c); err != nil {
			utils2.RespondWithMessage(w, "Failed to add client", http.StatusInternalServerError)
			return
		}
		utils2.ResponseWithSuccess(w, "Success")
	case "getVotes":
		pollId := r.URL.Query().Get("poll_id")
		if pollIdInt, err := strconv.Atoi(pollId); err != nil {
			utils2.RespondWithMessage(w, "Invalid poll ID", http.StatusBadRequest)
			return
		} else if err := s.addClientToGetVotes(pollIdInt, c); err != nil {
			utils2.RespondWithMessage(w, "Failed to add client", http.StatusInternalServerError)
			return
		}
		utils2.ResponseWithSuccess(w, "Success")
	default:
		return
	}

	s.clientsData.countClients++
	log.Printf("count connected clients: %d\n", s.clientsData.countClients)
}
