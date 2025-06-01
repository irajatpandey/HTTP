package handlers

import (
	"HTTP/router"
	"fmt"
	"net"
	"strings"
)



func HandleRequest(lines []string, connection net.Conn) {
	if len(lines) == 0 {
		fmt.Println("Empty request")
		return
	}


	for _, line := range lines {
    fmt.Println(line)
	}
	fmt.Println("******************")
	// Parse the request line
	requestLine := lines[0]
	fields := strings.Fields(requestLine)
	if len(fields) < 3 {
		fmt.Println("Invalid request line:", requestLine)
		return
	}

	method := fields[0]
	path := fields[1]
	version := fields[2]

	// Normalize path
	if len(path) > 1 && strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	fmt.Println("Method:", method)
	fmt.Println("Path:", path)
	fmt.Println("Version:", version)


	headers := make(map[string]string)
	for _, line := range lines[1:] {
		if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	for k, v := range headers {
		fmt.Printf("%s : %s\n", k, v)
	}

     router.RouteRequest(path, connection)
}
