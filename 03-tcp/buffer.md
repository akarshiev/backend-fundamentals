# Buffer

Paketlar darhol o'qilmaydi. Ular buffer'da turadi.

---

## Nazariya

Ma'lumot yuborilganda va qabul qilinganda, u darhol application ga yetib bormaydi. Buffer'da turadi.

### Yuborish tomoni (Send Buffer)

```text
Application
    ↓ (yozadi)
Socket Send Buffer
    ↓ (kernel oladi)
Network
```

### Qabul qilish tomoni (Receive Buffer)

```text
Network
    ↓ (paket keladi)
Socket Receive Buffer
    ↓ (application o'qiydi)
Application
```

---

## Diagram

```text
Sender                              Receiver

Application                        Application
    |                                  ^
    v                                  |
Send Buffer                      Receive Buffer
    |                                  ^
    v                                  |
Kernel                           Kernel
    |                                  ^
    v                                  |
Network  =========================> Network
```

---

## Buffer o'lchami

### Linux'da buffer o'lchamini ko'rish

```bash
# Receive buffer
sysctl net.core.rmem_default
sysctl net.core.rmem_max

# Send buffer
sysctl net.core.wmem_default
sysctl net.core.wmem_max
```

### Socket buffer o'lchamini o'zgartirish

```go
conn, _ := net.Dial("tcp", "localhost:8080")
tcpConn := conn.(*net.TCPConn)

// Set send buffer size
tcpConn.SetWriteBuffer(65536)

// Set receive buffer size
tcpConn.SetReadBuffer(65536)
```

---

## Buffer to'lishi

Buffer to'lsa nima bo'ladi?

### Packet Drop

```text
Receive Buffer to'ldi
    ↓
Yangi paket keldi
    ↓
Buffer joy yo'q
    ↓
Packet tashlab tashlandi!
```

### Back Pressure

```text
Application sekin o'qiydi
    ↓
Receive Buffer to'ldi
    ↓
Sender Window = 0
    ↓
Server to'xtaydi
    ↓
Back Pressure!
```

---

## Real-world misollar

### Nginx

```nginx
# Nginx buffer sozlamalari
proxy_buffering on;
proxy_buffer_size 4k;
proxy_buffers 8 4k;
```

### PostgreSQL

```ini
# PostgreSQL shared buffers
shared_buffers = 128MB
```

### Redis

```ini
# Redis client-output-buffer-limit
client-output-buffer-limit normal 0 0 0
```

---

## Amaliyot

### Buffer stats

```bash
ss -tm
```

Natija:

```text
skmem:(r0,rb374400,t0,tb46080,f0,w0,o0,bl0,d0)
```

- `r` -- receive queue
- `rb` -- receive buffer
- `t` -- send queue
- `tb` -- send buffer

### Go orqali buffer o'lchash

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, _ := net.Dial("tcp", "localhost:8080")
    tcpConn := conn.(*net.TCPConn)

    // Buffer o'lchamini o'zgartirish
    tcpConn.SetReadBuffer(1024 * 1024)   // 1MB
    tcpConn.SetWriteBuffer(1024 * 1024)  // 1MB

    fmt.Println("Buffer size set")
    conn.Close()
}
```

---

## Xulosa

- Buffer paketlar saqlanadigan joy
- Send va Receive buffer mavjud
- Buffer to'lsa -- packet drop yoki back pressure
- Nginx, PostgreSQL, Redis ham buffer ishlatadi
