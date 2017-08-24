kill $(lsof -i :8000 -c webcoder-server -c ^Google -t)
nodemon --exec "go build && ./webcoder-server" server.go