# Browser Architecture

Browser qanday ishlashini tushunish.

---

## Nazariya

Browser quyidagi qadamlarni bajardi:

```text
1. URL parse
2. DNS lookup
3. TCP connection
4. TLS handshake (HTTPS)
5. HTTP request
6. HTTP response
7. Parse HTML/CSS/JS
8. Render
```

---

## Diagram

```text
User types: https://google.com

   |
   V

URL Parser
- Protocol: https
- Host: google.com
- Port: 443

   |
   V

DNS Resolver
- google.com -> 142.250.190.78

   |
   V

TCP Connection
- Connect to 142.250.190.78:443

   |
   V

TLS Handshake
- Encrypt connection

   |
   V

HTTP Request
- GET / HTTP/1.1
- Host: google.com

   |
   V

HTTP Response
- 200 OK
- HTML body

   |
   V

Parser
- HTML -> DOM tree
- CSS -> CSSOM tree
- JS -> Execute

   |
   V

Renderer
- Layout
- Paint
- Composite
```

---

## Amaliyot

### Browser DevTools

```text
1. Chrome oching
2. F12 bosing
3. Network tab
4. Sahifani yangilang
5. Har bir request/response ni ko'ring
```

### curl bilan test

```bash
curl -v http://example.com
```

Headerlarni ko'ring.

---

## Xulosa

- Browser = TCP + TLS + HTTP + Parser + Renderer
- Har bir qadamni CLI bilan sinab ko'rish mumkin
- Backend developer uchun TCP/HTTP qismi eng muhim
