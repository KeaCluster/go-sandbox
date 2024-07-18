package chat

import (
  "bufio"
  "fmt"
  "net"
  "os"
)

func StartClient(username string) {
  conn, err := net.Dial("tcp", "localhost:8080")
  if err != nil {
    fmt.Println("Error connecting to server:", err)
    return
  }
  defer conn.Close()

  fmt.Println("Connected to chat server")

  go listenForMessages(conn)

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    message := scanner.Text()
    if message == "" {
      continue
    }
    fmt.Fprintf(conn, "%s\n", message)
  }
}

func listenForMessages(conn net.Conn) {
  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}
