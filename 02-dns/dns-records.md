# DNS Records

DNS faqat IP bermaydi. Unda turli xil recordlar mavjud.

---

## A Record

IPv4 manzil.

Misol:

```text
example.com -> 93.184.216.34
```

DNS:

```text
example.com IN A 93.184.216.34
```

A = Address. IPv4 uchun ishlatiladi.

---

## AAAA Record

IPv6 manzil.

Misol:

```text
example.com -> 2606:2800:220:1:248:1893:25c8:1946
```

DNS:

```text
example.com IN AAAA 2606:2800:220:1:248:1893:25c8:1946
```

Nega AAAA? A record bo'lgani uchun IPv6 uchun AAAA tanlangan.

---

## CNAME Record

Canonical Name. Alias yaratadi.

Misol:

```text
www.example.com -> example.com
```

DNS:

```text
www.example.com IN CNAME example.com
```

Jarayon:

```text
www.example.com
       |
       V
example.com
       |
       V
93.184.216.34
```

---

## MX Record

Mail Exchange. Email qaysi serverga borishini aytadi.

Misol:

```text
company.com MX 10 mail1.company.com
company.com MX 20 mail2.company.com
```

10 ustuvor. Agar ishlamasa 20 ishlatiladi.

---

## TXT Record

Oddiy matn saqlaydi. Lekin amalda juda muhim.

### SPF

Qaysi server email yuborishi mumkin:

```text
v=spf1 include:_spf.google.com ~all
```

### Domain Verification

Google domenni tasdiqlash uchun:

```text
google-site-verification=abc123
```

### DKIM

Email imzosi:

```text
v=DKIM1; k=rsa; ...
```

### DMARC

Email spoofingdan himoya:

```text
v=DMARC1; p=reject;
```

---

## Amaliyot

### A record olish

```bash
dig example.com A
```

### AAAA record olish

```bash
dig example.com AAAA
```

### CNAME record olish

```bash
dig www.example.com CNAME
```

### MX record olish

```bash
dig example.com MX
```

### TXT record olish

```bash
dig example.com TXT
```

---

## Xulosa

| Record | Vazifasi                                         |
| ------ | ------------------------------------------------ |
| A      | Domain -> IPv4                                   |
| AAAA   | Domain -> IPv6                                   |
| CNAME  | Domain alias                                     |
| MX     | Email server                                     |
| TXT    | SPF, DKIM, DMARC, verification va boshqa matnlar |
