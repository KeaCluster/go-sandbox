package main

import "github.com/keacluster/go-sandbox/chat-rt/internal/chat"

func main() {
  server := chat.NewServer()
  server.Run()
}
