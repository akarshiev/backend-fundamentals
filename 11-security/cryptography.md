# Kriptografiya asoslari

Ma'lumotni shifrlash va himoya qilish usullari.

---

## Simmetrik Shifrlash

Bir xil kalit ishlatiladi:

```text
Plain Text
    ↓
AES Encrypt (kalit)
    ↓
Cipher Text
    ↓
AES Decrypt (kalit)
    ↓
Plain Text
```

### Misollar

| Algoritm | Kalit o'lchami | Xususiyati |
|----------|---------------|------------|
| AES-128 | 128 bit | Tez, keng tarqalgan |
| AES-256 | 256 bit | Juda xavfsiz |
| ChaCha20 | 256 bit | Mobil qurilmalar uchun |
| 3DES | 168 bit | Eski, sekin |

### AES-256

256 bitli kalit ishlatadi.

**Afzalliklari:**
- Juda tez
- Katta hajmdagi ma'lumotlar uchun mos
- Hardwer tezlashtirish qo'llab-quvvatlaydi

**Kamchiligi:**
- Kalitni xavfsiz uzatish kerak
- Kalit almashish muammosi

---

## Assimmetrik Shifrlash

Ikki kalit ishlatiladi:

```text
Plain Text
    ↓
Public Key (Encrypt)
    ↓
Cipher Text
    ↓
Private Key (Decrypt)
    ↓
Plain Text
```

### Misollar

| Algoritm | Xususiyati |
|----------|------------|
| RSA | Keng tarqalgan, 2048+ bit |
| ECC | Kalit kichikroq, tezroq |
| Ed25519 | Zamonaviy, xavfsiz |

### RSA ishlash prinsipi

```text
1. Ikkita tub son yaratiladi (p, q)
2. n = p × q
3. Public Key = (e, n)
4. Private Key = (d, n)
5. Encrypt: c = m^e mod n
6. Decrypt: m = c^d mod n
```

---

## Simmetrik vs Assimmetrik

| Xususiyat | Simmetrik | Assimmetrik |
|-----------|-----------|-------------|
| Kalit soni | 1 | 2 |
| Tezlik | Tez | Sekin |
| Kalit almashish | Qiyin | Oson |
| Ishlatilishi | Ma'lumot shifrlash | Kalit almashish |

### Amaliy usul

Odatda **ikkalasi birga** ishlatiladi:

```text
1. Assimmetrik bilan kalit almashish
2. Simmetrik bilan ma'lumot shifrlash
```

Bu TLS/HTTPS da ishlatiladigan usul.

---

## Amaliyot

### AES-256 shifrlash (Python)

```python
from cryptography.hazmat.primitives.ciphers import Cipher, algorithms, modes
from cryptography.hazmat.backends import default_backend
import os

# Kalit va IV yaratish
key = os.urandom(32)  # 256 bit
iv = os.urandom(16)   # 128 bit

# Shifrlash
cipher = Cipher(algorithms.AES(key), modes.CBC(iv), backend=default_backend())
encryptor = cipher.encryptor()

plaintext = b"Hello, World!"
# Padding qo'shish (AES blok o'lchami 16 byte)
padded = plaintext + b" " * (16 - len(plaintext) % 16)
ciphertext = encryptor.update(padded) + encryptor.finalize()

print(f"Ciphertext: {ciphertext.hex()}")

# Shifrdan chiqarish
decryptor = Cipher(algorithms.AES(key), modes.CBC(iv), backend=default_backend()).decryptor()
decrypted = decryptor.update(ciphertext) + decryptor.finalize()
print(f"Decrypted: {decrypted.strip().decode()}")
```

### RSA kalit yaratish (Python)

```python
from cryptography.hazmat.primitives.asymmetric import rsa, padding
from cryptography.hazmat.primitives import hashes

# RSA kalit yaratish
private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
)

public_key = private_key.public_key()

# Shifrlash (public key bilan)
message = b"Hello, RSA!"
ciphertext = public_key.encrypt(
    message,
    padding.OAEP(
        mgf=padding.MGF1(algorithm=hashes.SHA256()),
        algorithm=hashes.SHA256(),
        label=None
    )
)

# Shifrdan chiqarish (private key bilan)
plaintext = private_key.decrypt(
    ciphertext,
    padding.OAEP(
        mgf=padding.MGF1(algorithm=hashes.SHA256()),
        algorithm=hashes.SHA256(),
        label=None
    )
)
print(f"Decrypted: {plaintext.decode()}")
```

### Go da AES shifrlash

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	// Kalit yaratish (32 byte = 256 bit)
	key := make([]byte, 32)
	rand.Read(key)

	// IV yaratish
	plaintext := []byte("Hello, AES!")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	// Shifrlash
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// Shifrdan chiqarish
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		panic("ciphertext too short")
	}

	nonce2, ciphertext2 := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext2, err := gcm.Open(nil, nonce2, ciphertext2, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted: %s\n", plaintext2)
}
```

---

## Xulosa

- Simmetrik -- bir xil kalit, tez
- Assimmetrik -- ikki kalit, xavfsiz
- Amaliyda ikkalasi birga ishlatiladi
- AES-256 va RSA eng mashhur
