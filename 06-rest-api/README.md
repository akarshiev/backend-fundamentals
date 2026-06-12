# REST API

REST (Representational State Transfer) -- API yaratish usuli.

---

## Mavzular

- REST Principles
- HTTP Methods (GET, POST, PUT, PATCH, DELETE)
- Resource Naming
- Idempotency
- Pagination (Offset, Cursor)

---

## Nazariya

REST aslida:

```text
Resource-oriented API
```

Masalan:

```text
User
Order
Product
Payment
```

resource.

---

## Diagram

```text
Client                        Server

   |--- GET /users ----------->|
   |<-- 200 OK + JSON ---------|
   |                           |
   |--- POST /users ---------->|
   |<-- 201 Created -----------|
   |                           |
   |--- PATCH /users/1 ------->|
   |<-- 200 OK ----------------|
   |                           |
   |--- DELETE /users/1 ------>|
   |<-- 204 No Content --------|
```

---

## Amaliyot

### GET request

```bash
curl http://localhost:8080/users
```

### POST request

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Ali"}' http://localhost:8080/users
```

### PUT request

```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"name":"Ali","age":20}' http://localhost:8080/users/1
```

### DELETE request

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

## Xulosa

- REST resource-oriented API
- HTTP methods CRUD operatsiyalari
- To'g'ri resource naming muhim
