package main

import (
	"fmt"
	"net"
	"net/http"
	"web-server/internal/handler"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	address := listener.Addr().String()

	// construct the server url
	serverURL := fmt.Sprintf("http://%s", address)

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/user", handler.UserHandler)
	http.HandleFunc("/books", handler.BookHandler)

	fmt.Println("Server is runing on:", serverURL)
	if err := http.Serve(listener, nil); err != nil {
		panic(err)
	}
}
