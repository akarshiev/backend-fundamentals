# 11 - Xavfsizlik asoslari (Security Fundamentals)

Shifrlash, raqamli imzo, sertifikatlar, tarmoq xavfsizligi va hujum turlari.

---

## Mavzular

### Kriptografiya
- [Kriptografiya asoslari](cryptography.md) -- simmetrik va assimmetrik shifrlash
- [Raqamli imzo](digital-signature.md) -- xabarni imzolash va tekshirish
- [Certificate Authority](certificate-authority.md) -- HTTPS ishonch zanjiri
- [OpenSSL amaliyoti](openssl.md) -- amaliy buyruqlar

### Tarmoq xavfsizligi
- [ARP](arp.md) -- Address Resolution Protocol
- [MAC Address](mac-address.md) -- Layer 2 identifikator
- [ARP Spoofing](arp-spoofing.md) -- hujum va himoya
- [MITM](mitm.md) -- Man In The Middle
- [DNS Spoofing](dns-spoofing.md) -- DNS soxtalashtirish
- [SSL Stripping](ssl-stripping.md) -- HTTPS ni HTTP ga tushirish

### Hujum turlari
- [DoS va DDoS](dos-ddos.md) -- xizmat ko'rsatishni bloklash
- [SYN Flood](syn-flood.md) -- TCP handshake hujumi
- [UDP Flood](udp-flood.md) -- UDP paketlar hujumi
- [Slowloris](slowloris.md) -- sekin connection hujumi
- [Botnet](botnet.md) -- zararlangan qurilmalar to'plami

### Himoya
- [Reverse Proxy](reverse-proxy.md) -- Nginx, Cloudflare
- [Amaliy xavfsizlik](security-practices.md) -- HTTPS, certificate, traffic monitoring

### Amaliy hujumlar
- [Hujum amaliyotlari](attacks-practical.md) -- o'rganish uchun hujum kodlari va yo'riqnomalar

---

## Qisqacha

| Mavzu | Tushuntirish |
|-------|--------------|
| Kriptografiya | Ma'lumotni shifrlash |
| Simmetrik | Bir xil kalit |
| Assimmetrik | Ikki kalit (public/private) |
| Digital Signature | Xabarni imzolash |
| CA | Sertifikat beruvchi tashkilot |
| ARP | IP → MAC aniqlash |
| MITM | Hujumchi orasida turib |
| DoS/DDoS | Serverni bloklash |
| SYN Flood | TCP handshake hujumi |
| Slowloris | Sekin HTTP connection |
| Reverse Proxy | Himoya qatlam |

---

## O'rganish tartibi

```text
Kriptografiya
-> Simmetrik / Assimmetrik
-> Digital Signature
-> Certificate Authority
-> OpenSSL
-> ARP / MAC
-> ARP Spoofing
-> MITM
-> DNS Spoofing
-> SSL Stripping
-> DoS / DDoS
-> SYN Flood
-> UDP Flood
-> Slowloris
-> Reverse Proxy
-> Himoya usullari
```

---

## ⚠️ Ogohlantirish

**Ushbu materiallar FAQAT ta'lim maqsadida yozilgan.**

Hujum kodlari va yo'riqnomalar:
- Faqat o'z tarmog'ingizda sinash uchun
- Ruxsatsiz kirish NOJOYI -- jinoyat
- Boshqa tizimlarga hujum qilish qonunan taqiqlangan
- Faqat o'rganish va himoya qilishni tushunish uchun

---

## How to use

1. Kriptografiyadan boshlang
2. Nazariyani o'qing
3. Amaliy buyruqlarni sinang
4. Hujum kodlarini o'z tarmog'ingizda sinang
5. Himoya usullarini o'rganing
