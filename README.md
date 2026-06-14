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
