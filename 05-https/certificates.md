# TLS Certificates

Serverni tasdiqlash uchun sertifikatlar ishlatiladi.

---

## Nazariya

Sertifikat = server shaxsiyati.

```text
Server Certificate:
- Domain: example.com
- Public Key: ...
- Issuer: Let's Encrypt
- Valid: 2024-01-01 - 2024-12-31
```

---

## Certificate turlari

### DV (Domain Validation)

```text
Oddiy
Tez
Bepul (Let's Encrypt)
```

### OV (Organization Validation)

```text
Kompaniya tasdiqlangan
Qimmatroq
```

### EV (Extended Validation)

```text
Eng qimmat
Eng ishonchli
Banklar uchun
```

---

## Diagram

```text
Certificate Chain:

Root CA
   |
   V
Intermediate CA
   |
   V
Server Certificate
```

---

## Amaliyot

### Sertifikatni ko'rish

```bash
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -text
```

### Sertifikat muddatini tekshirish

```bash
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -dates
```

### Chain tekshirish

```bash
openssl s_client -connect google.com:443 -showcerts
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
)

func main() {
	conn, err := tls.Dial("tcp", "google.com:443", &tls.Config{})
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	for _, cert := range state.PeerCertificates {
		fmt.Printf("Subject: %s\n", cert.Subject.CommonName)
		fmt.Printf("Issuer: %s\n", cert.Issuer.CommonName)
		fmt.Printf("Not Before: %s\n", cert.NotBefore)
		fmt.Printf("Not After: %s\n", cert.NotAfter)
		fmt.Println("---")
	}
}
```

### Python

```python
import ssl
import socket

context = ssl.create_default_context()

with socket.create_connection(("google.com", 443)) as sock:
    with context.wrap_socket(sock, server_hostname="google.com") as ssock:
        cert = ssock.getpeercert()
        print("Subject:", dict(x[0] for x in cert["subject"]))
        print("Issuer:", dict(x[0] for x in cert["issuer"]))
        print("Valid from:", cert["notBefore"])
        print("Valid to:", cert["notAfter"])
```

---

## Xulosa

- Sertifikat serverni tasdiqlaydi
- DV oddiy va bepul, EV eng qimmat
- Let's Encrypt bepul sertifikat beradi
