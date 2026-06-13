# Congestion Control

Tarmoqni himoya qiladi. Internetni tiqilib qolishidan saqlash.

---

## Nazariya

TCP packet loss yoki RTT oshishini **congestion** deb qabul qiladi.

### Congestion nimani anglatadi?

```text
Agar paket yo'qolsa:
- Tarmoq tiqilib qolgan bo'lishi mumkin
- Router buffer to'lib qolgan bo'lishi mumkin
- Boshqa clientlar ham yuklanayotgan bo'lishi mumkin
```

### TCP qanday javob beradi?

```text
1. Yuborish tezligini kamaytiradi
2. Congestion window'ni kichraytiradi
3. Sekin-asta qayta oshiradi
```

---

## Congestion Window (cwnd)

TCP har bir ulanish uchun **congestion window** saqlaydi. Bu server qancha paketni bir vaqtda yubora olishini belgilaydi.

```text
cwnd = 1   --> 1 paket yuborish mumkin
cwnd = 4   --> 4 paket yuborish mumkin
cwnd = 10  --> 10 paket yuborish mumkin
```

---

## Diagram

```text
Sender                              Network

cwnd = 1  ----------------------> [Router] --> Receiver
                                (1 paket o'tdi)
cwnd = 2  ----------------------> [Router] --> Receiver
                                (2 paket o'tdi)
cwnd = 4  ----------------------> [Router] --> Receiver
                                (4 paket o'tdi)
cwnd = 8  ----------------------> [Router] --> X (congestion!)
                                (packet loss)
cwnd = 1  ----------------------> [Router] --> Receiver
                                (qayta boshladi)
```

---

## Congestion Algorithms

### Tahoe

Eng qadimgi algoritm:

```text
1. Slow start
2. Congestion detected (3 duplicate ACK)
3. cwnd = 1
4. Slow start qayta
```

### Reno

Tahoe'ning yaxshilangan versiyasi:

```text
1. Slow start
2. 3 duplicate ACK
3. cwnd = cwnd / 2 (half)
4. Fast recovery
```

### CUBIC

Hozirgi Linux standart algoritmi:

```text
1. Cubic function orqali cwnd oshiradi
2. RTT ga bog'liq emas
3. Yuqori tezlikda yaxshi ishlaydi
```

### BBR

Google tomonidan yaratilgan:

```text
1. RTT va bandwidth ni o'lchaydi
2. Optimal tezlikni topadi
3. Packet loss ga kamroq bog'liq
```

---

## Amaliyot

### Hozirgi algoritmni ko'rish

```bash
sysctl net.ipv4.tcp_congestion_control
```

### Algoritmni o'zgartirish

```bash
# BBR o'rnatish
sudo sysctl -w net.ipv4.tcp_congestion_control=bbr

# CUBIC (default)
sudo sysctl -w net.ipv4.tcp_congestion_control=cubic
```

### TCP metrics ko'rish

```bash
ss -ti
```

---

## Xulosa

- Congestion Control tarmoqni tiqilib qolishdan saqlaydi
- cwnd orqali boshqariladi
- CUBIC -- Linux default algoritmi
- BBR -- Google'ning zamonaviy algoritmi
