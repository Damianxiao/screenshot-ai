package network

import (
	"screenshot-ai/config"
	"screenshot-ai/pkg/log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebSocketConn(t *testing.T) {
	config.InitConfig()
	log.InitLogger()
	NewHub()
	go HubServer.Start()
	engine := NewRouter()
	port := config.Cfg.Server.Port
	err := engine.Run(port)
	HubServer.Broadcast <- WebSocketMessage{}
	assert.Nil(t, err)
}
