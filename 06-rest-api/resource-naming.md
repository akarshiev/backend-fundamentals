# Resource Naming

REST API da resource nomlash qoidalari.

---

## Nazariya

Resource = noun (ot).

```text
/users
/products
/orders
```

---

## Yomon misol

```http
GET /getUser
POST /createUser
DELETE /deleteUser
```

REST emas.

---

## Yaxshi misol

```http
GET /users/123
POST /users
PATCH /users/123
DELETE /users/123
```

---

## Qoidalar

### 1. Plural ishlatish

```text
/users (to'g'ri)
/user (noto'g'ri)
```

### 2. Nested resources

```http
GET /users/123/orders
GET /users/123/orders/456
```

### 3. Query parameters

```http
GET /users?age=20&city=Tashkent
```

---

## Diagram

```text
/users              -> Collection
/users/123          -> Specific item
/users/123/orders   -> Nested collection
/users/123/orders/456 -> Nested item
```

---

## Amaliyot

### Collection

```bash
curl http://localhost:8080/users
```

### Specific item

```bash
curl http://localhost:8080/users/123
```

### Nested

```bash
curl http://localhost:8080/users/123/orders
```

---

## Xulosa

- Resource = noun (ot)
- Plural ishlatish: /users
- Nested resources: /users/123/orders
- Query parameters: /users?age=20
