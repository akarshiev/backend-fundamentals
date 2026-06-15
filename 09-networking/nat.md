# NAT (Network Address Translation)

## NAT nima?

**NAT = Network Address Translation**

Router ichki private IP'larni bitta public IP ortiga yashiradi.

---


![NAT](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/NAT_Concept-en.svg/500px-NAT_Concept-en.svg.png)

*NAT ishlash prinsipi: Private IP → Public IP*

## NAT ishlash jarayoni

### Uy tarmog'i

```
Laptop (192.168.1.10)
      ↓
   Router
      ↓
  84.54.12.8 (Public IP)
      ↓
    Google
```

### Qanday ishlaydi?

1. Siz `https://google.com` ochasiz
2. Request `192.168.1.10` dan chiqadi
3. Lekin internet private IP'ni tanimaydi
4. Router `192.168.1.10` → `84.54.12.8` ga almashtiradi
5. Google `84.54.12.8` ni ko'radi

---

## Javob qaytganda nima bo'ladi?

Router jadval saqlaydi:

```
192.168.1.10:52341 → 84.54.12.8:52341
```

Response kelganda:

1. Router: "Bu kim yuborgandi?" deb jadvaldan qaraydi
2. Keyin `192.168.1.10` ga yuboradi

---

## NAT bo'lmaganida nima bo'lardi?

Har bir qurilma uchun Public IP kerak bo'lardi:

- Telefon
- Laptop
- Printer
- Smart TV

IPv4 allaqachon tugagan bo'lardi.

---

## Amaliyot

### NAT kuzatish

```bash
traceroute google.com
```
