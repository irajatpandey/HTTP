# Simple HTTP Server in Go

This project implements a basic HTTP server from scratch in Go, demonstrating how to handle TCP connections, parse HTTP requests, and serve static files or API responses without using any external web frameworks.

---

## Project Structure

```
HTTP/
├── handlers/
│   ├── http.go         # Handles parsing and processing HTTP requests
│   └── static.go       # (Reserved for static file handlers)
├── router/
│   └── route.go        # Routes HTTP requests to appropriate handlers
├── utils/
│   └── contents.go     # Serves static files from the filesystem
├── static/
│   └── index.html      # Example static file (not shown here)
├── main.go             # Entry point, starts the TCP server
├── go.mod              # Go module definition
```

---

## Module & Function Overview

### `main.go`
**Purpose:** Entry point of the application. Sets up the TCP server, listens for incoming connections, and delegates request handling.

**Key Functions:**
- **`main()`**: Initializes the server, listens on a specified port, and spawns a goroutine for each client connection.
- **`handleConnection(connection net.Conn)`**: Reads the raw HTTP request from the client, parses it into lines, and passes it to the request handler.

---

### `handlers/http.go`
**Purpose:** Parses HTTP requests and delegates routing.

**Key Functions:**
- **`HandleRequest(lines []string, connection net.Conn)`**:
  - Parses the request line to extract the HTTP method, path, and version.
  - Normalizes the path (removes trailing slash except for root).
  - Parses HTTP headers into a map.
  - Prints request details for debugging.
  - Delegates the request to the router via `router.RouteRequest(path, connection)`.

---

### `router/route.go`
**Purpose:** Routes HTTP requests to the appropriate handler based on the request path.

**Key Functions:**
- **`RouteRequest(path string, connection net.Conn)`**:
  - Maintains a map of routes (paths) to handler functions.
  - Calls the corresponding handler if the path matches.
  - Returns a 404 Not Found response if the path is not registered.

**Example routes:**
- `/get` and `/` serve the `index.html` file.
- `/api` returns a JSON API response.

---

### `utils/contents.go`
**Purpose:** Serves static files from the `static` directory.

**Key Functions:**
- **`ServeFile(filename string, connection net.Conn)`**:
  - Reads the requested file from the `static` directory.
  - Determines the MIME type based on the file extension.
  - Sends the file content with appropriate HTTP headers.
  - Handles file-not-found errors gracefully.

---

## How It Works

1. **Server Startup:**  
   Run `main.go` to start the TCP server on the specified port (default: 8080).

2. **Connection Handling:**  
   For each incoming connection, the server reads the HTTP request, parses it, and passes it to the handler.

3. **Request Parsing:**  
   The handler extracts the HTTP method, path, version, and headers.

4. **Routing:**  
   The router matches the path to a handler function (serving a file or returning an API response).

5. **Response:**  
   The handler sends the appropriate HTTP response back to the client.

---

## Example Usage

- **GET /** or **GET /get**  
  Returns the `index.html` file from the `static` directory.

- **GET /api**  
  Returns a JSON response:  
  ```json
  {"message": "Hello from API"}
  ```

- **Any other path**  
  Returns a `404 Not Found` response.

---

## Extending the Server

- Add more routes in `router/route.go` by updating the `routes` map.
- Add more static files in the `static` directory.
- Implement additional handlers in `handlers/` as needed.

---

## Requirements

- Go 1.18 or newer

---

## Running the Server

```sh
go run main.go
```

The server will listen on port `8080` by default.  
You can set the `PORT` environment variable to change the port.

---
