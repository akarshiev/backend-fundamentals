# TCP vs UDP

Ikki asosiy transport protokolini taqqoslash.

---

## Umumiy taqqoslash

```text
                    TCP                     UDP
Connection:         ✓ Bor                   ✓ Yo'q
Reliable:           ✓ Ishonchli             ✗ Ishonchsiz
Ordered:            ✓ Tartibli              ✗ Tartibsiz
Retransmission:     ✓ Bor                   ✗ Yo'q
Flow Control:       ✓ Bor                   ✗ Yo'q
Congestion Control: ✓ Bor                   ✗ Yo'q
Speed:              ✗ Sekinroq              ✓ Tez
Overhead:           ✓ Katta                 ✓ Kichik
Latency:            ✓ Yuqori                ✓ Past
Header Size:        20 bytes                8 bytes
```

---

## TCP kim uchun?

```text
✓ Web (HTTPS)
✓ Email (SMTP)
✓ Fayl uzatish (FTP)
✓ Ma'baza (PostgreSQL, Redis)
✓ SSH

Sababi: Ma'lumot yetib borishi kerak.
```

---

## UDP kim uchun?

```text
✓ DNS so'rovlar
✓ VoIP (ovozli qo'ng'iroq)
✓ Online o'yinlar
✓ Video streaming
✓ QUIC (HTTP/3)

Sababi: Tezlik muhim, ayrim packet yo'qolsa ham bo'ladi.
```

---

## Real-world misollar

### TCP misollari

```text
HTTPS:
  Client --> SYN --> Server
         <-- SYN-ACK -->
         --> ACK -->
         --> GET /index.html -->
         <-- 200 OK + Data -->

SSH:
  Client --> SYN --> Server
         <-- SYN-ACK -->
         --> ACK -->
         --> Encryption -->
         <-- Encryption -->
```

### UDP misollari

```text
DNS:
  Client --> DNS Query (UDP) --> DNS Server
         <-- DNS Response (UDP) <-- DNS Server
  (Connection yo'q, tez!)

Online Game:
  Player --> Position Update (UDP) --> Server
  (Agar paket yo'qolsa, keyingisi keladi)
```

---

## Qachon TCP, qachon UDP?

### TCP tanlash kerak:

```text
✗ Ma'lumot yetib borishi shart bo'lsa
✗ Tartib muhim bo'lsa
✗ Xatolik yo'q bo'lishi kerak bo'lsa
✗ Connection kerak bo'lsa
```

### UDP tanlash kerak:

```text
✓ Tezlik muhim bo'lsa
✓ Real-time kerak bo'lsa
✓ Ayrim packet yo'qolsa ham bo'lsa
✓ Connection kerak bo'lmasa
```

---

## HTTP/2 vs HTTP/3

```text
HTTP/2: TCP + TLS
  ✓ Ishonchli
  ✗ HOL (Head Of Line Blocking)

HTTP/3: QUIC (UDP) + TLS
  ✓ Ishonchli
  ✓ HOL yo'q
  ✓ Tezroq
```

---

## Amaliyot

### TCP va UDP ni taqqoslash

```bash
# TCP
time curl http://example.com

# UDP (DNS)
time dig google.com
```

### Socket turlarini ko'rish

```bash
# TCP socketlar
ss -tan

# UDP socketlar
ss -uan
```

---

## Xulosa

- TCP ishonchli, lekin sekinroq
- UDP tez, lekin ishonchsiz
- Qaysi birini tanlash -- vazifaga bog'liq
- HTTP/3 QUIC (UDP) orqali ikkala afzallikni birlashtiradi
