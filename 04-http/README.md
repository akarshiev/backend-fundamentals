# HTTP

HTTP (HyperText Transfer Protocol) -- web'ning asosiy protokoli.

---

## Mavzular

- HTTP Request/Response
- Status Codes
- Headers
- Compression (Gzip, Brotli)

---

## Nazariya

HTTP request/response modeli:

```text
Browser
   |
   V
HTTP Request
   |
   V
Server
   |
   V
HTTP Response
```

---

## Amaliyot

### curl bilan test

```bash
curl -v http://example.com
```

### Headerlarni ko'rish

```bash
curl -I http://example.com
```

---

## Xulosa

- HTTP request/response modelida ishlaydi
- Status codes javob holatini ko'rsatadi
- Headers qo'shimcha ma'lumot beradi
