package router

import (
	"HTTP/utils"
	"fmt"
	"net"
)

func RouteRequest(path string, connection net.Conn) {
	var routes = map[string]func(net.Conn){
		"/get": func(conn net.Conn) { utils.ServeFile("index.html", conn) },
		"/":    func(conn net.Conn) { utils.ServeFile("index.html", conn) },
		"/api": func(conn net.Conn) {
			response := `{"message": "Hello from API"}`
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: %d\r\n\r\n%s", len(response), response)
		},
	}

	// Call the handler if path exists, else return 404
	if handler, ok := routes[path]; ok {
		handler(connection)
	} else {
		fmt.Fprint(connection, "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\n404 page not found")
	}
}
