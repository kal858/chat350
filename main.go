/*
	Author: Kyle LeBlanc
	Class:  CMPT 350
	File:   main.go

	Description: Main executable for chat350 real time chat web app.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var connections map[*websocket.Conn]bool

func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	connections[conn] = true
	for {
		// Read messages from the connection
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			return
		}
		log.Println(string(msg))
		sendAll(msg)

	}
}

func main() {
	// Parse the command line options for the web server
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("directory", "www/", "Web directory")
	flag.Parse()

	connections = make(map[*websocket.Conn]bool)

	// Set up the request handlers
	fs := http.Dir(*dir)
	filehandler := http.FileServer(fs)
	http.Handle("/", filehandler)
	http.HandleFunc("/ws", wsHandler)

	log.Printf("Listening on port: %d\n", *port)
	addr := fmt.Sprintf("localhost:%d", *port)

	// Enter into blocking mode
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())

}
