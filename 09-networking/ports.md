# Portlar

## Port nima?

IP qurilmani topadi.

**Port** esa qurilmadagi dasturga olib boradi.

---

## Misol

Kompyuter: `192.168.1.10`

Ichida bir vaqtda ishlayotgan dasturlar:

| Dastur | Port |
|--------|------|
| Chrome | (client port - avtomatik) |
| PostgreSQL | 5432 |
| Redis | 6379 |
| Spring Boot | 8080 |

Hammasi bitta IP'da. Qaysi dastur ekanini **port** bildiradi.

---

## Mashhur portlar

| Port | Xizmat |
|------|--------|
| 22 | SSH |
| 53 | DNS |
| 80 | HTTP |
| 443 | HTTPS |
| 5432 | PostgreSQL |
| 6379 | Redis |
| 8080 | Spring Boot |

---

## Port bo'lmaganida nima bo'lardi?

Har bir dastur uchun alohida IP kerak bo'lardi:

```
Chrome     → 192.168.1.10
PostgreSQL → 192.168.1.11
Redis      → 192.168.1.12
```

Bu absurd bo'lar edi. Portlar shu muammoni hal qiladi.

---

## Router qaysi qurilmaga yuborishni qanday biladi?

Router **IP + Port** kombinatsiyasini eslab qoladi.

Misol:

```
192.168.1.10:52341 → Google
```

Response `52341` port bilan qaytadi.

Router: "Aha, bu Abdukarimning browser requesti edi" deydi.

---

## Amaliyot

### Portlarni ko'rish

```bash
ss -tulpn
```
