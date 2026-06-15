# Reverse Proxy Vazifalari

Reverse Proxy bir necha muhim vazifani bajaradi.

---

## 1. Load Balancing

Bir nechta backend serverga trafikni taqsimlaydi.

```text
Client
   ↓
Nginx
 ├── App1
 ├── App2
 └── App3
```

### Nginx misoli

```nginx
upstream backend {
    server 127.0.0.1:8080;
    server 127.0.0.1:8081;
    server 127.0.0.1:8082;
}

server {
    location / {
        proxy_pass http://backend;
    }
}
```

---

## 2. Security

Backend IP'larni yashiradi.

```text
Internet:
- Nginx ni ko'radi (masalan, 1.2.3.4)
- Backend IP'larni ko'rmaydi

Backend serverlar:
- Faqat ichki tarmoqdan kiriladi
- Internetdan to'g'ridan-to'g'ri kirish mumkin emas
```

---

## 3. Caching

Static fayllarni cache qiladi. Backendga kamroq so'rov tushadi.

```text
Client → logo.png → Nginx → Backend (birinchi marta)
Client → logo.png → Nginx → Cache dan (keyingi marta)
```

### Nginx cache misoli

```nginx
proxy_cache_path /tmp/nginx levels=1:2 keys_zone=my_cache:10m;

server {
    location / {
        proxy_cache my_cache;
        proxy_cache_valid 200 10m;
        proxy_pass http://backend;
    }
}
```

---

## 4. Compression

Katta fayllarni siqadi. Trafik tezligini oshiradi.

```text
Backend: 10MB JSON qaytardi
Proxy: gzip, brotli bilan siqdi
Client: 2MB oldi (80% kam)
```

### Nginx gzip misoli

```nginx
http {
    gzip on;
    gzip_types text/plain application/json text/css;
    gzip_min_length 1000;
}
```

---

## 5. Rate Limiting

Ma'lum miqdordandan ortiq so'rov kelganda bloklaydi.

```text
100 req/min dan oshsa → 429 Too Many Requests
```

### Nginx rate limiting misoli

```nginx
http {
    limit_req_zone $binary_remote_addr zone=one:10m rate=10r/s;

    server {
        location /api/ {
            limit_req zone=one burst=20 nodelay;
            proxy_pass http://backend;
        }
    }
}
```

---

## 6. SSL Termination

HTTPS juda qimmat — har request TLS encryption/decryption qiladi.

Shuning uchun TLS faqat proxy'da ishlaydi:

```text
Client
   ↓ HTTPS (TLS)
Nginx
   ↓ HTTP (oddiy)
Backend
```

### Foydasi

| Oldingi holat | SSL Termination bilan |
|---------------|----------------------|
| Har bir server TLS ishlatadi | Faqat Nginx TLS ishlatadi |
| Har serverda sertifikat kerak | Faqat Nginx'da sertifikat |
| Katta resurs sarflanadi | Backend resurs tejaydi |

### SSL Termination qiluvchilar

| Vosita | Ishlatilishi |
|--------|-------------|
| Cloudflare | CDN + SSL |
| AWS ALB | AWS Load Balancer |
| Nginx | Eng ko'p ishlatiladigan |

---

## Backend nima qilmasligi kerak?

Backend faqat **Business Logic** yozishi kerak.

```text
Backend zimmasiga tashlamang:

❌ TLS
❌ Compression
❌ Caching
❌ Rate Limiting
❌ Load Balancing
❌ Static Files
❌ DDoS Protection

Bularni reverse proxy ga topshiring ✅
```

---

## Xulosa

| Vazifa | Tushuntirish |
|--------|-------------|
| Load Balancing | Trafikni taqsimlash |
| Security | Backend IP'larni yashirish |
| Caching | Static fayllarni saqlash |
| Compression | Trafikni siqish |
| Rate Limiting | So'rovlar sonini cheklash |
| SSL Termination | TLS ni proxy'da boshqarish |
