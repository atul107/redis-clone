/*
Package main implements a basic Redis client in Go.

This client connects to a Redis server running on localhost at port 6378 and sends commands to interact with the server.

Usage:
1. Run the executable to start the client.
2. Enter commands at the prompt to interact with the Redis server. Supported commands include GET and SET for retrieving and setting key-value pairs, respectively.

Response Handling:
- The client sends commands to the server and prints the server's response to the console.
- If there is an error during command execution or response handling, an error message is printed to the console.
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the Redis server running on localhost at port 6378
	conn, err := net.Dial("tcp", "localhost:6378")
	if err != nil {
		fmt.Println("Error connecting to Redis server:", err)
		return
	}
	defer conn.Close() // Close the connection when the main function exits

	// Create reader and writer for reading user input and writing to the server
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		// Prompt the user to enter a command
		fmt.Print("Enter command (e.g., GET key or SET key value): ")
		// Read the command entered by the user
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		// Write the command to the server
		_, err = writer.WriteString(cmd)
		if err != nil {
			fmt.Println("Error writing command to server:", err)
			return
		}
		// Flush the writer to ensure the command is sent immediately
		writer.Flush()

		// Read and print server response
		resp, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response from server:", err)
			return
		}
		fmt.Println("Server response:", resp)
	}
}
