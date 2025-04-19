package websocket_notifier

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"sync"
	"vote-broadcast-server/proto/websocket"
	"vote-broadcast-server/services/vote/pkg/models"
)

type WebSocketNotifierService struct {
	grpcConnections map[string]*grpcConnection
	dataChannel     models.DataChannel
	config          models.WebSocketGRPCServer
	wg              *sync.WaitGroup
}

type grpcConnection struct {
	conn           *grpc.ClientConn
	clientInstance interface{}
}

func NewWebSocketNotifierService(config models.WebSocketGRPCServer) *WebSocketNotifierService {
	return &WebSocketNotifierService{
		config:          config,
		dataChannel:     make(models.DataChannel),
		grpcConnections: make(map[string]*grpcConnection),
		wg:              &sync.WaitGroup{},
	}
}

func (s *WebSocketNotifierService) GetDataChannel() models.DataChannel {
	return s.dataChannel
}

type WebSocketNotifier interface {
	Start()
	GetDataChannel() models.DataChannel
}

func (s *WebSocketNotifierService) Start() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	s.wg.Add(1)
	go s.listenUpdatedServerData(ctx)

	log.Println("WebSocketNotifierService started")
	<-interrupt
	cancel()
	s.wg.Wait()
	log.Println("WebSocketNotifierService stopped")
}

func (s *WebSocketNotifierService) getGRPCServer(serviceName string) (interface{}, *grpc.ClientConn, error) {
	if data, exists := s.grpcConnections[serviceName]; exists {
		return data.clientInstance, data.conn, nil
	}

	conn, err := grpc.NewClient(s.config.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	var clientInstance interface{}

	switch serviceName {
	case "websocket":
		clientInstance = websocket.NewWebSocketServiceClient(conn)
	default:
		return nil, nil, errors.New("service not found: " + serviceName)
	}

	s.grpcConnections[serviceName] = &grpcConnection{
		clientInstance: clientInstance,
		conn:           conn,
	}

	return clientInstance, conn, nil
}
