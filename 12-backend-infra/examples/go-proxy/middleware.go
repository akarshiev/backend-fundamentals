// ============================================
// Middleware — Request qayta ishlash zanjiri
// ============================================
//
// Middleware = request serverga yetib kelishdan oldin
// o'tadigan "filtrlar" ketma-ketligi.
//
// Har bir middleware requestni qayta ishlaydi va keyingisiga uzatadi.
//
// Bizning zanjirimiz:
//
//   Request → Logging → Rate Limiting → Proxy Handler
//              ↓            ↓                ↓
//          Log yozadi   Limit tekshiradi   Backend ga yo'naltiradi
//
// Bu model Nginx, Express.js, va boshqa frameworklarda ham ishlatiladi.

package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ============================================
// Rate Limiter — Fixed Window Counter algoritmi
// ============================================
//
// Har bir IP uchun vaqt oynasida nechta request bo'lganini sanaydi.
//
// Algoritm:
//   1. IP manzilini tekshir
//   2. Agar yangi IP bo'lsa, 1 dan boshla
//   3. Agar vaqt oynasi tugagan bo'lsa, qayta boshla
//   4. Agar limitdan oshsa, 429 Too Many Requests qaytar
//
// Bu algoritmning kamchiligi — Burst Problem:
//
//   10:00:59 → 5 request (limit bo'yicha OK)
//   10:01:01 → 5 request (yangi oyna, yana OK)
//   2 soniyada 10 request o'tib ketadi!
//
// Yechim: Token Bucket yoki Leaky Bucket ishlatish.
// Biz hozir sodda Fixed Window ishlatamiz.

// clientInfo — bitta IP manzili uchun ma'lumotlar.
// Har bir clientning request soni va boshlanish vaqti saqlanadi.
type clientInfo struct {
	Count     int       // Vaqt oynasidagi request soni
	StartTime time.Time // Vaqt oynasining boshlanish vaqti
}

// RateLimiter — rate limiting mantiqini saqlaydi.
// Sync.Map ishlatamiz — chunki Go'da concurrent xotira kirish muhim.
// Oddiy map ishlatsak, race condition bo'ladi.
type RateLimiter struct {
	mu      sync.RWMutex          // O'qish/yozish kilidi (concurrency xavfsizligi)
	clients map[string]*clientInfo // IP → client info xaritasi
	config  RateLimitConfig        // Limit sozlamalari
}

// NewRateLimiter — yangi rate limiter yaratadi.
func NewRateLimiter(cfg RateLimitConfig) *RateLimiter {
	return &RateLimiter{
		clients: make(map[string]*clientInfo), // Bo'sh xarita yaratildi
		config:  cfg,
	}
}

// Allow — so'rovga ruxsat berilganini tekshiradi.
//
// Qaytaradi:
//   true  — so'rov o'tadi ✅
//   false — so'rov bloklanadi (429) ⛔
//
// Bu funksiya har bir request uchun chaqiriladi.
// Bir nechta goroutine (parallel request) bir vaqtda ishlashi mumkin,
// shuning uchun mutex ishlatamiz.
func (rl *RateLimiter) Allow(ip string) bool {
	// Agar rate limiting o'chirilgan bo'lsa, hamma so'rovlar o'tadi
	if !rl.config.Enabled {
		return true
	}

	now := time.Now()

	// Write lock — chunki biz xotirani o'zgartiramiz
	rl.mu.Lock()
	defer rl.mu.Unlock() // Funksiya tugagach lock ochiladi

	// IP mavjudmi tekshiramiz
	client, exists := rl.clients[ip]

	if !exists {
		// Yangi client — 1-dan boshlaymiz
		rl.clients[ip] = &clientInfo{
			Count:     1,
			StartTime: now,
		}
		return true // ✅ Birinchi so'rov — ruxsat
	}

	// Vaqt oynasi tugaganini tekshiramiz
	elapsed := now.Sub(client.StartTime).Seconds()
	if elapsed > float64(rl.config.Window) {
		// Yangi oyna boshlandi — hisoblagichni qayta boshlaymiz
		client.Count = 1
		client.StartTime = now
		return true // ✅ Yangi oyna — ruxsat
	}

	// Limitdan oshganini tekshiramiz
	if client.Count >= rl.config.Rate {
		return false // ⛔ Limit tugadi — bloklandi
	}

	// Hisoblagichni oshiramiz
	client.Count++
	return true // ✅ Ruxsat berildi
}

// ============================================
// Middleware — Request zanjiri
// ============================================

// Middleware — bir funksiyani boshqasiga o'radi.
// Bu Go'da middleware patternining standart usuli.
//
// Misol:
//
//	loggingMiddleware(rateLimitMiddleware(finalHandler))
//
// Natija:
//
//	Request → Logging → Rate Limiting → finalHandler
//
type Middleware func(http.Handler) http.Handler

