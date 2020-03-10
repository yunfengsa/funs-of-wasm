#/bin/sh
GOOS=js GOARCH=wasm go build -o main.wasm src/main.go 
go run server.go