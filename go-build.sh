kill $(lsof -i :8000 -c go-code-together -c ^Google -t)
nodemon --exec "go build && ./go-code-together" server.go