# SSL vs TLS

SSL eski, TLS yangi protokol.

---

## Nazariya

### SSL (Secure Sockets Layer)

```text
SSL 1.0: xavfli, ishlatilmaydi
SSL 2.0: xavfli, ishlatilmaydi
SSL 3.0: xavfli, ishlatilmaydi
```

### TLS (Transport Layer Security)

```text
TLS 1.0: eski, ishlatilmaydi
TLS 1.1: eski, ishlatilmaydi
TLS 1.2: hali ishlatiladi
TLS 1.3: eng yangi, tezroq
```

---

## Diagram

```text
SSL                    TLS

SSL 1.0                TLS 1.0
SSL 2.0                TLS 1.1
SSL 3.0                TLS 1.2
                       TLS 1.3 (eng yangi)
```

---

## Amaliyot

### TLS versiyasini tekshirish

```bash
openssl s_client -connect google.com:443
```

Natijada:

```text
Protocol  : TLSv1.3
```

### TLS 1.2 ni majburlash

```bash
openssl s_client -connect google.com:443 -tls1_2
```

---

## Xulosa

- SSL eski va xavfli
- TLS 1.3 eng yangi va tez
- Hozir hammasi TLS ishlatadi