// Chain — bir nechta middleware'ni birlashtiradi.
// Middleware'lar teskari tartibda qo'llaniladi (matematik compose).
//
// Misol:
//
//	Chain(handler, m1, m2, m3)
//	Tekshiriladi: m1(m2(m3(handler)))
//
// Bu Express.js'dagi app.use() ga o'xshaydi.
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	// Teskari tartibda qo'llanadi
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// LoggingMiddleware — har bir requestni logga yozadi.
//
// Nima qiladi:
//   1. Request kelganini logga yozadi
//   2. Keyingi middleware/handler ga uzatadi
//   3. Javob status kodini logga yozadi
//
// Bu middleware barcha middleware'lardan birinchi bo'lib ishlaydi.
//
// Log chiqishi:
//
//	[2024-01-15 10:30:45] GET /api/users → 200 (15ms)
//	[2024-01-15 10:30:46] GET /api/posts → 429 (1ms)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Request boshlanish vaqti

		// ResponseWriter ni wrapper qilamiz — status kodni olish uchun.
		// Go'da http.ResponseWriter orqali status kodni to'g'ridan-to'g'ri
		// olish mumkin emas, shuning uchun wrapper ishlatamiz.
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Keyingi middleware/handler ni chaqiramiz
		next.ServeHTTP(rw, r)

		// Request tugadi — logga yozamiz
		duration := time.Since(start)

		// Konsolga chiroyli formatda chiqaramiz
		// Misol: [2024-01-15 10:30:45] GET /api/users → 200 (15ms)
		fmt.Printf(
			"[%s] %s %s → %d (%s)\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method,
			r.URL.Path,
			rw.statusCode,
			duration.Round(time.Millisecond),
		)
	})
}

// RateLimitMiddleware — rate limiting qo'llaydi.
//
// Har bir request uchun:
//   1. Client IP manzilini aniqlaydi
//   2. RateLimiter orqali limitni tekshiradi
//   3. Agar limit oshsa — 429 Too Many Requests qaytaradi
//   4. Agar limit ichida bo'lsa — keyingi handler ga uzatadi
//
// IP manzilini aniqlash:
//   - X-Forwarded-For header — proxy orqali kelgan requestlar uchun
//   - X-Real-IP header — ba'zi proxy'lar ishlatadi
//   - RemoteAddr — to'g'ridan-to'g'ri client manzili
//
// Bu header'lar Reverse Proxy (Nginx, Cloudflare) tomonidan qo'shiladi.
func RateLimitMiddleware(limiter *RateLimiter) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Client IP manzilini aniqlaymiz
			ip := getClientIP(r)

			// Rate limitni tekshiramiz
			if !limiter.Allow(ip) {
				// ⛔ Limit oshdi — 429 Too Many Requests
				// HTTP/1.1 RFC 6585 ga binoan 429 status kodi ishlatiladi
				http.Error(
					w,
					"429 Too Many Requests",
					http.StatusTooManyRequests,
				)
				return // Request to'xtatildi — keyingi handler ga o'tmaydi
			}

			// ✅ Limit ichida — keyingi handler ga uzatamiz
			next.ServeHTTP(w, r)
		})
	}
}

// getClientIP — client IP manzilini aniqlaydi.
//
// Kiritish manbalari (ustuvorlik tartibida):
//  1. X-Forwarded-For — proxy chain'dagi manzillar
//  2. X-Real-IP — oddiy proxy manzili
//  3. RemoteAddr — to'g'ridan-to'g'ri TCP manzili
//
// Misol:
//
//	X-Forwarded-For: 203.0.113.195, 70.41.3.18, 120.0.0.1
//	Bu yerda birinchi IP — asl client (203.0.113.195)
func getClientIP(r *http.Request) string {
	// X-Forwarded-For header ni tekshiramiz
	// Bu header proxy orqali kelgan requestlarda bo'ladi
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		// Komma bilan ajratilgan IP lar — birinchisi asl client
		// Masalan: "203.0.113.195, 70.41.3.18"
		// Birinchi IP ni olamiz: "203.0.113.195"
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' {
				return xff[:i]
			}
		}
		return xff
	}

	// X-Real-IP header ni tekshiramiz
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// RemoteAddr dan olamiz — "192.168.1.1:54321" formatda
	// Portni ajratib tashlaymiz
	addr := r.RemoteAddr
	for i := len(addr) - 1; i >= 0; i-- {
		if addr[i] == ':' {
			return addr[:i]
		}
	}
	return addr
}

// ============================================
// ResponseWriter Wrapper
// ============================================
//
// Go'da http.ResponseWriter interfeysi WriteHeader() va Write() metodlarini
// ta'minlaydi, lekin status kodni saqlamaydi. Biz o'z wrapper yaratamiz.
//
// Bu pattern Go web dasturlashda juda mashhur.

// responseWriter — http.ResponseWriter ni wrapper qiladi.
// Status kodni saqlab qoladi.
type responseWriter struct {
	http.ResponseWriter              // Asl ResponseWriter (embedded)
	statusCode         int           // HTTP status kodi (200, 404, 500, ...)
}

// WriteHeader — status kodni saqlab, asl ResponseWriter ga uzatadi.
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
