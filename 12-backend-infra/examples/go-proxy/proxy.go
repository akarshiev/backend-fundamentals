// ============================================
// Proxy — Reverse Proxy + Round Robin Load Balancer
// ============================================
//
// Bu faylda asl proxy server mantiqiy joylashgan.
// Python versiyasidagi main.py faylining Go versiyasi.
//
// Asosiy vazifalar:
//   1. Client so'rovini qabul qilish
//   2. Round Robin algoritmi bilan backend server tanlash
//   3. So'rovni backend serverga yo'naltirish
//   4. Javobni clientga qaytarish
//
// Arxitektura:
//
//	Client → Proxy (Round Robin) → Backend Server
//
// Python versiyasidan farqi:
//   - Go'da http.Reverse ishlatiladi (kutubxona)
//   - Concurrent xotira kirish uchun sync.Mutex ishlatiladi
//   - Goroutines orqali parallel requestlarni qayta ishlaydi
//

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

// ============================================
// Load Balancer — Round Robin algoritmi
// ============================================
//
// Round Robin — eng sodda load balancing algoritmi.
// Backend serverlarni ketma-ket taqsimlaydi.
//
// Misol:
//
//	UPSTREAMS = ["github.com", "example.com"]
//
//	Request 1 → github.com
//	Request 2 → example.com
//	Request 3 → github.com
//	Request 4 → example.com
//
// Afzalligi: Juda sodda, tez ishlaydi.
// Kamchiligi: Server kuchlarini hisobga olmaydi.
// Yechim: Weighted Round Robin ishlatish.
type LoadBalancer struct {
	mu          sync.Mutex          // Concurrent xotira kirish uchun kilid
	upstreams   []UpstreamServer   // Backend serverlar ro'yxati
	currentIndex int               // Hozirgi indeks
	proxies     []*httputil.ReverseProxy // Har bir backend uchun proxy
}

// NewLoadBalancer — yangi load balancer yaratadi.
//
// Bu funksiya:
//   1. Har bir upstream server uchun ReverseProxy yaratadi
//   2. Proxy'lar ro'yxatini saqlab qo'yadi
//   3. Boshlang'ich indeksni 0 ga o'rnatadi
func NewLoadBalancer(upstreams []UpstreamServer) *LoadBalancer {
	lb := &LoadBalancer{
		upstreams:   upstreams,
		currentIndex: 0,
		proxies:     make([]*httputil.ReverseProxy, len(upstreams)),
	}

	// Har bir upstream server uchun ReverseProxy yaratamiz
	for i, upstream := range upstreams {
		// URL parse qilamiz — "https://github.com" → *url.URL
		parsedURL, err := url.Parse(upstream.URL)
		if err != nil {
			// Agar URL noto'g'ri bo'lsa, logga yozamiz va davom ettiramiz
			log.Printf("⚠️ URL parse xatolik: %s — %v", upstream.URL, err)
			continue
		}

		// httputil.ReverseProxy — Go kutubxonasidagi reverse proxy
		// Bu ob'ekt so'rovni backend serverga yo'naltiradi
		lb.proxies[i] = httputil.NewSingleHostReverseProxy(parsedURL)
	}

	return lb
}

// NextServer — keyingi backend serverni tanlaydi (Round Robin).
//
// Bu funksiya mutex bilan himoyalangan — chunki bir vaqtda
// bir nechta goroutine (parallel request) chaqirishi mumkin.
//
// Agar mutex bo'lmasa, indeks bir vaqtda ikki goroutine tomonidan
// o'qilishi va o'zgartirilishi mumkin — bu Race Condition deyiladi.
//
// Misol:
//
//	lb.NextServer()  // goroutine 1: currentIndex = 0 → github.com
//	lb.NextServer()  // goroutine 2: currentIndex = 1 → example.com
//	lb.NextServer()  // goroutine 1: currentIndex = 0 → github.com (qayta)
func (lb *LoadBalancer) NextServer() (string, *httputil.ReverseProxy) {
	lb.mu.Lock()         // Kilidni olamiz
	defer lb.mu.Unlock() // Funksiya tugagach kilidni tashlaymiz

	// Hozirgi indeksdagi serverni olamiz
	upstream := lb.upstreams[lb.currentIndex]
	proxy := lb.proxies[lb.currentIndex]

	// Indeksnii keyingisiga o'tkazamiz
	// Modulo (%) — indeks oxiriga yetganda boshiga qaytadi
	// Masalan: 0, 1, 0, 1, 0, 1, ...
	lb.currentIndex = (lb.currentIndex + 1) % len(lb.upstreams)

	return upstream.URL, proxy
}

// ============================================
// Proxy Handler — Client so'rovlarini qayta ishlaydi
// ============================================
//
// Bu handler har bir client so'rovini qabul qiladi va:
//   1. Keyingi backend serverni tanlaydi (Round Robin)
//   2. So'rovni backend serverga yo'naltiradi
//   3. Javobni clientga qaytaradi
//
// Go'da http.Handler interfeysi ishlatiladi:
//
//	type Handler interface {
//	    ServeHTTP(ResponseWriter, *Request)
//	}
//
// Bizning ProxyHandler ham shu interfeysi bajaradi.

// ProxyHandler — proxy so'rovlarini qayta ishlaydi.
type ProxyHandler struct {
	lb *LoadBalancer // Load balancer — backend serverlarni tanlaydi
}

// NewProxyHandler — yangi proxy handler yaratadi.
func NewProxyHandler(lb *LoadBalancer) *ProxyHandler {
	return &ProxyHandler{lb: lb}
}

// ServeHTTP — HTTP so'rovlarini qayta ishlaydi.
//
// Bu metodd http.Handler interfeysini bajaradi.
// Har bir client so'rovida chaqiriladi.
//
// Jarayon:
//   1. Keyingi backend serverni tanlash (Round Robin)
//   2. Backend server URL'sini logga chiqarish
//   3. httputil.ReverseProxy orqali so'rovni yo'naltirish
//
// Go'da ReverseProxy avtomatik ravishda:
//   - Client so'rovini backend serverga uzatadi
//   - Backend server javobini clientga qaytaradi
//   - Xatoliklarni (timeout, connection refused) hal qiladi
func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Keyingi backend serverni tanlaymiz (Round Robin)
	targetURL, proxy := ph.lb.NextServer()

	// Agar proxy yaratilmagan bo'lsa (URL parse xatolik), 500 qaytaramiz
	if proxy == nil {
		http.Error(w, "Backend server mavjud emas", http.StatusInternalServerError)
		return
	}

	// Logga yozamiz — qaysi serverga jo'natilayotganini ko'rish uchun
	// Misol: [Proxy] /api/users → https://github.com
	log.Printf("[Proxy] %s → %s", r.URL.Path, targetURL)

	// ReverseProxy orqali so'rovni backend serverga yo'naltiramiz
	// Go kutubxonasi avtomatik ravishda barcha ishni bajaradi:
	//   - Client so'rovini klonlaydi
	//   - Backend serverga yuboradi
	//   - Javobni clientga qaytaradi
	proxy.ServeHTTP(w, r)
}
