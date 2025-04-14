package websocket_server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"vote-broadcast-server/services/websocket/pkg/models"
)

type ServerManager struct {
	config       models.WebSocketServer
	clientsData  *clientsData
	dataChannels models.DataChannels
}

type clientsData struct {
	getPollsClients map[*websocket.Conn]bool
	getVotesClients map[int]map[*websocket.Conn]bool
	countClients    int
	pollsMutex      sync.RWMutex
	votesMutex      sync.RWMutex
}

func NewServerManager(config models.WebSocketServer, dataChannels models.DataChannels) *ServerManager {
	return &ServerManager{
		config: config,
		clientsData: &clientsData{
			getPollsClients: make(map[*websocket.Conn]bool),
			getVotesClients: make(map[int]map[*websocket.Conn]bool),
			pollsMutex:      sync.RWMutex{},
			votesMutex:      sync.RWMutex{},
		},
		dataChannels: dataChannels,
	}
}

type Server interface {
	Start(waitGroup *sync.WaitGroup, ctx context.Context)
}

func (s *ServerManager) Start(waitGroup *sync.WaitGroup, ctx context.Context) {
	privateCtx, privateCancel := context.WithCancel(context.Background())

	// start the websocket server
	go s.startWebSocketServer(privateCtx)
	// start listening for updated server data
	go s.listenUpdatedServerData(privateCtx)

	<-ctx.Done()
	privateCancel()
	waitGroup.Done()
	log.Println("WebSocket server is stopped.")
}

func (s *ServerManager) startWebSocketServer(ctx context.Context) {
	server := &http.Server{
		Addr: fmt.Sprintf(":%s", s.config.Port),
	}

	http.HandleFunc("/getPolls", s.subscribeClientToMethod)
	http.HandleFunc("/getVotes/{poll_id}", s.subscribeClientToMethod)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln("websocket error listen and serve: ", err)
		}
	}()

	log.Printf("websocket server is running on port: %s", s.config.Port)
	<-ctx.Done()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("websocket server shutdown failed: ", err)
	}
}
