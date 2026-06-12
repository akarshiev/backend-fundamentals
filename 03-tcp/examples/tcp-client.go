package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server")
	fmt.Println("Type messages (Ctrl+C to quit):")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "" {
			continue
		}

		conn.Write([]byte(message + "\n"))

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Server closed connection")
			return
		}

		fmt.Println(strings.TrimSpace(response))
	}
}
