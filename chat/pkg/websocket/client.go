package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Conn  *websocket.Conn
	Pool  *Pool
	Rooms map[*Room]bool
	mu    sync.Mutex
}

// type Username struct {
// 	Name string `json:"name"`
// }

// type allUsername []Username // list of usernames

// var usernames = allUsername{
// 	{

// 		Name: "@bot",
// 	},
// }

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)

	}
}
