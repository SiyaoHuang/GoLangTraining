package websocket

import (
	"fmt"
)

type Pool struct{
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client]bool
	Broadcast chan Message
}

func NewPool() *Pool{
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start(){
	for{
		select{
		case client:= <- pool.Register:
			pool.Clients[client] = true
			fmt.Printf("new Client: %+v, size %v\n", client, len(pool.Clients))
			for client, _ := range(pool.Clients){
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "new client entered"})
			}
		case client:= <- pool.Unregister:
			delete(pool.Clients, client)
			fmt.Printf("Client leave: %+v, size %v\n", client, len(pool.Clients))
			for client, _ := range(pool.Clients){
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "client left"})
			}
		case message:= <- pool.Broadcast:
			fmt.Printf("sending message: %+v\n", message)
			for client, _ := range(pool.Clients){
				fmt.Println(client)
				if err := client.Conn.WriteJSON(message); err != nil{
					fmt.Println(err)
					return
				}
			}
		}
	}
}