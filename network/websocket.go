package network

import (
	"encoding/json"
	"fmt"
	"screenshot-ai/api"
	"screenshot-ai/pkg/log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var HubServer *Hub

const (
	Type = iota
	TypeScreenShot
	TypeModel
	TypeResp
)

type Client struct {
	Id   string
	Conn *websocket.Conn
	Send chan WebSocketMessage
	Hub  *Hub
}

type WebSocketMessage struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Model   string `json:"model,omitempty"`
}

type Hub struct {
	Clients    map[string]Client
	Broadcast  chan WebSocketMessage
	Register   chan *Client
	UnRegister chan *Client
	Lock       sync.Mutex
}

func NewHub() *Hub {
	HubServer = &Hub{
		Clients:    make(map[string]Client),
		Broadcast:  make(chan WebSocketMessage),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
	}
	return HubServer
}

func (h *Hub) Start() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.Id] = *client
			fmt.Printf("Conection come in:%s,at %s,now clients num:%d\n", client.Id, time.Now().Format("15:04:05"), len(h.Clients))
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client.Id]; ok {
				close(client.Send)
				delete(h.Clients, client.Id)
				fmt.Printf("Conection disconect:%s,at %s,now clients num:%d\n", client.Id, time.Now().Format("15:04:05"), len(h.Clients))
			}
		case message := <-h.Broadcast:
			for _, client := range h.Clients {
				if message.Type == TypeModel {
					resp, err := api.CallApiWithimg(message.Content, message.Model)
					if err != nil {
						fmt.Printf("broadcast received err %s\n", err)
						continue
					}
					client.Send <- WebSocketMessage{Type: TypeResp, Content: resp}
				} else if message.Type == TypeScreenShot {
					//  screenshot data , to client
					fmt.Printf("send screenshot to phone")
					client.Send <- message
				}
			}
			// fmt.Printf("Hub receive msg: %s", message)
		}

	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Logobj.Errorf("client read failed: %s", err)
			}
			break
		}
		var wm WebSocketMessage
		// log.Logobj.Info("client receiced :%s")
		if err := json.Unmarshal(msg, &wm); err != nil {
			log.Logobj.Errorf("client read failed: %s", err)
			continue
		}
		// from client (model img tokenslimit...)
		c.Hub.Broadcast <- wm
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// fmt.Println("there no msg to send to hub")
				return
			}
			jsonMsg, _ := json.Marshal(message)
			if err := c.Conn.WriteMessage(websocket.TextMessage, jsonMsg); err != nil {
				log.Logobj.Errorf("send message failed,%s", err)
				return
			}
		}
	}
}
