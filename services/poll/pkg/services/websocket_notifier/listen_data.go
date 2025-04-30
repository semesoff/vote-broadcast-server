package websocket_notifier

import (
	"context"
	"log"
	"poll-service/pkg/models"
)

func (s *WebSocketNotifierService) listenUpdatedServerData(ctx context.Context) {
	defer s.wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Println("listenUpdatedServerData is stopped.")
			return
		case data := <-s.dataChannel:
			switch data.Method {
			case models.GetPolls:
				s.SendPolls(data.Data.([]models.Poll))
			}
		}
	}
}
