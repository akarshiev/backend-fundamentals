# Circuit Breaker

Servis yiqilganda butun tizimning yiqilishini to'xtatish. Elektr avtomati kabi ishlaydi.

---

## Muammo: Cascading Failure

```text
User
 ↓
API Gateway
 ↓
Auth Service
 ↓
User Service
 ↓
SMS Service

SMS Provider: Eskiz / PlayMobile / Twilio
```

Bir kuni **SMS Service** o'lib qoldi.

### Normal developer nima qiladi?

```text
sendSMS() chaqiraveradi.

Request #1  → Timeout
Request #2  → Timeout
Request #3  → Timeout
Request #10000 → Timeout

Threadlar to'lib ketadi.
Connection pool to'lib ketadi.
CPU kutishda qoladi.
Memory oshadi.

Keyin:
SMS → User Service → Auth Service → Gateway → Whole System

HAMMA YIQILADI.
```

Bu **Cascading Failure** deyiladi. Bir servis yiqilib, domino effekti bilan qolganlarini ham yiqitadi.

---

## Latency nima?

```text
Request yuborildi → Javob keldi orasidagi vaqt.

10ms   ✅ Yaxshi
50ms   ✅ Normal
200ms  ⚠️ Sekin
1s     ❌ Juda sekin
5s     ❌ Xato
```

SMS provider **50ms** o'rniga **10 sec** javob bersa, sizning thread **10 sec** bekor kutadi.

Shuning uchun **Timeout** qo'yiladi.

---

## Circuit Breaker

Elektrdagi avtomatni bilasiz. Tok ko'payib ketsa, **avtomat tushadi**.

Backend'da ham xuddi shunday. Agar servis ishlamayotgan bo'lsa, **u bilan gaplashmaymiz**.

---

## State 1: CLOSED (Normal)

```text
Request → Service → Javob keladi

Hammasi ishlayapti.
CLOSED holatida barcha requestlar o'tadi.
```

```text
Client
  ↓
[CLOSED] ✅
  ↓
SMS Provider
```

---

## State 2: OPEN (Xatolar ko'paydi)

```text
50% requestlar yiqildi.

Circuit Breaker: "Bu servis o'libdi" deydi.
OPEN holatiga o'tadi.

Endi requestlar SMS Provider ga umuman bormaydi.
Darhol FAIL FAST qiladi.
```

```text
Client
 ↓
[OPEN] ❌ BLOCK
```

Bu juda muhim. **Timeout kutmaydi**, darhol xatolik qaytaradi.

---

## State 3: HALF-OPEN (Sinov)

```text
Bir oz kutamiz (masalan, 5 second).
Keyin 1 ta request yuboramiz.

Agar ishlasa → CLOSED ga qaytamiz.
Agar yiqilsa → OPEN ga qaytamiz.
```

```text
OPEN
 ↓
5 sec kutish
 ↓
HALF OPEN
 ↓
Test Request
 ↓
Success?
 ├── Yes → CLOSED ✅
 └── No  → OPEN ❌
```

---


