# Redis Server and Client in Go

This repository contains a basic implementation of a Redis server and client in Go. The server and client support the GET and SET commands for storing and retrieving key-value pairs.

## Usage

### Running the Server
1. Navigate to the `server` directory.
2. Run the following command to build the server:
   ```bash
   go build server.go
   ```
3. Start the server by running the executable:
   ```bash
   ./server
   ```
   The server will start listening for connections on port 6378.

### Running the Client
1. Navigate to the `client` directory.
2. Run the following command to build the client:
   ```bash
   go build client.go
   ```
3. Start the client by running the executable:
   ```bash
   ./client
   ```
4. Enter commands at the prompt to interact with the server. Supported commands include GET and SET for retrieving and setting key-value pairs, respectively.

## Example Usage
Here's an example session demonstrating how to interact with the Redis server using the client:
```
Enter command (e.g., GET key or SET key value): SET mykey hello
Server response: OK
Enter command (e.g., GET key or SET key value): GET mykey
Server response: hello
Enter command (e.g., GET key or SET key value):
```

## Notes
- The server and client communicate over TCP/IP using a simple protocol.