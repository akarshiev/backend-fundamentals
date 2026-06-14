# QUIC Protocol

## QUIC nima?

**QUIC = Quick UDP Internet Connections**

Google tomonidan ishlab chiqilgan protocol.

Keyinchalik **IETF standardiga** aylandi.

---

## Nega UDP ustiga qurilgan?

| TCP | QUIC |
|-----|------|
| Kernel Space ichida | User Space ichida |
| OS yangilanishi kerak | Kernel patch kutmaydi |
| O'zgartirish qiyin | Application ichida yozish mumkin |

UDP:

- Fast
- Simple
- Connectionless
- TCP kabi kernelga bog'lanmagan

Shuning uchun QUIC UDP ustiga o'zining mexanizmlarini qurdi:

- Reliable delivery
- Ordering
- Retransmission
- Flow control
- Congestion control

---

## User Space va Kernel Space

### TCP

```
TCP → Kernel Space ichida ishlaydi
OS yangilanmasdan TCP'ni o'zgartirish qiyin
```

### QUIC

```
QUIC → User Space ichida ishlaydi
Chrome, Cloudflare, Nginx o'z implementatsiyasini chiqarishi mumkin
Kernel patch kutmaydi
```

---

## QUIC va Head Of Line Blocking

### TCP

Connection darajasida HOL bor.

Bitta packet yo'qolsa → butun connection kutadi.

### QUIC

Stream darajasida ishlaydi.

```
Stream 1 — ishlayapti ✓
Stream 2 — kutmoqda ✗ (packet yo'qoldi)
Stream 3 — ishlayapti ✓
```

Stream 2 da packet yo'qolsa, faqat Stream 2 kutadi. Stream 1 va Stream 3 ishlashda davom etadi.

---

## QUIC va TLS 1.3

### Oddiy HTTPS

```
TCP Handshake → TLS Handshake → HTTP Request
```

3 ta qadam, kamida 2 RTT.

### QUIC

```
QUIC Handshake (TLS 1.3 built-in) → HTTP Request
```

TLS 1.3 protocol ichiga qurilgan. Kamroq RTT, tezroq ulanish.
