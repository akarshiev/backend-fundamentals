# OpenSSL Amaliyoti

OpenSSL orqali SSL/TLS bilan ishlash.

---

## SSL ulanish

```bash
openssl s_client -connect google.com:443
```

---

## Sertifikat tafsilotlari

```bash
openssl x509 -in cert.pem -text -noout
```

---

## RSA key yaratish

```bash
openssl genrsa -out private.key 2048
```

---

## Public key chiqarish

```bash
openssl rsa -in private.key -pubout -out public.pem
```

---

## Sertifikat tekshirish

```bash
openssl verify -CAfile ca.pem cert.pem
```

---

## SSL test

```bash
openssl s_client -connect google.com:443 -tls1_2
openssl s_client -connect google.com:443 -tls1_3
```

---

## Xulosa

- OpenSSL SSL/TLS amaliyotlari uchun asosiy vosita
- Sertifikatlar, kalitlar va ulanishlarni boshqarish mumkin
