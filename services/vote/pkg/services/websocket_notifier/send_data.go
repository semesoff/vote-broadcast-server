package websocket_notifier

import (
	"context"
	"vote-service/pkg/models"
	"vote-service/pkg/utils"
	"vote-service/proto/websocket"
)

func (s *WebSocketNotifierService) SendVotes(pollVotes models.PollVotes) {
	serviceClient, _, err := s.getGRPCServer("websocket")
	if err != nil {
		return
	}

	websocketClient, ok := serviceClient.(websocket.WebSocketServiceClient)
	if !ok {
		return
	}

	var request websocket.VotesRequest
	request.PollId = int64(pollVotes.PollId)
	request.Options = utils.ConvertToProtoWebsocketVotes(pollVotes)

	// gRPC Request
	_, err = websocketClient.GetVotes(context.Background(), &request)
	if err != nil {
		return
	}
}
