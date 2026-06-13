package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// TCP Client -- serverga ulanib, xabar yuborish va javob olish
//
// Xususiyatlari:
// - Serverga ulanish
// - Foydalanuvchi kiritgan xabarlarni yuborish
// - Server javoblarini o'qish
// - "quit" yoki "exit" bilan uzilish
// - Timeout tekshirish

func main() {
	// Serverga ulanish
	conn, err := net.DialTimeout("tcp", "localhost:8080", 5*time.Second)
	if err != nil {
		fmt.Println("Connection error:", err)
		fmt.Println("Make sure server is running: go run tcp-server.go")
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to TCP server (localhost:8080)")
	fmt.Println("Type messages (Ctrl+C to quit, 'quit' to disconnect)")
	fmt.Println(strings.Repeat("-", 50))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nGoodbye!")
			return
		}

		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		// Serverga xabar yuborish
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		// "quit" -- client uziladi
		if message == "quit" || message == "exit" {
			// Server javobini o'qish
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			response, err := bufio.NewReader(conn).ReadString('\n')
			if err == nil {
				fmt.Println(strings.TrimSpace(response))
			}
			fmt.Println("Disconnected from server.")
			return
		}

		// Server javobini o'qish
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			fmt.Println("Server closed connection.")
			return
		}

		fmt.Println(strings.TrimSpace(response))
	}
}
