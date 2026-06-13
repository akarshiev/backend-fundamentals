# TCP (Transmission Control Protocol)

Ishonchli, tartibli va to'liq aloqa protokoli.

---

## Mavzular

- [TCP Basics](tcp-basics.md) -- asosiy tushunchalar
- [TCP 3-Way Handshake](tcp-handshake.md) -- ulanish o'rnatish
- [TCP Flags](tcp-flags.md) -- SYN, ACK, FIN, RST, PSH, URG va boshqalar
- [Sequence va ACK Numbers](sequence-ack.md) -- paket tartibini boshqarish
- [TCP Checksum](checksum.md) -- ma'lumot yaxlitligini tekshirish
- [RTT](rtt.md) -- Round Trip Time
- [TCP Fast Open](tcp-fast-open.md) -- TFO, RTT ≈ 0
- [Head Of Line Blocking](hol-blocking.md) -- HOL muammosi
- [TCP Retransmission](packet-loss.md) -- yo'qolgan paketlarni qayta yuborish
- [Flow Control](flow-control.md) -- receive window
- [Congestion Control](congestion-control.md) -- tarmoqni himoya qilish
- [TCP Slow Start](slow-start.md) -- sekin boshlash algoritmi
- [Buffer](buffer.md) -- paketlar saqlanadigan joy
- [MTU](mtu.md) -- maksimal paket o'lchami
- [MSS](mss.md) -- maksimal segment o'lchami
- [PMTUD](pmtud.md) -- Path MTU Discovery
- [TCP Packet Structure](tcp-packet-structure.md) -- paket tuzilishi

---

## TCP nima uchun kerak?

Internetda ma'lumotlar paketlarga bo'linib yuboriladi. Transport layer'da eng mashhur ikki protocol:

- **TCP** -- ishonchli, tartibli, ulanish asosida
- **UDP** -- tez, yengil, ulanishsiz

TCP ma'lumot yetib borganini **kafolatlaydi**.

---

## TCP ishlatiladigan joylar

| Protokol | TCP ishlatilishi |
|----------|-----------------|
| HTTPS    | Web trafik       |
| SSH      | Masofadan boshqarish |
| PostgreSQL | Bazaga ulanish |
| Redis    | Cache ulanish    |
| SMTP     | Email yuborish   |

---

## TCP Lifecycle

```text
1. Connection -- ulanish o'rnatish
2. ACK       -- paketlar qabul qilinishini tasdiqlash
3. Ordering  -- paketlarni tartibda saqlash
4. Retransmission -- yo'qolgan paketlarni qayta yuborish
```

---

## Amaliyotlar

### 1. TCP Connections

```bash
ss -tan
```

### 2. Listening Ports

```bash
ss -tulpn
```

### 3. Packet Capture

```bash
sudo tcpdump -i any port 8080
```

### 4. TCP Server (Go)

```bash
cd examples && go run tcp-server.go
```

### 5. TCP Client (Go)

```bash
cd examples && go run tcp-client.go
```

---

## Xulosa

- TCP ishonchli aloqa protokoli
- 3-Way Handshake orqali ulanish ochiladi
- Flags, Sequence, ACK, Checksum mehanizmlari ishonchni ta'minlaydi
- Flow Control va Congestion Control tarmoqni himoya qiladi
- RTT, MTU, MSS, PMTUD paket o'lchamlarini boshqaradi
