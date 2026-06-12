# REST API Compression

API response larini siqish.

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
Client                    Server

   |--- Accept-Encoding: br, gzip -->|
   |                                 |
   |<-- Content-Encoding: br --------|
   |<-- Compressed JSON ------------|
```

---

## Amaliyot

### Gzip test

```bash
curl -H "Accept-Encoding: gzip" \
  -H "Accept: application/json" \
  http://localhost:8080/users --compressed
```

### Brotli test

```bash
curl -H "Accept-Encoding: br" \
  -H "Accept: application/json" \
  http://localhost:8080/users --compressed
```

---

## Kod

### Go

```go
package main

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g gzipResponseWriter) Write(b []byte) (int, error) {
	return g.Writer.Write(b)
}

func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")

		gz, _ := gzip.NewWriterLevel(w, gzip.BestSpeed)
		defer gz.Close()

		next.ServeHTTP(gzipResponseWriter{w, gz}, r)
	})
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[{"id":1,"name":"Ali"},{"id":2,"name":"Bob"}]`)
	})

	handler := gzipMiddleware(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
```

---

## Xulosa

- Compression = trafikni kamaytirish
- Gzip tez, Brotli ko'proq siqadi
- API larda compression ishlatish kerak
