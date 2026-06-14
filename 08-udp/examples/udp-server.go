package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("UDP Server listening on :8080")

	buffer := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Printf("Received from %s: %s\n", remoteAddr, string(buffer[:n]))

	// Javob yuborish
	response := fmt.Sprintf("Server received: %s", string(buffer[:n]))
	conn.WriteToUDP([]byte(response), remoteAddr)
}
