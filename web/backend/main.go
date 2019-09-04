package main

import (
	"fmt"
	"net/http"

	"github.com/SiyaoHuang/GoLangTraining/web/backend/pkg/mongodb"
	"github.com/SiyaoHuang/GoLangTraining/web/backend/pkg/websocket"
	"github.com/gorilla/mux"
)

func serverWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(ws)
	}
	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	mongodb.Init()
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server\n")
	})
	pool := websocket.NewPool()
	go pool.Start()
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { serverWs(pool, w, r) })

	router.HandleFunc("/person", mongodb.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/person", mongodb.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/person/id/{id}", mongodb.GetPersonEndpointByID).Methods("GET")
	router.HandleFunc("/person/lastname/{lastname}", mongodb.GetPersonEndpointByLastname).Methods("GET")
	router.HandleFunc("/person/lastname/{lastname}", mongodb.DeletePersonEndpointByLastname).Methods("DELETE")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		fmt.Println("err", err)
	}
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()

}
