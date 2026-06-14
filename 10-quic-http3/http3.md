# HTTP/3

## HTTP/3 nima?

HTTP/3 — QUIC ustida ishlaydigan yangi HTTP versiyasi.

TCP o'rniga **QUIC** ishlatiladi.

---

## HTTP/3 afzalliklari

| Xususiyat | Tushuntirish |
|-----------|--------------|
| QUIC | UDP ustida, tezroq |
| Built-in TLS 1.3 | Alohida TLS handshake kerak emas |
| Stream-level HOL | Bir stream to'xtasa, boshqalar ishlaydi |
| Connection Migration | WiFi → Mobile o'tganda uzilmaydi |
| 0-RTT | Birinchi request'da ham tez ulanish |

---

## Oddiy HTTPS vs QUIC

### Oddiy HTTPS (HTTP/1.1 yoki HTTP/2)

```
1. TCP Handshake     — SYN, SYN-ACK, ACK
2. TLS Handshake
3. HTTP Request
```

### QUIC (HTTP/3)

```
1. QUIC Handshake (TLS 1.3 bilan)
2. HTTP Request
```

Kamroq RTT. Tezroq ulanish.

---

## HTTP/3 ishlatayotgan xizmatlar

- Google
- YouTube
- Cloudflare
- Facebook
- Vercel
