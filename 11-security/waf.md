# WAF (Web Application Firewall)

Application darajasidagi (Layer 7) himoya. SQL injection, XSS kabi hujumlarni to'xtatadi.

---

## WAF nima?

```text
WAF = Web Application Firewall

Firewall (L3/L4) — portlarni nazorat qiladi
WAF (L7) — HTTP so'rovlarini tekshiradi
```

---

## WAF qanday ishlaydi?

```text
Client → HTTP so'rov → WAF → Application Server

WAF:
- Request body ni tekshiradi
- SQL injection ni bloklaydi
- XSS ni bloklaydi
- Rate limiting qiladi
```

---

## SQL Injection

### Xavfli kod

```sql
-- User input: ' OR '1'='1
SELECT * FROM users WHERE username = '' OR '1'='1';
-- Barcha foydalanuvchilarni qaytaradi!
```

### Juda xavfli kod

```sql
-- User input: '; DROP TABLE users; --
DELETE * FROM users;
-- Jadvalni butunlay o'chiradi!
```

### SQL Injection turlari

| Tur | Usul |
|-----|------|
| In-band | Natija to'g'ridan-to'g'ri ko'rinadi |
| Blind | Natija bilvosita aniqlanadi |
| Union | UNION SELECT bilan ma'lumot olish |
| Time-based | Vaqt kechikishi orqali aniqlash |

---

## SQL Injection dan himoya

### 1. Parameterized Queries (Prepared Statements)

```python
# Xato ❌
query = f"SELECT * FROM users WHERE username = '{username}'"

# To'g'ri ✅
cursor.execute("SELECT * FROM users WHERE username = %s", (username,))
```

```go
// Xato ❌
query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", username)

// To'g'ri ✅
row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
```

### 2. ORM ishlatish

```python
# Django ORM
User.objects.filter(username=username)

# SQLAlchemy
session.query(User).filter(User.username == username).first()
```

### 3. Input Validation

```python
# Faqat ruxsat etilgan belgilarni qabul qilish
import re

def validate_username(username):
    if not re.match(r'^[a-zA-Z0-9_]+$', username):
        raise ValueError("Noto'g'ri username")
    return username
```

---

## WAF qoidalari

### Cloudflare WAF

```text
1. Cloudflare dashboard → Security → WAF
2. Custom rules → Create rule
3. Expression: http.request.uri.path contains "wp-admin"
4. Action: Block
```

### Nginx ModSecurity

```nginx
# ModSecurity ni yoqish
load_module modules/ngx_http_modsecurity_module.so;

server {
    modsecurity on;
    modsecurity_rules_file /etc/nginx/modsecurity/main.conf;
}
```

---

## XSS (Cross-Site Scripting)

### Xavfli kod

```html
<!-- User input: <script>alert('XSS')</script> -->
<div>Hello, <script>alert('XSS')</script></div>
```

### Himoya

```python
# HTML encoding
from markupsafe import escape

safe_output = escape(user_input)
```

```javascript
// JavaScript da
element.textContent = userInput;  // ✅ To'g'ri
element.innerHTML = userInput;    // ❌ Xato
```

---

## CSRF (Cross-Site Request Forgery)

### Xavfli kod

```html
<!-- Foydalanuvchi saytiga yuklangan yashirin forma -->
<form action="https://bank.com/transfer" method="POST">
  <input type="hidden" name="to" value="hacker">
  <input type="hidden" name="amount" value="10000">
</form>
<script>document.forms[0].submit();</script>
```

### Himoya

```python
# CSRF token
@app.route('/transfer', methods=['POST'])
def transfer():
    if request.form['csrf_token'] != session['csrf_token']:
        abort(403)
```

---

## Amaliyot

### SQL Injection test

```bash
# sqlmap bilan tekshirish (faqat o'z saytingizda!)
sqlmap -u "http://localhost/page?id=1" --batch

# Manual test
curl "http://localhost/page?id=1' OR '1'='1"
```

### WAF test

```bash
# ModSecurity test
curl -H "User-Agent: sqlmap/1.0" http://localhost/

# Cloudflare WAF test
curl -H "User-Agent: sqlmap/1.0" https://example.com/
```

---

## Xulosa

- WAF Layer 7 himoya — HTTP so'rovlarini tekshiradi
- SQL Injection parameterized queries bilan oldini olish
- XSS HTML encoding bilan himoya
- CSRF token bilan himoya
- Cloudflare, ModSecurity eng mashhur WAF vositalari
