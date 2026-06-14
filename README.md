# Backend Fundamentals

Backend dasturlash asoslarini bosqichma-bosqich o'rganish uchun loyiha.

---

## What is this project?

A structured learning path for backend development fundamentals -- from Linux internals to building a CLI browser.

---

## Topics

### 01 Linux Fundamentals
- Everything is a file
- File Descriptors
- Sockets
- ulimit

### 02 DNS
- Domain Name System
- DNS Records (A, AAAA, CNAME, MX, TXT)
- Recursive vs Iterative DNS
- TTL and DNS Propagation

### 03 TCP
- TCP Basics
- Client/Server Sockets
- 3-Way Handshake
- TCP Flags (SYN, ACK, FIN, RST, PSH, URG)
- Sequence & ACK Numbers
- TCP Checksum
- RTT, MTU, MSS, PMTUD
- Flow Control & Congestion Control
- TCP Slow Start
- Head Of Line Blocking
- TCP Fast Open (TFO)
- Buffer
- TCP Packet Structure

### 04 HTTP
- HTTP Request/Response
- Status Codes
- Headers
- Compression (Gzip, Brotli)

### 05 HTTPS
- SSL/TLS
- Certificates
- HTTPS Handshake

### 06 REST API
- REST Principles
- HTTP Methods
- Resource Naming
- Idempotency
- Pagination (Offset, Cursor)

### 07 CLI Browser
- TCP Connection
- HTTP Request/Response
- Header/Body Parsing
- Minimal Browser

### 08 UDP
- UDP Basics
- TCP vs UDP
- DNS Client (Go)
- UDP Server/Client (Go)

### 09 Networking
- IP Address (Public/Private)
- NAT (Network Address Translation)
- Ports
- Port Forwarding
- Dynamic vs Static IP
- OSI Model

### 10 QUIC & HTTP/3
- TCP Head Of Line Blocking
- HTTP Version Evolution (1.1 → 2 → 3)
- QUIC Protocol
- HTTP/3
- Connection Migration

### 11 Security
- Cryptography (Symmetric/Asymmetric)
- Digital Signatures
- Certificate Authority (CA)
- OpenSSL
- ARP & MAC Address
- ARP Spoofing, MITM, DNS Spoofing, SSL Stripping
- DoS / DDoS, SYN Flood, UDP Flood, Slowloris
- Botnet
- Reverse Proxy, Cloudflare
- [Firewall (L3/L4)](11-security/firewall.md) — UFW, iptables, port control
- [WAF (L7)](11-security/waf.md) — SQL Injection, XSS, CSRF
- [NAC & VPC](11-security/nac-vpc.md) — Private/Public zones, DB hiding
- [SSH Tunneling](11-security/ssh-tunneling.md) — Port forwarding, ssh -L/-R/-D
- [OWASP Top 10](11-security/owasp-top10.md) — Common vulnerabilities
- [Security Layers](11-security/security-layers.md) — Defense in depth
- Practical Attack Lab (educational code)

---

## Learning Path

```text
Linux
-> File Descriptor
-> Socket
-> TCP
-> Handshake
-> Flags, Sequence, ACK
-> MTU, MSS, PMTUD
-> Flow Control, Congestion Control
-> UDP
-> DNS Client
-> HTTP
-> HTTPS
-> REST
-> Compression
-> Connection Pooling
-> IP Address, NAT, Ports
-> OSI Model
-> TCP HOL Blocking
-> QUIC & HTTP/3
-> Security Fundamentals
-> Cryptography, Digital Signatures, CA
-> ARP Spoofing, MITM
-> DDoS, SYN Flood, Slowloris
-> Reverse Proxy, Defense
-> Firewall (L3/L4)
-> WAF (L7), SQL Injection
-> NAC, VPC, Private Networks
-> OWASP Top 10
-> Security Layers (Defense in Depth)
```

Each topic builds on the previous one.

---

## How to use

1. Start from 01 and go in order
2. Read the theory in each README
3. Run the example code
4. Try the exercises yourself

---

## Languages

- Go (primary for TCP/HTTP examples)
- Python (alternative examples)
