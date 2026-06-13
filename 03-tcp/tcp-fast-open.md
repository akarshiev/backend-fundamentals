# TCP Fast Open (TFO)

TCP ulanishini tezlashtirish usuli. RTT ≈ 0 ga erishish mumkin.

---

## Oddiy TCP

Oddiy TCP ulanishda:

```text
1. SYN          --> Clientdan Serverga
2. SYN-ACK      <-- Serverdan Clientga
3. ACK          --> Clientdan Serverga
4. DATA         --> Clientdan Serverga

1 RTT -- faqat ulanish o'rnatish
1 RTT -- ma'lumot yuborish
Jami: 2 RTT
```

---

## TCP Fast Open

TFO cookie ishlatadi. Birinchi ulanishda cookie olinadi, keyingilarida ishlatiladi.

### Birinchi ulanish

```text
Client -----------------> SYN + TFO Request
     <----------------- SYN-ACK + Cookie

Cookie olindi!
```

### Keyingi ulanish

```text
Client -----------------> SYN + Cookie + DATA
     <----------------- SYN-ACK + DATA

Darhol ma'lumot yuborildi!
```

### Natija

```text
RTT ≈ 0

Oddiy TCP: 2 RTT
TFO: 1 RTT
```

---

## Diagram

```text
Oddiy TCP:

Client                    Server
  |--- SYN ------------------>|
  |<-- SYN-ACK ---------------|
  |--- ACK ------------------>|
  |--- DATA ----------------->|  (1 RTT oldin)
  |<-- DATA ------------------|

TCP Fast Open:

Client                    Server
  |--- SYN + Cookie -------->|  (birinchi marta)
  |<-- SYN-ACK + Cookie -----|
  |                           |
  |--- SYN + Cookie + DATA ->|  (keyingi marta)
  |<-- ACK + DATA ------------|  (darhol!)
```

---

## Kamchiliklari

TFO hamma joyda ishlatilmaydi:

```text
✗ Hamma OS qo'llamaydi
✗ Hamma load balancer qo'llamaydi
✗ Internetda to'liq ishlatilmaydi
✗ Ba'zi firewall'lar bloklaydi
```

---

## RFC

TCP Fast Open uchun RFC:

```text
https://datatracker.ietf.org/doc/html/rfc7413
```

---

## Amaliyot

### TFO ni tekshirish

```bash
# Linux'da TFO mavjudligini tekshirish
sysctl net.ipv4.tcp_fastopen

# 0 = TFO o'chirilgan
# 1 = Client uchun yoqilgan
# 2 = Server uchun yoqilgan
# 3 = Hammasi yoqilgan
```

### TFN ni yoqish

```bash
# Server uchun
sudo sysctl -w net.ipv4.tcp_fastopen=3
```

### Go orqali

```go
package main

import (
    "fmt"
    "net"
    "syscall"
)

func main() {
    // TFO bilan ulanish
    dialer := &net.Dialer{
        Control: func(network, address string, c syscall.RawConn) error {
            return c.Control(func(fd uintptr) {
                // TFO flag'ni o'rnatish
                syscall.SetsockoptInt(int(fd), syscall.IPPROTO_TCP, 12, 1)
            })
        },
    }

    conn, err := dialer.Dial("tcp", "example.com:80")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Connected with TFO")
}
```

---

## Xulosa

- TFO ulanish vaqtini kamaytiradi
- Cookie ishlatadi
- RTT ≈ 0 ga erishish mumkin
- Hamma joyda ishlatilmaydi
- RFC 7413
