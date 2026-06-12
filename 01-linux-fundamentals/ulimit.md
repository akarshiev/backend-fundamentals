# ulimit

Linux'da ochiq fayl/socket sonini boshqarish uchun ishlatiladi.

---

## Nazariya

`ulimit` buyrug'i process uchun resurs limitlarini ko'rsatadi va o'zgartiradi.

Asosiy parametrlar:

```text
-n  Open files limiti (FD limit)
-u  Max user processes
-s  Stack size
-c  Core file size
```

---

## Amaliyot

### Hozirgi limitni ko'rish

```bash
ulimit -n
```

Natija:

```text
1024
```

Demak, 1024 ta ochiq file/socket dan ko'p bo'lmaydi.

### Vaqtincha o'zgartirish

```bash
ulimit -n 65535
```

### Doimiy o'zgartirish

`/etc/security/limits.conf` fayliga qo'shing:

```text
* soft nofile 65535
* hard nofile 65535
```

### Tekshirish

```bash
ulimit -a
```

---

## Diagram

```text
Process

   |
   V

+------------------+
| FD Limit         |
| soft: 1024       |
| hard: 65535      |
+------------------+

   |
   V

Kernel

   |
   V

Resources
- Files
- Sockets
- Pipes
```

---

## Xulosa

- `ulimit -n` -- ochiq fayl limiti
- Serverlar uchun 65535 ga ko'tarish kerak
- Doimiy o'zgartirish uchun `/etc/security/limits.conf` ishlatiladi
