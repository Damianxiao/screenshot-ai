package network

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/", "./static/index.html")
	ws := r.Group("/v1")
	{
		// upgrade
		ws.GET("/ws", handlerWebSocket())
	}

	return r
}
