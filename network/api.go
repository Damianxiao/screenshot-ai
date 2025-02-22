package network

import (
	"fmt"
	"net/http"
	"screenshot-ai/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func newUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  4096, // 增大缓冲区
		WriteBufferSize: 4096,
		// allow all cross orign
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func handlerWebSocket() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		upgrader := newUpgrader()
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Logobj.Errorf("WebSocket 升级失败: %v", err)
			return
		}
		fmt.Println("websocket connection success")
		client := &Client{
			Id:   conn.LocalAddr().String(),
			Hub:  HubServer,
			Conn: conn,
			Send: make(chan WebSocketMessage),
		}
		HubServer.Register <- client

		go client.ReadPump()
		go client.WritePump()
	}
}
