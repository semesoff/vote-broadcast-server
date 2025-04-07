package websocket_server

import (
	"context"
	"log"
	"vote-broadcast-server/services/websocket/pkg/models"
)

func (s *ServerManager) listenUpdatedServerData(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("listenUpdatedServerData is stopped.")
			return
		case data := <-s.dataChannels:
			switch data.Method {
			case models.GetPolls:
				s.SendPolls(data.Data, data.Method)
			case models.GetVotes:
				s.SendVotes(data.Data, data.Method)
			}
		}
	}
}
