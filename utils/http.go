package utils

import (
	"fmt"
	"net"
	"strings"
	"os"
)

// header
// GET /find/ HTTP/1.1
// Host: localhost:8080
// Connection: keep-alive
// Cache-Control: max-age=0
// sec-ch-ua: "Chromium";v="136", "Google Chrome";v="136", "Not.A/Brand";v="99"
// sec-ch-ua-mobile: ?0
// sec-ch-ua-platform: "Windows"
// Upgrade-Insecure-Requests: 1
// DNT: 1
// User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36
// Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
// Sec-Fetch-Site: cross-site
// Sec-Fetch-Mode: navigate
// Sec-Fetch-User: ?1
// Sec-Fetch-Dest: document
// Accept-Encoding: gzip, deflate, br, zstd
// Accept-Language: en-US,en;q=0.9

func serveFile(filename string, connection net.Conn) {
	filePath := "static/" + filename 
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprint(connection, "HTTP/1.1 500 Internal Server Error\r\n\r\nCould not read file")
		return
	}

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/html\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(content)) +
		"\r\n" +
		string(content)
	fmt.Fprint(connection, response)
}

func HandleRequest(lines []string, connection net.Conn) {
	if len(lines) == 0 {
		fmt.Println("Empty request")
		return
	}


	for _, line := range lines {
    fmt.Println(line)
	}
	fmt.Println("******************\n")
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

    switch path {
    case "/get":
        fmt.Println("/get/", path)
        serveFile("index.html", connection)
    case "/find":
        fmt.Println("/find", path)
    default:
        fmt.Println("Page not found")
    }

}