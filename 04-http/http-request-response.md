# HTTP Request/Response

HTTP asosiy request/response modelida ishlaydi.

---

## Nazariya

### HTTP Request

```text
Method URI HTTP/Version
Header1: Value1
Header2: Value2

Body (optional)
```

Misol:

```http
GET /users/123 HTTP/1.1
Host: api.example.com
Accept: application/json
```

### HTTP Response

```text
HTTP/Version Status Code Reason
Header1: Value1

Body
```

Misol:

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 123,
  "name": "Abdukarim"
}
```

---

## Diagram

```text
Client                           Server

   |                                |
   |--- Request ------------------>|
   |                                |
   |<--- Response ------------------|
   |                                |
```

---

## Amaliyot

### curl bilan request

```bash
curl http://example.com
```

### Headers bilan

```bash
curl -H "Accept: application/json" http://example.com/api/users
```

### POST request

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Ali"}' http://example.com/api/users
```

---

## Kod

### Go

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Headers:", resp.Header)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Body:", string(body))
}
```

### Python

```python
import requests

response = requests.get("http://example.com")

print("Status:", response.status_code)
print("Headers:", response.headers)
print("Body:", response.text)
```

---

## Xulosa

- HTTP request: Method + URI + Headers + Body
- HTTP response: Status + Headers + Body
- `curl` va `requests` kutubxonasi ishlatiladi
