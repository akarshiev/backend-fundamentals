# TCP Checksum

TCP segment buzilgan yoki buzilmaganini tekshiradi.

---

## Nazariya

TCP checksum butun segment ustida hisoblanadi. Header va data qamrab olinadi.

### Sender

1. Checksum hisoblaydi
2. Header ichiga joylashtiradi
3. Paketni yuboradi

### Receiver

1. Paketni oladi
2. Yana hisoblaydi
3. Mos kelmasa -- packet tashlab yuboriladi

---

## Qanday ishlaydi

```text
Sender:

1. TCP header + data ni o'qiydi
2. 16-bit qismlarga bo'ladi
3. Hammasini qo'shadi (one's complement)
4. Natijani o'chiradi (one's complement)
5. Checksum maydoniga yozadi

Receiver:

1. TCP header + data ni o'qiydi
2. 16-bit qismlarga bo'ladi
3. Hammasini qo'shadi (one's complement)
4. Agar 0 chiqsa -- xatolik yo'q
5. Agar 0 bo'lmasa -- xatolik bor, packet tashlab tashlanadi
```

---

## Diagram

```text
Sender                              Receiver

[data] + [checksum] ---------->
                               Checksum hisoblaydi
                                0 chiqdi --> qabul qilindi
                                != 0 --> tashlab tashlandi
```

---

## Pseudo Header

TCP checksum pseudo header ni ham qamrab oladi:

```text
+------------------+
| Source IP        |
+------------------+
| Destination IP   |
+------------------+
| Zero | Protocol | TCP Length |
+------------------+
| TCP Header       |
+------------------+
| TCP Data         |
+------------------+
```

---

## Amaliyot

### Tcpdump bilan checksum ko'rish

```bash
sudo tcpdump -i any -nn port 8080 -v
```

Verbose mode'da checksum ko'rinadi.

### Go orqali

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

    // TCP socket avtomatik checksum hisoblaydi
    conn.Write([]byte("Hello"))
    fmt.Println("Data sent with automatic checksum")
}
```

---

## Xulosa

- TCP checksum ma'lumot yaxlitligini ta'minlaydi
- Sender hisoblaydi, receiver tekshiradi
- Xatolik aniqlansa, packet tashlab tashlanadi
- Pseudo header IP manzillarini ham qamrab oladi
