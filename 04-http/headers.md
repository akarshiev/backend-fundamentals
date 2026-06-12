# HTTP Headers

HTTP request/response uchun qo'shimcha ma'lumot beradi.

---

## Nazariya

Headers key: value formatida.

```text
Header-Name: value
```

---

## Muhim headerlar

### Request

```text
Host: example.com
User-Agent: Mozilla/5.0
Accept: application/json
Content-Type: application/json
Authorization: Bearer token123
```

### Response

```text
Content-Type: application/json
Content-Length: 1234
Cache-Control: max-age=3600
Set-Cookie: session=abc123
```

---

## Diagram

```text
Request:
+------------------+
| Host             |
| User-Agent       |
| Accept           |
| Content-Type     |
| Authorization    |
+------------------+

Response:
+------------------+
| Content-Type     |
| Content-Length   |
| Cache-Control    |
| Set-Cookie        |
+------------------+
```

---

## Amaliyot

### Headerlarni ko'rish

```bash
curl -I http://example.com
```

### Custom header

```bash
curl -H "Authorization: Bearer token123" http://example.com/api
```

### Response header

```bash
curl -D- http://example.com
```

---

## Kod

### Go

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	req.Header.Set("Authorization", "Bearer token123")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
}
```

### Python

```python
import requests

headers = {
    "Authorization": "Bearer token123",
    "Accept": "application/json"
}

response = requests.get("http://example.com", headers=headers)
print("Content-Type:", response.headers["Content-Type"])
```

---

## Xulosa

- Headers key: value formatida
- Request va response headerlari turlicha
- Authorization, Content-Type, Accept eng ko'p ishlatiladi
