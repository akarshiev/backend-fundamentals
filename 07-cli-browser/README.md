# CLI Browser

O'zimizning minimal CLI browser yaratish.

---

## Mavzular

- TCP orqali ulanish
- HTTP request yuborish
- HTTP response olish
- Header va body parse qilish
- Minimal browser

---

## Nazariya

Browser ham aslida shu ishni bajaradi:

```text
1. DNS -> IP topish
2. TCP -> ulanish ochish
3. TLS -> shifrlash (HTTPS bo'lsa)
4. HTTP -> request yuborish
5. Response -> parse qilish
6. Render -> sahifani ko'rsatish
```

---

## Diagram

```text
CLI Browser

   |
   V

DNS Lookup
   |
   V

TCP Connect
   |
   V

TLS Handshake (HTTPS)
   |
   V

HTTP Request
   |
   V

HTTP Response
   |
   V

Parse & Display
```

---

## Amaliyot

### Python versiyasi

```bash
python http-client.py example.com
```

### Go versiyasi

```bash
go run http-client.go example.com
```

---

## Xulosa

- CLI browser TCP va HTTP ni tushunishga yordam beradi
- Har bir qadamni amalda ko'rish mumkin
- Spring Boot'dagi RestTemplate ortida shu jarayonlar ishlaydi
