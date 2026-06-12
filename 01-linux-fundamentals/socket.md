# Socket

Socket -- operatsion tizim bergan aloqa nuqtasi.

---

## Nazariya

TCP socket identifikatori:

```text
Protocol
Source IP
Source Port
Destination IP
Destination Port
```

Misol:

```text
TCP
192.168.1.10
52341
142.250.190.78
443
```

---

## Portlar

```text
0-65535
```

---

## Mashhur portlar

```text
22   SSH
25   SMTP
53   DNS
80   HTTP
443  HTTPS
5432 PostgreSQL
6379 Redis
8080 App
```

---

## Amaliyot

### Portlarni ko'rish

```bash
ss -tulpn
```

yoki

```bash
sudo lsof -i
```

---

## Kod

### Go

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    // Server socket
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Server listening on :8080")

    // Client socket
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Client connected")
}
```

### Python

```python
import socket

# Server socket
server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('', 8080))
server.listen(1)
print("Server listening on :8080")

# Client socket
client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect(('localhost', 8080))
print("Client connected")

client.close()
server.close()
```

---

## Xulosa

- Socket = aloqa nuqtasi
- TCP socket 5 ta elementdan iborat
- Portlar 0-65535 oralig'ida
