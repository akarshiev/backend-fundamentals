# SSH Tunneling

SSH orqali xavfsiz tunnel yaratish. Port forwarding bilan xavfsiz ulanish.

---

## SSH Tunnel nima?

```text
SSH Tunnel = SSH orqali boshqa portlarga yo'l ochish

Masalan:
- Remote database ga xavfsiz ulanish
- Public Wi-Fi da trafikni shifrlash
- Boshqa serverlarga yashirin kirish
```

---

## SSH Tunnel turlari

| Tur | Buyruq | Vazifa |
|-----|--------|--------|
| Local | `ssh -L` | Local port → Remote server |
| Remote | `ssh -R` | Remote port → Local server |
| Dynamic | `ssh -D` | SOCKS proxy yaratish |

---

## Local Port Forwarding (`ssh -L`)

### Oddiy misol

```bash
# Local port 5432 → Remote database 5432
ssh -L 5432:localhost:5432 user@remote-server
```

### Qanday ishlaydi?

```text
Local:5432 → SSH Tunnel → Remote:5432

1. SSH tunnel ochish:
   ssh -L 5432:db-server:5432 user@jump-server

2. Local dan ulanish:
   psql -h localhost -p 5432 -U myuser mydb

3. Natija:
   Local:5432 → Jump Server → DB Server:5432
```

### Redis uchun tunnel

```bash
# Local port 6379 → Remote Redis 6379
ssh -L 6379:localhost:6379 user@redis-server

# Redis CLI dan ulanish
redis-cli -h localhost -p 6379
```

---

## Remote Port Forwarding (`ssh -R`)

### Oddiy misol

```bash
# Remote port 8080 → Local port 80
ssh -R 8080:localhost:80 user@remote-server
```

### Qanday ishlaydi?

```text
Remote:8080 → SSH Tunnel → Local:80

1. SSH tunnel ochish:
   ssh -R 8080:localhost:80 user@remote-server

2. Remote server dan ulanish:
   curl http://localhost:8080

3. Natija:
   Remote:8080 → SSH Tunnel → Local:80
```

### Web server ni publish qilish

```bash
# Local web server ni remote ga publish
ssh -R 8080:localhost:3000 user@vps-server

# Endi http://vps-server:8080 → localhost:3000
```

---

## Dynamic Port Forwarding (`ssh -D`)

### SOCKS proxy yaratish

```bash
# SOCKS proxy ochish
ssh -D 1080 user@remote-server
```

### Brauzer sozlash

```text
1. Brauzer → Settings → Network → Proxy
2. SOCKS Host: localhost
3. Port: 1080
4. Barcha trafik SSH orqali o'tadi
```

---

## SSH Config sozlash

### `~/.ssh/config`

```text
# Jump Server
Host jump
    HostName jump.example.com
    User admin
    Port 22

# Database tunnel
Host db-tunnel
    HostName jump.example.com
    User admin
    LocalForward 5432 db-server:5432

# Web server tunnel
Host web-tunnel
    HostName vps.example.com
    User root
    RemoteForward 8080 localhost:3000
```

### Ishlatish

```bash
# Config dan foydalanish
ssh db-tunnel

# Avtomatik tunnel ochiladi
```

---

## SSH o'girish (ProxyJump)

```text
Internet → Jump Server → Private Server

1. SSH config:
Host private-server
    HostName 10.0.1.100
    User admin
    ProxyJump jump-server

2. Ishlatish:
ssh private-server

# Avtomatik: Internet → Jump Server → Private Server
```

---

## Amaliyot

### Tunnel test

```bash
# SSH tunnel ochish
ssh -L 5432:localhost:5432 user@jump-server

# Boshqa terminal da test
psql -h localhost -p 5432 -U myuser mydb
```

### Tunnel holatini tekshirish

```bash
# Faol ulanishlarni ko'rish
ss -tulpn | grep 5432

# yoki
netstat -tulpn | grep 5432
```

### Tunnel to'xtatish

```bash
# SSH session dan chiqish
exit

# yoki
# Ctrl + D
```

---

## Xulosa

- SSH Tunnel orqali xavfsiz ulanish
- Local (`-L`), Remote (`-R`), Dynamic (`-D`) turlari bor
- SSH Config bilan oson boshqarish
- ProxyJump orqali bir necha serverdan o'tish
