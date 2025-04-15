package websocket_server

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"vote-broadcast-server/services/websocket/pkg/models"
)

func (s *ServerManager) SendPolls(polls []*models.Poll) {
	jsonData := &bytes.Buffer{}
	if err := json.NewEncoder(jsonData).Encode(polls); err != nil {
		return
	}

	countSentData := 0
	for c := range s.clientsData.getPollsClients {
		err := c.WriteMessage(websocket.BinaryMessage, jsonData.Bytes())
		if err != nil {
			// delete the client from the map
			s.unsubscribeClientFromGetPolls(c)
			continue
		}
		countSentData++
	}

	log.Printf("[sendPolls]: count sent data: %d", countSentData)
}

func (s *ServerManager) SendVotes(pollVotes models.PollVotes) {
	jsonData := &bytes.Buffer{}
	if err := json.NewEncoder(jsonData).Encode(pollVotes); err != nil {
		return
	}

	countSentData := 0
	for c := range s.clientsData.getVotesClients[pollVotes.ID] {
		err := c.WriteMessage(websocket.BinaryMessage, jsonData.Bytes())
		if err != nil {
			// delete the client from the map
			s.unsubscribeClientFromGetVotes(c, pollVotes.ID)
			continue
		}
		countSentData++
	}

	log.Printf("[sendVotes]: count sent data: %d", countSentData)
}
