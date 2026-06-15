//go:build ignore

// ============================================
// 04-body.go — HTTP Body Streaming
// ============================================
//
// HTTP Body — request/responsening ma'lumot qismi.
// GET so'rovda body bo'lmaydi, POST/PUT so'rovda body bo'ladi.
//
// Ikki usul:
//   1. Content-Length — body uzunligi oldindan ma'lum
//   2. Chunked — body uzunligi noma'lum, qismma-qism keladi
//
// Ishga tushirish:
//
//	go run 04-body.go
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
	// 1. POST request bilan body yuboramiz
	// ============================================
	//
	// POST request formati:
	//
	//   POST /post HTTP/1.1        ← Request Line
	//   Host: httpbin.org          ← Header
	//   Content-Type: application/json  ← Body formati
	//   Content-Length: 42         ← Body uzunligi (byte)
	//                              ← Bo'sh qator
	//   {"name": "Otabek"}         ← Body

	conn, err := net.Dial("tcp", "httpbin.org:80")
	if err != nil {
		fmt.Printf("Xatolik: %v\n", err)
		return
	}
	defer conn.Close()

	// Body — JSON formatda ma'lumot
	body := `{"name": "Otabek", "role": "Developer"}`

	// POST request yozamiz
	request := "POST /post HTTP/1.1\r\n" +
		"Host: httpbin.org\r\n" +
		"Content-Type: application/json\r\n" +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) +
		"Connection: close\r\n" +
		"\r\n" +
		body // Body — bo'sh qatordan keyin

	fmt.Println("--- POST Request ---")
	fmt.Println("Request Line: POST /post HTTP/1.1")
	fmt.Println("Headers:")
	fmt.Println("  Host: httpbin.org")
	fmt.Println("  Content-Type: application/json")
	fmt.Printf("  Content-Length: %d\n", len(body))
	fmt.Println("  Connection: close")
	fmt.Println("Body:")
	fmt.Printf("  %s\n", body)
	fmt.Println()

	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Printf("Yozish xatolik: %v\n", err)
		return
	}
	fmt.Println("Request yuborildi ✅")
	fmt.Println()

	// ============================================
	// 2. Response'ni o'qiymiz
	// ============================================
	fmt.Println("--- Server Javobi ---")
	fmt.Println(strings.Repeat("=", 60))

	scanner := bufio.NewScanner(conn)
	lineNum := 0
	inBody := false
	bodyContent := []string{}

	// Katta body uchun buffer
	scanner.Buffer(make([]byte, 0), 1024*1024) // 1MB buffer

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		// Status Line
		if lineNum == 1 {
			fmt.Printf("Status: %s\n\n", line)
			continue
		}

		// Bo'sh qator — headers tugadi
		if line == "" && !inBody {
			fmt.Println("--- Headers tugadi, Body boshlandi ---")
			fmt.Println()
			inBody = true
			continue
		}

		// Body
		if inBody {
			bodyContent = append(bodyContent, line)
		} else {
			// Header'larni chiroyli chiqaramiz
			fmt.Printf("  %s\n", line)
		}
	}

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	// ============================================
	// 3. Body tahlili
	// ============================================
	fmt.Println("--- Body Tahlili ---")
	fmt.Println()
	fmt.Println("Body qabul qilindi!")
	fmt.Println("Uzunligi:", len(strings.Join(bodyContent, "\n")), "qator")
	fmt.Println()

	// Body ni to'liq chiqaramiz (ko'pi bilan 20 qator)
	fmt.Println("--- Body Content ---")
	for i, line := range bodyContent {
		if i >= 20 {
			fmt.Println("  ... (davom etmoqda)")
			break
		}
		fmt.Printf("  %s\n", line)
	}
	fmt.Println()

	// ============================================
	// 4. Content-Length tushuntirish
	// ============================================
	fmt.Println("--- Content-Length Nima Uchun Muhim? ---")
	fmt.Println()
	fmt.Println("Content-Length yo'q bo'lsa:")
	fmt.Println("  Server qachon body tugaganini bilmaydi!")
	fmt.Println("  → Timeout kerak yoki Connection: close")
	fmt.Println()
	fmt.Println("Content-Length bor bo'lsa:")
	fmt.Println("  Server aniq biladi — shuncha byte o'qish kerak")
	fmt.Println("  → Tez va samarali")
	fmt.Println()

	// ============================================
	// 5. Chunked Transfer Encoding
	// ============================================
	fmt.Println("--- Chunked Transfer Encoding ---")
	fmt.Println()
	fmt.Println("Agar body uzunligi oldindan noma'lum bo'lsa:")
	fmt.Println("  Transfer-Encoding: chunked")
	fmt.Println()
	fmt.Println("Format:")
	fmt.Println("  SIZE\\r\\n     ← Har bir chunkning uzunligi")
	fmt.Println("  DATA\\r\\n     ← Ma'lumot")
	fmt.Println("  SIZE\\r\\n     ← Keyingi chunk")
	fmt.Println("  DATA\\r\\n     ← Ma'lumot")
	fmt.Println("  0\\r\\n         ← Tugadi!")
	fmt.Println("  \\r\\n")
	fmt.Println()
	fmt.Println("Misol:")
	fmt.Println("  1a\\r\\n                    ← 26 byte")
	fmt.Printf("  %s\\r\\n", `{"name": "Otabek"}`)
	fmt.Println("  0\\r\\n                    ← Tugadi")
}
