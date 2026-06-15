# Redis

In-Memory Key-Value Database. Disk emas, RAM da ishlaydi.

---

## Redis nima?

```text
Redis = Remote Dictionary Server

Ma'lumot tuzilmasi:
user:1 -> Otabek
user:2 -> Akbar
session:abc123 -> { user_id: 1, role: "admin" }
```

---

## Nega tez?

| Saqlash | Tezlik |
|---------|--------|
| RAM | ~100 ns |
| Disk (HDD) | ~10 ms |
| Disk (SSD) | ~100 μs |

```text
RAM ~ 100,000 marta tezroq!
```

---

## Redis ishlatilishi

### 1. Cache

```text
Savol: "Otabek ismi qancha?"

Oldin: Database dan ol (sekin)
Keyin: Redis dan ol (tez)

Cache miss bo'lsa → Database dan olib Redis ga yoz
```

### 2. Session

```text
Login → Session ID yaratish → Redis ga saqlash

Session ID: abc123
Data: { user_id: 1, name: "Otabek", role: "admin" }
TTL: 3600 sekund (1 soat)
```

### 3. Rate Limiting

```text
Client 1 daqiqada 100 ta so'rov yubordi.

Redis:
INCR rate:user:123
EXPIRE rate:user:123 60

Agar 100 dan oshsa → 429 Too Many Requests
```

### 4. Queue

```text
Email yuborish:

Redis List:
LPUSH email_queue { to: "user@example.com", subject: "Salom" }

Worker:
RPOP email_queue → Email yuborish
```

### 5. Pub/Sub

```text
Chat ilovasi:

Publisher: PUBLISH chat "Salom!"
Subscriber 1: SUBSCRIBE chat → "Salom!" oladi
Subscriber 2: SUBSCRIBE chat → "Salom!" oladi
```

### 6. OTP

```text
SMS kodi yuborish:

SET otp:998901234567 543210 EXPIRE 300

300 sekund ichida tekshirish:
GET otp:998901234567
```

### 7. Load Balancing

```text
Session-based load balancing:

Session ID: abc123
Redis: { server: "app2" }

Client har doim app2 ga boradi (Sticky Session)
```

---

## Redis buyruqlari

```bash
# Qiymat qo'yish
SET user:1 "Otabek"

# Qiymat olish
GET user:1

# O'chirish
DEL user:1

# Muddat bilan
SET session:abc "data" EX 3600

# List
LPUSH queue "task1"
RPOP queue

# Hash
HSET user:1 name "Otabek" age 25
HGET user:1 name

# Counter
INCR page_views
DECR page_views
```

---

## Xulosa

| Xususiyat | Tushuntirish |
|-----------|-------------|
| Tezlik | RAM da ishlaydi (~100 ns) |
| Cache | Ma'lumotlarni qayta ishlatish |
| Session | Foydalanuvchi sessiyalarini saqlash |
| Rate Limiting | So'rovlar sonini cheklash |
| Queue | Vazifalarni navbatga qo'yish |
| Pub/Sub | Real-time xabarlashish |
| OTP | SMS/email kodlarni saqlash |
