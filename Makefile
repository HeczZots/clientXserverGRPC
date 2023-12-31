CC = go build

CLIENT = $(wildcard ./cmd/client/*.go)
SERVER = $(wildcard ./cmd/server/*.go)

all: clean client server ./server ./client

start:./server ./client

mod: 
	go mod tidy

build: mod client server

client:
	$(CC) $(CLIENT)

server:
	$(CC) $(SERVER)

clean:
	rm client server