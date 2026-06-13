# MSS (Maximum Segment Size)

TCP segmentning maksimal payload o'lchami.

---

## Nazariya

MSS -- TCP payload'ning maksimal o'lchami. IP va TCP header'lar kirilmaydi.

### Formula

```text
MSS = MTU - IP Header - TCP Header

MSS = 1500 - 20 - 20
MSS = 1460 bytes
```

---

## Diagram

```text
+--------------------------------------------------+
| IP Header | TCP Header | Payload                 |
| 20 bytes  | 20 bytes   | 1460 bytes (MSS)        |
+--------------------------------------------------+
| <------------ 1500 bytes (MTU) ----------------->|
```

---

## MSS vs MTU

```text
MTU  = IP Header + TCP Header + Payload = 1500 bytes
MSS  = Payload = 1460 bytes

MTU: butun paket
MSS: faqat ma'lumot qismi
```

---

## MSS Negotiation

MSS ulanish o'rnatish paytida SYN paketlarida uzatiladi:

```text
Client --> SYN (MSS=1460) --> Server
Client <-- SYN-ACK (MSS=1400) <-- Server

Natija: MSS = min(1460, 1400) = 1400
```

---

## Amaliyot

### MSS ni ko'rish

```bash
sudo tcpdump -i any port 8080 -nn -v 2>/dev/null | grep mss
```

### ss bilan

```bash
ss -ti dst 93.184.216.34
```

Natija:

```text
mss:1448
```

### Go orqali

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    conn, _ := net.Dial("tcp", "example.com:80")
    tcpConn := conn.(*net.TCPConn)

    fmt.Println("Local addr:", tcpConn.LocalAddr())
    fmt.Println("Remote addr:", tcpConn.RemoteAddr())

    // TCP avtomatik MSS ni negotiate qiladi
    conn.Close()
}
```

---

## Xulosa

- MSS = MTU - IP Header - TCP Header
- Standart: 1460 bytes
- Ulanish o'rnatish paytida negotiate qilinadi
- MSS payload o'lchamini belgilaydi
