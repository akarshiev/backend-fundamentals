# 12 - Backend Infratuzilmasi (Backend Infrastructure)

Proxy, Redis, Load Balancing va zamonaviy backend arxitekturasi.

---

## Mavzular

### Proxy
- [Proxy asoslari](proxy.md) -- Forward va Reverse Proxy
- [Reverse Proxy Vazifalari](reverse-proxy-features.md) -- Load Balancing, Caching, Compression, Rate Limiting, SSL Termination

### Redis
- [Redis](redis.md) -- In-Memory Key-Value Database, Cache, Session, Rate Limiting, Queue, Pub/Sub

### Rate Limiting
- [Rate Limiting & Throttling](rate-limiting.md) -- Noisy Neighbor, Token Bucket, Leaky Bucket, Distributed Rate Limiting

### Load Balancing
- [Load Balancing](load-balancing.md) -- Round Robin, Least Connections, IP Hashing, Weighted, L4/L7, Mirroring

### Arxitektura
- [Zamonaviy Backend Arxitekturasi](modern-architecture.md) -- DNS, TCP, TLS, HTTP, Reverse Proxy, Load Balancer, Application, Redis, Database

### Resilience
- [Circuit Breaker](circuit-breaker.md) -- Cascading Failure, CLOSED/OPEN/HALF-OPEN, Fallback, Node.js amaliyot

---

## Qisqacha

| Mavzu | Tushuntirish |
|-------|-------------|
| Forward Proxy | Clientni yashiradi (VPN) |
| Reverse Proxy | Serverni yashiradi (Nginx) |
| SSL Termination | TLS faqat proxy'da |
| Redis | RAM da tez database (~100 ns) |
| Rate Limiting | So'rovlar sonini cheklash (STOP) |
| Throttling | So'rovlar tezligini sekinlashtirish (SLOW DOWN) |
| Token Bucket | Burst ga ruxsat, eng mashhur |
| Leaky Bucket | Tekis chiqish, Nginx ishlatadi |
| Round Robin | Ketma-ket taqsimlash |
| Least Connections | Eng kam connection ga |
| L4 Load Balancing | TCP/UDP (tez) |
| L7 Load Balancing | HTTP (aqlli) |
| Mirroring | Traffic ni nusxalash |
| Circuit Breaker | Servis yiqilishini to'xtatish |
| Fallback | Zaxira servis ishlatish |
| Fail Fast | Kutmasdan darhol javob |

---

## O'rganish tartibi

```text
Proxy
-> Forward Proxy (VPN)
-> Reverse Proxy (Nginx)
-> Reverse Proxy Vazifalari
-> SSL Termination
-> Load Balancing
-> Round Robin
-> Least Connections
-> IP Hashing
-> Weighted
-> L4 / L7
-> Mirroring
-> Redis
-> Cache, Session, Queue, Pub/Sub
-> Rate Limiting
-> Fixed Window, Token Bucket, Leaky Bucket
-> Distributed Rate Limiting (Redis + Lua)
-> Zamonaviy Backend Arxitekturasi
-> Circuit Breaker
-> Cascading Failure
-> CLOSED / OPEN / HALF-OPEN
-> Fallback, Fail Fast
```

---

## How to use

1. Proxy asoslaridan boshlang
2. Reverse Proxy vazifalarini o'qing
3. Load Balancing algoritmlarini tushuning
4. Redis ni o'rganing
5. Rate Limiting va Throttling ni tushuning
6. Circuit Breaker ni o'rganing
7. Zamonaviy arxitekturani tahlil qiling
