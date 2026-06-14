# Reverse Proxy

Client va Application o'rtasida turib, xavfsizlik va optimizatsiya ta'minlaydi.

---

## Tushuntirish

```text
Normal:
Client → Application Server

Reverse Proxy:
Client → Nginx / Cloudflare → Application Server
```

---

## Reverse Proxy vazifalari

| Vazifa | Tushuntirish |
|--------|-------------|
| TLS | HTTPS SSL/TLS boshqaruvi |
| Cache | Static fayllarni saqlash |
| Compression | Gzip, Brotli |
| Load Balancing | Trafikni taqsimlash |
| DDoS Protection | Hujumlarni to'xtatish |
| Rate Limiting | So'rovlar cheklash |
| WAF | Web Application Firewall |

---

## Nginx Reverse Proxy

### Asosiy sozlash

```nginx
# /etc/nginx/nginx.conf

http {
    # Upstream serverlar
    upstream backend {
        server 127.0.0.1:8080;
        server 127.0.0.1:8081;
    }
    
    server {
        listen 443 ssl;
        server_name example.com;
        
        # SSL sozlash
        ssl_certificate /path/to/cert.pem;
        ssl_certificate_key /path/to/key.pem;
        
        # Reverse Proxy
        location / {
            proxy_pass http://backend;
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }
        
        # Static fayllar
        location /static/ {
            alias /var/www/static/;
            expires 30d;
        }
    }
}
```

### Rate Limiting

```nginx
http {
    limit_req_zone $binary_remote_addr zone=one:10m rate=10r/s;
    
    server {
        location / {
            limit_req zone=one burst=20 nodelay;
            proxy_pass http://backend;
        }
    }
}
```

### DDoS Himoya

```nginx
http {
    # Connection limit
    limit_conn_zone $binary_remote_addr zone=conn:10m;
    
    server {
        location / {
            limit_conn conn 10;
            limit_req zone=one burst=20;
            proxy_pass http://backend;
        }
    }
}
```

---

## Cloudflare

### Foydalar

| Foyda | Tushuntirish |
|-------|-------------|
| DDoS Himoya | 30+ Tbps quvvat |
| Cache | Global CDN |
| WAF | Web Application Firewall |
| TLS | Free SSL/TLS |
| Analytics | Trafik tahlili |

### Sozlash

```text
1. Cloudflare ga ro'yxatdan o'ting
2. Domain qo'shing
3. DNS serverlarini o'zgartiring
4. SSL/TLS rejimini tanlang (Full Strict)
5. Cache Level: Standard
6. Browser Integrity Check: ON
```

---

## Amaliyot

### Nginx o'rnatish (Ubuntu)

```bash
sudo apt update
sudo apt install nginx
sudo systemctl start nginx
sudo systemctl enable nginx
```

### Config tekshirish

```bash
sudo nginx -t
sudo systemctl reload nginx
```

### Reverse Proxy test

```bash
curl -I https://example.com
curl -H "X-Forwarded-For: 1.2.3.4" http://localhost
```

---

## Xulosa

- Reverse Proxy xavfsizlik va optimizatsiya ta'minlaydi
- Nginx, Cloudflare, HAProxy eng mashhur
- TLS, Cache, Rate Limiting, DDoS himoya vazifalari bor
