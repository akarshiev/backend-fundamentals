# Proxy

O'rtakash — client va server o'rtasida turib, so'rovlar yo'naltiradi.

---

## Proxy nima?

```text
Oddiy holat:

Client
   ↓
Server

Proxy bilan:

Client
   ↓
Proxy
   ↓
Server

Client server bilan emas, proxy bilan gaplashadi.
```

---

## Forward Proxy

Forward Proxy client tarafida turadi. Clientni serverdan yashiradi.

```text
Laptop
   ↓
Forward Proxy
   ↓
Internet
```

### VPN

VPN aslida Forward Proxy'ga o'xshaydi.

```text
Siz
 ↓
VPN
 ↓
Google

Google sizning IP emas, VPN IP ni ko'radi.
```

### Corporate Network

Katta kompaniyalarda ishlatiladi:

```text
Employee
   ↓
Proxy
   ↓
Internet
```

Nima uchun?

| Vazifa | Tushuntirish |
|--------|-------------|
| Monitoring | Qaysi saytlarga kirilayotganini kuzatish |
| Logging | Barcha so'rovlar loglanadi |
| Content Filtering | Ruxsatsiz saytlarni bloklash |
| Security | Xavfli saytlardan himoya |

---

## Reverse Proxy

Forward proxy clientni yashiradi. Reverse proxy serverni yashiradi.

```text
Client
   ↓
Reverse Proxy
   ↓
Backend
```

### Misollar

| Reverse Proxy | Ishlatilishi |
|---------------|-------------|
| Cloudflare | CDN + DDoS himoya |
| Nginx |eng ko'p ishlatiladigan |
| HAProxy | Yuqori trafigli loyihalar |
| Traefik | Docker/Kubernetes bilan |

### Qanday ishlaydi?

```text
Client api.company.com ga so'rov yuboradi.

Reverse Proxy (Nginx):
- SSL/TLS ni hal qiladi
- Trafikni backend serverlarga taqsimlaydi
- Caching, compression, rate limiting qiladi

Backend serverlar:
- Faqat business logic yozadi
- Reverse proxy orqali clientni ko'rmaydi
```

---

## Forward vs Reverse Proxy

| Xususiyat | Forward Proxy | Reverse Proxy |
|-----------|---------------|---------------|
| Qayerda turadi | Client tarafida | Server tarafida |
| Kimni yashiradi | Clientni | Serverni |
| Misol | VPN, korporativ proxy | Nginx, Cloudflare |
| Maqsad | Maxfiylik, filtrlash | Xavfsizlik, optimizatsiya |

---

## Xulosa

- Proxy = o'rtakash
- Forward proxy clientni, reverse proxy serverni yashiradi
- VPN — Forward Proxy'ning bir turi
- Reverse Proxy — zamonaviy backend arxitekturasining ajralmas qismi
