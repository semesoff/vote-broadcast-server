package websocket_notifier

import (
	"context"
	"google.golang.org/grpc"
	"poll-service/pkg/models"
	"poll-service/pkg/utils"
	"poll-service/proto/websocket"
)

func (s *WebSocketNotifierService) SendPolls(polls []models.Poll) {
	serviceClient, conn, err := s.getGRPCServer("websocket")
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
		}
	}(conn)

	if err != nil {
		return
	}

	websocketClient, ok := serviceClient.(websocket.WebSocketServiceClient)
	if !ok {
		return
	}

	var request websocket.PollsRequest
	request.Polls = utils.ConvertToProtoWebSocketPolls(polls)

	// gRPC Request
	_, err = websocketClient.GetPolls(context.Background(), &request)
	if err != nil {
		return
	}
}
