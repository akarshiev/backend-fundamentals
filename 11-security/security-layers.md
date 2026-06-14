# Himoyalash qatlamlari (Security Layers)

Internetdan kelayotgan trafik bir necha qatlamda filtrlanadi.

---

## Qatlamlar sxemasi

```text
Internet
   ↓
Firewall (L3/L4)
   ↓
WAF (L7)
   ↓
Application
   ↓
Database
```

---

## Har bir qatlam

### 1. Firewall (Layer 3/4)

```text
Vazifasi:
- IP manzillarni filtrlash
- Portlarni boshqarish
- TCP/UDP protokollarni nazorat qilish

Misol:
- Port 22 (SSH) faqat ma'lum IP lardan ochiq
- Port 80, 443 hamma uchun ochiq
- Port 5432 (DB) faqat ichki tarmoqdan ochiq
```

### 2. WAF (Layer 7)

```text
Vazifasi:
- HTTP so'rovlarini tekshirish
- SQL injection ni bloklash
- XSS ni bloklash
- Rate limiting

Misol:
- GET /page?id=1' OR '1'=1 → BLOCK
- POST <script>alert('XSS')</script> → BLOCK
```

### 3. Application

```text
Vazifasi:
- Autentifikatsiya (kirish)
- Autorizatsiya (ruxsat)
- Ma'lumotlarni tozalash
- Xatolarni boshqarish

Misol:
- JWT token tekshirish
- Role-based access control
- Input validation
```

### 4. Database

```text
Vazifasi:
- Ma'lumotlarni shifrlash
- Backup olish
- Access control
- Logging

Misol:
- SQL injection dan himoya (parameterized queries)
- Ma'lumotlarni shifrlash (encryption at rest)
- Audit log
```

---

## Himoya strategiyasi

### Defense in Depth

```text
Bir qatlam buzilsa, keyingisi himoya qiladi.

Masalan:
1. Firewall WAF ni himoya qiladi
2. WAF Application ni himoya qiladi
3. Application Database ni himoya qiladi
```

### Least Privilege

```text
Faqat kerakli ruxsatlarni bering.

Masalan:
- Web server faqat READ ruxsatga ega
- Database admin faqat kerakli operatsiyalar
```

### Fail Securely

```text
Xatolik yuz berganda xavfsiz holatga o'ting.

Masalan:
- Autentifikatsiya xatosi → Kirishni rad etish
- Database xatosi → 500 Server Error
```

---

## Amaliyot

### Firewall tekshirish

```bash
# UFW holati
sudo ufw status

# Ochiq portlar
ss -tulpn
```

### WAF test

```bash
# SQL Injection test
curl "http://localhost/page?id=1' OR '1'=1"

# XSS test
curl -d "name=<script>alert('XSS')</script>" http://localhost/submit
```

### Application test

```bash
# JWT token test
curl -H "Authorization: Bearer <token>" http://localhost/api

# Role-based test
curl -H "Authorization: Bearer <user-token>" http://localhost/admin
```

---

## Xulosa

| Qatlam | OSI Layer | Vazifa |
|--------|-----------|--------|
| Firewall | L3/L4 | IP va port filtrlash |
| WAF | L7 | HTTP so'rovlarini tekshirish |
| Application | L7 | Autentifikatsiya va autorizatsiya |
| Database | - | Ma'lumotlarni himoya qilish |
