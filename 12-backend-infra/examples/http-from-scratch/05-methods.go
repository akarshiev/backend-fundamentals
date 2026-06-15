//go:build ignore

// ============================================
// 05-methods.go — HTTP Methods
// ============================================
//
// HTTP Methods — serverga qanday ish qilish kerakligini aytadi.
//
// GET    — Ma'lumot olish (idempotent)
// POST   — Yangi resurs yaratish (NOT idempotent)
// PUT    — Resursni to'liq yangilash (idempotent)
// DELETE — Resursni o'chirish (idempotent)
// PATCH  — Resursni qisman yangilash (NOT idempotent)
//
// Ishga tushirish:
//
//	go run 05-methods.go
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
	// 1. GET — Ma'lumot olish
	// ============================================
	//
	// GET so'rovda body bo'lmaydi.
	// Faqat URL parametrlar orqali ma'lumot yuboriladi.
	//
	// GET /users?page=1&limit=10 HTTP/1.1
	// Host: api.example.com
	//
	fmt.Println("=== 1. GET — Ma'lumot olish ===")
	fmt.Println()
	doRequest("GET", "/get?page=1&limit=10", "", map[string]string{
		"Accept": "application/json",
	})
	fmt.Println()

	// ============================================
	// 2. POST — Yangi resurs yaratish
	// ============================================
	//
	// POST so'rovda body bo'ladi.
	// Server yangi resurs yaratadi va 201 Created qaytaradi.
	//
	// POST /users HTTP/1.1
	// Host: api.example.com
	// Content-Type: application/json
	//
	// {"name": "Otabek"}
	//
	fmt.Println("=== 2. POST — Yangi resurs yaratish ===")
	fmt.Println()
	doRequest("POST", "/post", `{"name": "Otabek", "role": "Developer"}`, map[string]string{
		"Content-Type": "application/json",
	})
	fmt.Println()

	// ============================================
	// 3. PUT — Resursni to'liq yangilash
	// ============================================
	//
	// PUT so'rovda body bo'ladi.
	// Resurs to'liq yangilanadi (qisman emas).
	// Idempotent — bir necha marta bersang ham bir xil natija.
	//
	// PUT /users/1 HTTP/1.1
	// Host: api.example.com
	// Content-Type: application/json
	//
	// {"name": "Otabek Nurmuhammad"}
	//
	fmt.Println("=== 3. PUT — Resursni to'liq yangilash ===")
	fmt.Println()
	doRequest("PUT", "/put", `{"name": "Otabek Nurmuhammad"}`, map[string]string{
		"Content-Type": "application/json",
	})
	fmt.Println()

	// ============================================
	// 4. DELETE — Resursni o'chirish
	// ============================================
	//
	// DELETE so'rovda body bo'lmaydi (odatiy).
	// Server resursni o'chiradi va 204 No Content qaytaradi.
	//
	// DELETE /users/1 HTTP/1.1
	// Host: api.example.com
	//
	fmt.Println("=== 4. DELETE — Resursni o'chirish ===")
	fmt.Println()
	doRequest("DELETE", "/delete", "", nil)
	fmt.Println()

	// ============================================
	// 5. Idempotent tushuntirish
	// ============================================
	fmt.Println("=== Idempotent Nima? ===")
	fmt.Println()
	fmt.Println("GET /users/1    → Har safar bir xil natija ✅")
	fmt.Println("PUT /users/1    → Har safar bir xil natija ✅")
	fmt.Println("DELETE /users/1 → Har safar bir xil natija ✅")
	fmt.Println("POST /users     → Har safar YANGI user yaratadi ❌")
	fmt.Println()
	fmt.Println("Idempotent = Bir necha marta bersang ham natija o'zgarmaydi")
	fmt.Println()

	// ============================================
	// 6. HTTP Methods jadvali
	// ============================================
	fmt.Println("=== HTTP Methods Jadvali ===")
	fmt.Println()
	fmt.Println("Method   | Body | Idempotent | Maqsad")
	fmt.Println("---------|------|------------|------------------")
	fmt.Println("GET      | ❌   | ✅         | Ma'lumot olish")
	fmt.Println("POST     | ✅   | ❌         | Yangi resurs")
	fmt.Println("PUT      | ✅   | ✅         | To'liq yangilash")
	fmt.Println("DELETE   | ❌   | ✅         | O'chirish")
	fmt.Println("PATCH    | ✅   | ❌         | Qisman yangilash")
	fmt.Println("HEAD     | ❌   | ✅         | Faqat header'lar")
	fmt.Println("OPTIONS  | ❌   | ✅         | Qo'llanilishi")
}

// doRequest — HTTP request yuboradi va javobni chiqaradi.
// Bu funksiya har bir method uchun ishlatiladi.
func doRequest(method, path, body string, headers map[string]string) {
	// TCP Connection
	conn, err := net.Dial("tcp", "httpbin.org:80")
	if err != nil {
		fmt.Printf("Xatolik: %v\n", err)
		return
	}
	defer conn.Close()

	// Request yozamiz
	fmt.Printf("Request: %s %s HTTP/1.1\n", method, path)
	fmt.Println("Headers:")

	// Asl header'lar
	reqHeaders := "Host: httpbin.org\r\n" +
		"Connection: close\r\n"

	for key, value := range headers {
		reqHeaders += fmt.Sprintf("%s: %s\r\n", key, value)
		fmt.Printf("  %s: %s\n", key, value)
	}

	// Body bo'lsa, Content-Length qo'shamiz
	if body != "" {
		reqHeaders += fmt.Sprintf("Content-Length: %d\r\n", len(body))
		fmt.Printf("  Content-Length: %d\n", len(body))
	}

	// To'liq request
	request := fmt.Sprintf("%s %s HTTP/1.1\r\n%s\r\n%s", method, path, reqHeaders, body)

	_, err = conn.Write([]byte(request))
	if err != nil {
		fmt.Printf("Yozish xatolik: %v\n", err)
		return
	}

	// Response
	fmt.Println()
	fmt.Println("Response:")
	scanner := bufio.NewScanner(conn)
	scanner.Buffer(make([]byte, 0), 1024*1024)
	lineNum := 0
	inBody := false
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		if lineNum == 1 {
			fmt.Printf("  %s\n", line)
			continue
		}

		if line == "" && !inBody {
			fmt.Println("  --- Body ---")
			inBody = true
			continue
		}

		if inBody && lineNum < 25 {
			fmt.Printf("  %s\n", line)
		}
	}
	fmt.Println(strings.Repeat("-", 50))
}
