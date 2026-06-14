# 10 - QUIC va HTTP/3

TCP Head Of Line Blocking muammosi va yangi avlod protocoli.

---

## Mavzular

- [TCP Head Of Line Blocking](hol-blocking.md)
- [HTTP Versiyalar Evolyutsiyasi](http-evolution.md)
- [QUIC Protocol](quic.md)
- [HTTP/3](http3.md)
- [Connection Migration](connection-migration.md)

---

## Qisqacha

| Mavzu | Tushuntirish |
|-------|--------------|
| HOL Blocking | TCP paket ketma-ketligi buzilganda barcha stream to'xtaydi |
| HTTP/1.1 | Multiplexing yo'q, har request uchun navbat |
| HTTP/2 | Multiplexing bor, lekin TCP sabab HOL qolgan |
| QUIC | UDP ustida, user-space, stream-level HOL yo'q |
| HTTP/3 | QUIC ustida, built-in TLS 1.3, connection migration |
| Connection Migration | IP o'zgarsa ham connection davom etadi |

---

## Amaliyot

### HTTP versiyasini tekshirish

```bash
curl -I --http3 https://google.com
```

### HTTP/2 tekshirish

```bash
curl -I --http2 https://google.com
```

### Socketlarni ko'rish

```bash
# UDP socketlar
ss -u

# TCP socketlar
ss -t
```

### Wireshark

Wireshark orqali UDP, QUIC, TLS paketlarini kuzatishingiz mumkin.
