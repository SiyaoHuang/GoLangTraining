package main

import(
	"fmt"
	"net/http"

	"github.com/SiyaoHuang/GoLangTraining/web/backend/pkg/websocket"
)

func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil{
		fmt.Println(ws)
	}
	client := &websocket.Client{
		Conn : ws,
		Pool : pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes(){
	http.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w, "Simple Server\n")
	})
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func (w http.ResponseWriter, r *http.Request){serverWs(pool, w, r)})
}

func main(){
	fmt.Println("Chat App v0.01")
	setupRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println("err",err)
	}
}