![Circuit Breaker States](https://martinfowler.com/bliki/images/circuitBreaker/state.png)

*Circuit Breaker holatlari: CLOSED → OPEN → HALF-OPEN*

## Circuit Breaker Lifecycle

```text
CLOSED
   ↓
Error Rate High (50%+)
   ↓
OPEN
   ↓
Reset Timeout (5 sec)
   ↓
HALF OPEN
   ↓
Test Request
   ↓
Success?
   ├── Yes → CLOSED
   └── No  → OPEN
```

---

## Fallback

SMS ketmadi, lekin tizim ishlashi kerak.

```text
SMS → Email ga o'tamiz.

Bu Fallback deyiladi.
```

### Fallback turlari

| Tur | Misol |
|-----|-------|
| Email fallback | SMS o'rniga email yuborish |
| Cache fallback | Database o'rniga cache dan olish |
| Default fallback | Oldindan belgilangan qiymat qaytarish |
| Retry fallback | Boshqa providerga urinish |

---

## Nima uchun kerak?

### Circuit Breaker bo'lmasa

```text
10 000 login
 ↓
10 000 timeout
 ↓
Auth API ham o'ladi ❌
```

### Circuit Breaker bilan

```text
10 000 login
 ↓
OPEN (FAIL FAST)
 ↓
Email OTP yuboriladi ✅

Tizim ishlashda davom etadi.
```

---

## Node.js Amaliyot

### O'rnatish

```bash
mkdir circuit-breaker-example && cd circuit-breaker-example
npm init -y
npm install opossum
```

### Asosiy kod: `index.js`

```javascript
const CircuitBreaker = require("opossum");

// ❌ Noto'g'ri SMS servisi (tasodifiy xato)
function unstableSmsService(message) {
  return new Promise((resolve, reject) => {
    const random = Math.random();

    if (random > 0.5) {
      resolve(`SMS yuborildi: ${message}`);
    } else {
      reject(new Error("SMS provider ishlamayapti"));
    }
  });
}

// ⚙️ Circuit Breaker sozlash
const options = {
  timeout: 3000,              // 3 sekund kutish
  errorThresholdPercentage: 50, // 50% xato bo'lsa → OPEN
  resetTimeout: 5000          // 5 sekund kutib, HALF OPEN ga o'tish
};

// 🔌 Circuit Breaker yaratish
const breaker = new CircuitBreaker(unstableSmsService, options);

// 🔄 Fallback: SMS ishlamasa, Email yuborish
breaker.fallback(() => {
  return "Fallback: Email yuborildi";
});

// 📢 Event listenerlar
breaker.on("open", () => {
  console.log("🔴 OPEN — Circuit Breaker ochildi!");
});

breaker.on("halfOpen", () => {
  console.log("🟡 HALF OPEN — Sinov so'rovi yuborilmoqda...");
});

breaker.on("close", () => {
  console.log("🟢 CLOSED — Servis qayta ishlamoqda!");
});

breaker.on("reject", () => {
  console.log("⛔ FAIL FAST — So'rov rad etildi!");
});

// Test: Har soniyada 1 ta request
let counter = 1;

setInterval(async () => {
  console.log(`\n--- Request #${counter++} ---`);

  try {
    const result = await breaker.fire("Salom, dunyo!");
    console.log("✅", result);
  } catch (err) {
    console.log("❌", err.message);
  }
}, 1000);
```

### Ishga tushirish

```bash
node index.js
```

### Namuna chiqishi

```text
--- Request #1 ---
✅ SMS yuborildi: Salom, dunyo!

--- Request #2 ---
❌ SMS provider ishlamayapti

--- Request #3 ---
🔴 OPEN — Circuit Breaker ochildi!
⛔ FAIL FAST — So'rov rad etildi!

--- Request #4 ---
⛔ FAIL FAST — So'rov rad etildi!

--- Request #5 ---
⛔ FAIL FAST — So'rov rad etildi!

--- Request #6 ---
🟡 HALF OPEN — Sinov so'rovi yuborilmoqda...
✅ SMS yuborildi: Salom, dunyo!
🟢 CLOSED — Servis qayta ishlamoqda!
```

---

## Timeout Misoli

```javascript
function slowService() {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve("OK");
    }, 10000); // 10 sekund
  });
}

const breaker = new CircuitBreaker(slowService, {
  timeout: 3000 // 3 sekund kutish
});

// 3 sekund kutadi, keyin TIMEOUT deb hisoblaydi.
```

---

## Production Misol

```text
Login:
Auth API → OTP API → SMS Provider

SMS provider o'ldi.

Circuit Breaker bo'lmasa:
10 000 login → 10 000 timeout → Auth API ham o'ladi ❌

Circuit Breaker bilan:
10 000 login → OPEN → Email OTP → Tizim ishlaydi ✅
```

---

## Xulosa

| Xususiyat | Tushuntirish |
|-----------|-------------|
| Cascading Failure | Bir servis yiqilsa, hamma yiqiladi |
| CLOSED | Normal holat — hamma o'tadi |
| OPEN | Xato ko'p — darhol bloklaydi |
| HALF-OPEN | Sinov — 1 ta request yuboradi |
| Fallback | Asosiy servis o'lsa, zaxira ishlaydi |
| Timeout | Kutish vaqtini cheklash |
| Fail Fast | Kutmasdan darhol javob berish |
