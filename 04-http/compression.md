# HTTP Compression

Trafikni kamaytirish uchun ma'lumotni siqish.

---

## Nazariya

Browser:

```text
Accept-Encoding: br, gzip
```

deydi.

Server:

```text
Content-Encoding: br
```

deb javob beradi.

---

## Gzip vs Brotli

### Gzip

```text
Eski
Juda tez
CPU kam ishlatadi
```

### Brotli (br)

```text
Yangi
Ko'proq siqadi
Kamroq trafik
CPU ko'proq ishlatadi
```

---

## Diagram

```text
Browser                  Server

   |--- Accept-Encoding: br, gzip -->|
   |                                 |
   |<-- Content-Encoding: br --------|
   |                                 |
   |<-- Compressed body ------------|
```

---

## Amaliyot

### Gzip test

```bash
curl -H "Accept-Encoding: gzip" http://example.com --compressed
```

### Brotli test

```bash
curl -H "Accept-Encoding: br" http://example.com --compressed
```

### Siqishni ko'rish

```bash
curl -o /dev/null -s -w "Size: %{size_download}\n" \
  -H "Accept-Encoding: gzip" http://example.com
```

---

## Kod

### Go

```go
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Accept-Encoding", "gzip, br")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Content-Encoding:", resp.Header.Get("Content-Encoding"))

	var reader io.Reader = resp.Body
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		reader, _ = gzip.NewReader(resp.Body)
	}

	body, _ := io.ReadAll(reader)
	fmt.Println("Body length:", len(body))
}
```

### Python

```python
import requests

response = requests.get("http://example.com", headers={"Accept-Encoding": "gzip"})

print("Content-Encoding:", response.headers.get("Content-Encoding"))
print("Body length:", len(response.content))
```

---

## Xulosa

- Compression = trafikni kamaytirish
- Gzip tez, Brotli ko'proq siqadi
- Cloudflare, Vercel, Google Brotli ishlatadi
