# Port Forwarding

## Muammo

NAT sabab internet tashqarisidan sizning laptop'ingiz ko'rinmaydi.

Masalan:

- Spring Boot `localhost:8080` ishda turibdi
- Tashqaridan `84.54.12.8:8080` ochilmaydi

---

## Yechim - Port Forwarding

Routerga qoida qo'shadi:

```
84.54.12.8:8080 → 192.168.1.10:8080
```

Shunda internet sizning serveringizga kira oladi.

---

## Qanday qilinadi?

1. Router admin paneliga kiring (odatda `192.168.1.1`)
2. Port Forwarding bo'limini toping
3. Qoida qo'shing:
   - External port: `8080`
   - Internal IP: `192.168.1.10`
   - Internal port: `8080`
   - Protocol: TCP

---

## Nega server bo'la olmaysiz?

Sabablar:

1. **Dynamic IP** - IP doim o'zgaradi
2. **NAT ortidasiz** - tashqaridan ko'rinmaysiz
3. **ISP portlarni yopgan** bo'lishi mumkin
4. **Router port forwarding** qilmagan

---

## Uyda server ishlatish uchun

Kerakli narsalar:

- **Static IP** yoki **DDNS**
- **Port Forwarding**
