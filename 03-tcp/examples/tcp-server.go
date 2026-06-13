package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

// TCP Server -- Echo server with goroutine, scanner, graceful shutdown
//
// Xususiyatlari:
// - Har bir client uchun alohida goroutine
// - Echo: client nima yuborsa, server qaytaradi
// - Graceful shutdown: Ctrl+C bilan tozalash
// - Client ulanish/uzilish loglari
// - Timestamp har xabarda

func main() {
	// Server yaratish
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Listen error:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP Server listening on :8080")
	fmt.Println("Press Ctrl+C to stop")

	// Graceful shutdown uchun signal handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Clientlarni track qilish uchun WaitGroup
	var wg sync.WaitGroup

	// Shutdown flag
	done := make(chan bool)

	// Signal handler goroutine
	go func() {
		<-sigChan
		fmt.Println("\nShutting down server...")
		close(done)
		listener.Close()
	}()

	// Asosiy loop -- clientlarni qabul qilish
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-done:
				fmt.Println("Waiting for active connections to close...")
				wg.Wait()
				fmt.Println("Server stopped.")
				return
			default:
				fmt.Println("Accept error:", err)
				continue
			}
		}

		wg.Add(1)
		go handleConnection(conn, &wg, done)
	}
}

// handleConnection -- har bir client uchun alohida goroutine
func handleConnection(conn net.Conn, wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()
	defer conn.Close()

	clientAddr := conn.RemoteAddr().String()
	timestamp := time.Now().Format("15:04:05")

	fmt.Printf("[%s] Client connected: %s\n", timestamp, clientAddr)

	// Client uzilishini tekshirish uchun goroutine
	disconnected := make(chan bool, 1)
	go func() {
		<-done
		disconnected <- true
	}()

	// Scanner -- satr-satr o'qish
	scanner := bufio.NewScanner(conn)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024) // 1MB buffer

	for {
		// Shutdown tekshirish
		select {
		case <-disconnected:
			fmt.Printf("[%s] Client %s: server shutting down\n",
				time.Now().Format("15:04:05"), clientAddr)
			return
		default:
		}

		// Scanner timeout
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		if scanner.Scan() {
			message := strings.TrimSpace(scanner.Text())
			if message == "" {
				continue
			}

			timestamp = time.Now().Format("15:04:05")
			fmt.Printf("[%s] %s: %s\n", timestamp, clientAddr, message)

			// Echo -- xabarni qaytarish
			response := fmt.Sprintf("[%s] Echo: %s\n",
				time.Now().Format("15:04:05"), message)
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Printf("[%s] Write error to %s: %v\n",
					time.Now().Format("15:04:05"), clientAddr, err)
				return
			}

			// "quit" yoki "exit" -- client uziladi
			if message == "quit" || message == "exit" {
				fmt.Printf("[%s] Client %s disconnected (quit)\n",
					time.Now().Format("15:04:05"), clientAddr)
				conn.Write([]byte("Goodbye!\n"))
				return
			}
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Printf("[%s] Client %s: read error: %v\n",
					time.Now().Format("15:04:05"), clientAddr, err)
			}
			return
		}
	}
}
