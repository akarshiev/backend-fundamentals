# Idempotency

Bir xil request 100 marta yuborilsa ham natija bir xil bo'lsa -- idempotent.

---

## Nazariya

> Idempotency: Bir xil request necha marta yuborilsa ham, natija bir xil.

---

## Methodlar

### GET -- Idempotent

```http
GET /users/123
```

100 marta yuborsangiz ham user o'zgarmaydi.

### PUT -- Idempotent

```http
PUT /users/123

{
  "name": "Ali"
}
```

100 marta yuborsangiz ham:

```text
Ali
```

bo'lib qoladi.

### DELETE -- Idempotent

```http
DELETE /users/123
```

1-marta:

```text
user deleted
```

100-marta:

```text
already deleted
```

Oxirgi holat bir xil.

### POST -- Odatda idempotent emas

```http
POST /payments
```

Ikki marta bossangiz:

```text
100$
100$
```

ikki marta yechiladi.

---

## Stripe yechimi

Stripe ishlatadi:

```http
POST /payments

Idempotency-Key: f4c58c39-....
```

Server:

```python
if alreadyProcessed(key):
    return previousResponse

processPayment()
saveKey()

return success
```

Shunda:

```text
2 marta bosildi
```

lekin:

```text
1 marta to'lov
```

amalga oshadi.

---

## Diagram

```text
Idempotent:
GET    -> Read (har doim bir xil)
PUT    -> Replace (har doim bir xil)
DELETE -> Delete (har doim bir xil)

Non-Idempotent:
POST   -> Create (har safar yangi)
```

---

## Amaliyot

### Idempotency key bilan

```bash
curl -X POST \
  -H "Idempotency-Key: 123e4567-e89b-12d3-a456-426614174000" \
  -H "Content-Type: application/json" \
  -d '{"amount": 100}' \
  http://localhost:8080/payments
```

---

## Xulosa

- GET, PUT, DELETE = idempotent
- POST = odatda idempotent emas
- Stripe Idempotency Key ishlatadi
