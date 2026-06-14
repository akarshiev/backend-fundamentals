# MAC Address

Layer 2 identifikator, qurilmaga biriktirilgan.

---

## MAC nima?

MAC = Media Access Control Address

Har bir network kartasi (NIC) uchun yagona fizik manzil.

```text
00:1A:2B:3C:4D:5E
```

---

## Tuzilishi

```text
00:1A:2B:3C:4D:5E
│        │           │
│        │           └── Vendor Specific (3 byte)
│        └────────────── Vendor OUI (3 byte)  
└─────────────────────── Format identifier
```

---

## O'zgarish mumkinmi?

| Xususiyat | IP | MAC |
|-----------|-----|------|
| O'zgarishi | Ha | Yo'q (lekin spoofing mumkin) |
| Scope | Global | Lokal tarmoq |
| OSI Layer | 3 | 2 |

---

## Amaliyot

### MAC manzilni ko'rish

```bash
# Linux
ip link show

# macOS
ifconfig en0 | grep ether

# Windows
getmac
```

### MAC spoofing (test uchun)

```bash
# Linux (vaqtinchalik)
sudo ip link set dev eth0 down
sudo ip link set dev eth0 address 00:11:22:33:44:55
sudo ip link set dev eth0 up
```

---

## Xulosa

- MAC Layer 2 fizik manzil
- Vendor va device identifikatsiya qiladi
- Lokal tarmoqda ishlatiladi
