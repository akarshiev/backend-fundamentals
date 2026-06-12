# REST Principles

REST asosiy qoidalari.

---

## Nazariya

REST (Representational State Transfer) 6 ta asosiy qoidaga asoslangan.

---

## 1. Client-Server

```text
Client: Frontend, Mobile, CLI
Server: Backend, Database
```

Alohida ishlaydi.

---

## 2. Stateless

Server client haqida hech narsa eslamaydi.

Har bir request mustaqil.

```text
Request 1: GET /users/1
Request 2: GET /users/1
```

Ikki request ham bir xil javob beradi.

---

## 3. Cache

Server javobini cache qilish mumkin.

```text
GET /users -> Cache-Control: max-age=3600
```

---

## 4. Uniform Interface

Barcha resource lar bir xil interface.

```text
GET /users
GET /products
GET /orders
```

---

## 5. Layered System

```text
Client -> Load Balancer -> Server -> Database
```

Client faqat serverni ko'radi.

---

## 6. Code on Demand (optional)

Server clientga kod yuborishi mumkin.

```text
JavaScript
```

---

## Diagram

```text
REST Architecture:

+------------------+
| Client           |
+------------------+
        |
        V
+------------------+
| Server           |
+------------------+
        |
        V
+------------------+
| Database         |
+------------------+

Principles:
- Client-Server
- Stateless
- Cache
- Uniform Interface
- Layered System
```

---

## Xulosa

- REST 6 ta qoidaga asoslangan
- Stateless va Client-Server eng muhim
- Uniform Interface barcha resource lar uchun bir xil
