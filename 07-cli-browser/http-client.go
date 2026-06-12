package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run http-client.go <host>")
		fmt.Println("Example: go run http-client.go example.com")
		return
	}

	host := os.Args[1]

	// Step 1: TCP connection
	fmt.Printf("Connecting to %s:80...\n", host)
	conn, err := net.Dial("tcp", host+":80")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected!")

	// Step 2: Build HTTP request
	request := fmt.Sprintf(
		"GET / HTTP/1.1\r\nHost: %s\r\nConnection: close\r\n\r\n",
		host,
	)

	// Step 3: Send request
	fmt.Printf("\nSending request:\n%s\n", request)
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	// Step 4: Read response
	fmt.Println("Response:")
	fmt.Println("--------")
	body, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println(string(body))
}
