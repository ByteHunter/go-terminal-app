BINARY_NAME=go-terminal-app

build:
	go build -o ${BINARY_NAME} main.go canvas.go chronos.go keys.go terminal.go utils.go

run:
	go run main.go canvas.go chronos.go keys.go terminal.go utils.go
