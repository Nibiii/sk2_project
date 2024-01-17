package main

import (
	"encoding/json"
	"fmt"
	"net"
	"sk2-project/types"
	"strings"
)

const BUFFER_SIZE = 1024

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	files := new(types.Files)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleClient(conn, files)
	}
}

func handleClient(conn net.Conn, files *types.Files) {
	defer conn.Close()

	var v int
	var requestLine string
	buffer := make([]byte, BUFFER_SIZE)
	var requestString strings.Builder
	request := types.Request{}

	for {
		// Read data from the client
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Read data from buffer to string builder because its easier to convert to string
		if _, err := requestString.Write(buffer[:n]); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// If less than buffer size was read, assume that there's nothing left to read and proceed handling received request
		if n < BUFFER_SIZE {
			break
		}
	}

	// Split the request into strings by newline character
	split := strings.Split(requestString.String(), "\n")
	if len(split) < 1 {
		fmt.Println("Error: not enough fields passed in http request")
	}

	// The first line of the request will always be the request line
	requestLine = split[0]
	for v = 1; v < len(split); v++ {
		if split[v] == "\r" {
			break
		}
		split[v] = strings.ReplaceAll(split[v], "\r", "")
	}

	// Add double quotes to every header's key and value so that we can unmarhsal it to json
	for i := 1; i < v; i++ {
		split[i] = strings.ReplaceAll(split[i], "\"", "\\\"")
		temp := strings.Split(split[i], ": ")
		temp[0] = fmt.Sprintf("\"%s\"", temp[0])
		temp[1] = fmt.Sprintf("\"%s\"", temp[1])
		split[i] = strings.Join(temp, ": ")
	}

	// Check if the request line contains all of the required info (as in rfc2616)
	requestSplit := strings.Fields(requestLine)
	if len(split) < 3 {
		fmt.Println("Error: not enough fields in the request line")
	}

	// Unmarhsal the request's data to json object
	requestData := fmt.Sprintf("{\"Method\": \"%s\", \"Uri\": \"%s\", \"Http_Version\": \"%s\", \"Header\": {%s}, \"Body\": \"%s\"}", requestSplit[0], requestSplit[1], requestSplit[2], strings.Join(split[1:v], ", "), strings.Join(split[v+1:], "\\n"))
	if err := json.Unmarshal([]byte(requestData), &request); err != nil {
		fmt.Println(requestData)
		fmt.Println("Error:", err)
		return
	}

	// print request for debug
	fmt.Println("Received: ", requestData)

	// handle request
	res := request.Handle(conn, files)

	// send response
	request.Respond(conn, res)
}
