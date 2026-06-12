# HTTP Status Codes

Server javobining holatini ko'rsatadi.

---

## Nazariya

Status codes 5 ta kategoriyaga bo'linadi.

---

## 1xx -- Informational

Kam ishlatiladi.

```text
100 Continue
101 Switching Protocols
```

---

## 2xx -- Success

```text
200 OK
201 Created
204 No Content
```

---

## 3xx -- Redirect

```text
301 Moved Permanently
302 Found
307 Temporary Redirect
```

---

## 4xx -- Client Error

```text
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
409 Conflict
429 Too Many Requests
```

---

## 5xx -- Server Error

```text
500 Internal Server Error
502 Bad Gateway
503 Service Unavailable
```

---

## Diagram

```text
1xx: Info
2xx: OK
3xx: Redirect
4xx: Client xatosi
5xx: Server xatosi
```

---

## Best Practice

Yomon:

```http
HTTP 200 OK
{
  "error": "User not found"
}
```

To'g'ri:

```http
HTTP 404 Not Found
{
  "message": "User not found"
}
```

Status code va body bir xil ma'noda bo'lishi kerak.

---

## Amaliyot

### Status code ko'rish

```bash
curl -I http://example.com
```

### Redirect tekshirish

```bash
curl -L http://example.com
```

---

## Xulosa

- 2xx = muvaffaqiyat
- 4xx = client xatosi
- 5xx = server xatosi
- Status code body bilan mos bo'lishi kerak
