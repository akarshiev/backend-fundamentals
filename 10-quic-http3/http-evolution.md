# HTTP Versiyalar Evolyutsiyasi

## Yulak

```
HTTP/1.1 → HTTP/2 → HTTP/3
```

---

## HTTP/1.1

| Xususiyat | Qiymat |
|-----------|--------|
| Transport | TCP |
| Multiplexing | Yo'q |
| HOL muammosi | Kuchli |

---

## HTTP/2

| Xususiyat | Qiymat |
|-----------|--------|
| Transport | TCP |
| Multiplexing | Ha |
| HOL muammosi | Kamaygan, lekin TCP sabab qolgan |

---

## HTTP/3

| Xususiyat | Qiymat |
|-----------|--------|
| Transport | QUIC (UDP ustida) |
| TLS | Built-in TLS 1.3 |
| Multiplexing | Stream-level |
| HOL muammosi | Yo'q (stream darajasida) |
| Connection Migration | Ha |

---

## Bugungi Internet

Google, YouTube, Cloudflare, Facebook, Vercel — hammasi **HTTP/3 va QUIC** ishlatmoqda.
