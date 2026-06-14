# OWASP TOP 10

Eng ko'p uchraydigan xavfsizlik zaifliklari.

---

## OWASP TOP 10 nima?

```text
OWASP = Open Web Application Security Project

TOP 10 = Eng ko'p uchraydigan 10 ta xavfsizlik zaiflik

Har yili yangilanadi (2021 versiyasi)
```

---

## A01: Broken Access Control

### Tushuntirish

```text
Foydalanuvchi o'ziga ruxsati yo'q resurslarga kirishi mumkin.

Masalan:
- Admin paneliga oddiy foydalanuvchi kirishi
- Boshqa foydalanuvchi ma'lumotlarini ko'rish
- IDOR (Insecure Direct Object Reference)
```

### Misol

```python
# Xato ❌
@app.route('/user/<int:user_id>')
def get_user(user_id):
    user = User.query.get(user_id)  # Har kim istalgan user ni ko'ra oladi
    return user.to_dict()

# To'g'ri ✅
@app.route('/user/<int:user_id>')
@login_required
def get_user(user_id):
    if current_user.id != user_id and not current_user.is_admin:
        abort(403)
    user = User.query.get(user_id)
    return user.to_dict()
```

### Himoya

```python
# 1. Authorization check
if not user.has_permission('read', resource):
    abort(403)

# 2. Role-based access
@app.route('/admin')
@admin_required
def admin_panel():
    pass

# 3. IDOR oldini olish
@app.route('/order/<int:order_id>')
@login_required
def get_order(order_id):
    order = Order.query.get(order_id)
    if order.user_id != current_user.id:
        abort(403)
```

---

## A03: Injection

### Tushuntirish

```text
Foydalanuvchi kiritgan ma'lumot serverda kod sifatida bajariladi.

Turlari:
- SQL Injection
- NoSQL Injection
- Command Injection
- LDAP Injection
```

### SQL Injection misol

```sql
-- Xavfli kod
SELECT * FROM users WHERE username = '' OR '1'='1';
-- Barcha foydalanuvchilarni qaytaradi!

-- Juda xavfli kod
DELETE * FROM users;
-- Jadvalni butunlay o'chiradi!
```

### Himoya

```python
# 1. Parameterized Queries
cursor.execute("SELECT * FROM users WHERE username = %s", (username,))

# 2. ORM
User.objects.filter(username=username)

# 3. Input Validation
import re
if not re.match(r'^[a-zA-Z0-9_]+$', username):
    raise ValueError("Noto'g'ri username")
```

---

## A07: Identification and Authentication Failures

### Tushuntirish

```text
Autentifikatsiya tizimidagi zaifliklar.

Masalan:
- Kuchsiz parollar
- MFA yo'q
- Session boshqaruvi zaif
```

### Kuchsiz parol misol

```python
# Xato ❌
password = "123456"
password = "password"
password = "admin"

# To'g'ri ✅
# Kuchli parol talablari:
# - Kamida 8 ta belgi
# - Katta-kichik harflar
# - Raqamlar
# - Maxsus belgilar
```

### LocalStorage xavfsizligi

```text
⚠️ LocalStorage saqlanmaydi!

JavaScript kirishi mumkin:
localStorage.getItem('token')  # Oson o'g'irlash

Xavfsiz alternativlar:
1. HttpOnly Cookie (eng yaxshi)
2. In-memory (SPA uchun)
3. Session storage (vaqtincha)
```

### Himoya

```python
# 1. MFA (Multi-Factor Authentication)
import pyotp
totp = pyotp.TOTP('base32secret')
if not totp.verify(user_code):
    return "Xato kod!"

# 2. Parolni xashlash
from werkzeug.security import generate_password_hash, check_password_hash
hashed = generate_password_hash(password)
if not check_password_hash(hashed, input_password):
    return "Xato parol!"

# 3. JWT token
import jwt
token = jwt.encode({'user_id': user.id}, secret_key, algorithm='HS256')
```

### JWT Token xavfsiz saqlash

```text
❌ Xato: localStorage ga saqlash
   localStorage.setItem('token', token)
   → XSS hujumida o'g'irlash mumkin

✅ To'g'ri: HttpOnly Cookie
   Set-Cookie: token=xxx; HttpOnly; Secure; SameSite=Strict
   → JavaScript o'qiy olmaydi
```

### HttpOnly Cookie sozlash

```python
# Flask da
@app.route('/login', methods=['POST'])
def login():
    token = create_token(user)
    response = make_response(redirect('/dashboard'))
    response.set_cookie(
        'token',
        token,
        httponly=True,      # JavaScript kirish mumkin emas
        secure=True,       # Faqat HTTPS
        samesite='Strict', # CSRF dan himoya
        max_age=3600       # 1 soat
    )
    return response
```

---

## A05: Security Misconfiguration

### Tushuntirish

```text
Xavfsizlik sozlamalaridagi xatolar.

Masalan:
- Default parollar
- Xato portlar
- Debug mode yoqilgan
```

### Xato misol

```python
# Xato ❌
DEBUG = True  # Production da!
SECRET_KEY = "secret"  # Oson taxmin qilinadi

# To'g'ri ✅
DEBUG = False
SECRET_KEY = os.environ.get('SECRET_KEY')
```

### Himoya

```bash
# 1. Debug mode ni o'chirish
export FLASK_ENV=production

# 2. Secret key
export SECRET_KEY=$(openssl rand -hex 32)

# 3. Portlarni tekshirish
ss -tulpn

# 4. Firewall sozlash
sudo ufw default deny incoming
sudo ufw allow 22/tcp
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
```

---

## A06: Vulnerable and Outdated Components

### Tushuntirish

```text
Eski va zaif komponentlar ishlatish.

Masalan:
- Eski kutubxonalar
- Tanish zaifliklar
```

### Tekshirish

```bash
# npm audit
npm audit

# Python
pip-audit

# Go
govulncheck ./...
```

### Himoya

```bash
# 1. Dependency ni yangilash
npm update

# 2. Zaifliklarni tekshirish
npm audit fix

# 3. Lock file saqlash
npm install --save-exact
```

---

## A04: Insecure Design

### Tushuntirish

```text
Xavfsizlik loyihasidagi zaifliklar.

Masalan:
- Threat modeling yo'q
- Xavfsizlik talablari aniqlanmagan
```

### Himoya

```text
1. Threat Modeling
   - STRIDE (Spoofing, Tampering, Repudiation, Info Disclosure, DoS, Elevation)
   
2. Security Requirements
   - Autentifikatsiya
   - Autorizatsiya
   - Ma'lumotlarni shifrlash
   
3. Secure Design Patterns
   - Defense in depth
   - Least privilege
   - Fail securely
```

---

## Amaliyot

### OWASP ZAP test

```bash
# OWASP ZAP ni ishga tushirish
docker run -u zap -p 8080:8080 -p 8090:8090 -i ghcr.io/zaproxy/zaproxy:stable zap-webswing.sh

# Zap CLI
zap-cli quick-scan http://localhost:8080
```

### npm audit test

```bash
# Zaifliklarni tekshirish
npm audit

# Avtomatik tuzatish
npm audit fix

# Qo'lda tuzatish
npm audit fix --force
```

---

## Xulosa

| OWASP | Muammo | Yechim |
|-------|--------|--------|
| A01 | Broken Access Control | Authorization check |
| A03 | Injection | Parameterized queries |
| A05 | Security Misconfiguration | Debug mode off, strong secrets |
| A06 | Vulnerable Components | npm audit, update |
| A07 | Auth Failures | MFA, strong passwords |
