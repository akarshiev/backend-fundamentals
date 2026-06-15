//go:build ignore

// ============================================
// 03-headers.go — HTTP Headers Parsing
// ============================================
//
// Headers — HTTP request/response metadata'si.
// Key: Value formatida, bir qatorga joylashadi.
//
// Misol:
//
//	Content-Type: application/json; charset=utf-8
//	     ↑          ↑                    ↑
//	   Name       Value              Parameters
//
// Ishga tushirish:
//
//	go run 03-headers.go
//

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// ============================================
	// 1. Turli header'lar bilan request yozamiz
	// ============================================
	//
	// Header'lar turli ma'lumotlar beradi:
	//   - Host: Qaysi serverga murojaat
	//   - Accept: Qanday format qabul qilamiz
	//   - User-Agent: Qaysi dastur ishlatilmoqda
	//   - Authorization: Autentifikatsiya tokeni
	//   - Content-Type: Body formati
	//   - Connection: Connection turu

	conn, err := net.Dial("tcp", "httpbin.org:80")
	if err != nil {
		fmt.Printf("Xatolik: %v\n", err)
		return
	}
	defer conn.Close()

	// Turli header'lar bilan GET request
	request := "GET /headers HTTP/1.1\r\n" +
		"Host: httpbin.org\r\n" +
		"Accept: application/json\r\n" +
		"User-Agent: Go-Scratch/1.0\r\n" +
		"X-Custom-Header: Salom Dunyo\r\n" +
		"Connection: close\r\n" +
		"\r\n"

	fmt.Println("--- Yuborilgan Header'lar ---")
	// Request qatorlarini chiqaramiz
	for _, line := range strings.Split(request, "\r\n") {
		if line == "" {
			fmt.Println("  (bo'sh qator — headers tugadi)")
		} else {
			fmt.Printf("  %s\n", line)
		}
	}
	fmt.Println()

	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Printf("Yozish xatolik: %v\n", err)
		return
	}

	// ============================================
	// 2. Response header'larini parse qilamiz
	// ============================================
	fmt.Println("--- Qabul qilingan Header'lar ---")
	fmt.Println(strings.Repeat("=", 50))

	scanner := bufio.NewScanner(conn)
	lineNum := 0
	inHeaders := true
	parsedHeaders := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Status Line
		if lineNum == 1 {
			fmt.Printf("Status: %s\n\n", line)
			continue
		}

		// Bo'sh qator — headers tugadi
		if line == "" {
			fmt.Println("--- Headers tugadi ---")
			inHeaders = false
			continue
		}

		// Header parse qilish
		if inHeaders {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				parsedHeaders[key] = value
				fmt.Printf("  %-25s → %s\n", key, value)
			}
		}
	}

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()

	// ============================================
	// 3. Header'lar tahlili
	// ============================================
	fmt.Println("--- Header'lar Tahlili ---")
	fmt.Println()

	// Content-Type tahlili
	if ct, ok := parsedHeaders["Content-Type"]; ok {
		fmt.Printf("Content-Type: %s\n", ct)
		// "application/json; charset=utf-8" → tahlil
		parts := strings.Split(ct, ";")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "charset=") {
				fmt.Printf("  Charset: %s\n", strings.TrimPrefix(part, "charset="))
			} else {
				fmt.Printf("  Media Type: %s\n", part)
			}
		}
		fmt.Println()
	}

	// Connection tahlili
	if conn, ok := parsedHeaders["Connection"]; ok {
		fmt.Printf("Connection: %s\n", conn)
		if conn == "close" {
			fmt.Println("  → Server javobdan keyin connection ni yopadi")
		} else if conn == "keep-alive" {
			fmt.Println("  → Server connection ni saqlab qo'yadi")
		}
		fmt.Println()
	}

	// Server tahlili
	if server, ok := parsedHeaders["Server"]; ok {
		fmt.Printf("Server: %s\n", server)
		fmt.Println("  → Server dasturi haqida ma'lumot")
		fmt.Println()
	}

	// ============================================
	// 4. Header'larning ahamiyati
	// ============================================
	fmt.Println("--- Header'larning Ahamiyati ---")
	fmt.Println()
	fmt.Println("Content-Type: Body formati (JSON, HTML, XML, ...)")
	fmt.Println("Content-Length: Body uzunligi (byte)")
	fmt.Println("Host: Qaysi domain (virtual hosting uchun)")
	fmt.Println("Authorization: Auth token (Bearer, Basic, ...)")
	fmt.Println("Cache-Control: Keshlash qoidalari")
	fmt.Println("Transfer-Encoding: Body uzatish usuli (chunked)")
	fmt.Println("Connection: Connection turu (keep-alive, close)")
}
