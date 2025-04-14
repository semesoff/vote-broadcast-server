package websocket_server

import (
	"github.com/gorilla/websocket"
)

func (s *ServerManager) addClientToGetPolls(c *websocket.Conn) error {
	s.clientsData.pollsMutex.Lock()
	defer s.clientsData.pollsMutex.Unlock()

	s.clientsData.getPollsClients[c] = true
	s.clientsData.countClients++

	return nil
}

func (s *ServerManager) addClientToGetVotes(pollID int, c *websocket.Conn) error {
	s.clientsData.votesMutex.Lock()
	defer s.clientsData.votesMutex.Unlock()

	if _, ok := s.clientsData.getVotesClients[pollID]; !ok {
		s.clientsData.getVotesClients[pollID] = make(map[*websocket.Conn]bool)
	}
	s.clientsData.getVotesClients[pollID][c] = true
	s.clientsData.countClients++

	return nil
}
