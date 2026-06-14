# UDP (User Datagram Protocol)

Tez, yengil, ulanishsiz transport protokoli.

---

## Mavzular

- [UDP Basics](udp-basics.md) -- asosiy tushunchalar
- [TCP vs UDP](tcp-vs-udp.md) -- taqqoslash

---

## UDP nima?

UDP = User Datagram Protocol

Asosiy falsafa: **Fire and Forget** -- Yubor va unut.

TCP dan farqi:

```text
TCP: Ishonchli, tartibli, ulanish asosida
UDP: Tez, yengil, ulanishsiz
```

---

## UDP xususiyatlari

```text
✗ Connection yo'q
✗ ACK yo'q
✗ Ordering yo'q
✗ Retransmission yo'q
✗ Flow control yo'q
✗ Congestion control yo'q

✓ Juda tez
✓ Kam overhead
✓ Kam latency
```

---

## UDP ishlatiladigan joylar

| Protokol | UDP ishlatilishi |
|----------|-----------------|
| DNS      | Domain so'rovlar |
| VoIP     | Ovozli qo'ng'iroq |
| Online Games | Tezlik kerak |
| Video Streaming | Real-time |
| QUIC     | HTTP/3 asosi |

---

## Amaliyotlar

### 1. UDP Server (Go)

```bash
cd examples && go run udp-server.go
```

### 2. UDP Client (Go)

```bash
cd examples && go run udp-client.go
```

### 3. DNS Client (Go)

```bash
cd examples && go run dns-client.go google.com
```

### 4. UDP diagnostikasi

```bash
# UDP socketlarni ko'rish
ss -uan

# UDP traffic ni tekshirish
sudo tcpdump -i any udp port 53
```

---

## Xulosa

- UDP tez va yengil protokol
- Fire and Forget falsafasi
- DNS, VoIP, Gaming, Streaming da ishlatiladi
- TCP ga nisbatan tezroq, lekin ishonchsiz
