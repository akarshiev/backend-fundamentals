# UDP Basics

UDP -- eng oddiy transport protokoli.

---

## Nazariya

UDP ma'lumotni yuboradi va unutadi. Qabul qilindi yoki yo'qmi -- bilmasligi kerak.

### TCP vs UDP

```text
TCP:
Client --> SYN --> Server
     <-- SYN-ACK -->
     --> ACK -->
     --> DATA -->
     <-- ACK -->
Connection o'rnatildi, ma'lumot yetdi.

UDP:
Client --> DATA --> Server
Yubor va unut.
```

---

## UDP Header (8 bytes)

```text
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          Source Port          |       Destination Port        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|            Length             |           Checksum            |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

TCP headerdan farqi:

```text
TCP Header:  20 bytes (minimum)
UDP Header:  8 bytes (doim)

TCP: Sequence, ACK, Flags, Window...
UDP: Faqat Port, Length, Checksum
```

---

## Diagram

```text
Sender                              Receiver

UDP Packet ----------------------> UDP Packet
| Src Port | Dst Port | Len | Chk |
|          Data                   |
                                  Qabul qildi yoki yo'qmi
                                  --- muhim emas!
```

---

## Go orqali UDP Server

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    // UDP server
    addr, _ := net.ResolveUDPAddr("udp", ":8080")
    conn, _ := net.ListenUDP("udp", addr)
    defer conn.Close()

    fmt.Println("UDP Server listening on :8080")

    buffer := make([]byte, 1024)
    n, remoteAddr, _ := conn.ReadFromUDP(buffer)

    fmt.Printf("Received from %s: %s\n", remoteAddr, string(buffer[:n]))

    // Javob yuborish
    conn.WriteToUDP([]byte("ACK"), remoteAddr)
}
```

---

## Go orqali UDP Client

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    // UDP client
    addr, _ := net.ResolveUDPAddr("udp", "localhost:8080")
    conn, _ := net.DialUDP("udp", nil, addr)
    defer conn.Close()

    // Ma'lumot yuborish
    conn.Write([]byte("Hello UDP!"))
    fmt.Println("Sent: Hello UDP!")

    // Javob olish
    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)
    fmt.Printf("Received: %s\n", string(buffer[:n]))
}
```

---

## Amaliyot

### UDP server ishga tushirish

```bash
# Terminal 1
go run udp-server.go
```

### UDP client yuborish

```bash
# Terminal 2
go run udp-client.go
```

### UDP socketlarni ko'rish

```bash
ss -uan
```

### UDP traffic ni tekshirish

```bash
sudo tcpdump -i any udp port 8080
```

---

## Xulosa

- UDP eng oddiy transport protokoli
- Header faqat 8 bytes
- Connection, ACK, Ordering yo'q
- Fire and Forget falsafasi
