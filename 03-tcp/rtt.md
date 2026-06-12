# RTT (Round Trip Time)

Paketning borib kelish vaqti.

---

## Nazariya

RTT = borish vaqti + qaytish vaqti

```text
Laptop -> Google
Google -> Laptop
```

---

## Amaliyot

### RTT ni o'lchash

```bash
ping google.com
```

Natija:

```text
64 bytes from google.com: icmp_seq=1 ttl=116 time=24.1 ms
```

Bu:

```text
RTT = 24.1 ms
```

---

## Diagram

```text
Client            Server

     |------------>|
     |   borish     |
     |              |
     |<------------|
     |  qaytish    |
     |              |
<--------- RTT --------->
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
    start := time.Now()

    conn, err := net.DialTimeout("tcp", "google.com:80", 5*time.Second)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    elapsed := time.Since(start)
    fmt.Printf("RTT: %v\n", elapsed)
}
```

### Python

```python
import socket
import time

start = time.time()
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(("google.com", 80))
elapsed = time.time() - start

print(f"RTT: {elapsed*1000:.2f} ms")
sock.close()
```

---

## Xulosa

- RTT = borib kelish vaqti
- `ping` bilan o'lchash mumkin
- RTT qancha kam bo'lsa, ulanish shuncha tez
