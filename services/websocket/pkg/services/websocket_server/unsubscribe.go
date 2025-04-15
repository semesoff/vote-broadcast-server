package websocket_server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func (s *ServerManager) unsubscribeClientFromGetPolls(conn *websocket.Conn) {
	s.clientsData.pollsMutex.Lock()
	defer s.clientsData.pollsMutex.Unlock()
	if err := conn.Close(); err != nil {
	}
	delete(s.clientsData.getPollsClients, conn)
	s.clientsData.countClients--
	log.Printf("count connected clients: %d\n", s.clientsData.countClients)
}

func (s *ServerManager) unsubscribeClientFromGetVotes(conn *websocket.Conn, pollId int) {
	fmt.Println("unsubscribeClientFromGetVotes: ", pollId)
	s.clientsData.votesMutex.Lock()
	defer s.clientsData.votesMutex.Unlock()
	if err := conn.Close(); err != nil {
	}
	delete(s.clientsData.getVotesClients[pollId], conn)
	s.clientsData.countClients--
	log.Printf("count connected clients: %d\n", s.clientsData.countClients)
}
