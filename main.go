package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/polling", polling)
	r.GET("/sse", sse)
	r.GET("/ws", ws)

	r.Run("0.0.0.0:9000")
}

func polling(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": time.Now().Format(time.RFC3339),
	})
}

func sse(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		message := fmt.Sprintf("data: Event %s\n\n", time.Now().Format(time.RFC3339))
		_, err := c.Writer.WriteString(message)
		if err != nil {
			log.Println("Error writing SSE:", err)
			break
		}
		c.Writer.Flush()
		time.Sleep(1 * time.Second)
	}
}

func ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}

	defer conn.Close()

	for {
		message := fmt.Sprintf("WebSocket message: %s", time.Now().Format(time.RFC3339))

		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("WebSocket write error:", err)
			break
		}

		_, _, err = conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}
