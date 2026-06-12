# TTL (Time To Live)

DNS cache necha soniya saqlanishini belgilaydi.

---

## Nazariya

TTL = DNS record cache'da qoladigan vaqt (soniyalarda).

Misol:

```text
example.com A 93.184.216.34 TTL=3600
```

Bu degani, DNS resolver bu recordni 3600 soniya (1 soat) cache'da saqlaydi.

---

## Diagram

```text
Resolver

   |
   V

+------------------+
| Cache            |
| example.com      |
| 93.184.216.34    |
| TTL: 3600        |
+------------------+

   |
   V

3600 soniya o'tgach

   |
   V

Cache'dan o'chiriladi
```

---

## Amaliyot

### TTL ni ko'rish

```bash
dig example.com | grep -i ttl
```

yoki

```bash
dig +noall +answer example.com
```

### TTL ni o'zgartirish

DNS provayderning panelida o'zgartiring.

Masalan, Cloudflare'da:

```text
DNS -> Records -> TTL
```

---

## Xulosa

- TTL = cache saqlash vaqti
- Kam TTL = ko'proq DNS query, lekin tez o'zgarish
- Ko'p TTL = kamroq DNS query, lekin sekin o'zgarish
