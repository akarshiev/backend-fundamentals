# MTU (Maximum Transmission Unit)

Internet orqali yuboriladigan paketning maksimal o'lchami.

---

## Nazariya

Internet 1GB faylni bitta bo'lak yubormaydi. Paketlarga bo'ladi.

Standart Ethernet:

```text
MTU = 1500 bytes
```

---

## Paket tuzilishi

```text
[IP Header][TCP Header][Payload]
   20          20        1460 bytes
```

---

## Diagram

```text
+--------------------------------------------------+
| IP Header | TCP Header | Payload                 |
| 20 bytes  | 20 bytes   | 1460 bytes              |
+--------------------------------------------------+
| <------------ 1500 bytes (MTU) ----------------->|
```

---

## Amaliyot

### MTU ni ko'rish

```bash
ip link
```

yoki

```bash
ip addr
```

Natija:

```text
mtu 1500
```

### Paketlarni ko'rish

```bash
sudo tcpdump -i any -nn
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
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // TCP socket o'ziga xos MTU ni biladi
    tcpConn := conn.(*net.TCPConn)
    fmt.Println("TCP Connection established")
    fmt.Println("Local addr:", tcpConn.LocalAddr())
    fmt.Println("Remote addr:", tcpConn.RemoteAddr())
}
```

---

## Xulosa

- MTU = paketning maksimal o'lchami
- Standart Ethernet: 1500 bytes
- IP Header + TCP Header + Payload = MTU
