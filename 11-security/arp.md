# ARP (Address Resolution Protocol)

IP manzildan MAC manzilga o'tish protokoli.

---

## ARP nima?

ARP = Address Resolution Protocol

Lokal tarmoqda IP manzilni MAC manzilga aniqlash uchun ishlatiladi.

```text
192.168.1.1  →  AA:BB:CC:DD:EE:FF
```

---

## Ishlash prinsipi

```text
1. Device A: "192.168.1.1 ning MAC adresi kim?"
2. Broadcast: hamma ga yuboriladi
3. Device B: "Men 192.168.1.1, mening MAC: AA:BB:CC:DD:EE:FF"
4. Device A: ARP jadvaliga saqlaydi
```

---

## ARP jadvali

```bash
arp -a
# yoki
ip neigh
```

---

## Amaliyot

### ARP jadvalini ko'rish

```bash
# Linux
ip neigh

# macOS
arp -a

# Windows
arp -a
```

### ARP so'rov yuborish

```bash
# Ping orqali ARP ni trigger qilish
ping -c 1 192.168.1.1

# Keyin ARP jadvalini tekshirish
ip neigh show
```

### ARP packet capture

```bash
sudo tcpdump -i any arp
```

---

## ARP turlari

| Tur | Tushuntirish |
|-----|-------------|
| Request | "Bu IP ning MAC si kim?" |
| Reply | "Men shu IP, menga qarang" |
| Gratuitous | O'zini e'lon qilish |

---

## Xulosa

- ARP lokal tarmoqda IP → MAC aniqlash uchun
- Broadcast so'rov, unicast javob
- ARP jadvali cache sifatida saqlanadi
