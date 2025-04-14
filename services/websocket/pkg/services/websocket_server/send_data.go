package websocket_server

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
)

func (s *ServerManager) SendPolls(data interface{}) {
	jsonData := &bytes.Buffer{}
	if err := json.NewEncoder(jsonData).Encode(data); err != nil {
		return
	}

	for c := range s.clientsData.getPollsClients {
		err := c.WriteMessage(websocket.BinaryMessage, jsonData.Bytes())
		if err != nil {
			// delete the client from the map
			if err := c.Close(); err != nil {
			}
			s.unsubscribeClientFromGetPolls(c)
		}
	}
}

func (s *ServerManager) SendVotes(data interface{}, pollId int) {
	jsonData := &bytes.Buffer{}
	if err := json.NewEncoder(jsonData).Encode(data); err != nil {
		return
	}

	for c := range s.clientsData.getVotesClients[pollId] {
		err := c.WriteMessage(websocket.BinaryMessage, jsonData.Bytes())
		if err != nil {
			// delete the client from the map
			if err := c.Close(); err != nil {
			}
			s.unsubscribeClientFromGetVotes(c, pollId)
		}
	}
}
