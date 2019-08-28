package main

import(
	"fmt"
	"net/http"

	// "backend/pkg/websocket"
)

func serverWs(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(ws)
	}
	go websocket.writer(ws)
	websocket.reader(ws)
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