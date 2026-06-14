# Raqamli Imzo (Digital Signature)

Xabarni imzolash va tekshirish usuli.

---

## Maqsad

1. **Kim yuborganini isbotlash** -- autentifikatsiya
2. **Ma'lumot o'zgarmaganini isbotlash** -- yaxlitlik

---

## Jarayon

### Imzo qo'yish (sign)

```text
Message
    ↓
Hash (SHA-256)
    ↓
Hash Value
    ↓
Private Key bilan sign
    ↓
Digital Signature
```

### Tekshirish (verify)

```text
Message
    ↓
Hash (SHA-256)
    ↓
Hash Value (1)
    ↓
Digital Signature
    ↓
Public Key bilan verify
    ↓
Hash Value (2)
    ↓
1 == 2 ? → Valid / Invalid
```

---

## Amaliyot

### RSA imzo yaratish (Python)

```python
from cryptography.hazmat.primitives.asymmetric import rsa, padding
from cryptography.hazmat.primitives import hashes, serialization

# Kalit yaratish
private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
)
public_key = private_key.public_key()

# Xabar
message = b"This is an important message."

# Imzo qo'yish
signature = private_key.sign(
    message,
    padding.PSS(
        mgf=padding.MGF1(hashes.SHA256()),
        salt_length=padding.PSS.MAX_LENGTH
    ),
    hashes.SHA256()
)

print(f"Signature: {signature.hex()[:64]}...")

# Tekshirish
try:
    public_key.verify(
        signature,
        message,
        padding.PSS(
            mgf=padding.MGF1(hashes.SHA256()),
            salt_length=padding.PSS.MAX_LENGTH
        ),
        hashes.SHA256()
    )
    print("✅ Valid signature")
except Exception:
    print("❌ Invalid signature")

# Xabar o'zgarsa
try:
    public_key.verify(
        signature,
        b"This is a DIFFERENT message.",
        padding.PSS(
            mgf=padding.MGF1(hashes.SHA256()),
            salt_length=padding.PSS.MAX_LENGTH
        ),
        hashes.SHA256()
    )
    print("✅ Valid signature")
except Exception:
    print("❌ Invalid signature (message changed)")
```

### Go da imzo

```go
package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func main() {
	// Kalit yaratish
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Xabar
	message := []byte("Important message to sign")

	// Hash
	hash := sha256.Sum256(message)

	// Imzo qo'yish
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		panic(err)
	}

	fmt.Printf("Signature: %x\n", signature[:32])

	// Tekshirish
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		fmt.Println("❌ Invalid signature")
	} else {
		fmt.Println("✅ Valid signature")
	}
}
```

---

## Xulosa

- Digital Signature xabarni imzolash va tekshirish usuli
- Private Key bilan imzo qo'yiladi
- Public Key bilan tekshiriladi
- Hash algoritmi (SHA-256) ishlatiladi
