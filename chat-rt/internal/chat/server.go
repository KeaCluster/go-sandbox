package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/keacluster/go-sandbox/chat-rt/internal/models"
)

type Server struct {
  clients   map[net.Conn]string
  broadcast chan models.Message
  mu sync.Mutex
}

func NewServer() *Server {
  return &Server {
    clients: make(map[net.Conn]string),
    broadcast: make(chan models.Message),
  }
}

func (s *Server) Run() {
  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println("Error starting server", err)
    return
  }
  defer listener.Close()
  fmt.Println("Chat server started on port 8080")

  go s.handleBroadCasts()

  for {
    conn, err := listener.Accept()
    if err != nil {
      fmt.Println("Error accepting connection", err)
      continue
    }
    go s.handleNewClient(conn)
  }
}

func (s *Server) handleNewClient(conn net.Conn) {
  s.mu.Lock()
  s.clients[conn] = conn.RemoteAddr().String()
  s.mu.Unlock()

  fmt.Println("New client connected:", conn.RemoteAddr().String())

  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    msgContent := scanner.Text()
    if strings.TrimSpace(msgContent) == "" {
      continue
    }

    msg := models.Message {
      Username: s.clients[conn],
      Content: msgContent,
    }
    s.broadcast <- msg
  }
} 

func (s *Server) handleBroadCasts() {
  for msg := range s.broadcast {
    s.mu.Lock()
    for client := range s.clients {
      _, err := fmt.Fprintf(client, "%s: %s\n", msg.Username, msg.Content)
      if err != nil {
        client.Close()
        delete(s.clients, client)
      }
    }
    s.mu.Unlock()
  }
}
