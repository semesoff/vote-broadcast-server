package websocket_server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"strings"
	utils2 "vote-broadcast-server/services/websocket/pkg/utils"
)

var upgrader = websocket.Upgrader{}

func (s *ServerManager) subscribeClientToMethod(w http.ResponseWriter, r *http.Request) {
	method := strings.Split(r.URL.Path[1:], "/")[0]
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	switch method {
	case "getPolls":
		if err := s.addClientToGetPolls(c); err != nil {
			utils2.RespondWithError(c, "Error adding a client", websocket.CloseMessage)
			return
		}
		utils2.RespondWithSuccess(c, fmt.Sprintf("Connected successfully to %s", method), websocket.TextMessage)
	case "getVotes":
		pollId := mux.Vars(r)["poll_id"]
		if pollIdInt, err := strconv.Atoi(pollId); err != nil {
			utils2.RespondWithError(c, "Invalid server data", websocket.CloseMessage)
			return
		} else if err := s.addClientToGetVotes(pollIdInt, c); err != nil {
			utils2.RespondWithError(c, "", websocket.CloseMessage)
			return
		}
		utils2.RespondWithSuccess(c, fmt.Sprintf("Connected successfully to %s", method), websocket.TextMessage)
	default:
		return
	}

	log.Printf("count connected clients: %d\n", s.clientsData.countClients)
}
