# HTTP From Scratch — Go bilan HTTP ni 0 dan tushunish

HTTP ni kutubxona ishlatmasdan, oddiy TCP socket orqali qanday ishlashini ko'ramiz.

---

## Nimani o'rganamiz?

```text
1. TCP connection — server bilan ulanish
2. Raw HTTP request — qo'lda so'rov yozish
3. Raw HTTP response — javobni qo'lda parse qilish
4. Headers — header'lar qanday ishlaydi
5. Body — body streaming va chunked encoding
6. HTTP Methods — GET, POST, PUT, DELETE
```

---

## HTTP aslida nima?

HTTP = **HyperText Transfer Protocol**

Aslida bu oddiy **matn protokoli**. Siz serverga matn yuborasiz, server matn qaytaradi.

### HTTP Request formati

```http
GET /api/users HTTP/1.1
Host: example.com
User-Agent: curl/7.88.1
Accept: */*

```

Bu oddiy **matn**! Qo'lda yozish mumkin. Shunday qilamiz.

### HTTP Response formati

```http
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 42

{"name": "Otabek", "role": "Backend Developer"}
```

Yana bir **matn**. Hamma narsa matn formatida.

---

## 1. TCP Connection — Server bilan ulanish

Har bir HTTP request **TCP connection** dan boshlanadi.

### Jarayon

```text
1. DNS Resolution
   example.com → 93.184.216.34

2. TCP 3-Way Handshake
   Client → SYN     → Server
   Client ← SYN-ACK ← Server
   Client → ACK     → Server

3. TCP Connection Established ✅

4. HTTP Request yuborish
   Client → "GET / HTTP/1.1\r\n..." → Server

5. HTTP Response olish
   Client ← "HTTP/1.1 200 OK\r\n..." ← Server

6. TCP Connection yopish (yoki keep-alive)
```

### Go'da TCP connection

```go
// net package — TCP socket ochamiz
conn, err := net.Dial("tcp", "example.com:80")
// Bu yerda 3-Way Handshake avtomatik bajariladi
```

---

## 2. Raw HTTP Request — Qo'lda so'rov yozish

HTTP request **4 qismdan** iborat:

```text
┌─────────────────────────────────┐
│ 1. Request Line                  │
│    GET /path HTTP/1.1            │
├─────────────────────────────────┤
│ 2. Headers                       │
│    Host: example.com             │
│    Content-Type: text/html       │
├─────────────────────────────────┤
│ 3. Blank Line (bo'sh qator)      │
│                                  │
├─────────────────────────────────┤
│ 4. Body (ixtiyoriy)              │
│    {"name": "Otabek"}            │
└─────────────────────────────────┘
```

### Muhim qoidalar

```text
- Har bir qator \r\n (CR+LF) bilan tugaydi
- Headers va Body orasida BO'SH QATOR kerak
- Request Line: METHOD PATH VERSION
```

### Go'da qo'lda request yozish

```go
// Format string bilan request yozamiz
request := "GET / HTTP/1.1\r\n" +
    "Host: example.com\r\n" +
    "Accept: */*\r\n" +
    "\r\n"  // Bo'sh qator — headers tugadi

// TCP socket orqali yuboramiz
conn.Write([]byte(request))
```

---

## 3. Raw HTTP Response — Javobni parse qilish

Response ham **4 qismdan** iborat:

```text
┌─────────────────────────────────┐
│ 1. Status Line                   │
│    HTTP/1.1 200 OK               │
├─────────────────────────────────┤
│ 2. Headers                       │
│    Content-Type: text/html       │
│    Content-Length: 1234           │
├─────────────────────────────────┤
│ 3. Blank Line (bo'sh qator)      │
│                                  │
├─────────────────────────────────┤
│ 4. Body                          │
│    <html>...</html>              │
└─────────────────────────────────┘
```

### Status Code turlari

```text
1xx — Informational (100 Continue)
2xx — Muvaffaqiyat (200 OK, 201 Created, 204 No Content)
3xx — Yo'naltirish (301 Moved, 304 Not Modified)
4xx — Client xatosi (400 Bad Request, 404 Not Found, 429 Too Many)
5xx — Server xatosi (500 Internal Error, 502 Bad Gateway, 503 Unavailable)
```

---

## 4. Headers — Header'lar qanday ishlaydi

Header'lar **key: value** formatida.

### Muhim header'lar

