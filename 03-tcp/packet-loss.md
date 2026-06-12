# Packet Loss

Internetda paket yo'qolishi mumkin. TCP buni tuzatadi.

---

## Nazariya

TCP Sequence Number saqlaydi.

```text
1
2
3
4
5
```

paket ketdi.

4 yo'qoldi.

Server:

```text
4 kelmadi
```

deydi.

TCP:

```text
4 ni qayta yubor
```

qiladi.

---

## Diagram

```text
Client            Server

1 -------------->
2 -------------->
3 -------------->
4 --X (yo'qoldi)
5 -------------->

     <----------- 4 ni qayta yubor

4 -------------->
```

---

## Amaliyot

### Packet loss ni tekshirish

```bash
ping -c 100 google.com
```

### TCP retransmission ko'rish

```bash
sudo tcpdump -i any -nn 'tcp[tcpflags] & (tcp-syn|tcp-fin) != 0'
```

---

## Kod

### Go

```go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.DialTimeout("tcp", "google.com:80", 5*time.Second)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    // Large data yuborish
    data := make([]byte, 1024*1024) // 1MB
    for i := range data {
        data[i] = byte(i % 256)
    }

    _, err = conn.Write(data)
    if err != nil {
        fmt.Println("Write error:", err)
        return
    }

    fmt.Println("Data sent successfully")
}
```

---

## Xulosa

- Packet loss internetda bo'lishi mumkin
- TCP Sequence Number orqali yo'qolgan paketlarni aniqlaydi
- Retransmission orqali paketlarni qayta yuboradi
