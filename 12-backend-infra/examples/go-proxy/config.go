// ============================================
// Config — Proxy konfiguratsiyasi
// ============================================
//
// Bu faylda proxy serverning barcha sozlamalari saqlanadi.
// Python versiyasidagi config.py faylining Go versiyasi.
//
// Upstreams — backend serverlar ro'yxati (Round Robin tartibida taqsimlanadi)
// RateLimit — har bir IP uchun request limiti
// Port — proxy server qaysi portda ishlaydi
//

package main

// UpstreamServer — bitta backend serverni ifodalovchi struct.
// Har bir serverning URL manzili va og'irligi (weight) bor.
type UpstreamServer struct {
	URL    string // Masalan: "https://github.com"
	Weight int    // Og'irlik — qancha og'ir bo'lsa, shuncha ko'p request oladi
}

// RateLimitConfig — rate limiting sozlamalari.
// Fixed Window Counter algoritmini ishlatadi.
type RateLimitConfig struct {
	Enabled bool    // Rate limiting yoqilganmi?
	Rate    int     // Maksimal request soni (masalan: 5)
	Window  int     // Vaqt oynasi sekundlarda (masalan: 60 = 1 daqiqa)
}

// ProxyConfig — umumiy proxy konfiguratsiyasi.
// Barcha sozlamalar shu struct ichida yig'iladi.
type ProxyConfig struct {
	Port      int              // Proxy server porti (masalan: 8080)
	Upstreams []UpstreamServer // Backend serverlar ro'yxati
	RateLimit RateLimitConfig  // Rate limiting sozlamalari
}

// DefaultConfig — standart konfiguratsiya.
// Python versiyasidagi config.py faylidagi qiymatlar bilan bir xil.
//
// Misol:
//
//	Upstreams = [
//	  "https://github.com",   // Server 1
//	  "https://example.com"   // Server 2
//	]
//
// Rate Limit = 5 request / 60 soniya
func DefaultConfig() *ProxyConfig {
	return &ProxyConfig{
		Port: 8080,
		Upstreams: []UpstreamServer{
			{URL: "https://github.com", Weight: 1}, // Server 1
			{URL: "https://example.com", Weight: 1}, // Server 2
		},
		RateLimit: RateLimitConfig{
			Enabled: true,
			Rate:    5,    // 5 ta request
			Window:  60,   // 60 soniya (1 daqiqa) ichida
		},
	}
}
