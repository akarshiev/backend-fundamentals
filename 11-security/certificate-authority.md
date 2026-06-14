# Certificate Authority (CA)

HTTPS ishonch CA orqali quriladi.

---

## Zanjir

```text
Root CA (Implicitly trusted by OS/Browser)
    ↓
Intermediate CA
    ↓
Leaf Certificate (server's cert)
```

---

## Ishlash prinsipi

```text
Browser → Server
    ↓
Server → Leaf Certificate
    ↓
Browser → Intermediate CA tekshirish
    ↓
Intermediate → Root CA tekshirish
    ↓
Root CA → Trusted (OS/Browser da saqlangan)
    ↓
✅ Connection established
```

---

## Sertifikat tarkibiy qismi

| Qism | Tushuntirish |
|------|-------------|
| Subject | Kimga berilgan (domain) |
| Issuer | Kim bergan (CA) |
| Valid From/To | Amal qilish muddati |
| Public Key | Serverning public key |
| Signature | CA ning imzosi |
| SAN | Qo'shimcha domainlar |

---

## Amaliyot

### Sertifikat ko'rish

```bash
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -text
```

### Sertifikat zanjirini ko'rish

```bash
openssl s_client -connect google.com:443 -showcerts
```

### Sertifikat muddatini tekshirish

```bash
openssl s_client -connect google.com:443 </dev/null 2>/dev/null | openssl x509 -noout -dates
```

### O'z sertifikatingizni yaratish

```bash
# 1. Private key yaratish
openssl genrsa -out private.key 2048

# 2. CSR (Certificate Signing Request) yaratish
openssl req -new -key private.key -out request.csr

# 3. Self-signed sertifikat yaratish (test uchun)
openssl x509 -req -days 365 -in request.csr -signkey private.key -out cert.pem
```

---

## Self-signed vs CA-signed

| Xususiyat | Self-signed | CA-signed |
|-----------|-------------|-----------|
| Kim imzolaydi | O'zingiz | CA |
| Browser ishonchi | ❌ | ✅ |
| Ishlatilishi | Test/Development | Production |
| Narx | Bepul | Pullik |

---

## Xulosa

- CA zanjiri ishonch ta'minlaydi
- Root CA → Intermediate CA → Leaf Certificate
- Self-signed test uchun, CA-signed production uchun
