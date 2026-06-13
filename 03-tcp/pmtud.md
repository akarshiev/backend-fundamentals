# PMTUD (Path MTU Discovery)

Yo'ldagi eng kichik MTU ni topadi.

---

## Nazariya

Internet yo'li turli qurilmalardan o'tadi. Har bir qurilmaning MTU si har xil bo'lishi mumkin.

```text
Server -> Router -> VPN -> Client

Server MTU:  1500
Router MTU:  1500
VPN MTU:     1400
Client MTU:  1500

Path MTU = min(1500, 1500, 1400, 1500) = 1400
```

---

## Diagram

```text
Server                    Router                    Client
  |--- 1500 bytes -------->|                         |
  |                        |--- 1500 bytes --------->|
  |                        |    (VPN: MTU 1400)       |
  |                        |    Fragment! ✗           |
  |<-- ICMP "too big" -----|                         |
  |                        |                         |
  |--- 1400 bytes -------->|                         |
  |                        |--- 1400 bytes --------->|
  |                        |    (o'tdi!) ✓           |
```

---

## Qanday ishlaydi

1. Sender katta paket yuboradi
2. Agar paket yo'ldagi MTU'dan katta bo'lsa, router uni tashlab tashlaydi
3. Router ICMP "fragmentation needed" xabarini yuboradi
4. Sender kichikroq paket yuboradi

---

## Fragmentation

Agar PMTUD ishlamasa, paket fragment bo'ladi:

```text
1500 byte paket
    ↓
Router (MTU 1400)
    ↓
2 ta fragment:
  - 1400 byte
  - 100 byte (20 byte header bilan 120 byte)
```

Fragmentation yomon:

```text
✗ Tezlikni kamaytiradi
✗ Xatolik ehtimolini oshiradi
✗ Qayta yig'ish kerak
```

---

## PMTUD Process

```text
1. Sender MSS = 1460 (MTU 1500)
2. Paket yuboradi (1500 bytes)
3. Router: "MTU 1400, bu katta!"
4. Router: ICMP "fragmentation needed"
5. Sender: "Yaxshi, 1400 ga tushaman"
6. Sender: MSS = 1400 - 40 = 1360
7. Keyingi paketlar 1400 bytes
```

---

## Amaliyot

### Path MTU ni tekshirish

```bash
# MTU yo'lini tekshirish
ping -M do -s 1472 google.com

# 1472 + 28 (ICMP header) = 1500 (MTU)
```

### Traceroute bilan MTU topish

```bash
traceroute -M do -s 1472 google.com
```

### ss bilan

```bash
ss -ti
```

---

## Xulosa

- PMTUD yo'ldagi eng kichik MTU'ni topadi
- Fragmentation oldini oladi
- ICMP "fragmentation needed" ishlatiladi
- VPN va tunnellarda muhim
