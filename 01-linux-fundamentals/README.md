# Linux Fundamentals

Linux asoslarini tushunish backend developer uchun zarur.

---

## Mavzular

- Everything is a file
- File Descriptor
- Socket
- ulimit

---

## Nazariya

Linux operatsion tizimining asosiy falsafasi:

```text
Everything is a file
```

Bu degani, Linux'da hamma narsa fayl sifatida ko'riladi:

```text
Disk file
Socket
Pipe
Terminal
Keyboard
Mouse
```

Kernel uchun bularning hammasi -- fayl.

---

## Amaliyot

### File Descriptor ni ko'rish

```bash
ls -l /proc/$$/fd
```

yoki

```bash
lsof -p $$
```

### FD limitini tekshirish

```bash
ulimit -n
```

Natija:

```text
1024
```

Demak, 1024 ta ochiq file/socket dan ko'p bo'lmaydi.

Serverlar uchun:

```bash
ulimit -n 65535
```

---

## Kod

### Go

```go
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    // Fayl ochish
    file, err := os.Open("users.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Socket ochish
    conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer conn.Close()

    fmt.Println("Files opened successfully")
}
```

---

## Xulosa

- Linux'da hamma narsa fayl
- Har bir ochiq fayl/socket uchun FD ajratiladi
- FD limiti `ulimit -n` bilan tekshiriladi
