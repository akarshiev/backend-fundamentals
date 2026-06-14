# TCP Head Of Line Blocking (HOL)

## Muammo

TCP packetlarni qat'iy tartibda yetkazadi.

### Misol

Paketlar shunday ketma-ketlikda kelishi kerak:

```
P1 → P2 → P3 → ... → P26 → P27 → P28 → ... → P100
```

Faraz qilaylik: **P27 yo'qolib qoldi.**

Lekin P28, P29, P30 yetib kelgan.

### TCP javobi

```
"P27 kelmaguncha keyingilarni application'ga bera olmayman"
```

Natija: P28, P29, P30 bekor kutib qoladi.

Bu: **Head Of Line Blocking** deyiladi.

---

## TCP nega shunday?

TCP ning asosiy vazifasi:

- **Reliable** — ishonchli yetkazish
- **Ordered** — tartibni saqlash

Shuning uchun `1 → 2 → 3 → 4 → 5` ketma-ketligini buzolmaydi.

Bu TCP dizaynining bir qismi.

---

## HTTP/1.1 dagi muammo

HTTP/1.1 da:

- Har request uchun **alohida connection** yoki
- Bir connection ichida **navbat**

ishlatilgan.

Natija: **Request 1** tugamaguncha **Request 2** kutadi.

Bu ham HOL muammosini kuchaytirgan.

---

## HTTP/2 yechimi

HTTP/2 **Multiplexing** olib keldi.

Bitta TCP connection ichida bir nechta stream parallel ishlaydi:

```
TCP Connection
├── Stream 1
├── Stream 2
└── Stream 3
```

---

## HTTP/2 nega to'liq yechim emas?

Multiplexing application darajasida ishlaydi.

Lekin tagida **TCP** hali ham ishlayapti.

### Misol

Stream 2 packeti yo'qolsa → TCP packetni qayta kutadi.

Shu paytda Stream 1 va Stream 3 ham to'xtab qolishi mumkin.

### Xulosa

- HTTP/2 HOL muammosini **kamaytirdi**
- Lekin **yo'qota olmadi**
- Sabab: **TCP**
