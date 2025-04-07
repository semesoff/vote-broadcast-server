package websocket_server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func (s *ServerManager) subscribeClientToMethod(w http.ResponseWriter, r *http.Request) {
	method := r.URL.Path[1:]
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	s.clientsData.clientMutex.Lock()
	switch method {
	case "getPolls":
		s.clientsData.clients["getPolls"] = append(s.clientsData.clients["getPolls"], c)
	case "getVotes":
		s.clientsData.clients["getVotes"] = append(s.clientsData.clients["getVotes"], c)
	}
	s.clientsData.countClients++
	log.Printf("count connected clients: %d\n", s.clientsData.countClients)
	s.clientsData.clientMutex.Unlock()
}