| Header | vazifa | Misol |
|--------|--------|-------|
| Host | Qaysi domain | Host: example.com |
| Content-Type | Body turi | Content-Type: application/json |
| Content-Length | Body uzunligi | Content-Length: 42 |
| Accept | Qanday format qabul qiladi | Accept: application/json |
| Authorization | Auth token | Authorization: Bearer xxx |
| Connection | Connection turu | Connection: keep-alive |
| Transfer-Encoding | Body uzatish usuli | Transfer-Encoding: chunked |

### Header parse qilish

```text
Content-Type: application/json; charset=utf-8
     ↑         ↑                    ↑
   Name      Value              Parameters
```

---

## 5. Body — Body streaming va Chunked Encoding

### Content-Length bilan

```text
Content-Length: 42

{"name": "Otabek"}
```

Server biladi — body 42 byte. Shuncha o'qiydi.

### Chunked Transfer Encoding

Katta fayllar uchun body uzunligi oldindan noma'lum bo'lishi mumkin.

```text
Transfer-Encoding: chunked

1a\r\n
{"name": "Otabek", "role":\r\n
13\r\n
 "Developer"}\r\n
0\r\n
\r\n
```

```text
- Har bir chunk: SIZE\r\nDATA\r\n
- Oxirgi chunk: 0\r\n (tugadi)
```

---

## 6. HTTP Methods

| Method | Maqsad | Idempotent | Body |
|--------|--------|------------|------|
| GET | Ma'lumot olish | ✅ | ❌ |
| POST | Yangi resurs yaratish | ❌ | ✅ |
| PUT | Resursni to'liq yangilash | ✅ | ✅ |
| DELETE | Resursni o'chirish | ✅ | ❌ |
| PATCH | Resurni qisman yangilash | ❌ | ✅ |
| HEAD | Faqat header'lar (body'siz) | ✅ | ❌ |

### Idempotent nima?

```text
GET /users/1   → Har safar bir xil natija
PUT /users/1   → Har safar bir xil natija (idempotent)
POST /users    → Har safar yangi user yaratadi (NOT idempotent)
```

---

## 7. Connection Management

### HTTP/1.0 — Har request uchun yangi connection

```text
Request 1 → TCP Connect → Response → TCP Close
Request 2 → TCP Connect → Response → TCP Close
```

### HTTP/1.1 — Keep-Alive (default)

```text
TCP Connect
Request 1 → Response
Request 2 → Response
Request 3 → Response
TCP Close
```

### HTTP/2 — Multiplexing

```text
TCP Connect (1 marta)
Request 1, 2, 3 → Bir vaqtda parallel
Response 2, 1, 3 → Tartibsiz qaytadi
```

---

## 8. Go'da HTTP — Ikki usul

### Usul 1: net/http kutubxonasi (oddiy)

```go
resp, err := http.Get("https://example.com")
// Kutubxona avtomatik:
// - DNS resolution
// - TCP connection
// - TLS handshake (HTTPS bo'lsa)
// - HTTP request yozish
// - Response parse qilish
```

### Usul 2: net.Dial + qo'lda (o'rganish uchun)

```go
conn, _ := net.Dial("tcp", "example.com:80")
// Faqat TCP connection

fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: example.com\r\n\r\n")
// Qo'lda HTTP request yozamiz

scanner := bufio.NewScanner(conn)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
// Qo'lda response o'qiyamiz
```

---

## Amaliyot: Go misollari

```text
examples/http-from-scratch/
├── 01-tcp.go         — TCP connection from scratch
├── 02-raw-http.go    — Raw HTTP request/response
├── 03-headers.go     — Headers parsing
├── 04-body.go        — Body streaming va chunked
└── 05-methods.go     — HTTP Methods (GET, POST, PUT, DELETE)
```

Har bir faylni alohida ishga tushirishingiz mumkin:

```bash
cd examples/http-from-scratch
go run 01-tcp.go
go run 02-raw-http.go
go run 03-headers.go
go run 04-body.go
go run 05-methods.go
```

---

## Xulosa

```text
HTTP = TCP + Matn protokoli

Client:
  1. DNS → IP manzil
  2. TCP 3-Way Handshake
  3. "GET / HTTP/1.1\r\nHost: ...\r\n\r\n"  ← Qo'lda yoziladi
  4. Response o'qiladi

Server:
  1. TCP connection qabul qiladi
  2. Request matnini parse qiladi
  3. "HTTP/1.1 200 OK\r\nContent-Type: ...\r\n\r\nBody"  ← Qo'lda yoziladi
  4. Response yuboradi

Go:
  - net.Dial("tcp", ...)  → TCP connection
  - conn.Write([]byte(...))  → HTTP request yozish
  - bufio.Scanner  → Response o'qish
  - http.Get()  → Hammasi avtomatik
```
