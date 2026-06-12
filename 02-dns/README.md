# DNS

Domain Name System -- internetning telefon kitobi.

---

## Mavzular

- DNS nima
- TLD (Top Level Domain)
- DNS Records (A, AAAA, CNAME, MX, TXT)
- Recursive vs Iterative DNS
- TTL
- DNS Propagation

---

## Nazariya

Foydalanuvchi google.com yozadi. Lekin internet IP manzil bilan ishlaydi. Shuning uchun:

```text
Domain -> IP
```

ga aylantirish kerak. Buni DNS bajaradi.

Jarayon:

```text
Browser
   |
   V
DNS Resolver
   |
   V
DNS Server
   |
   V
IP qaytadi
   |
   V
Browser TCP ulanish ochadi
```

---

## Amaliyot

### DNS query

```bash
dig google.com
```

### NS lookup

```bash
nslookup google.com
```

### DNS recordlarni ko'rish

```bash
dig google.com A
dig google.com AAAA
dig google.com MX
dig google.com TXT
```

---

## Xulosa

- DNS domenni IP ga aylantiradi
- Turli xil recordlar mavjud (A, AAAA, CNAME, MX, TXT)
- TTL vaqti bor, shuning uchun DNS kechikishi mumkin
