# Firewall (Layer 3/4 Himoya)

Tarmoq darajasidagi qo'riqchi. IP va TCP/UDP portlarni nazorat qiladi.

---

## Firewall nima?

```text
Firewall — tarmoq darajasidagi qo'riqchi.

OSI model:
Layer 3 → IP
Layer 4 → TCP / UDP

bilan ishlaydi.
```

---

## Portlarni nazorat qiladi

| Port | Protokol | Xizmat |
|------|----------|--------|
| 22 | TCP | SSH |
| 80 | TCP | HTTP |
| 443 | TCP | HTTPS |
| 5432 | TCP | PostgreSQL |
| 6379 | TCP | Redis |

---

## Firewall nima muammoni hal qiladi?

### Masalan: SSH brute force

```text
SSH → port 22 internetga ochiq.

Botlar:
- root
- admin
- ubuntu
- test

username'lar bilan bruteforce qiladi.
```

### Yechim: Firewall qoidalari

```text
Firewall:
"faqat mening IP'mdan ruxsat ber" deyishi mumkin.
```

---

## ⚠️ Xato: `listen-address="*"`

```text
listen-address="*" — bu juda ham katta xato!

Bu degani:
5432 port: "hey, men hamma uchun ochiman!"

Buni hacker tez aniqlab brute force attack qiladi.
```

### To'g'ri sozlash

```text
# Xato ❌
listen-address="*"

# To'g'ri ✅
listen-address="127.0.0.1"
```

---

## Linux Firewall (UFW)

### Ulanishlarni ko'rish

```bash
ss -tulpn
```

### Firewall holatini tekshirish

```bash
sudo ufw status
```

### SSH ruxsat berish

```bash
sudo ufw allow 22/tcp
```

### HTTP ruxsat berish

```bash
sudo ufw allow 80/tcp
```

### HTTPS ruxsat berish

```bash
sudo ufw allow 443/tcp
```

### Barcha kiruvchi trafikni bloklash

```bash
sudo ufw default deny incoming
```

### Barcha chiquvchi trafikni ruxsat berish

```bash
sudo ufw default allow outgoing
```

### Firewall ni yoqish

```bash
sudo ufw enable
```

### Qoidani o'chirish

```bash
sudo ufw delete allow 22/tcp
```

---

## Iptables (kengaytirilgan)

```bash
# Kiruvchi SSH trafikni cheklash (10 soatda 6 ta urinish)
sudo iptables -A INPUT -p tcp --dport 22 -m recent --set --name SSH
sudo iptables -A INPUT -p tcp --dport 22 -m recent --update --seconds 3600 --hitcount 6 --name SSH -j DROP

# Bitta IP dan kirishni bloklash
sudo iptables -A INPUT -s 1.2.3.4 -j DROP

# Port 80 uchun rate limiting
sudo iptables -A INPUT -p tcp --dport 80 -m limit --limit 100/minute --limit-burst 200 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 80 -j DROP
```

---

## Amaliyot

### Portlarni tekshirish

```bash
# Ochiq portlarni ko'rish
ss -tulpn

# Natijada:
# LISTEN  0  128  0.0.0.0:22    0.0.0.0:*    users:(("sshd"))
# LISTEN  0  128  0.0.0.0:80    0.0.0.0:*    users:(("nginx"))
# LISTEN  0  128  127.0.0.1:5432 0.0.0.0:*    users:(("postgres"))
```

### Firewall qoidalari

```bash
# Qoidalarni ko'rish
sudo ufw status numbered

# Natija:
# [ 1] 22/tcp    ALLOW IN    Anywhere
# [ 2] 80/tcp    ALLOW IN    Anywhere
# [ 3] 443/tcp   ALLOW IN    Anywhere
```

---

## Xulosa

- Firewall tarmoq darajasidagi qo'riqchi
- Layer 3 (IP) va Layer 4 (TCP/UDP) bilan ishlaydi
- Portlarni nazorat qiladi
- UFW va iptables Linux uchun asosiy vositalar
- `listen-address="*"` xato — faqat kerakli IP larni ruxsat bering
