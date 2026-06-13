# TCP Flags

TCP header ichida flaglar mavjud. Har bir flag ma'lum bir vazifani bajaradi.

---

## Flags turlari

### SYN (Synchronization)

Connection boshlash. 3-Way Handshake'ning birinchi va ikkinchi paketida ishlatiladi.

```text
Client --> SYN --> Server
```

Sequence number ni sinxronlashtiradi. Ya'ni, qarama-qarshi tomon qaysi sequence number ni qabul qilishi kerakligini bildiradi.

### ACK (Acknowledgement)

Paket qabul qilinganini bildiradi. Acknowledgement number maydoni to'g'ri bo'lsa, bu flag o'rnatiladi.

```text
Server --> SYN + ACK --> Client
```

Ulanish o'rnatilgandan keyin barcha paketlarda ACK flag bo'ladi.

### FIN (Finish)

Connection'ni normal yopish. Sender tomonidan yuboriladi. Zaxiralangan resurslarni bo'shatadi va aloqani shirinroq tugatadi.

```text
Client --> FIN --> Server
     <-- ACK -->
     <-- FIN -->
     <-- ACK -->
```

### RST (Reset)

Connection'ni majburan uzish. Agar TCP aloqada noto'g'ri narsa sezilsa yoki suhbat bo'lmasligi kerak bo'lsa, RST yuboriladi.

```text
Client --> RST --> Server
```

### PSH (Push)

Buffer kutmasdan darhol yuborishni so'raydi. Real-time audio yoki video streaming kabi ilovalarda ishlatiladi.

```text
Client --> PSH + ACK --> Server
```

### URG (Urgent)

Urgent ma'lumotni ko'rsatadi. Urgent Pointer maydoni bilan birgalikda paketdagi urgent ma'lumotning joylashuvini aniqlash uchun ishlatiladi.

```text
Client --> URG + ACK --> Server
```

### WND (Window)

Receive window o'lchamini senderga bildiradi. Qabul qiluvchi host qancha ma'lumotni qabul qila olishini ko'rsatadi. Sender shu o'lchamga qarab ma'lumot yuboradi.

```text
Receiver --> WND=65535 --> Sender
```

### CHK (Checksum)

TCP segment yaxlitligini tekshiradi. Checksum butun segment ustida hisoblanadi (header va data). Har bir hop'da qayta hisoblanadi.

### SEQ (Sequence Number)

Har bir segmentga unikal raqam beriladi. Paketlar qaysi tartibda qabul qilinishi kerakligini aniqlash uchun ishlatiladi.

### ACK Number

Qabul qilingan TCP segmentni tasdiqlash va keyingi kutilayotgan sequence number ni bildirish uchun ishlatiladi.

---

## Flag combinations

### Ulanish o'rnatish

```text
Step 1: SYN=1, ACK=0     --> Clientdan Serverga
Step 2: SYN=1, ACK=1     --> Serverdan Clientga
Step 3: ACK=1             --> Clientdan Serverga
```

### Ulanish yopish

```text
Step 1: FIN=1, ACK=1     --> Clientdan Serverga
Step 2: ACK=1             --> Serverdan Clientga
Step 3: FIN=1, ACK=1     --> Serverdan Clientga
Step 4: ACK=1             --> Clientdan Serverga
```

### Ma'lumot yuborish

```text
PSH=1, ACK=1             --> Darhol yetkazish kerak
ACK=1                     --> Oddiy ma'lumot
```

---

## Diagram

```text
TCP Header (20 bytes minimum)
+--+--+--+--+--+--+--+--+
|CWR|ECE|URG|ACK|PSH|RST|SYN|FIN|
+--+--+--+--+--+--+--+--+
|     Sequence Number     |
+-------------------------+
|  Acknowledgement Number |
+-------------------------+
|Offset| Reserved |Flags  |
+-------------------------+
|      Window Size        |
+-------------------------+
|      Checksum           |
+-------------------------+
|      Urgent Pointer     |
+-------------------------+
```

---

## Amaliyot

### TCP dump bilan flaglarni ko'rish

```bash
sudo tcpdump -i any -nn 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0'
```

### SYN flood tekshirish

```bash
sudo tcpdump -i any 'tcp[tcpflags] & tcp-syn != 0 and tcp[tcpflags] & tcp-ack == 0'
```

---

## Xulosa

- TCP flaglar paket tuzilmasida muhim rol o'ynaydi
- SYN, ACK, FIN -- asosiy flaglar
- RST, PSH, URG -- maxsus holatlar uchun
- WND, CHK, SEQ, ACK -- ishonchni ta'minlaydi
