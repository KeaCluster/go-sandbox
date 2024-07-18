package chat

import (
  "flag"
  "fmt"

  "github.com/keacluster/go-sandbox/chat-rt/internal/chat"
)

func main() {
  username := flag.String("username", "Anonymous", "Your chat username")
  flag.Parse()

  if *username == "" {
    fmt.Println("Username cannot be empty")
    return
  }

  chat.StartClient(*username)
}
