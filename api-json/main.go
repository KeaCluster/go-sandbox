package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	var err error
	client, err = ConnectToDB("connection_here")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	r := router.NewRouter()
	http.ListenAndServe(":8000", r)
}
