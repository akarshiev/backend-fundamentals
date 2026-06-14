# 09 - Networking (Tarmoq asoslari)

IP address, NAT, portlar va OSI modeli.

---

## Mavzular

- [IP Address (Public va Private)](ip-address.md)
- [NAT - Network Address Translation](nat.md)
- [Portlar](ports.md)
- [Port Forwarding](port-forwarding.md)
- [Dynamic va Static IP](dynamic-static-ip.md)
- [OSI Model](osi-model.md)

---

## Qisqacha

| Mavzu | Tushuntirish |
|-------|--------------|
| IP Address | Qurilmalarning internetdagi manzili |
| Public IP | Internetda yagona, ISP tomonidan beriladi |
| Private IP | Lokal tarmoqda ishlatiladi |
| NAT | Private IP'larni bitta Public IP ortiga yashiradi |
| Port | Dasturni aniqlash uchun raqam |
| Port Forwarding | Tashqaridan internal serverga kirish |
| Dynamic IP | Vaqtinchalik IP, o'zgarishi mumkin |
| Static IP | Doimiy IP, serverlar uchun |
| OSI Model | Tarmoqni 7 qatlamga bo'lish |

---

## Amaliyot

### Private IP ko'rish

```bash
ifconfig | grep inet
```

### Public IP ko'rish

```bash
curl ifconfig.me
```

### Portlarni ko'rish

```bash
ss -tulpn
```

### NAT kuzatish

```bash
traceroute google.com
```

### Local server ishga tushirish

```bash
python3 -m http.server 8080
curl localhost:8080
```

---

## Keyingi qadam

10-darsga o'ting yoki amaliyot bajaring.
