# web-server

## About

A small web-server-like application written in go

### Why

Just because

### How

```bash
go run web-server/cmd/main.go
```

### Code

- Main
  - Entry point for everything else
- Handlers
  - They handle the `/user` and `/books` endpoints.
  - Creating instances of the models
  - Handling errors
- Models
  - Only `user` and `book`
  - No relationships
- Utils
  - Error handling for `StatusInternalServerError`
  - In case something goes wrong

### Endpoints

Just three

- "/"
  - this will return a small string saying the server is running
- "/user"
  - This will return a user json with Nacho libre information (no phishing)
- "/books"
  - This will return a collection of book objects in the likes of `json`
