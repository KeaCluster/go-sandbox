package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	address := listener.Addr().String()

	// construct the server url
	serverURL := fmt.Sprintf("http://%s", address)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Web server running")
	})

	fmt.Println("Server is runing on:", serverURL)
	if err := http.Serve(listener, nil); err != nil {
		panic(err)
	}
}
