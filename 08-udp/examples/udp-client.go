package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Connected to UDP server")
	fmt.Println("Type messages (Ctrl+C to quit):")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "" {
			continue
		}

		// Ma'lumot yuborish
		conn.Write([]byte(message))
		fmt.Printf("Sent: %s\n", message)

		// Javob olish
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}

		fmt.Printf("Received: %s\n", string(buffer[:n]))
	}
}
