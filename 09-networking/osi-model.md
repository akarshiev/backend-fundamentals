# OSI Model

## OSI nima?

**OSI = Open Systems Interconnection**

Tarmoqni qatlamlarga bo'lish modeli.

---

## Nima muammoni hal qiladi?

Agar qatlamlar bo'lmasa:

- HTTP
- TCP
- IP
- Ethernet

hammasi aralashib ketadi.

OSI: har qatlamni alohida mas'ul qiladi.

---

## OSI 7 Layer

### 7. Application

HTTP, DNS, SMTP

### 6. Presentation

TLS, Encryption, Compression

### 5. Session

Session Management

### 4. Transport

TCP, UDP, Portlar

### 3. Network

IP Address, Routing

### 2. Data Link

MAC Address, Switch, Ethernet

### 1. Physical

Kabel, Wi-Fi Signal, Elektr impulslari

---

## Eng muhimlari

Backend uchun:

| Layer | Texnologiya |
|-------|-------------|
| Layer 7 | HTTP |
| Layer 4 | TCP / UDP |
| Layer 3 | IP |

---

## Amaliyot

Tarmoq qatlamlarini ko'rish:

```bash
# HTTP (Layer 7) headerlarini ko'rish
curl -v https://example.com 2>&1 | grep -i "< "

# TCP (Layer 4) ulanishlarini ko'rish
ss -tulpn

# IP (Layer 3) route jadvalini ko'rish
ip route
```
