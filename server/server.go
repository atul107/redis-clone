/*
Package main implements a basic Redis server in Go.

This server listens for connections on port 6378 and handles GET and SET commands for storing and retrieving key-value pairs.

Usage:
- Run the executable to start the server.
- The server will listen for incoming connections and handle commands from clients.

Response Handling:
- The server responds to clients based on the received commands.
- If the command is valid, the server performs the requested operation and sends a response.
- If the command is invalid or incomplete, the server sends an error message to the client.
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var data = make(map[string]string)

func handleConnection(conn net.Conn) {
	defer conn.Close() // Close the connection when the function exits

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		// Read the command sent by the client
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from client:", err)
			return
		}

		// Trim whitespace and print the received command
		cmd = strings.TrimSpace(cmd)
		fmt.Println("Received command:", cmd)

		// Split the command into parts
		parts := strings.Fields(cmd)
		if len(parts) < 1 {
			// If the command is empty, send an error message
			writer.WriteString("ERR wrong number of arguments for 'get' command\r\n")
			writer.Flush()
		}

		// Handle different types of commands
		switch strings.ToUpper(parts[0]) {
		case "GET":
			// If the command is GET, retrieve the value for the specified key
			if len(parts) < 2 {
				writer.WriteString("ERR wrong number of arguments for 'get' command\r\n")
				writer.Flush()
			} else {
				key := parts[1]
				value := data[key]
				writer.WriteString(fmt.Sprintf("%s\r\n", value))
				writer.Flush()
			}
		case "SET":
			// If the command is SET, set the value for the specified key
			if len(parts) < 3 {
				writer.WriteString("ERR wrong number of arguments for 'set' command\r\n")
				writer.Flush()
			} else {
				key := parts[1]
				value := parts[2]
				data[key] = value
				writer.WriteString("OK\r\n")
				writer.Flush()
			}
		default:
			// If the command is unknown, send an error message
			writer.WriteString("Unknown command\r\n")
			writer.Flush()
		}
	}
}

func main() {
	// Start listening for connections on port 6378
	listener, err := net.Listen("tcp", ":6378")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close() // Close the listener when the function exits

	fmt.Println("Redis server listening on port 6378")

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle connections concurrently
		go handleConnection(conn)
	}
}
