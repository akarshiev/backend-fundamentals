# Head Of Line Blocking (HOL)

TCP tartibni saqlaydi. Bu ba'zan muammo tug'diradi.

---

## Nazariya

TCP barcha paketlarni tartibda yetkazishi kerak. Agar bitta paket kechiksa, keyingilarini kutish kerak.

### Misol

```text
Paket 1 --> Yetib keldi ✓
Paket 2 --> Yetib keldi ✓
Paket 3 --> Yo'qoldi ✗
Paket 4 --> Yetib keldi, lekin 3 kutmoqda
Paket 5 --> Yetib keldi, lekin 3 kutmoqda
```

### Natija

TCP 3 kelmaguncha 4 va 5 ni berolmaydi. Bu:

```text
Head Of Line Blocking
```

muammosi deyiladi.

---

## Diagram

```text
Sender                              Receiver

P1 ------------------------------> P1 ✓ (app)
P2 ------------------------------> P2 ✓ (app)
P3 --X (yo'qoldi)
P4 ------------------------------> P4 ⏳ (kutmoqda)
P5 ------------------------------> P5 ⏳ (kutmoqda)
     (P3 yo'q, P4 va P5 blocklangan)

     <--- P3 ni qayta yuborish so'rovi ---
P3 ------------------------------> P3 ✓ (app)
                                  P4 ✓ (app)
                                  P5 ✓ (app)
```

---

## HTTP/2 da HOL

HTTP/2 da bu juda mashhur muammo:

```text
HTTP/2 Stream 1: /index.html
HTTP/2 Stream 2: /style.css
HTTP/2 Stream 3: /app.js

Agar /style.css paketi yo'qolsa:
- /index.html kutmoqda
- /app.js kutmoqda
- Hamma stream blocklangan
```

---

## HTTP/3 yechimi

HTTP/3 QUIC protokolini ishlatadi. QUIC UDP asosida ishlaydi va har bir stream mustaqil:

```text
QUIC Stream 1: /index.html  --> mustaqil
QUIC Stream 2: /style.css   --> mustaqil
QUIC Stream 3: /app.js      --> mustaqil

/style.css yo'qolsa ham:
- /index.html davom etadi
- /app.js davom etadi
```

---

## Amaliyot

### HOL ni tekshirish

```bash
sudo tcpdump -i any port 443 -nn
```

### HTTP/2 vs HTTP/3

```bash
# HTTP/2
curl -v --http2 https://example.com

# HTTP/3
curl -v --http3 https://example.com
```

---

## Xulosa

- HOL -- TCP tartib saqlashning kamchiligi
- Bitta paket kechiksa, barcha keyingilar kutadi
- HTTP/2 da juda sezilarli muammo
- HTTP/3 (QUIC) bu muammoni hal qiladi
