package utils

import (
	"fmt"
	"mime"
	"net"
	"os"
	"path/filepath"
)

func ServeFile(filename string, connection net.Conn) {
	filePath := "static/" + filename
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprint(connection, "HTTP/1.1 500 Internal Server Error\r\n\r\nCould not read file")
		return
	}

	ext := filepath.Ext(filename)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	response := "HTTP/1.1 200 OK\r\n" +
		fmt.Sprintf("Content-Type: %s\r\n", contentType) +
		fmt.Sprintf("Content-Length: %d\r\n", len(content)) +
		"\r\n" + string(content)

	fmt.Fprint(connection, response)
}
