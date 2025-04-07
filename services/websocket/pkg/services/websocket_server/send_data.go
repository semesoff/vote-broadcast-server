package websocket_server

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"vote-broadcast-server/services/websocket/pkg/models"
)

func (s *ServerManager) SendPolls(data interface{}, method models.MethodType) {
	jsonData := &bytes.Buffer{}
	if err := json.NewEncoder(jsonData).Encode(data); err != nil {
		return
	}

	for _, c := range s.clientsData.clients[method.String()] {
		err := c.WriteMessage(websocket.BinaryMessage, jsonData.Bytes())
		if err != nil {
			// TODO: unsubscribe client
			c.Close()
			continue
		}
	}
}

func (s *ServerManager) SendVotes(data interface{}, method models.MethodType) {
	// logic
}
