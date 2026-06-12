# DNS Nima?

DNS (Domain Name System) -- internetning telefon kitobi.

---

## Nazariya

DNS domenni IP ga aylantiradi.

Misol:

```text
google.com -> 142.250.190.78
github.com -> 140.82.121.4
```

Browser serverga ulanishdan oldin DNS'dan IP so'raydi.

Jarayon:

```text
Browser
   |
   V
DNS Resolver
   |
   V
DNS Server
   |
   V
IP qaytadi
   |
   V
Browser TCP ulanish ochadi
```

---

## TLD (Top Level Domain)

Domenning oxirgi qismi.

Mashhur TLDlar:

```text
.com
.org
.net
.edu
.gov
.io
.dev
.uz
```

Misol:

```text
42.uz

42 = Second Level Domain
uz = TLD
```

---

## Amaliyot

### DNS query

```bash
dig google.com
```

### NS lookup

```bash
nslookup google.com
```

### DNS serverni o'zgartirish

```bash
dig @8.8.8.8 google.com
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
    ips, err := net.LookupIP("google.com")
    if err != nil {
        panic(err)
    }

    for _, ip := range ips {
        fmt.Println(ip)
    }
}
```

### Python

```python
import socket

ips = socket.getaddrinfo("google.com", 80)
for ip in ips:
    print(ip[4][0])
```

---

## Xulosa

- DNS domenni IP ga aylantiradi
- TLD = domenning oxirgi qismi
- `dig` va `nslookup` buyruqlari ishlatiladi
