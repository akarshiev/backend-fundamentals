# Rate Limiting & Throttling

Server resurslarini himoya qilish. Bitta foydalanuvchining haddan tashqari so'rovlari boshqalarga xalaqit bermasligi kerak.

---

## Noisy Neighbor Problem

```text
Cloud va distributed systemlarda juda mashhur muammo.

Server
├── User A
├── User B
├── User C
└── User D

hammasi bir xil resurslardan foydalanmoqda.
```

### Oddiy holat

```text
A → 10 req/sec
B → 15 req/sec
C → 8 req/sec
D → 12 req/sec

hamma baxtli.
```

### Muammo

```text
A → 10000 req/sec qilib yubordi.

Natija:
- CPU 100%
- RAM to'ldi
- DB connection tugadi

B va C foydalanuvchilar:
- 503 Service Unavailable
- Timeout

ola boshlaydi.
```

### Hayotiy misol

```text
Ko'p qavatli uyda yashaysiz.
Hamma xonadonlar bitta Wi-Fi routerdan foydalanadi.

Oddiyda: kimdir yangilik o'qiydi, kimdir Instagram ko'radi.

Bir kuni 5-qavatdagi "serhovli qo'shni" (Noisy Neighbor)
4K formatdagi 10 ta filmni birdaniga yuklab olishga qo'ydi.

Natija: Wi-Fi tezligi hamma uchun o'lib qoldi.
Siz oddiy Google sahifasini ham ocholmaysiz.
```

---

## Yechim: Rate Limiting & Throttling

Tizimda shunday "Politsiya" kerakki, u har bir foydalanuvchiga aytsin:

> "Sekundiga 5 ta so'rov mumkin. Agar oshirib yuborsang, seni vaqtincha bloklayman, toki boshqalarga xalaqit bermagin."

Bu nafaqat xavfsizlik (DDoS dan himoya), balki **adolat (Fairness)** masalasidir.

---

## Rate Limiting vs Throttling

### Rate Limiting

**Maqsad:** Qancha request yuborishi mumkinligini cheklash.

```text
Limit = 100 req/min

101-request → 429 Too Many Requests
```

```text
Rate Limiting:
Request → Limit?
              ↓
            No → STOP (429)
```

### Throttling

**Maqsad:** To'xtatib qo'yish emas, **sekinlashtirish**.

```text
Rate Limiting:
100 req/min oshdi → 429 (STOP)

Throttling:
100 req/min oshdi → Kut → Keyin ishlat
```

```text
Throttling:
Request → Limit?
              ↓
            No → Queue → Delay → Process
```

### Misol: Netflix

```text
Video: 4K o'rniga 720p ga tushadi
API: 100ms o'rniga 2 sec delay qo'yadi
```

### Farqi

| Xususiyat | Rate Limiting | Throttling |
|-----------|---------------|------------|
| Natija | STOP (429) | SLOW DOWN |
| Maqsad | Bloklash | Sekinlashtirish |
| Misol | Login API bloklangan | Video sifati pastlangan |

### Qo'llanilish joylari

| API | Rate Limit |
|-----|-----------|
| Login API | 5 urinish / 15 minut |
| OTP API | 3 urinish / 5 minut |
| Search API | 100 req / minute |
| Payment API | 10 req / minute |

---

## Rate Limiting Algoritmlari

### 1. Fixed Window Counter

Eng sodda algoritm.

```text
Limit = 5
Window = 60 sec

10:00 → 10:01 oralig'ida nechta request bo'ldi?
```

#### Kamchiligi: Burst Problem

```text
10:00:59 → 5 request (limit bo'yicha OK)
10:01:01 → 5 request (yangi oyna, yana OK)

2 soniyada: 10 request o'tib ketadi!
Bu Burst Problem deyiladi.
```

---

### 2. Token Bucket

Eng mashhur algoritm.

```text
Bucket ichida tokenlar bor.

Har request: 1 token oladi.

Token qolsa → Allow ✅
Token tugasa → 429 ❌
```

#### Misol

```text
Capacity = 10
Refill = 5/sec

10 soniya hech kim kelmadi.
Bucket to'ldi.

Birdan 10 request keldi.
Hammasi o'tadi.

Bu Burst Traffic ni qo'llab-quvvatlaydi.
```

#### Qo'llanilishi

| Xizmat | Model |
|--------|-------|
| Google Cloud API | Token Bucket |
| AWS API Gateway | Token Bucket |
| Stripe | Token Bucket |

---

### 3. Leaky Bucket

Teshik chelak.

```text
Input: 1000 req/sec tez kelishi mumkin
Output: 10 req/sec qat'iy

[Bucket]
   ↓
   ↓
   ↓
Output: 10 req/sec

Chelak to'lsa → Drop
```

#### Afzalligi

```text
Server yuklamasi tekis.
```

#### Qo'llanilishi

```text
Nginx limit_req ichida shunga o'xshash model ishlatiladi.
```

---

### Algoritmlarni taqqoslash

