//go:build ignore

// ============================================
// 02-raw-http.go — Raw HTTP Request/Response
// ============================================
//
// Bu faylda HTTP request'ni QO'lda yozamiz va response'ni parse qilamiz.
//
// HTTP Request formati:
//
//	GET /path HTTP/1.1\r\n    ← Request Line
//	Host: example.com\r\n     ← Header
//	\r\n                      ← Bo'sh qator (headers tugadi)
//
// HTTP Response formati:
//
//	HTTP/1.1 200 OK\r\n       ← Status Line
//	Content-Type: ...\r\n     ← Header
//	\r\n                      ← Bo'sh qator
//	<body>                    ← Body
//
// Ishga tushirish:
//
//	go run 02-raw-http.go
//

package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	// ============================================
	// 1. TCP Connection ochamiz
	// ============================================
	// Port 80 — HTTP (HTTPS uchun 443)
	conn, err := net.Dial("tcp", "httpbin.org:80")
	if err != nil {
		fmt.Printf("Xatolik: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("TCP Connection tayyor ✅")
	fmt.Println()

	// ============================================
	// 2. HTTP Request'ni QO'lda yozamiz
	// ============================================
	//
	// Bu qo'lda yozilgan HTTP request:
	//
	//   GET /get HTTP/1.1          ← Request Line (Method, Path, Version)
	//   Host: httpbin.org          ← Header (majburiy)
	//   Accept: */*                ← Header (qanday format qabul qiladi)
	//   Connection: close          ← Header (javobdan keyin connection yopiladi)
	//                              ← Bo'sh qator (headers tugadi)
	//
	// Har bir qator \r\n (CR+LF) bilan tugaydi.
	// Headers va Body orasida BO'SH QATOR kerak!
	request := "GET /get HTTP/1.1\r\n" +
		"Host: httpbin.org\r\n" +
		"Accept: */*\r\n" +
		"Connection: close\r\n" +
		"\r\n" // Bo'sh qator — headers tugadi, body boshlanadi

	fmt.Println("--- Yuborilgan HTTP Request ---")
	fmt.Println(request)

	// TCP socket orqali yuboramiz
	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Printf("Yozish xatolik: %v\n", err)
		return
	}
	fmt.Println("Request yuborildi ✅")
	fmt.Println()

	// ============================================
	// 3. HTTP Response'ni o'qiymiz
	// ============================================
	//
	// Server javobini qatorma-qator o'qiymiz.
	// Birinchi qator — Status Line (HTTP/1.1 200 OK)
	// Keyingi qatorlar — Headers
	// Bo'sh qator — Headers tugadi
	// Qolgan qatorlar — Body
	fmt.Println("--- Server Javobi (Raw) ---")
	fmt.Println(strings.Repeat("=", 60))

	scanner := bufio.NewScanner(conn)
	lineNum := 0
	statusCode := 0
	contentLength := 0
	headersDone := false
	bodyLines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Birinchi qator — Status Line
		if lineNum == 1 {
			fmt.Printf("Status Line: %s\n", line)
			// "HTTP/1.1 200 OK" → 200
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				statusCode, _ = strconv.Atoi(parts[1])
			}
			continue
		}

		// Bo'sh qator — headers tugadi
		if line == "" {
			fmt.Println("--- Headers tugadi ---")
			headersDone = true
			continue
		}

		// Headers
		if !headersDone {
			fmt.Printf("Header: %s\n", line)
			// Content-Length header'ini topamiz
			if strings.HasPrefix(strings.ToLower(line), "content-length:") {
				val := strings.TrimSpace(strings.SplitN(line, ":", 2)[1])
				contentLength, _ = strconv.Atoi(val)
			}
			continue
		}

		// Body
		bodyLines = append(bodyLines, line)
	}

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	// ============================================
	// 4. Tahlil
	// ============================================
	fmt.Println("--- Tahlil ---")
	fmt.Printf("Status Code:    %d\n", statusCode)
	fmt.Printf("Content-Length: %d bytes\n", contentLength)
	fmt.Printf("Body Lines:     %d\n", len(bodyLines))
	fmt.Println()

	// Body ni chiqaramiz (ko'pi bilan 10 qator)
	fmt.Println("--- Body ---")
	for i, line := range bodyLines {
		if i >= 10 {
			fmt.Println("  ... (davom etmoqda)")
			break
		}
		fmt.Printf("  %s\n", line)
	}
	fmt.Println()

	// ============================================
	// 5. Xulosa
	// ============================================
	fmt.Println("--- Xulosa ---")
	fmt.Println("HTTP = Matn protokoli")
	fmt.Println("Request:  METHOD PATH VERSION + Headers + Body")
	fmt.Println("Response: STATUS + Headers + Body")
	fmt.Println("Hammasi \\r\\n bilan ajratilgan")
}
