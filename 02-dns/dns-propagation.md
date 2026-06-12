# DNS Propagation

DNS o'zgarishlari butun internetga tarqalishi uchun vaqt kerak.

---

## Nazariya

DNS record o'zgarganda, barcha resolverlar darhol yangilanmaydi. TTL vaqtiga qarab kechikish bo'ladi.

Misol:

```text
DNS o'zgardi
   |
   V
Resolver 1 (cache'da): 3600 soniya qoldi
Resolver 2 (cache'da): 1800 soniya qoldi
Resolver 3 (yangilandi): darhol yangilandi
```

---

## Diagram

```text
DNS Server
   |
   V
+------------------+
| TTL: 3600        |
+------------------+

   |
   V

Barcha resolverlar

Resolver 1: 3600 soniya kutadi
Resolver 2: 1800 soniya kutadi
Resolver 3: 0 soniya (yangilandi)
```

---

## Amaliyot

### DNS propagation tekshirish

```bash
# Turli DNS serverlardan so'rang
dig @8.8.8.8 example.com
dig @1.1.1.1 example.com
dig @208.67.222.222 example.com
```

### Online tekshirish

```text
whatsmydns.net
```

---

## Xulosa

- DNS propagation TTL vaqtiga qarab kechikadi
- Kam TTL = tezroq tarqaladi
- Ko'p TTL = sekinroq tarqaladi
