package websocket

import "fmt"

type Room struct {
	Name       string
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
	Private    bool `json:"private"`
}

// NewRoom creates a new Room
func NewRoom(name string) *Room {
	return &Room{
		Name:       name,
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (room *Room) Start() {
	for {
		select {
		case client := <-room.Register:
			room.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(room.Clients))
			for client := range room.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-room.Unregister:
			delete(room.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(room.Clients))
			for client := range room.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-room.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range room.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
