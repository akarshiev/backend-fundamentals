// ============================================
// Main — Proxy serverning kirish nuqtasi
// ============================================
//
// Bu fayl barcha komponentlarni birlashtiradi:
//   1. Konfiguratsiyani yuklaydi
//   2. LoadBalancer yaratadi
//   3. Rate Limiter yaratadi
//   4. Middleware zanjirini tuzadi
//   5. Serverni ishga tushiradi
//
// Ishga tushirish:
//
//	go run .
//
// Natija:
//
//	Proxy server ishga tushdi: http://localhost:8080
//	Backend serverlar: https://github.com, https://example.com
//	Rate Limit: 5 request / 60 soniya
//
// Sinov:
//
//	curl http://localhost:8080
//	curl http://localhost:8080/api/users
//
// Arxitektura:
//
//	Client Request
//	    ↓
//	[LoggingMiddleware]     ← Request logga yoziladi
//	    ↓
//	[RateLimitMiddleware]   ← Limit tekshiriladi
//	    ↓
//	[ProxyHandler]          ← Backend serverga yo'naltiriladi
//	    ↓
//	Backend Server (Round Robin)
//	    ↓
//	Client Response
//

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ============================================
	// 1. Konfiguratsiyani yuklaymiz
	// ============================================
	// DefaultConfig() — standart sozlamalarni qaytaradi.
	// Bu yerda UPSTREAMS va RATE_LIMIT o'zgartirilishi mumkin.
	config := DefaultConfig()

	// ============================================
	// 2. Load Balancer yaratamiz
	// ============================================
	// NewLoadBalancer — backend serverlar ro'yxatini oladi
	// va har biri uchun ReverseProxy yaratadi.
	//
	// Round Robin algoritmi bilan ishlaydi:
	//   Request 1 → Server 1
	//   Request 2 → Server 2
	//   Request 3 → Server 1 (qayta)
	lb := NewLoadBalancer(config.Upstreams)

	// ============================================
	// 3. Rate Limiter yaratamiz
	// ============================================
	// NewRateLimiter — Fixed Window Counter algoritmini ishlatadi.
	// Har bir IP uchun vaqt oynasida request sonini sanaydi.
	//
	// Misol: 5 request / 60 soniya
	//   → 10:00:00 dan 10:01:00 gacha 5 ta request ruxsat
	//   → 6-chi request → 429 Too Many Requests
	limiter := NewRateLimiter(config.RateLimit)

	// ============================================
	// 4. Middleware zanjirini tuzamiz
	// ============================================
	// Go'da middleware pattern:
	//
	//   Chain(handler, middleware1, middleware2, ...)
	//
	// Bu yerda:
	//   - handler: ProxyHandler (backend serverga yo'naltiradi)
	//   - middleware1: LoggingMiddleware (request logga yozadi)
	//   - middleware2: RateLimitMiddleware (limit tekshiradi)
	//
	// Zanjir tartibi:
	//
	//   Request → Logging → Rate Limiting → Proxy → Backend
	//
	// Express.js versiyasi:
	//
	//   app.use(loggingMiddleware)
	//   app.use(rateLimitMiddleware)
	//   app.use(proxyHandler)
	//
	proxyHandler := NewProxyHandler(lb)
	handler := Chain(
		proxyHandler,               // Asl handler — backend serverga yo'naltiradi
		LoggingMiddleware,          // 1-chi middleware — log yozadi
		RateLimitMiddleware(limiter), // 2-chi middleware — limit tekshiradi
	)

	// ============================================
	// 5. Serverni ishga tushiramiz
	// ============================================
	// http.ListenAndServe — HTTP serverni ishga tushiradi.
	//
	// Birinchi parametr: ":8080" — port (barcha interface'larda)
	// Ikkinchi parametr: handler — so'rovlar qayta ishlanadigan handler
	//
	// Server to'xtatish uchun: Ctrl+C (SIGINT)
	addr := fmt.Sprintf(":%d", config.Port)

	// Boshlang'ich ma'lumotlarni chiqaramiz
	fmt.Println("========================================")
	fmt.Println("  Go Reverse Proxy — Middleware bilan")
	fmt.Println("========================================")
	fmt.Printf("Server ishga tushdi: http://localhost:%d\n", config.Port)
	fmt.Println("Backend serverlar:")
	for i, upstream := range config.Upstreams {
		fmt.Printf("  %d. %s\n", i+1, upstream.URL)
	}
	if config.RateLimit.Enabled {
		fmt.Printf("Rate Limit: %d request / %d soniya\n",
			config.RateLimit.Rate, config.RateLimit.Window)
	} else {
		fmt.Println("Rate Limit: o'chirilgan")
	}
	fmt.Println("========================================")
	fmt.Println("Sinov uchun:")
	fmt.Printf("  curl http://localhost:%d\n", config.Port)
	fmt.Printf("  curl http://localhost:%d/api/users\n", config.Port)
	fmt.Println("========================================")

	// Serverni ishga tushiramiz
	// Agar port band bo'lsa yoki boshqa xatolik bo'lsa, logga yoziladi
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server xatolik: %v", err)
	}
}
