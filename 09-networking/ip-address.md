# IP Address

## Nega IP Address kerak?

Internetdagi har bir qurilma bir-birini topishi uchun manzil kerak.

Bu manzil:

**IP Address**

deyiladi.

Misol:

```
8.8.8.8
```

yoki

```
192.168.1.10
```

---

## Public va Private IP

IP addresslar ikkiga bo'linadi:

1. **Public IP**
2. **Private IP**

---

## Private IP

Uy ichidagi lokal tarmoqda ishlatiladi.

### Mashhur diapazonlar

| Diapazon | CIDR |
|----------|------|
| `10.0.0.0` | `10.0.0.0/8` |
| `172.16.0.0 - 172.31.255.255` | `172.16.0.0/12` |
| `192.168.0.0` | `192.168.0.0/16` |

### Misollar

```
192.168.1.10
192.168.1.11
192.168.1.12
```

---

## Public IP

Internetda yagona bo'lishi kerak.

### Misol

```
84.54.12.8
95.46.200.1
```

Public IP orqali sizni internet topadi.

---

## Nega aynan Public va Private?

Internet boshida har qurilmaga Public IP berish rejalashtirilgan.

Lekin muammo paydo bo'ldi:

- **IPv4**: 32 bit
- **≈ 4.3 milliard IP**
- Internet esa milliardlab qurilmalarga yetmay qoldi.

Shuning uchun:

**Private IP + NAT**

o'ylab topildi.

---

## Amaliyot

### Private IP ko'rish

```bash
ifconfig | grep inet
```

yoki

```bash
ip addr
```

### Public IP ko'rish

```bash
curl ifconfig.me
```
