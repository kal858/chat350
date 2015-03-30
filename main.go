package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Parse the command line options for the web server
	port := flag.Int("port", 8080, "Port to listen on")
	dir := flag.String("directory", "www/", "Web directory")
	flag.Parse()

	// Set up the request handlers
	fs := http.Dir(*dir)
	handler := http.FileServer(fs)
	http.Handle("/", handler)

	log.Printf("Listening on port: %d\n", *port)
	addr := fmt.Sprintf("localhost:%d", *port)

	// Enter into blocking mode
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())

}
