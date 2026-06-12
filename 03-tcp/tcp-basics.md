# TCP Basics

TCP (Transmission Control Protocol) -- ishonchli aloqa protokoli.

---

## Nazariya

TCP ikki qurilma o'rtasida ishonchli aloqa o'rnatadi.

Asosiy xususiyatlar:

```text
- Ishonchli: paketlar yetib kelishini ta'minlaydi
- Tartibli: paketlar to'g'ri tartibda keladi
- Xatolarni tuzatadi: yo'qolgan paketlarni qayta yuboradi
```

---

## Client va Server Socket

### Client Socket

```go
conn, _ := net.Dial("tcp", "localhost:8080")
```

### Server Socket

```go
listener, _ := net.Listen("tcp", ":8080")
conn, _ := listener.Accept()
```

---

## Diagram

```text
Client            Server

Connect --------> Listen
     <---------- Accept

Data  ---------> Data
     <--------- Data
```

---

## Amaliyot

### TCP server yaratish

```bash
nc -l 8080
```

### TCP client

```bash
nc localhost 8080
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
    // Server
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    fmt.Println("Server listening on :8080")

    conn, err := listener.Accept()
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Received:", string(buffer[:n]))
}
```

### Python

```python
import socket

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('', 8080))
server.listen(1)
print("Server listening on :8080")

conn, addr = server.accept()
print(f"Connection from {addr}")

data = conn.recv(1024)
print(f"Received: {data.decode()}")

conn.close()
server.close()
```

---

## Xulosa

- TCP ishonchli aloqa protokoli
- Client socket: net.Dial
- Server socket: net.Listen + Accept
