//go:build ignore

// ============================================
// 01-tcp.go — TCP Connection from Scratch
// ============================================
//
// HTTP aslida TCP ustida ishlaydi.
// Bu faylda TCP connection qanday ochilishini ko'ramiz.
//
// Jarayon:
//   1. DNS Resolution: example.com → 93.184.216.34
//   2. TCP 3-Way Handshake: SYN → SYN-ACK → ACK
//   3. TCP Connection Established ✅
//
// Ishga tushirish:
//
//	go run 01-tcp.go
//

package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// ============================================
	// 1. TCP Connection ochamiz
	// ============================================
	//
	// net.Dial("tcp", "host:port") — TCP connection ochadi.
	//
	// Bu yerda avtomatik bajariladi:
	//   - DNS resolution: example.com → IP manzil
	//   - TCP 3-Way Handshake: SYN → SYN-ACK → ACK
	//   - Connection established ✅
	//
	// Bu oddiy HTTP uchun port 80 (HTTPS uchun 443)
	fmt.Println("TCP Connection ochilmoqda...")
	fmt.Println("Server: example.com:80")

	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		fmt.Printf("Xatolik: %v\n", err)
		return
	}
	defer conn.Close() // Funksiya tugagach connection yopiladi

	fmt.Println("TCP Connection tayyor! ✅")
	fmt.Println()

	// ============================================
	// 2. TCP Connection ma'lumotlari
	// ============================================
	//
	// LocalAddr — client manzili (sizning IP + port)
	// RemoteAddr — server manzili (example.com IP + port)
	fmt.Println("--- TCP Connection Ma'lumotlari ---")
	fmt.Printf("Local Address:  %s\n", conn.LocalAddr())
	fmt.Printf("Remote Address: %s\n", conn.RemoteAddr())
	fmt.Println()

	// ============================================
	// 3. Oddiy matn yuboramiz (HTTP emas!)
	// ============================================
	//
	// TCP — faqat matn/bayt yuboradi.
	// HTTP — matnning MAXSUS formati.
	//
	// Biz hozir oddiy matn yuboramiz, ko'ramiz server nima qaytaradi.
	fmt.Println("--- Oddiy matn yuborilmoqda ---")
	message := "Hello from TCP!\r\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Yozish xatolik: %v\n", err)
		return
	}
	fmt.Printf("Yuborildi: %q\n", message)
	fmt.Println()

	// ============================================
	// 4. Server javobini o'qiymiz
	// ============================================
	//
	// bufio.Scanner — TCP socket dan qatorma-qator o'qiydi.
	// Har bir qator (\r\n) bilan ajratilgan.
	fmt.Println("--- Server javobi ---")
	scanner := bufio.NewScanner(conn)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("  → %s\n", line)
		lineCount++
		if lineCount > 20 { // Ko'p qator kerak emas
			break
		}
	}
	fmt.Println()

	// ============================================
	// 5. Xulosa
	// ============================================
	fmt.Println("--- Xulosa ---")
	fmt.Println("TCP connection oddiy matn uzatadi.")
	fmt.Println("HTTP — bu matnning MAXSUS formati.")
	fmt.Println("Keyingi faylda HTTP requestni qo'lda yozamiz!")
	fmt.Println()
	fmt.Println("HTTP Request formati:")
	fmt.Println("  GET / HTTP/1.1\\r\\n")
	fmt.Println("  Host: example.com\\r\\n")
	fmt.Println("  \\r\\n")

	// Connection yopiladi (defer conn.Close())
}
