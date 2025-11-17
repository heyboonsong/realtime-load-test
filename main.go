package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/polling", polling)
	r.GET("/sse", sse)

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
