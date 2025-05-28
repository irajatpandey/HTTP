package main

import (
	"HTTP/utils" // Importing the utils package for handling requests
	"bufio"
	"fmt"
	"net"
	"os"
)

	   
const (
	defaultPort = "8080" // Agar koi port nahi bataya toh yeh use karenge
)
/*HTTP Protocol in Go Language from scratch*/
func main() {


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Error Starting listener:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Server listening on port %s...\n", port)	

	for {

		connection, error := listener.Accept()
		if error != nil {
			fmt.Println("Error accepting connection:", error)
			continue // check for new connection
		}

		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	fmt.Println("Client Conneted:", connection.RemoteAddr())
	reader := bufio.NewReader(connection)

	// Read request line and headers
	var requestLines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		line = line[:len(line)-1] // Remove '\n'
		if line == "\r" || line == "" { // End of headers
			break
		}
		requestLines = append(requestLines, line)
	}

	if len(requestLines) == 0 {
		fmt.Println("Empty request")
		return
	}
	// Now pass the full request (line + headers) to your handler
	utils.HandleRequest(requestLines, connection)
	
}