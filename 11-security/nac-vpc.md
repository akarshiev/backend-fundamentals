# NAC va VPC

Tarmoq kirishni boshqarish va xususiy tarmoq yaratish.

---

## NAC (Network Access Control)

### NAC nima?

```text
NAC = Network Access Control

Qurilmalarning tarmoqqa kirishini boshqaradi.

Masalan:
- DB serverlarni yashirish
- Faqat kerakli qurilmalarni ruxsat berish
```

---

## ⚠️ Xato: Barcha portlarni ochish

```text
listen-address="*" — bu juda ham katta xato!

Bu degani:
5432 port: "hey, men hamma uchun ochiman!"

Buni hacker tez aniqlab brute force attack qiladi.
```

### To'g'ri sozlash

```text
# Xato ❌
listen-address="*"

# To'g'ri ✅
listen-address="127.0.0.1"
```

---

## Public va Private Zone

```text
Internet
   ↓
Public Zone (DMZ)
   - Web server
   - API server
   - Load Balancer
   ↓
Private Zone
   - Database
   - Cache
   - Internal services
```

### Public Zone

| Xizmat | Port | Tavsif |
|--------|------|--------|
| Web Server | 80, 443 | HTTP/HTTPS |
| API Server | 8080, 8443 | REST API |
| Load Balancer | 80, 443 | Trafik taqsimlash |

### Private Zone

| Xizmat | Port | Tavsif |
|--------|------|--------|
| PostgreSQL | 5432 | Database |
| Redis | 6379 | Cache |
| MongoDB | 27017 | NoSQL Database |

---

## VPC (Virtual Private Cloud)

### VPC nima?

```text
VPC = Virtual Private Cloud

Xususiy tarmoq yaratish imkoniyati.

Cloud da:
- AWS VPC
- Google Cloud VPC
- Azure VPC

On-premise:
- VLAN
- VPN
- Firewall
```

### VPC拓撲

```text
Internet Gateway
       ↓
   Public Subnet
   - Web Server (10.0.1.0/24)
   - Bastion Host (10.0.1.0/24)
       ↓
   Private Subnet
   - App Server (10.0.2.0/24)
       ↓
   Database Subnet
   - PostgreSQL (10.0.3.0/24)
   - Redis (10.0.3.0/24)
```

---

## Private Network yaratish

### 1. VLAN (On-premise)

```bash
# VLAN yaratish
sudo ip link add link eth0 name eth0.100 type vlan id 100
sudo ip addr add 10.0.100.1/24 dev eth0.100
sudo ip link set dev eth0.100 up
```

### 2. VPN (Remote access)

```bash
# OpenVPN server
sudo apt install openvpn

# Client config
client
dev tun
proto udp
remote vpn.example.com 1194
```

### 3. AWS VPC

```bash
# VPC yaratish
aws ec2 create-vpc --cidr-block 10.0.0.0/16

# Subnet yaratish
aws ec2 create-subnet --vpc-id vpc-xxx --cidr-block 10.0.1.0/24

# Security Group
aws ec2 create-security-group --group-name my-sg --description "My SG"
```

---

## SSH Tunneling

Batafsil: [SSH Tunneling](ssh-tunneling.md)

```text
SSH Tunnel = SSH orqali xavfsiz tunnel yaratish

ssh -L 5432:localhost:5432 user@jump-server
```

---

## Bastion Host (Jump Server)

```text
Internet → Bastion Host → Private Network

Bastion Host:
- Faqat SSH ruxsat
- Boshqa portlar bloklangan
- Log yoziladi
```

### Sozlash

```bash
# Bastion Host da
sudo ufw default deny incoming
sudo ufw allow 22/tcp
sudo ufw enable

# SSH config
# /etc/ssh/sshd_config
PermitRootLogin no
PasswordAuthentication no
```

---

## Amaliyot

### Portlarni tekshirish

```bash
# Ochiq portlarni ko'rish
ss -tulpn

# Natijada:
# LISTEN  0  128  127.0.0.1:5432  0.0.0.0:*  users:(("postgres"))
# LISTEN  0  128  127.0.0.1:6379  0.0.0.0:*  users:(("redis-server"))
```

Batafsil: [SSH Tunneling](ssh-tunneling.md)

---

## Xulosa

- NAC tarmoq kirishini boshqaradi
- Public va Private zone larni ajratish kerak
- VPC xususiy tarmoq yaratish imkoniyati
- SSH tunneling orqali xavfsiz ulanish
- Bastion Host orqali private network ga kirish
