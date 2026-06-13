# Flow Control

Mijoz qanchalik tez qabul qila olsa, server shunchalik yuboradi.

---

## Nazariya

Agar server juda tez yuborsa, client buffer to'lib ketishi mumkin. Flow Control buni oldini oladi.

### Receive Window

Buni **Receive Window** nazorat qiladi. Client serverga "men shuncha ma'lumotni qabul qila olaman" deb bildiradi.

```text
Client --> Server: Window Size = 65535
Server: "Yaxshi, 65535 baytdan ko'p yubormayman"
```

---

## Diagram

```text
Sender                              Receiver

                  Window = 65535
     <----------------------------
[data: 1000] ---------------------->
     <----------------------------  Window = 64535
[data: 1000] ---------------------->
     <----------------------------  Window = 63535
     ...                            ...
[data: 63535] --------------------->  (window to'ldi)
     <----------------------------  Window = 0
     (yuborish to'xtaydi)
     ... (client o'qiydi) ...
     <----------------------------  Window = 30000
[data: 1000] --------------------->  (qayta boshlaydi)
```

---

## Zero Window

Agar client window = 0 bo'lsa, server to'xtaydi va vaqti-vaqti bilan tekshiradi:

```text
Sender                              Receiver

[data] -------------------------->  Window = 0
     <----------------------------  Window = 0
(data kutmoqda)
     <----------------------------  Window = 0
(data kutmoqda)
     <----------------------------  Window = 1000
[data] -------------------------->  (qayta boshlaydi)
```

---

## Back Pressure

Buffer to'lsa, back pressure bo'ladi. Server yuborishni sekinlashtiradi yoki to'xtatadi.

```text
Application --> Socket Buffer --> Kernel --> Network
                         ↑
                    Buffer to'ldi
                    Back pressure!
```

---

## Amaliyot

### Socket buffer o'lchamini ko'rish

```bash
ss -tm
```

### Window size ni tekshirish

```bash
sudo tcpdump -i any port 8080 -nn -v 2>/dev/null | grep window
```

### Go orqali

```go
package main

import (
    "fmt"
    "net"
)

func main() {
    listener, _ := net.Listen("tcp", ":8080")
    conn, _ := listener.Accept()
    defer conn.Close()

    // TCP avtomatik flow control bajaradi
    buffer := make([]byte, 1024)
    n, _ := conn.Read(buffer)
    fmt.Printf("Received %d bytes\n", n)
}
```

---

## Xulosa

- Flow Control client buffer'ni himoya qiladi
- Receive Window orqali boshqariladi
- Zero Window -- server to'xtaydi
- Back Pressure -- buffer to'lganda yuzaga keladi
