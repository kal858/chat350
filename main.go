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

// Global map for storing the websocket connections
var connections map[*websocket.Conn]bool

// Struct needed for upgrading websocket connection
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Function to broadcast incoming message to all connections
func sendAll(msg []byte) {
	for conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			delete(connections, conn)
			return
		}
	}
}

// Function that handles the websocket connections. Upgrades the connection,
// reads in the message contents from the websocket connection, then calls
// sendAll to broadcast the message to all active connections
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
	host := flag.String("host", "localhost", "IP address to listen on")
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("directory", "www/", "Web directory")
	flag.Parse()

	// Initialize the connections data structure
	connections = make(map[*websocket.Conn]bool)

	// Set up the request handlers
	fs := http.Dir(*dir)
	filehandler := http.FileServer(fs)

	// HTTP handler to serve back files in www/ directory
	http.Handle("/", filehandler)

	// Handle incoming websocket connection
	http.HandleFunc("/ws", wsHandler)

	log.Printf("Listening on %s:%d\n", *host, *port)
	addr := fmt.Sprintf("%s:%d", *host, *port)

	// Enter into blocking mode
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
