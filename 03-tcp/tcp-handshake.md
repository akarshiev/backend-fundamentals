# TCP 3-Way Handshake

TCP ulanish o'rnatish jarayoni.

---

## Nazariya

TCP ulanish o'rnatish uchun 3 ta qadam kerak:

```text
1. SYN
2. SYN-ACK
3. ACK
```

---

## Diagram

```text
Client            Server

SYN ------------>
     |            |
     |<-----------| SYN ACK
     |            |
ACK ------------>
     |            |
     V            V
ESTABLISHED
```

---

## Amaliyot

### TCP dump

```bash
sudo tcpdump -i any port 8080
```

### Boshqa terminalda

```bash
nc localhost 8080
```

SYN paketlarini ko'rasiz.

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
    go func() {
        listener, _ := net.Listen("tcp", ":8080")
        conn, _ := listener.Accept()
        defer conn.Close()

        buffer := make([]byte, 1024)
        n, _ := conn.Read(buffer)
        fmt.Println("Server received:", string(buffer[:n]))
    }()

    // Client
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    conn.Write([]byte("Hello TCP!"))
    fmt.Println("Client sent message")
}
```

### Python

```python
import socket
import threading

def server():
    server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server.bind(('', 8080))
    server.listen(1)
    conn, addr = server.accept()
    data = conn.recv(1024)
    print(f"Server received: {data.decode()}")
    conn.close()
    server.close()

def client():
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(('localhost', 8080))
    client.send(b'Hello TCP!')
    print("Client sent message")
    client.close()

# Server va client ni parallel ishga tushiring
server_thread = threading.Thread(target=server)
server_thread.start()

import time
time.sleep(1)  # Server tayyor bo'lishini kutish

client()
server_thread.join()
```

---

## Xulosa

- 3-Way Handshake: SYN -> SYN-ACK -> ACK
- Ulanish ochilgandan keyin ma'lumot yuboriladi
- `tcpdump` orqali paketlarni ko'rish mumkin
