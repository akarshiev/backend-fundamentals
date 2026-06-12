# Recursive vs Iterative DNS

DNS so'rovlarining ikki xil turi.

---

## Nazariya

### Recursive DNS

Client (browser) birinchi DNS serverga so'raydi. Server barcha ishni o'zi bajaradi.

```text
Client
   |
   V
DNS Resolver
   |
   V
Root DNS -> TLD DNS -> Authoritative DNS
   |
   V
IP qaytadi
```

Client faqat yakuniy javobni oladi.

### Iterative DNS

DNS server har qadamda clientga javob beradi. Client o'zi so'raydi.

```text
Client -> Root DNS
   |
   V
Client -> TLD DNS
   |
   V
Client -> Authoritative DNS
   |
   V
IP qaytadi
```

---

## Diagram

```text
Recursive:
Client -> Resolver -> [Root -> TLD -> Auth] -> IP

Iterative:
Client -> Root
Client -> TLD
Client -> Auth
Client -> IP
```

---

## Amaliyot

### Recursive query

```bash
dig google.com
```

### Iterative query

```bash
dig google.com +norecurse
```

### Root DNS serverlarni ko'rish

```bash
dig . NS
```

---

## Xulosa

- Recursive: Client faqat javob oladi, server hamma ishni bajaradi
- Iterative: Client har qadamda o'zi so'raydi
- ISP'lar odatda recursive DNS ishlatadi
