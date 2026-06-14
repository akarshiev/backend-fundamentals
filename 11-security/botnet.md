# Botnet

Zararlangan qurilmalar to'plami — DDoS hujumi uchun asosiy qurol.

---

## Tushuntirish

```text
10,000 ta kompyuter
    ↓
Zararlangan (malware orqali)
    ↓
Botnet (hujumchi nazoratida)
    ↓
DDoS hujumi
```

---

## Botnet tuzilishi

```text
Attacker (C&C Server)
    ↓
Command & Control
    ↓
Bot 1  (zombie)
Bot 2  (zombie)
Bot 3  (zombie)
...
Bot N  (zombie)
    ↓
Target Server
```

---

## Botnet yaratish jarayoni

```text
1. Malware yaratish
   ↓
2. Foydalanuvchiga yuborish (phishing, drive-by)
   ↓
3. Qurilmaga o'rnatish
   ↓
4. C&C serverga ulanish
   ↓
5. Buyruq kutish
   ↓
6. DDoS boshlash
```

---

## Botnet turlari

| Tur | Xususiyati |
|-----|------------|
| IRC Botnet | Eski, IRC orqali boshqariladi |
| HTTP Botnet | Web orqali boshqariladi |
| P2P Botnet | Markazsiz, qiyin yo'q qilish |
| IoT Botnet | Router, kamera, smart qurilmalar |

---

## Mashhur botnetlar

| Nom | Yil | Xususiyat |
|-----|-----|-----------|
| Mirai | 2016 | IoT botnet, 600 Gbps |
| Emotet | 2018 | Banking trojan |
| TrickBot | 2020 | Banking, ransomware |
| RapperBot | 2023 | SSH/DLL botnet |

---

## Himoya

### 1. Endpoint Protection

```text
- Antivirus / EDR
- Behavior monitoring
- Application control
```

### 2. Network Security

```text
- Firewall
- IDS/IPS
- Traffic analysis
```

### 3. User Awareness

```text
- Phishing training
- Software updates
- Strong passwords
```

---

## Xulosa

- Botnet = zararlangan qurilmalar guruhi
- C&C server orqali boshqariladi
- DDoS, spam, ma'lumot o'g'irlash uchun ishlatiladi
- Endpoint protection va foydalanuvchi bilimi bilan himoya
