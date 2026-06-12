# HTTPS

HTTP + TLS = HTTPS.

---

## Mavzular

- SSL vs TLS
- TLS Handshake
- Certificates

---

## Nazariya

HTTP:

```text
Plain Text
```

HTTPS:

```text
HTTP + TLS
```

---

## Diagram

```text
Browser
   |
   V
TCP
   |
   V
TLS Handshake
   |
   V
HTTP (encrypted)
   |
   V
Server
```

---

## Amaliyot

### HTTPS ulanish

```bash
curl https://example.com
```

### TLS handshake ko'rish

```bash
openssl s_client -connect google.com:443
```

### Certificate ko'rish

```bash
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -dates
```

---

## Kod

### Go

```go
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func main() {
	// TLS config
	config := &tls.Config{
		InsecureSkipVerify: false,
	}

	// TLS dial
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 5 * 1e9},
		"tcp",
		"google.com:443",
		config,
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// HTTP request
	request := "GET / HTTP/1.1\r\nHost: google.com\r\nConnection: close\r\n\r\n"
	conn.Write([]byte(request))

	// Read response
	body, _ := io.ReadAll(conn)
	fmt.Println(string(body))
}
```

### Python

```python
import ssl
import socket

context = ssl.create_default_context()

with socket.create_connection(("google.com", 443)) as sock:
    with context.wrap_socket(sock, server_hostname="google.com") as ssock:
        request = "GET / HTTP/1.1\r\nHost: google.com\r\nConnection: close\r\n\r\n"
        ssock.send(request.encode())

        response = ssock.recv(4096)
        print(response.decode(errors="ignore"))
```

---

## Xulosa

- HTTPS = HTTP + TLS
- SSL eski, TLS yangi
- Sertifikatlar serverni tasdiqlaydi
