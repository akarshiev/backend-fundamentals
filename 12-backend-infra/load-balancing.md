# Load Balancing

Bir server yetmay qolganda ishlatiladi. Trafikni bir nechta serverga taqsimlaydi.

---

## Load Balancing nima?

```text
Oldin:

Client
 ↓
App (1 ta server)

Keyin:

Client
 ↓
LB (Load Balancer)
 ├── App1
 ├── App2
 └── App3
```

---

## Algoritmlar

### 1. Round Robin

Eng oddiy algoritm. Ketma-ket taqsimlaydi.

```text
Req1 → App1
Req2 → App2
Req3 → App3
Req4 → App1
Req5 → App2
Req6 → App3
...
```

| Afzalligi | Kamchiligi |
|-----------|-----------|
| Sodda, oson tushunish | Server kuchlarini hisobga olmaydi |

---

### 2. Least Connections

Eng kam aktiv connection bor serverga yuboradi.

```text
App1 = 200 conn
App2 = 30 conn
App3 = 20 conn

Yangi request → App3 ga ketadi (eng kam connection)
```

| Afzalligi | Kamchiligi |
|-----------|-----------|
| Serverlarni teng ushlaydi | Har bir connection ni hisoblash kerak |

Ko'pincha Round Robin'dan yaxshiroq ishlaydi.

---

### 3. IP Hashing

Bir xil IP har doim bir xil serverga boradi.

```text
192.168.1.10 → doim App2 ga tushadi
192.168.1.20 → doim App1 ga tushadi
```

| Afzalligi | Kamchiligi |
|-----------|-----------|
| Sticky Session — session'lar saqlanadi | Server tizilsa, session'lar yo'qoladi |

Eski sistemalarda ishlatilgan. Hozir Redis/session store o'rnini oldi.

---

### 4. Weighted Load Balancing

Serverlar kuchi har xil bo'lsa ishlatiladi.

```text
App1 = 16 CPU (kuchli)
App2 = 4 CPU
App3 = 4 CPU

Weight:
App1 = 4
App2 = 1
App3 = 1

Taqsimlash:
App1, App1, App1, App1, App2, App3
```

---

## L4 Load Balancing (Layer 4)

TCP/UDP darajasida ishlaydi.

```text
OSI: Layer 4 → TCP / UDP

LB faqat ko'radi:
- IP manzil
- Port raqami

HTTP ichiga qaramaydi.
```

| Afzalligi | Kamchiligi |
|-----------|-----------|
| JUDA TEZ | Aqlli routing yo'q |

### Misollar

| Vosita | Ishlatilishi |
|--------|-------------|
| AWS NLB | Yuqori tezlik kerak bo'lganda |
| HAProxy TCP Mode | TCP trafigi uchun |

---

## L7 Load Balancing (Layer 7)

HTTP darajasida ishlaydi.

```text
OSI: Layer 7 → HTTP

LB ko'radi:
- Path (/api/*, /static/*)
- Header
- Cookie
- Host
```

| Afzalligi | Kamchiligi |
|-----------|-----------|
| Aqlli routing | CPU ko'proq ishlatadi |

### Misol

```text
/api/*      → Backend
/static/*   → CDN
/images/*   → Image Server

HTTP header ga qarab yo'naltiradi.
```

---

## L4 vs L7

| Xususiyat | L4 | L7 |
|-----------|-----|-----|
| OSI Layer | TCP/UDP | HTTP |
| Tezlik | Juda tez | Sekinroq |
| Routing | IP + Port | Path, Header, Cookie |
| CPU | Kam | Ko'p |
| Misol | AWS NLB | Nginx, HAProxy |

---

## Mirroring (Traffic Mirroring)

Traffic ni nusxalash. Yangi versiyani test qilish uchun.

```text
Asosiy trafik:

Client
   ↓
Production (v1)

Mirror:

Client
   ↓
Production (v1) ← foydalanuvchi javobni shu yerdan oladi
   ↓
Shadow Server (v2) ← request nusxasi shu yerga ketadi
```

### Qanday ishlaydi?

```text
1. Client request yuboradi
2. Production (v1) javob beradi
3. Request nusxasi Shadow (v2) ga ham yuboriladi
4. Foydalanuvchi sezmaydi — javobni production dan oladi
```

### Nima uchun?

```text
Yangi versiyani real trafik bilan test qilish.

Misol:
- Prod v1 → ishlaydi
- Prod v2 → test qilinmoqda

Real trafik: v1 ga ishlaydi
Nusxa: v2 ga ketadi
```

### Nginx misoli

```nginx
# Traffic mirroring
server {
    location / {
        proxy_pass http://production;

        # Mirror to shadow
        mirror /mirror;
        mirror_request_body on;
    }

    location = /mirror {
        internal;
        proxy_pass http://shadow$request_uri;
    }
}
```

---

## Xulosa

| Algoritm | Tushuntirish |
|----------|-------------|
| Round Robin | Ketma-ket taqsimlash |
| Least Connections | Eng kam connection ga |
| IP Hashing | Bir xil IP → bir xil server |
| Weighted | Server kuchiga qarab |
| L4 | TCP/UDP darajasida (tez) |
| L7 | HTTP darajasida (aqlli) |
| Mirroring | Traffic ni nusxalash |
