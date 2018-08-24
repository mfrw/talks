all: server client

server: server.go common.go
	go build server.go common.go

client: client.go common.go
	go build client.go common.go

.PHONY: clean
clean:
	rm -f server client