| Algoritm | Xususiyat | Qo'llanilishi |
|----------|-----------|---------------|
| Fixed Window | Sodda, burst muammosi bor | Oddiy API limitlar |
| Token Bucket | Burst ga ruxsat beradi | Google, AWS, Stripe |
| Leaky Bucket | Tekis chiqish | Nginx, Shopify |

---

## Production Muammo: Distributed Rate Limiting

### Muammo

```text
1 server bo'lsa:

requests[ip] += 1 → ishlaydi ✅

10 server bo'lsa:

LB
├── App1 (o'z xotirasi)
├── App2 (o'z xotirasi)
├── App3 (o'z xotirasi)
└── App4 (o'z xotirasi)

Har birida alohida memory.
Limit ishlamay qoladi. ❌
```

### Yechim: Redis

```text
Client
   ↓
Load Balancer
   ↓
Application
   ↓
Redis (markazlashgan xotira)

Hammasi Redis ga qaraydi.
```

---

## Race Condition

```text
Limit: 5

Server A: count = 4 o'qidi
Server B: count = 4 o'qidi

Ikkalasi ham 5 yozdi.

Aslida: 6 bo'lishi kerak edi.
Bu Race Condition deyiladi.
```

### Yechim: Redis + Lua

Redis **Single Threaded**. Lua script tekshir → oshir → expire qo'y hammasini **1 operation** qilib bajaradi.

**Atomic** — Race Condition yo'q.

---

## Redis + Lua Buyruqlari

```bash
# Counter
INCR login:user:123

# TTL
EXPIRE login:user:123 60

# Tekshirish
TTL login:user:123
```

---

## Amaliyot: Python MVP (In-Memory)

```python
import time

# Xotira: { "IP_MANZIL": [SO'ROV_SONI, BOSHLANGAN_VAQT] }
buckets = {}

def is_allowed(ip_address, limit=5, window=60):
    current_time = time.time()
    
    # Agar bu IP birinchi marta kelsa
    if ip_address not in buckets:
        buckets[ip_address] = [1, current_time]
        return True

    count, start_time = buckets[ip_address]

    # Agar vaqt oynasi (1 daqiqa) tugagan bo'lsa, yangilaymiz
    if current_time - start_time > window:
        buckets[ip_address] = [1, current_time]
        return True

    # Agar limitdan oshib ketsa
    if count >= limit:
        return False  # ⛔ STOP

    # So'rovni sanaymiz
    buckets[ip_address][0] += 1
    return True

# Sinov
user_ip = "192.168.1.1"
for i in range(7):
    if is_allowed(user_ip):
        print(f"So'rov {i+1}: ✅ O'tdi")
    else:
        print(f"So'rov {i+1}: ⛔ 429 Too Many Requests")
```

### Muammo

```text
Bu kod faqat bitta serverda ishlaydi.
10 ta server bo'lsa, har bir serverning o'z xotirasi bo'ladi.
Foydalanuvchi jami 50 ta so'rov yuborishi mumkin.

Bizga Markazlashgan Xotima kerak → Redis.
```

---

## Amaliyot: Redis + Lua (Production)

```python
import redis

r = redis.Redis(host='localhost', port=6379, decode_responses=True)

# LUA SCRIPT: Redis ichida bajariladi!
lua_script = """
local key = KEYS[1]
local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])

-- 1. Hozirgi qiymatni oshiramiz (INCR)
local current = redis.call("INCR", key)

-- 2. Agar birinchi so'rov bo'lsa, vaqtni belgilaymiz
if current == 1 then
    redis.call("EXPIRE", key, window)
end

-- 3. Natijani tekshiramiz
if current > limit then
    return 0  -- Bloklandi
else
    return 1  -- Ruxsat
end
"""

# Scriptni Redisga yuklaymiz
limiter_script = r.register_script(lua_script)

def check_rate_limit(ip, limit=5, window=60):
    key = f"rate_limit:{ip}"
    allowed = limiter_script(keys=[key], args=[limit, window])
    return allowed == 1

# Sinov
ip = "10.0.0.5"
for i in range(7):
    if check_rate_limit(ip):
        print(f"So'rov {i+1}: ✅ Ruxsat")
    else:
        print(f"So'rov {i+1}: ⛔ Limit tugadi, 1 daqiqa kuting.")
```

### Nima uchun mukammal?

| Xususiyat | Tushuntirish |
|-----------|-------------|
| Tez | Redis xotirada ishlaydi (~100 ns) |
| Xavfsiz | Lua tufayli Race Condition yo'q |
| Tarqoq | 100 ta server bo'lsa ham bitta Redis ga qaraydi |

---

## Xulosa

| Mavzu | Tushuntirish |
|-------|-------------|
| Noisy Neighbor | Bitta foydalanuvchi hamma resursni olib ketishi |
| Rate Limiting | So'rovlar sonini cheklash (STOP) |
| Throttling | So'rovlar tezligini sekinlashtirish (SLOW DOWN) |
| Fixed Window | Sodda, burst muammosi bor |
| Token Bucket | Burst ga ruxsat, eng mashhur |
| Leaky Bucket | Tekis chiqish, Nginx ishlatadi |
| Distributed | Redis + Lua — atomik, tez, tarqoq |
