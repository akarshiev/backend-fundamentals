# TLS Handshake

TLS ulanish o'rnatish jarayoni.

---

## Nazariya

TLS handshake TCP dan keyin sodir bo'ladi.

```text
1. Client Hello
2. Server Hello
3. Certificate
4. Key Exchange
5. Finished
```

---

## Diagram

```text
Client                      Server

   |--- Client Hello --------->|
   |                           |
   |<--- Server Hello ----------|
   |<--- Certificate -----------|
   |<--- Server Hello Done ----|
   |                           |
   |--- Client Key Exchange -->|
   |--- Change Cipher Spec --->|
   |--- Finished ------------->|
   |                           |
   |<--- Change Cipher Spec ---|
   |<--- Finished -------------|
   |                           |
   |<=== Encrypted Data =====>|
```

---

## Amaliyot

### TLS handshake ko'rish

```bash
openssl s_client -connect google.com:443 -msg
```

### TLS versiyasini tekshirish

```bash
openssl s_client -connect google.com:443
```

---

## Kod

### Go

```go
package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func main() {
	start := time.Now()

	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 5 * time.Second},
		"tcp",
		"google.com:443",
		config,
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	elapsed := time.Since(start)
	fmt.Printf("TLS Handshake time: %v\n", elapsed)

	state := conn.ConnectionState()
	fmt.Printf("TLS Version: %x\n", state.Version)
	fmt.Printf("Cipher Suite: %x\n", state.CipherSuite)
}
```

### Python

```python
import ssl
import socket
import time

start = time.time()

context = ssl.create_default_context()

with socket.create_connection(("google.com", 443)) as sock:
    with context.wrap_socket(sock, server_hostname="google.com") as ssock:
        elapsed = time.time() - start
        print(f"TLS Handshake time: {elapsed*1000:.2f} ms")
        print(f"TLS Version: {ssock.version()}")
        print(f"Cipher: {ssock.cipher()}")
```

---

## Xulosa

- TLS handshake TCP dan keyin sodir bo'ladi
- Client va server shifrlash usulini kelishadi
- TLS 1.3 tezroq, 2 ta round trip yetarli
