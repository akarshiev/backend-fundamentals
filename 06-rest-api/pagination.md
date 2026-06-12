# Pagination

Katta datasetlarni qismlarga bo'lish.

---

## Nazariya

Database'da:

```text
10 million users
```

bor.

Yomon:

```http
GET /users
```

10 millionni qaytarish.

---

## Offset Pagination

```http
GET /users?page=2&size=20
```

SQL:

```sql
SELECT * FROM users
LIMIT 20 OFFSET 20;
```

Afzalligi:

```text
Oddiy
```

Kamchiligi:

```text
Katta datasetda sekin
Masalan: OFFSET 1000000
```

---

## Cursor Pagination

```http
GET /users?cursor=4857
```

SQL:

```sql
SELECT * FROM users
WHERE id > 4857
LIMIT 20;
```

Afzalligi:

```text
Juda tez
```

Kamchiligi:

```text
Page 57 ga o'tish qiyin
```

---

## Diagram

```text
Offset:
Page 1: users[0:20]
Page 2: users[20:40]
Page 3: users[40:60]

Cursor:
Cursor 1: users WHERE id > 0 LIMIT 20
Cursor 2: users WHERE id > 20 LIMIT 20
Cursor 3: users WHERE id > 40 LIMIT 20
```

---

## Qachon qaysi biri?

### Offset

```text
Admin panel
CRM
Dashboard
```

### Cursor

```text
Instagram
Twitter
Telegram
YouTube comments
```

---

## Amaliyot

### Offset

```bash
curl "http://localhost:8080/users?page=1&size=20"
curl "http://localhost:8080/users?page=2&size=20"
```

### Cursor

```bash
curl "http://localhost:8080/users?cursor=0&size=20"
curl "http://localhost:8080/users?cursor=20&size=20"
```

---

## Kod

### Go

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func offsetHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}

	offset := (page - 1) * size

	fmt.Printf("Offset: %d, Size: %d\n", offset, size)

	// Database query:
	// SELECT * FROM users LIMIT ? OFFSET ?

	users := []User{
		{ID: offset + 1, Name: "User 1"},
		{ID: offset + 2, Name: "User 2"},
	}

	json.NewEncoder(w).Encode(users)
}

func cursorHandler(w http.ResponseWriter, r *http.Request) {
	cursor, _ := strconv.Atoi(r.URL.Query().Get("cursor"))
	size, _ := strconv.Atoi(r.URL.Query().Get("size"))

	if size < 1 {
		size = 20
	}

	fmt.Printf("Cursor: %d, Size: %d\n", cursor, size)

	// Database query:
	// SELECT * FROM users WHERE id > ? LIMIT ?

	users := []User{
		{ID: cursor + 1, Name: "User 1"},
		{ID: cursor + 2, Name: "User 2"},
	}

	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Has("cursor") {
			cursorHandler(w, r)
		} else {
			offsetHandler(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
```

---

## Xulosa

- Offset: oddiy, lekin katta datasetda sekin
- Cursor: tez, lekin page ga saklash qiyin
- Admin panel uchun Offset, social media uchun Cursor
