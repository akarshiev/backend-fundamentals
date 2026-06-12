# File Descriptor

Linux'da har bir ochiq fayl, socket yoki I/O resursi uchun raqam ajratiladi.

---

## Nazariya

File Descriptor (FD) -- kernel tomonidan berilgan integer raqam.

Har bir ochiq resurs:

```text
0 = stdin
1 = stdout
2 = stderr
3 = birinchi ochiq fayl
4 = ikkinchi ochiq fayl/socket
```

---

## Amaliyot

### FD ni ko'rish

```bash
ls -l /proc/$$/fd
```

yoki

```bash
lsof -p $$
```

### Go da tekshirish

```go
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    file, _ := os.Open("users.txt")
    fmt.Println("File FD:", file.Fd())

    conn, _ := net.Dial("tcp", "google.com:80")
    fmt.Println("Socket FD:", conn.(*net.TCPConn).File())
    defer conn.Close()
}
```

---

## Xulosa

- FD = integer raqam
- Har bir ochiq resurs uchun unique FD
- FD limiti `ulimit -n` bilan belgilanadi
