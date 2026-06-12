# HTTP Methods

REST API da ishlatiladigan HTTP methodlari.

---

## GET

Ma'lumot olish.

```http
GET /users/123
```

Response:

```http
HTTP/1.1 200 OK

{
  "id": 123,
  "name": "Ali"
}
```

---

## POST

Yangi resource yaratish.

```http
POST /users

{
  "name": "Ali"
}
```

Response:

```http
HTTP/1.1 201 Created

{
  "id": 124,
  "name": "Ali"
}
```

---

## PUT

To'liq almashtirish.

```http
PUT /users/123

{
  "name": "Ali",
  "age": 20
}
```

Response:

```http
HTTP/1.1 200 OK

{
  "id": 123,
  "name": "Ali",
  "age": 20
}
```

---

## PATCH

Qisman update.

```http
PATCH /users/123

{
  "age": 21
}
```

Response:

```http
HTTP/1.1 200 OK

{
  "id": 123,
  "name": "Ali",
  "age": 21
}
```

---

## DELETE

O'chirish.

```http
DELETE /users/123
```

Response:

```http
HTTP/1.1 204 No Content
```

---

## Diagram

```text
GET    -> Read
POST   -> Create
PUT    -> Update (full)
PATCH  -> Update (partial)
DELETE -> Delete
```

---

## Amaliyot

### GET

```bash
curl http://localhost:8080/users
```

### POST

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Ali"}' http://localhost:8080/users
```

### PUT

```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"name":"Ali","age":20}' http://localhost:8080/users/1
```

### PATCH

```bash
curl -X PATCH -H "Content-Type: application/json" \
  -d '{"age":21}' http://localhost:8080/users/1
```

### DELETE

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

## Xulosa

- GET = o'qish
- POST = yaratish
- PUT = to'liq yangilash
- PATCH = qisman yangilash
- DELETE = o'chirish
