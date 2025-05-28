Requirements 

##Building HTTP Protocol in Go


## Functional Requirement ##

### 1. TCP Connection Handlinng
- The server must listen for incoming TCP connections on a configurable port (default: '8080').
- Each client connection should be handled concurrently using GoRoutines

### 2. HTTP Request Parsing
- The server must support the following HTTP methods:
- 'GET' (initially)
- 'POST' (optional for future expansion)
- It must parse:
- Request line (method, path, versions)
- Headers
- Query Parameters

### 3. Static File Handling
- The server should serve files from a predefined root directory (e.g. '/public')
- Must correctly set the 'Content-Type' based on file extension

### 4. Response Handling
- Respond with a valid HTTP response:
- Status line(e.g., 'HTTP/1.1 200 'OK')
- Headers('Content-Type', 'Content-Length', etc)
- Body(file content, HTML, text, etc)


### 5. Error Handling
- Return appropriate error response:
 - '404 Not Found' if file/path does not exist
 - '400 Bad Request' for malformed requests
 - '500 Internal Server Error' for server issues

 ### 6. Logging
 - Log Each request to the consolse with:
 - Timestamp
 - Method
 - URL
 - Response status

 ---

 ## Non Functional Requirements 
