const CircuitBreaker = require("opossum");

// ❌ Noto'g'ri SMS servisi (tasodifiy xato bilan)
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
  timeout: 3000, // 3 sekund kutish vaqti
  errorThresholdPercentage: 50, // 50% xato bo'lsa → OPEN
  resetTimeout: 5000, // 5 sekund kutib, HALF OPEN ga o'tish
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

// 🚀 Test: Har soniyada 1 ta request
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
