package main

import (
	"fmt"
	// "io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func (r *http.Request) bool {return true},
}

func reader(conn *websocket.Conn){
	for{
		messageType, p, err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		p  = []byte(`"`+ string(p) + `"` + " is recieved")
		if err := conn.WriteMessage(messageType, p); err != nil{
			log.Println(err)
			return
		}
	}
}

func serverWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(ws)
	}
	reader(ws)
}

func setupRoutes(){
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w, "Simple Server\n")
	})
	http.HandleFunc("/ws", serverWs)
}

func main(){
	fmt.Println("Chat App v0.01")
	setupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Println("err",err)
	}
	// fmt.Println()
}