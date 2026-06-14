# Connection Migration

## TCP muammosi

TCP da IP o'zgarsa **connection uziladi**.

### Misol

```
WiFi → Mobile Internet o'tdingiz
```

TCP reconnect qilishi kerak.

---

## QUIC yechimi

QUIC **Connection ID** ishlatadi.

IP o'zgarsa ham connection davom etishi mumkin.

### Misol

```
WiFi (192.168.1.10) → Mobile (10.0.0.5)
Connection ID: abc123
```

IP o'zgardi, lekin Connection ID bir xil → connection davom etadi.

---

## Qachon foydali?

- WiFi dan mobile internet ga o'tish
- IP o'zgarishi kerak bo'lganda
- Streaming yoki real-time ilovalarda
